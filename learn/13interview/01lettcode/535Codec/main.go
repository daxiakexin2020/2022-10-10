package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	c := Constructor()
	lognUrl := "www.baidu.com"

	encode := c.encode(lognUrl)
	fmt.Println("encode:", encode)

	decode := c.decode(encode)
	fmt.Println("decode:", decode)
}

type Codec struct {
	urls map[string]*url
}

type url struct {
	original string
	tiny     string
}

const tinyUrlPre = "http://tinyurl.com/"

func Constructor() Codec {
	return Codec{map[string]*url{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	key := generateKey(longUrl)
	if u, ok := this.urls[key]; ok {
		return u.tiny
	}
	tinyUrl := tinyUrlPre + key
	u := &url{tiny: tinyUrl, original: longUrl}
	this.urls[key] = u
	return tinyUrl
}

func (this *Codec) decode(shortUrl string) string {
	index := strings.LastIndex(shortUrl, "/")
	key := shortUrl[index+1:]
	if u, ok := this.urls[key]; ok {
		return u.original
	}
	return ""
}

func generateKey(url string) string {
	// 可以使用Hash接口进行操作
	b := []byte(url)
	hasher := md5.New()
	hasher.Write(b)
	return fmt.Sprintf("%x\n", hasher.Sum(nil))
}
