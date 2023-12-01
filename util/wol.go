package util

import (
	"bytes"
	"log"
	"net"
	"strconv"
)

func DoWake(macAddr string, ipAddr string, port int) error {
	// build packet
	packet, buildErr := buildMagicPacket(macAddr)
	if buildErr != nil {
		log.Printf("do wake failed build magic packet by %s cause of %v", macAddr, buildErr)
		return buildErr
	}

	// send packet
	conn, dialErr := net.Dial("udp", net.JoinHostPort(ipAddr, strconv.Itoa(port)))
	if dialErr != nil {
		log.Printf("do wake failed dial to %s %s:%d cause of %v", macAddr, ipAddr, port, dialErr)
		return dialErr
	}

	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("do wake failed close connection %s %s:%d cause of %v", macAddr, ipAddr, port, closeErr)
		}
	}()

	_, writeErr := conn.Write(packet)
	if writeErr != nil {
		log.Printf("do wake failed send data to %s %s:%d cause of %v", macAddr, ipAddr, port, writeErr)
		return writeErr
	}

	log.Printf("do wake send magic packet to %s %s:%d", macAddr, ipAddr, port)
	return nil
}

func buildMagicPacket(macAddr string) ([]byte, error) {
	macBytes, err := net.ParseMAC(macAddr)
	if err != nil {
		return nil, err
	}

	header := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	payload := bytes.Repeat(macBytes, 16)
	magicPacket := append(header, payload...)

	return magicPacket, nil
}
