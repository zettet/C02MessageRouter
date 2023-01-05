package main

import (
	"EmissionsMessageRouter/clients"
	"EmissionsMessageRouter/parsers/bytes"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("starting up processor")
	for true {
		fmt.Println("creating connection")
		connection, err := clients.GetTCPConnection()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error obtaining tcp connection: %v", err))
		} else {
			fmt.Println("received connection")
			err = handleMessage(connection)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error handling message: %v", err))
			}
		}

		// sleeping here so that there is some time in between requests, mainly so that terminal output can be easily read
		fmt.Println("sleeping for 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}

func handleMessage(connection *net.TCPConn) error {
	messageBytes, err := clients.FetchTCPData(connection)
	fmt.Println("Received message bytes")
	if err != nil {
		return fmt.Errorf("Error fetching data from tcp connection: %v", err)
	}

	emissionsMessage, err := bytes.ParseMessage(messageBytes)
	if err != nil {
		return fmt.Errorf("Error parsing message bytes: %v", err)
	}

	// could have a prettier output here...
	fmt.Println(emissionsMessage)
	return nil
}
