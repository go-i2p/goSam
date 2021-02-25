package goSam

import (
	"net"
	"time"
)

type DatagramConn interface {
	ReadFrom(p []byte) (n int, addr net.Addr, err error)
	Read(b []byte) (n int, err error)
	WriteTo(p []byte, addr net.Addr) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
