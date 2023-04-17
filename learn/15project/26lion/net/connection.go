package net

import (
	"26lion/config"
	"26lion/iface"
	"26lion/pack"
	"context"
	"errors"
	"log"
	"net"
	"sync"
)

type Connection struct {
	server      iface.IServer
	conn        net.Conn
	connId      uint64
	isClosed    bool
	msgHandler  iface.IMsgHandle
	ctx         context.Context
	cancel      context.CancelFunc
	msgBufferCh chan []byte
	msgLock     sync.RWMutex
}

var _ iface.IConnection = (*Connection)(nil)

func newConnection(s iface.IServer, c net.Conn, connId uint64) iface.IConnection {
	return &Connection{
		server: s,
		conn:   c,
		connId: connId,
	}
}

func (c *Connection) Start() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("connection Start err:", err)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancelFunc

	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.quit()
		return
	}
}

func (c *Connection) Send(data []byte) error {
	c.msgLock.RLock()
	defer c.msgLock.RUnlock()

	if c.isClosed {
		return errors.New("Connection is closed:" + string(c.connId))
	}
	_, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func (c *Connection) SendToQueue(data []byte) error {
	return nil
}

// 直接将Message数据发送数据给远程的TCP客户端(无缓冲)
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	return nil
}

// 直接将Message数据发送给远程的TCP客户端(有缓冲)
func (c *Connection) SendBuffMsg(msgID uint32, data []byte) error {
	return nil
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// 用于从客户端读取数据
func (c *Connection) StartReader() {
	defer c.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			buffer := make([]byte, config.GetConnctionConfigConfig().IoReaderBufferSize)
			n, err := c.conn.Read(buffer[:])
			if err != nil {
				return
			}
			msg := pack.NewMessage(uint32(n), buffer[0:n])
			request := NewRequest(msg, c)
			c.msgHandler.Execute(request)
		}
	}
}

func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) Connection() iface.IConnection {
	return c
}

func (c *Connection) Server() iface.IServer {
	return c.server
}

func (c *Connection) ConnId() uint64 {
	return c.connId
}

func (c *Connection) quit() {
	log.Printf("connection ：%d stop", c.connId)
	if c.isClosed {
		return
	}
	c.conn.Close()
	c.isClosed = true
}
