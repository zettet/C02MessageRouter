package clients

import (
	"net"
)

const (
	HOST = "localhost"
	PORT = "3333"
	TYPE = "tcp"
)

func GetTCPAddr() (*net.TCPAddr, error) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		return nil, err
	}
	return tcpServer, err
}

func FetchTCPData(tcpServer *net.TCPAddr) ([]byte, error) {
	connection, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		return nil, err
	}

	received := make([]byte, 1024)
	_, err = connection.Read(received)
	connection.Close()
	if err != nil {
		return nil, err
	}

	return received, nil
}
