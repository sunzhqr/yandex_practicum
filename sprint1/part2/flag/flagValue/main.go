package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n *NetAddress) Set(addrValue string) error {
	hp := strings.Split(addrValue, ":")
	if len(hp) != 2 {
		return errors.New("need address in host:port format")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	n.Host = hp[0]
	n.Port = port
	return nil
}

func main() {
	addr := new(NetAddress)
	_ = flag.Value(addr)
	flag.Var(addr, "addr", "network address host:port")
	flag.Parse()
	fmt.Println("Host:", addr.Host)
	fmt.Println("Port:", addr.Port)
}
