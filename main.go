package main

import (
	"EmissionsMessageRouter/clients"
	"EmissionsMessageRouter/parsers/bytes"
	"fmt"
	"net"
	"time"
)

func main() {

	for true {
		tcpServer, err := clients.GetTCPAddr()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error obtaining tcp addr: %v", err))
		}

		err = handleMessage(tcpServer)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error handling message: %v", err))
		}

		// sleeping here so that there is some time in between requests, mainly so that terminal output can be easily read
		fmt.Println("sleeping for 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

func handleMessage(tcpServer *net.TCPAddr) error {
	messageBytes, err := clients.FetchTCPData(tcpServer)
	if err != nil {
		return fmt.Errorf("Error fetching data from tcp server: %v", err)
	}

	emissionsMessage, err := bytes.ParseMessage(messageBytes)
	if err != nil {
		return fmt.Errorf("Error parsing message bytes: %v", err)
	}

	// could have a prettier output here...
	fmt.Println(emissionsMessage)
	return nil
}
