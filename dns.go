package main

import (
	"net"
)

func doResolveA(host string) (ip string) {
	var addrs []string
	var err error

	if len(host) < 3 || len(host) > 253 {
		return ""
	}

	addrs, err = net.LookupHost(host)
	if err != nil || len(addrs) < 1 {
		Info.Printf("failed to lookup host %s\n", host)
		return ""
	}

	return addrs[0]
}

func doResolvePTR(addr string) (name string) {
	var names []string
	var err error

Info.Printf("addr: %s\n", addr)
	names, err = net.LookupAddr(addr)
	if err != nil || len(names) < 1 {
		Info.Printf("failed to lookup addr %s\n", addr)
		return ""
	}

	return names[0]
}
