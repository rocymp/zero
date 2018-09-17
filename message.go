package zero

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/adler32"
)

// Message struct
type Message struct {
	msgSize  int32
	cmdID    int32
	data     []byte
	checksum uint32
}

// NewMessage create a new message
func NewMessage(cmdID int32, data []byte) *Message {
	msg := &Message{
		msgSize: int32(len(data)) + 4 + 4,
		cmdID:   cmdID,
		data:    data,
	}

	msg.checksum = msg.calcChecksum()
	return msg
}

// GetData get message data
func (msg *Message) GetData() []byte {
	return msg.data
}

// GetCMD get message ID
func (msg *Message) GetCMD() int32 {
	return msg.cmdID
}

// Verify verify checksum
func (msg *Message) Verify() bool {
	return msg.checksum == msg.calcChecksum()
}

func (msg *Message) calcChecksum() uint32 {
	if msg == nil {
		return 0
	}

	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, msg.cmdID)
	if err != nil {
		return 0
	}
	err = binary.Write(data, binary.LittleEndian, msg.data)
	if err != nil {
		return 0
	}

	checksum := adler32.Checksum(data.Bytes())
	return checksum
}

func (msg *Message) String() string {
	return fmt.Sprintf("Size=%d CMD=%d DataLen=%d Checksum=%d", msg.msgSize, msg.GetCMD(), len(msg.GetData()), msg.checksum)
}
