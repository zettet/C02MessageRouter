package bytes

import (
	"EmissionsMessageRouter/model"
	"encoding/binary"
	"fmt"
	"math"
)

func ParseMessage(messageBytes []byte) (model.EmissionsMessage, error) {
	message, error := _parseMessageBytes(messageBytes)
	if error != nil {
		errorMsg := fmt.Errorf("error parsing message with error: %v", error)
		fmt.Println(errorMsg)
		return model.EmissionsMessage{}, errorMsg
	}

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
func _parseMessageBytes(messageBytes []byte) (model.EmissionsMessage, error) {

	if !_validateHeader(messageBytes) {
		return model.EmissionsMessage{}, fmt.Errorf("Invalid Header")
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

	return model.EmissionsMessage{
		Tail_number:  tailNumberValue,
		Engine_count: engineCount,
		Engine_name:  engineNameValue,
		Latitude:     latitude,
		Longitude:    longitude,
		Altitude:     altitude,
		Temperature:  temperature,
	}, nil
}

func _validateHeader(messageBytes []byte) bool {
	var expectedHeader = []byte{65, 73, 82} //  0100 0001 0100 1001 0101 0010

	return messageBytes[0] == expectedHeader[0] && messageBytes[1] == expectedHeader[1] && messageBytes[2] == expectedHeader[2]
}
