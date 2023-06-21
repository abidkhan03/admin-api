package freeling

import (
	"bytes"
	"fmt"
	"net"
)

const (
	bufSize = 2048

	msgServerReady = "FL-SERVER-READY"
	msgReset       = "RESET_STATS"
	msgFlush       = "FLUSH_BUFFER"
)

func newSocket(host string, port int) (*socket, error) {
	sock, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	//err = sock.KeepAlive()
	//if err != nil {
	//	return nil, err
	//}

	c := &socket{
		conn: sock,
	}

	err = c.getReady()
	if err != nil {
		return nil, err
	}

	return c, nil
}

type socket struct {
	conn net.Conn
}

func (s *socket) Close() error {
	return s.conn.Close()
}

func (s *socket) getReady() error {
	msg, err := s.Request([]byte(msgReset))
	if err != nil {
		return err
	}

	if string(msg) != msgServerReady {
		return fmt.Errorf("SERVER NOT READY")
	}

	return s.flush()
}

func (s *socket) Request(msg []byte) ([]byte, error) {
	_, err := s.send(msg)
	if err != nil {
		return nil, err
	}

	resp, err := s.receive()
	if err != nil {
		return nil, err
	}

	return resp, s.flush()
}

func (s *socket) flush() error {
	_, err := s.send([]byte(msgFlush))
	if err != nil {
		return err
	}

	_, err = s.receive()
	return err
}

func (s *socket) receive() ([]byte, error) {
	buf := make([]byte, bufSize)

	var result bytes.Buffer
	for {
		n, err := s.conn.Read(buf)
		if err != nil {
			return nil, err
		}
		if n > 0 {
			result.Write(buf[:n])
		}
		if !(n > 0 && buf[n-1] != 0) {
			break
		}
	}

	// removing \0 from the end
	resp := result.Bytes()
	if len(resp) > 0 {
		resp = resp[:len(resp)-1]
	}

	return resp, nil
}

func (s *socket) send(message []byte) (int, error) {
	return s.conn.Write(append(message, 0))
}
