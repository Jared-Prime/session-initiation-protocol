package sip

type Peers map[string]Endpoint

type Endpoint struct {
  User    string
  Address string
}
