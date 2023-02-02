package protocol

var nullBulkBytes = []byte("$-1\r\n")
var emptyMultiBulkBytes = []byte("*0\r\n")

type NullBulkReply struct{}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

func (r *NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

type EmptyMultiBulkReply struct{}

func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

func MakeEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return &EmptyMultiBulkReply{}
}
