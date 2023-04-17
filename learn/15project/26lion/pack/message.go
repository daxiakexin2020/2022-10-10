package pack

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
	rawData []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
		rawData: data,
	}
}

func NewMessage(len uint32, data []byte) *Message {
	return &Message{
		DataLen: len,
		Data:    data,
		rawData: data,
	}
}

func NewMessageByMsgId(id uint32, len uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: len,
		Data:    data,
		rawData: data,
	}
}

func (msg *Message) Init(ID uint32, data []byte) {
	msg.Id = ID
	msg.Data = data
	msg.rawData = data
	msg.DataLen = uint32(len(data))
}

func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) GetMsgID() uint32 {
	return msg.Id
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) GetRawData() []byte {
	return msg.rawData
}

func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

func (msg *Message) SetMsgID(msgID uint32) {
	msg.Id = msgID
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
