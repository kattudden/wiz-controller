package controller

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// TODO Implement wiz functions.
// https://dev.to/santosh/how-to-control-philips-wiz-bulb-using-go-2ad9

func TurnOn(ip string, port string, temperature int, dimming int) {
	c, err := net.Dial("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		panic("Unable to connect to light bulp!")
	}
	// Erstelle den JSON-String mit den übergebenen Parametern
	jsonStr := fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "temp": %d, "dimming": %d}}`, temperature, dimming)

	c.Write([]byte(jsonStr))
}

func TurnOff(ip string, port string) {
	c, err := net.Dial("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		panic("Unable to connect to light bulp!")
	}

	c.Write([]byte(`{"method": "setPilot", "params":{"state": false}}`))
}

func GetStatus(ip string, port string) (*StatusResponse, error) {
	// Verbindung per UDP herstellen
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to light bulb: %w", err)
	}
	defer conn.Close()

	// Senden der Anfrage
	request := []byte(`{"method": "getPilot", "params":{}}`)
	_, err = conn.Write(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Setzen eines Lese-Timeouts, da UDP keine Garantie für Antwortlieferung gibt.
	deadline := time.Now().Add(3 * time.Second)
	if err := conn.SetReadDeadline(deadline); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	// Buffer, um die Antwort zu lesen
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parsen der JSON-Antwort in das Struct
	var status StatusResponse
	if err := json.Unmarshal(buffer[:n], &status); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &status, nil
}
