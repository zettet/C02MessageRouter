package EmissionsMessageRouter

import (
	"testing"
)

var testBytes = []byte{
	65, 73, 82, 0, 0, 0, 6, 78, 50, 48, 57, 48, 52, 0, 0, 0,
	2, 0, 0, 0, 7, 71, 69, 110, 120, 45, 49, 66, 64, 67, 142, 214,
	235, 255, 29, 96, 192, 80, 212, 192, 142, 99, 1, 101, 64, 226, 3, 240,
	0, 0, 0, 0, 192, 74, 153, 153, 153, 153, 153, 154,
}
var expectedMessage = EmissionsMessage{
	tail_number:  "N20904",
	engine_count: 2,
	engine_name:  "GEnx-1B",
	latitude:     39.115933894698856,
	longitude:    -67.32425269764424,
	altitude:     36895.5,
	temperature:  -53.2,
}

func TestMessageParserWithValidMessageBytesReturnsValidStruct(t *testing.T) {
	message, error := handleMessage(testBytes)

	if error != nil {
		t.Fatal("unexpected error thrown")
	}

	if message.tail_number != expectedMessage.tail_number {
		t.Fatalf("expected tail number: %v, got: %v", expectedMessage.tail_number, message.tail_number)
	}

	if message.engine_count != expectedMessage.engine_count {
		t.Fatalf("expected engine count: %v, got: %v", expectedMessage.engine_count, message.engine_count)
	}

	if message.engine_name != expectedMessage.engine_name {
		t.Fatalf("expected engine name: %v, got: %v", expectedMessage.engine_name, message.engine_name)
	}

	if message.latitude != expectedMessage.latitude {
		t.Fatalf("expected latitude: %v, got: %v", expectedMessage.latitude, message.latitude)
	}

	if message.longitude != expectedMessage.longitude {
		t.Fatalf("expected longitude: %v, got: %v", expectedMessage.longitude, message.longitude)
	}

	if message.altitude != expectedMessage.altitude {
		t.Fatalf("expected altitude: %v, got: %v", expectedMessage.altitude, message.altitude)
	}

	if message.temperature != expectedMessage.temperature {
		t.Fatalf("expected temperature: %v, got: %v", expectedMessage.temperature, message.temperature)
	}
}
