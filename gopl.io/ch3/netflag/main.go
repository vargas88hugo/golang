package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	fmt.Printf("FlagUp %d: %b\n", FlagUp, FlagUp)
	fmt.Printf("FlagBroadcast %d: %b\n", FlagBroadcast, FlagBroadcast)
	fmt.Printf("FlagLoopback %d: %b\n", FlagLoopback, FlagLoopback)
	fmt.Printf("FlagPointToPoint %d: %b\n", FlagPointToPoint, FlagPointToPoint)
	fmt.Printf("FlagMulticast %d: %b\n", FlagMulticast, FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
