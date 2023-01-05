package clients

import (
	"net"
)

const (
	HOST = "data.salad.com"
	PORT = "5000"
	TYPE = "tcp"
)

func GetTCPConnection() (*net.TCPConn, error) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		return nil, err
	}
	connection, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func FetchTCPData(conn *net.TCPConn) ([]byte, error) {
	received := make([]byte, 1024)
	_, err := conn.Read(received)
	conn.Close()
	if err != nil {
		return nil, err
	}

	return received, nil
}
