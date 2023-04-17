package iface

type IDataPack interface {
	GetHeadLen() uint32                //获取包头长度方法
	Pack(msg IMessage) ([]byte, error) //封包方法
	Unpack([]byte) (IMessage, error)   //拆包方法
}

const (
	LionDataPack string = "lion_pack"

	//自定义封包方式在此添加
)

const (
	//Zinx 默认标准报文协议格式
	LionMessage string = "lion_message"
)
