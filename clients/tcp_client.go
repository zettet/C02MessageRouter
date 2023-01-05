package clients

import (
	"net"
)

const (
	HOST = "data.salad.com"
	PORT = "5000"
	TYPE = "tcp"
)

var tcpConnection *net.TCPConn = nil

func getConnection() (*net.TCPConn, error) {
	if tcpConnection != nil {
		return tcpConnection, nil
	}
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

func fetchData(conn *net.TCPConn) ([]byte, error) {
	received := make([]byte, 1024)
	_, err := conn.Read(received)
	if err != nil {
		return nil, err
	}

	conn.Close()

	return received, nil
}
