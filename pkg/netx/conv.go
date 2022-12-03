package netx

import (
	"encoding/binary"
	"errors"
	"math/big"
	"net"
)

func IPv4ToInt(ip net.IP) (uint32, error) {
	if len(ip) == 16 {
		return 0, errors.New("IPv4ToInt: IPv6 address")
	}
	return binary.BigEndian.Uint32(ip), nil
}

func IntToIPv4(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func IPv6ToInt(ipv6 net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(ipv6)
	return IPv6Int
}

func IntToIPv6(nn *big.Int) net.IP {
	ip := nn.Bytes()
	var a net.IP = ip
	return a
}

func StringToIPv4(ip string) (uint32, error) {
	return IPv4ToInt(net.ParseIP(ip))
}

func StringToIPv6(ip string) (*big.Int, error) {
	return IPv6ToInt(net.ParseIP(ip)), nil
}
