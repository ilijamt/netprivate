package netprivate_test

import (
	"fmt"
	"github.com/ilijamt/netprivate"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestIsPrivateIP(t *testing.T) {
	var tests = []struct {
		ip      net.IP
		private bool
	}{
		{ip: net.ParseIP("127.0.0.1"), private: true},
		{ip: net.ParseIP("192.168.254.254"), private: true},
		{ip: net.ParseIP("10.255.0.3"), private: true},
		{ip: net.ParseIP("172.16.255.255"), private: true},
		{ip: net.ParseIP("172.16.255.255"), private: true},
		{ip: net.ParseIP("172.31.255.255"), private: true},
		{ip: net.ParseIP("128.0.0.1"), private: false},
		{ip: net.ParseIP("192.169.255.255"), private: false},
		{ip: net.ParseIP("9.255.0.255"), private: false},
		{ip: net.ParseIP("172.32.255.255"), private: false},
		{ip: net.ParseIP("::0"), private: true},
		{ip: net.ParseIP("::1"), private: true},
		{ip: net.ParseIP("::2"), private: false},
		{ip: net.ParseIP("fe80::1"), private: true},
		{ip: net.ParseIP("febf::1"), private: true},
		{ip: net.ParseIP("fec0::1"), private: false},
		{ip: net.ParseIP("feff::1"), private: false},
		{ip: net.ParseIP("ff00::1"), private: true},
		{ip: net.ParseIP("ff10::1"), private: true},
		{ip: net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), private: true},
		{ip: net.ParseIP("2002::"), private: true},
		{ip: net.ParseIP("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), private: true},
		{ip: net.ParseIP("0100::"), private: true},
		{ip: net.ParseIP("0100::0000:ffff:ffff:ffff:ffff"), private: true},
		{ip: net.ParseIP("0100::0001:0000:0000:0000:0000"), private: false},
	}

	for _, test := range tests {
		t.Logf("%s is in a private network %v", test.ip.String(), test.private)
		assert.EqualValues(t, test.private, netprivate.Is(test.ip))
	}
}

func ExampleIsV4_privateIP() {
	fmt.Println(netprivate.IsV4(net.ParseIP("127.0.0.1")))
	// Output: true
}

func ExampleIsV4_publicIP() {
	fmt.Println(netprivate.IsV4(net.ParseIP("1.1.1.1")))
	// Output: false
}

func ExampleIsV6_privateIP() {
	fmt.Println(netprivate.IsV6(net.ParseIP("fe80::1")))
	// Output: true
}

func ExampleIsV6_publicIP() {
	fmt.Println(netprivate.IsV6(net.ParseIP("feff::1")))
	// Output: false
}

func ExampleIs_privateIPv4() {
	fmt.Println(netprivate.Is(net.ParseIP("127.0.0.1")))
	// Output: true
}

func ExampleIs_publicIPv4() {
	fmt.Println(netprivate.IsV4(net.ParseIP("1.1.1.1")))
	// Output: false
}

func ExampleIs_privateIPv6() {
	fmt.Println(netprivate.IsV6(net.ParseIP("fe80::1")))
	// Output: true
}

func ExampleIs_publicIPv6() {
	fmt.Println(netprivate.IsV6(net.ParseIP("feff::1")))
	// Output: false
}
