package main

import (
  "github.com/Jared-Prime/session-initiation-protocol/sip"
  "fmt"
  "bytes"
)

func main() {
  req, err := sip.NewRequest("INVITE","user2@server2.com","user1@server2.com","12345678",0,"abc","catapult")
  if err != nil {
    fmt.Printf("error: %s", err)
    return
  }

  message := new(bytes.Buffer)
  sip.Write(*message)
  message = WriteFrom(sip.Write(*message))

  fmt.Printf(message.String())
}
