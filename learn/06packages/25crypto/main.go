package main

import (
	"crypto"
	"crypto/aes"
	"fmt"
	"hash"
)

func main() {
	/**
	crypto包搜集了常用的密码（算法）常量。
	*/

	crypto.RegisterHash(crypto.BLAKE2b_256, func() hash.Hash {
		return nil
	})

	//key的长度必须是16、24、32字节
	cipher, err := aes.NewCipher([]byte("1234567890123456"))
	fmt.Println("aes.NewCipher:", cipher, err)

	dst := []byte("abcdefghi12312345")
	src := []byte("1211111345678901234561")

	cipher.Encrypt(src, dst)
	fmt.Println("cipher.Encrypt dst:", dst, src, string(dst), string(src)) //[216 30 113 24 118 203 123 128 32 221 21 203 139 190 123 64 50 51 52 53 54 49] �qv�{� �ˋ�{@234561

	//cipher.Decrypt(src, dst)
	fmt.Println("cipher.Decrypt dst:", dst, src, string(dst), string(src))
}
