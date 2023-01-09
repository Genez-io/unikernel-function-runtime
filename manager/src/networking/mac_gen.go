package networking

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a random local MAC address and return it in byte array format
func GenerateMACAddr_B() []byte {
	buf := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	// Set the local bit
	buf[0] |= 2
	return buf
}

// Generate a random local MAC address and return it in string format
func GenerateMACAddr_S() string {
	buf := GenerateMACAddr_B()
	str_mac := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return str_mac
}
