// Copywright 2013 Jared Davis. All rights reserved.

package sip

import (
  "bufio"
  "io"
)

// SIP request parsing errors.
type ProtocolError struct {
	ErrorString string
}

func (err *ProtocolError) Error() string { return err.ErrorString }

var (
   RecipientMissing = &ProtocolError{"recipient cannot be empty"}
   InitiatorMissing = &ProtocolError{"initiator cannot be empty"}
   ToBranchMissing  = &ProtocolError{"branch required for 'To:' header"}
   FromTagMissing   = &ProtocolError{"tag required for 'From:' header"}
   CallIDMissing    = &ProtocolError{"callid cannot be empty"}
)

type Request struct {
  Method string            // eg. INVITE
  Branch string
  Tag    int64
  CallID string            // zxcvb12345
  Peers  Peers
  Header Header
}

// Return value if nonempty, def otherwise
func valueOrDefault(value, def string) string {
  if value != "" {
    return value
  }
  return def
}

func (r *Request) Write(w io.Writer) error {
  return r.write(w, nil)
}

func (req *Request) write(w io.Writer, extraHeaders Header) error {
  var bw *bufio.Writer

  if _, ok := w.(io.ByteWriter); !ok {
    bw = bufio.NewWriter(w)
    w = bw
  }

  return nil
}

func (req *Request) SetTo(toURI Endpoint) error {
  req.Header["To"] = toURI
  return nil
}

func (req *Request) GetTo() (Endpoint, error) {
  return req.Header["To"], nil
}

func (req *Request) SetFrom(fromURI Endpoint) error {
  req.Header["From"] = fromURI
  return nil
}

func (req *Request) GetFrom() (Endpoint, error) {
  return req.Header["From"], nil
}

func NewRequest(method, branch string, tag int64, callid string) (*Request, error) {

  req := &Request{
          Method:     method,
          Branch:     branch,
          Tag:        tag,
          CallID:     callid,
          Peers:      make(Peers),
          Header:     make(Header),
  }

  return req, nil
}
