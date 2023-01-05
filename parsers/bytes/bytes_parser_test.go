package bytes

import (
	"EmissionsMessageRouter/model"
	"testing"
)

var testBytes = []byte{
	65, 73, 82, 0, 0, 0, 6, 78, 50, 48, 57, 48, 52, 0, 0, 0,
	2, 0, 0, 0, 7, 71, 69, 110, 120, 45, 49, 66, 64, 67, 142, 214,
	235, 255, 29, 96, 192, 80, 212, 192, 142, 99, 1, 101, 64, 226, 3, 240,
	0, 0, 0, 0, 192, 74, 153, 153, 153, 153, 153, 154,
}
var expectedMessage = model.EmissionsMessage{
	Tail_number:  "N20904",
	Engine_count: 2,
	Engine_name:  "GEnx-1B",
	Latitude:     39.115933894698856,
	Longitude:    -67.32425269764424,
	Altitude:     36895.5,
	Temperature:  -53.2,
}

func Test_BytesParser_WithValidMessageBytes_ReturnsValidStruct(t *testing.T) {
	message, error := parseMessage(testBytes)

	if error != nil {
		t.Fatal("unexpected error thrown")
	}

	if message.Tail_number != expectedMessage.Tail_number {
		t.Fatalf("expected tail number: %v, got: %v", expectedMessage.Tail_number, message.Tail_number)
	}

	if message.Engine_count != expectedMessage.Engine_count {
		t.Fatalf("expected engine count: %v, got: %v", expectedMessage.Engine_count, message.Engine_count)
	}

	if message.Engine_name != expectedMessage.Engine_name {
		t.Fatalf("expected engine name: %v, got: %v", expectedMessage.Engine_name, message.Engine_name)
	}

	if message.Latitude != expectedMessage.Latitude {
		t.Fatalf("expected latitude: %v, got: %v", expectedMessage.Latitude, message.Latitude)
	}

	if message.Longitude != expectedMessage.Longitude {
		t.Fatalf("expected longitude: %v, got: %v", expectedMessage.Longitude, message.Longitude)
	}

	if message.Altitude != expectedMessage.Altitude {
		t.Fatalf("expected altitude: %v, got: %v", expectedMessage.Altitude, message.Altitude)
	}

	if message.Temperature != expectedMessage.Temperature {
		t.Fatalf("expected temperature: %v, got: %v", expectedMessage.Temperature, message.Temperature)
	}
}
