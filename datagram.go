package goSam

import (
	"net"
	"time"
)

// DatagramConn
type DatagramConn struct {
	RWC
	conn  net.PacketConn
	RAddr net.Addr
}

// WrapConn wraps a net.PacketConn in a DatagramConn.
func WrapPacketConn(c net.Conn) *Conn {
	wrap := Conn{
		conn: c,
	}
	wrap.Reader = NewReadLogger("<", c)
	wrap.Writer = NewWriteLogger(">", c)
	wrap.RWC.c = c
	return &wrap
}

func (d *DatagramConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return d.conn.ReadFrom(p)
}
func (d *DatagramConn) Read(b []byte) (n int, err error) {
	n, _, err = d.ReadFrom(b)
	return n, err
}
func (d *DatagramConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return d.conn.WriteTo(p, addr)
}
func (d *DatagramConn) Write(b []byte) (n int, err error) {
	n, err = d.WriteTo(b, d.RemoteAddr())
	return n, err
}
func (d *DatagramConn) Close() error {
	return d.conn.Close()
}
func (d *DatagramConn) LocalAddr() net.Addr {
	return d.conn.LocalAddr()
}
func (d *DatagramConn) RemoteAddr() net.Addr {
	return d.RAddr
}
func (d *DatagramConn) SetDeadline(t time.Time) error {
	return d.conn.SetDeadline(t)
}
func (d *DatagramConn) SetReadDeadline(t time.Time) error {
	return d.conn.SetReadDeadline(t)
}
func (d *DatagramConn) SetWriteDeadline(t time.Time) error {
	return d.conn.SetWriteDeadline(t)
}

var dgt net.PacketConn = &DatagramConn{}

//func (c *Client) DatagramSend()
