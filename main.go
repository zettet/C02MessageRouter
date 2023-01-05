package EmissionsMessageRouter

import (
	"EmissionsMessageRouter/clients"
	"EmissionsMessageRouter/parsers/bytes"
	"fmt"
	"net"
	"time"
)

func main() {

	for true {
		connection, err := clients.GetTCPConnection()
		if err != nil {
			fmt.Errorf("Error obtaining tcp connection: %v", err)
		} else {
			handleMessage(connection)
		}

		// sleeping here so that there is some time in between requests, mainly so that terminal output can be easily read
		time.Sleep(10 * time.Second)
	}

}

func handleMessage(connection *net.TCPConn) {
	messageBytes, err := clients.FetchTCPData(connection)
	if err != nil {
		fmt.Errorf("Error fetching data from tcp connection: %v", err)
	}

	emissionsMessage, err := bytes.ParseMessage(messageBytes)
	if err != nil {
		fmt.Errorf("Error parsing message bytes: %v", err)
	}

	// could have a prettier output here...
	fmt.Println(emissionsMessage)
}
