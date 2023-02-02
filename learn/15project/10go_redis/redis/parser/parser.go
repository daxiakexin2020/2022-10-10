package parser

import (
	"10go_redis/interface/redis"
	"10go_redis/redis/protocol"
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
)

type Payload struct {
	Data redis.Reply
	Err  error
}

/*
*
ParseStream 通过 io.Reader 读取数据并将结果通过 channel 将结果返回给调用者
流式处理的接口适合供客户端/服务端使用
*/
func ParseStream(reader io.Reader) <-chan *Payload {
	ch := make(chan *Payload)
	go parse0(reader, ch)
	return ch
}

/*
*
RESP 通过第一个字符来表示格式：

简单字符串：以"+" 开始， 如："+OK\r\n"
错误：以"-" 开始，如："-ERR Invalid Synatx\r\n"
整数：以":"开始，如：":1\r\n"
字符串：以 $ 开始
数组：以 * 开始
*/
func parse0(rawReader io.Reader, ch chan<- *Payload) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("parse0 err ", err, string(debug.Stack()))
		}
	}()

	//初始化读取状态
	reader := bufio.NewReader(rawReader)

	for {
		line, err := reader.ReadBytes('\n')

		//读取发生错误，直接终止，将错误返回管道
		if err != nil {
			ch <- &Payload{Err: err}
			close(ch)
			return
		}

		length := len(line)
		//在复制流量中有一些空行，忽略此错误
		if length < 2 || line[length-2] != '\r' {
			continue
		}
		//返回去除line中可能的后缀suffix的子切片。（共用底层数组），将\r \n  后缀去除
		line = bytes.TrimSuffix(line, []byte{'\r', '\n'})

		switch line[0] {

		//简单字符串,以"+" 开始， 如："+OK\r\n"
		case '+':
			content := string(line[1:])
			ch <- &Payload{
				Data: protocol.MakeStatusReply(content),
			}
			//全量同步，说明是RDB或者是AOF，RDB和后面的AOF之间没有CRLF，因此需要区别对待
			if strings.HasPrefix(content, "FULLRESYNC") {
				err = parseRDBBulkString(reader, ch)
				if err != nil {
					ch <- &Payload{Err: err}
					close(ch)
					return
				}
			}
			//错误， 以"-" 开始，如："-ERR Invalid Synatx\r\n"
		case '-':
			ch <- &Payload{
				Data: protocol.MakeErrReply(string(line[1:])),
			}
			//整数：以":"开始，如：":1\r\n"
		case ':':
			value, err := strconv.ParseInt(string(line[1:]), 10, 64)
			if err != nil {
				protocolError(ch, "illegal number "+string(line[1:]))
				continue
			}
			ch <- &Payload{
				Data: protocol.MakeIntReply(value),
			}
			//字符串：以 $ 开始
		case '$':
			err = parseBulkString(line, reader, ch)
			if err != nil {
				ch <- &Payload{Err: err}
				close(ch)
				return
			}
			//数组：以 * 开始
		case '*':
			err = parseArray(line, reader, ch)
			if err != nil {
				ch <- &Payload{Err: err}
				close(ch)
				return
			}
		default:
			args := bytes.Split(line, []byte{' '})
			ch <- &Payload{
				Data: protocol.MakeMultiBulkReply(args),
			}
		}
	}
}

func ParseOne(data []byte) (redis.Reply, error) {
	ch := make(chan *Payload)
	reader := bytes.NewReader(data)
	go parse0(reader, ch)
	payload := <-ch // parse0 will close the channel
	if payload == nil {
		return nil, errors.New("no protocol")
	}
	return payload.Data, payload.Err
}
func parseArray(header []byte, reader *bufio.Reader, ch chan<- *Payload) error {
	nStrs, err := strconv.ParseInt(string(header[1:]), 10, 64)
	if err != nil || nStrs < 0 {
		protocolError(ch, "illegal array header "+string(header[1:]))
		return nil
	} else if nStrs == 0 {
		ch <- &Payload{
			Data: protocol.MakeEmptyMultiBulkReply(),
		}
		return nil
	}
	lines := make([][]byte, 0, nStrs)
	for i := int64(0); i < nStrs; i++ {
		var line []byte
		line, err = reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		length := len(line)
		if length < 4 || line[length-2] != '\r' || line[0] != '$' {
			protocolError(ch, "illegal bulk string header "+string(line))
			break
		}
		strLen, err := strconv.ParseInt(string(line[1:length-2]), 10, 64)
		if err != nil || strLen < -1 {
			protocolError(ch, "illegal bulk string length "+string(line))
			break
		} else if strLen == -1 {
			lines = append(lines, []byte{})
		} else {
			body := make([]byte, strLen+2)
			_, err := io.ReadFull(reader, body)
			if err != nil {
				return err
			}
			lines = append(lines, body[:len(body)-2])
		}
	}
	ch <- &Payload{
		Data: protocol.MakeMultiBulkReply(lines),
	}
	return nil
}

// RDB和后面的AOF之间没有CRLF，因此需要区别对待
func parseRDBBulkString(reader *bufio.Reader, ch chan<- *Payload) error {
	header, err := reader.ReadBytes('\n')
	header = bytes.TrimSuffix(header, []byte{'\r', '\n'})
	if len(header) == 0 {
		return errors.New("empty header")
	}
	strLen, err := strconv.ParseInt(string(header[1:]), 10, 64)
	if err != nil || strLen <= 0 {
		return errors.New("illegal bulk header: " + string(header))
	}
	body := make([]byte, strLen)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return err
	}
	ch <- &Payload{
		Data: protocol.MakeBulkReply(body[:len(body)]),
	}
	return nil
}

func parseBulkString(header []byte, reader *bufio.Reader, ch chan<- *Payload) error {
	strLen, err := strconv.ParseInt(string(header[1:]), 10, 64)
	if err != nil || strLen < -1 {
		protocolError(ch, "illegal bulk string header: "+string(header))
		return nil
	} else if strLen == -1 {
		ch <- &Payload{
			Data: protocol.MakeNullBulkReply(),
		}
		return nil
	}
	body := make([]byte, strLen+2)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return err
	}
	ch <- &Payload{
		Data: protocol.MakeBulkReply(body[:len(body)-2]),
	}
	return nil
}

func protocolError(ch chan<- *Payload, msg string) {
	err := errors.New("protocol error: " + msg)
	ch <- &Payload{Err: err}
}
