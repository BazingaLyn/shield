package protocol

import (
	"encoding/binary"
	"errors"
	"github.com/vmihailenco/msgpack"
	"io"
	"shield/rpc/codec"
)

type MessageType byte

const (
	MessageTypeRequest MessageType = iota
	MessageTypeResponse
)

type CompressType byte

const (
	CompressTypeNone CompressType = iota
)

type StatusCode byte

const (
	StatusOK StatusCode = iota
	StatusError
)

type ProtocolType byte

const (
	Default ProtocolType = iota
)

type Header struct {
	Seq           uint64
	MessageType   MessageType
	CompressType  CompressType
	SerializeType codec.SerializeType
	StatusCode    StatusCode
	ServiceName   string
	MethodName    string
	Error         string
	MetaData      map[string]string
}

type Message struct {
	*Header
	Data []byte
}

func (m Message) Clone() *Message {
	header := *m.Header
	c := new(Message)
	c.Header = &header
	c.Data = m.Data
	return c
}

type Protocol interface {
	NewMessage() *Message
	DecodeMessage(r io.Reader) (*Message, error)
	EncodeMessage(message *Message) []byte
}

var protocols = map[ProtocolType]Protocol{
	Default: &RPCProtocol{},
}

func DecodeMessage(t ProtocolType, r io.Reader) (*Message, error) {
	return protocols[t].DecodeMessage(r)
}

type RPCProtocol struct {
}

func (RPCProtocol) NewMessage() *Message {
	return &Message{Header: &Header{}}
}

func EncodeMessage(t ProtocolType, m *Message) []byte {
	return protocols[t].EncodeMessage(m)
}

func (RPCProtocol) DecodeMessage(r io.Reader) (msg *Message, err error) {

	first3bytes := make([]byte, 3)
	_, err = io.ReadFull(r, first3bytes)
	if err != nil {
		return
	}
	if !checkMagic(first3bytes[:2]) {
		err = errors.New("wrong protocol")
		return
	}

	totalLenBytes := make([]byte, 4)
	_, err = io.ReadFull(r, totalLenBytes)

	if err != nil {
		return
	}

	totalLen := int(binary.BigEndian.Uint32(totalLenBytes))
	if totalLen < 4 {
		err = errors.New("invalid total length")
		return
	}

	data := make([]byte, totalLen)
	_, err = io.ReadFull(r, data)
	headerLen := int(binary.BigEndian.Uint32(data[:4]))
	headerBytes := data[4 : headerLen+4]
	header := &Header{}
	err = msgpack.Unmarshal(headerBytes, header)
	if err != nil {
		return
	}
	msg = new(Message)
	msg.Header = header
	msg.Data = data[headerLen+4:]

	return msg, err
}

func checkMagic(bytes []byte) bool {
	return bytes[0] == 0xab && bytes[1] == 0xba
}

func (RPCProtocol) EncodeMessage(message *Message) []byte {
	first3bytes := []byte{0xab, 0xba, 0x00}

	bytes, _ := msgpack.Marshal(message.Header)
	totalLen := 4 + len(bytes) + len(message.Data)
	totalLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(totalLenBytes, uint32(totalLen))

	data := make([]byte, totalLen+7)
	start := 0
	copyFullWithOffset(data, first3bytes, &start)
	copyFullWithOffset(data, totalLenBytes, &start)

	headerLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(headerLenBytes, uint32(len(bytes)))
	copyFullWithOffset(data, headerLenBytes, &start)
	copyFullWithOffset(data, bytes, &start)
	copyFullWithOffset(data, message.Data, &start)
	return data
}

func copyFullWithOffset(dst []byte, src []byte, start *int) {
	copy(dst[*start:*start+len(src)], src)
	*start = *start + len(src)
}
