package EmissionsMessageRouter

import (
	"encoding/binary"
	"fmt"
	"math"
)

func handleMessage(messageBytes []byte) (EmissionsMessage, error) {
	message, error := parseMessageBytes(messageBytes)
	if error != nil {
		errorMsg := fmt.Errorf("error parsing message with error: %v", error)
		fmt.Println(errorMsg)
		return EmissionsMessage{}, errorMsg
	}

	fmt.Println(message)
	return message, nil
}

/*
 * If messageBytes is not in the specified format below, the message will be rejected and an error will be returned.
 *
 * messageBytes format is the following:
 * header: 3 bytes - exactly: 0100 0001 0100 1001 0101 0010
 * tail_number_size: 4 bytes, unsigned int
 * tail_number_value: variable length; number of bytes specified by tail_number_size, UTF-8-encoded string
 * engine_count: 4 bytes, unsigned int
 * engine_name_size: 4 bytes, unsigned int
 * engine_name_value: variable length; number of bytes specified by tail_number_size, UTF-8-encoded string
 * latitude: 8 bytes, IEEE-754 64-bit floating-point number
 * longitude: 8 bytes, IEEE-754 64-bit floating-point number
 * altitude: 8 bytes, IEEE-754 64-bit floating-point number
 * temperature: 8 bytes, IEEE-754 64-bit floating-point number
 */
func parseMessageBytes(messageBytes []byte) (EmissionsMessage, error) {

	if !validateHeader(messageBytes) {
		return EmissionsMessage{}, fmt.Errorf("Invalid Header")
	}

	var headerIndex = 0

	var tailNumberSizeIndex = headerIndex + 3
	var tailNumberSize = int(binary.BigEndian.Uint32(messageBytes[tailNumberSizeIndex : tailNumberSizeIndex+4]))

	var tailNumberValueIndex = tailNumberSizeIndex + 4
	var tailNumberValue = string(messageBytes[tailNumberValueIndex : tailNumberValueIndex+tailNumberSize])

	var engineCountIndex = tailNumberValueIndex + tailNumberSize
	var engineCount = int(binary.BigEndian.Uint32(messageBytes[engineCountIndex : engineCountIndex+4]))

	var engineNameSizeIndex = engineCountIndex + 4
	var engineNameSize = int(binary.BigEndian.Uint32(messageBytes[engineNameSizeIndex : engineNameSizeIndex+4]))

	var engineNameValueIndex = engineNameSizeIndex + 4
	var engineNameValue = string(messageBytes[engineNameValueIndex : engineNameValueIndex+engineNameSize])

	var latitudeIndex = engineNameValueIndex + engineNameSize
	var latitude = math.Float64frombits(binary.BigEndian.Uint64(messageBytes[latitudeIndex : latitudeIndex+8]))

	var longitudeIndex = latitudeIndex + 8
	var longitude = math.Float64frombits(binary.BigEndian.Uint64(messageBytes[longitudeIndex : longitudeIndex+8]))

	var altitudeIndex = longitudeIndex + 8
	var altitude = math.Float64frombits(binary.BigEndian.Uint64(messageBytes[altitudeIndex : altitudeIndex+8]))

	var temperatureIndex = altitudeIndex + 8
	var temperature = math.Float64frombits(binary.BigEndian.Uint64(messageBytes[temperatureIndex : temperatureIndex+8]))

	return EmissionsMessage{
		tail_number:  tailNumberValue,
		engine_count: engineCount,
		engine_name:  engineNameValue,
		latitude:     latitude,
		longitude:    longitude,
		altitude:     altitude,
		temperature:  temperature,
	}, nil
}

func validateHeader(messageBytes []byte) bool {
	var expectedHeader = []byte{65, 73, 82} //  0100 0001 0100 1001 0101 0010

	return messageBytes[0] == expectedHeader[0] && messageBytes[1] == expectedHeader[1] && messageBytes[2] == expectedHeader[2]
}

type EmissionsMessage struct {
	tail_number  string  // The international aircraft registration. A unique code assigned to the aircraft.
	engine_count int     // The number of engines on the aircraft.
	engine_name  string  // The engine name.
	latitude     float64 // The latitude in degrees.
	longitude    float64 // The longitude in degrees.
	altitude     float64 // The altitude in degrees.
	temperature  float64 // The temperature in degrees Fahrenheit.
}
