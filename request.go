// Copywright 2013 Jared Davis. All rights reserved.

package sip

// SIP request parsing errors.
type ProtocolError struct {
	ErrorString string
}

// SIP request parsing errors.
type ProtocolError struct {
	ErrorString string
}

type Request struct {
  // INVITE sip:user2@server2.com SIP/2.0
  Method string // eg. INVITE
  Uri string // eg. sip:user2@server2.com
  Version string // eg. SIP/2.0
  // Via: SIP/2.0/UDP pc.server.com;branch=1234567890asdf Max-Forwards: 70
  Transport string // eg. UDP
  // To: user2 <sip:user2@server2.com>
  To string
  // From: user1 <sip:user1@server1.com>; tag=1234567890
  From string
  // Call-ID: zxcvb12345@pc.server1.com
  CallID string
  // CSeq: 314159 INVITE
  CSeq string
  // Contact: <sip:user1@pc.server1.com>
  Contact string
  // Content-Type: application/sdp
  ContentType int64
  // Content-Length: 42
  ContentLength string
  
}
