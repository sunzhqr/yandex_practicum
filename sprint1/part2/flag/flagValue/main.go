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

func (n *NetAddress) Set(addr string) error {
	hostPort := strings.Split(addr, ":")
	if len(hostPort) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(hostPort[1])
	if err != nil {
		return err
	}
	n.Host = hostPort[0]
	n.Port = port
	return nil
}

func main() {
	addr := new(NetAddress)
	_ = flag.Value(addr)
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	//Socket
	fmt.Printf("https://%s:%d\n", addr.Host, addr.Port)
}
