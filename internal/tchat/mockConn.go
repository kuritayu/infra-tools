package tchat

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"net"
	"time"
)

type MockConn struct {
	mock.Mock
}

func (m *MockConn) Close() error                       { return nil }
func (m *MockConn) LocalAddr() net.Addr                { return nil }
func (m *MockConn) RemoteAddr() net.Addr               { return nil }
func (m *MockConn) SetDeadline(t time.Time) error      { return nil }
func (m *MockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockConn) SetWriteDeadline(t time.Time) error { return nil }

func (m *MockConn) Read(b []byte) (n int, err error) {
	fmt.Println("Read() parameter --->", string(b))
	return 1, errors.New("dummy")
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	fmt.Println("Write() parameter --->", string(b))
	return 1, errors.New("dummy")
}
