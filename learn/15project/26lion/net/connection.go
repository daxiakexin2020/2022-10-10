package net

import (
	"26lion/iface"
	"context"
	"log"
	"net"
)

type Connection struct {
	server     iface.IServer
	conn       net.Conn
	connId     uint64
	isClosed   bool
	msgHandler iface.IMsgHandle
	ctx        context.Context
	cancel     context.CancelFunc
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

	select {
	case <-c.ctx.Done():
		c.quit()
		return
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
