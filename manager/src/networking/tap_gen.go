package networking

import (
	"fmt"
	"net"
	"os/exec"

	"github.com/lorenzosaino/go-sysctl"
)

type TapInterface struct {
	Name string
	IP   net.IP
	Net  net.IPNet
}

var intf_map map[string]*TapInterface = make(map[string]*TapInterface)
var Intf_count int

func InitNetworking() {
	intf_list, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	Intf_count = len(intf_list)
}

// Create a new TAP interface and return its name
func ConfigureTap(uuid string) (*TapInterface, error) {
	tap_if_name := "tap_" + uuid
	tap_ip_addr := "172.16." + "1" + ".1/30"

	// Create interface
	ip_tap_create_args := []string{
		"tuntap",
		"add",
		"dev",
		tap_if_name,
		"mode",
		"tap",
	}
	out, err := exec.Command("ip", ip_tap_create_args...).Output()
	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}
	sysctl.Set("net.ipv4.conf."+tap_if_name+".proxy_arp", "1")
	sysctl.Set("net.ipv6.conf."+tap_if_name+".disable_ipv6", "1")

	// Set interface to up state
	ip_tap_up := []string{
		"link",
		"set",
		"dev",
		tap_if_name,
		"up",
	}
	out, err = exec.Command("ip", ip_tap_up...).Output()
	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}
	// Set add address to interface
	tap_ip, tap_net, err := net.ParseCIDR(tap_ip_addr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ip_add_addr := []string{
		"addr",
		"add",
		tap_ip_addr,
		"dev",
		tap_if_name,
	}
	out, err = exec.Command("ip", ip_add_addr...).Output()
	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}
	var new_tap TapInterface = TapInterface{
		Name: tap_if_name,
		IP:   tap_ip,
		Net:  *tap_net,
	}
	fmt.Println("Created new TAP interface [" + tap_if_name + "] with addr [" + tap_ip_addr + "]")

	intf_map[uuid] = &new_tap

	return &new_tap, nil
}
