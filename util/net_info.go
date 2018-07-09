package util

import (
	"net"
)

func LocalIP() string {
	// type Interface {}  represents a mapping between network interface name and index. It also represents network interface facility information.
	// interfaces, _ := net.Interfaces() // 这个方法查询的是硬件网络接口信息
	// for i, v := range interfaces {} // 这个方法查询的是网络层的网络地址信息
	//
	// type Addr {}  represents a network end point address.
	addresses, _ := net.InterfaceAddrs()
	for _, v := range addresses {
		// 因为是接口，所以需要做接口询问
		// IPNet包含IP4和IP6两种地址格式
		if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
