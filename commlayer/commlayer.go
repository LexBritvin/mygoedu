package commlayer

import (
	"mygoedu/commlayer/myproto"
)

const (
	Protobuf uint8 = iota
)

type MyConnection interface {
	EncodeAndSend(obj interface{}, destinaton string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

func NewConnection(connType uint8) MyConnection {
	switch connType {
	case Protobuf:
		return myproto.NewProtoHandler()
	}
	return nil
}
