package netx

import (
	"encoding/binary"
	"math/big"
	"net"
)

func IPv4ToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		panic("no sane way to convert ipv6 into uint32")
	}
	return binary.BigEndian.Uint32(ip)
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
