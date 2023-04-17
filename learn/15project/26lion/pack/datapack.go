package pack

import (
	"26lion/iface"
	"bytes"
	"encoding/binary"
)

type DataPack struct {
}

var defaultHeaderLen uint32 = 8

var _ iface.IDataPack = (*DataPack)(nil)

func NewDataPack() iface.IDataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	return defaultHeaderLen
}

func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	dataBuffer := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuffer, binary.BigEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuffer, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuffer, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuffer.Bytes(), nil
}
func (dp *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {

	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.BigEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
