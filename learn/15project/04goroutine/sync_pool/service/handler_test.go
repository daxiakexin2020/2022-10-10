package service

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bufb := bufferPool.Get().(*bytes.Buffer)
		bufb.Write(data)
		bufb.Reset()
		bufferPool.Put(bufb)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var bufb bytes.Buffer
		bufb.Write(data)
	}
}
