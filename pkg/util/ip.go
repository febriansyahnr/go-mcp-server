package util

import (
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP(host string) (net.IP, error) {
	conn, err := net.Dial("udp", host)
	if err != nil {
		return []byte(``), err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
