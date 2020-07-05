package exercise

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr ip address
type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	ipAddrStr := make([]string, 4, 4)
	for i := 0; i < 4; i++ {
		ipAddrStr[i] = strconv.Itoa(int(ipAddr[i]))
	}
	return strings.Join(ipAddrStr, ".")
}

// Bird 鸟
type Bird struct {
	Name string
}

// Flyable 会飞的
type Flyable interface {
	Fly()
}

// Fly 普通的鸟
func (b Bird) Fly() {
	fmt.Printf("%v is flying\n", b.Name)
}

// Eagle 老鹰
type Eagle Bird

// Fly 老鹰
func (b *Eagle) Fly() {
	fmt.Printf("%v is a Eagle, it is flying\n", b.Name)
}
