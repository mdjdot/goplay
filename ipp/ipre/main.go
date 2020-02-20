package main

import (
	"fmt"
	"regexp"
)

func IsIP(ip string) (b bool) {
	// if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
	if m, _ := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, ip); !m {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsIP("192.168.1.1"))
	// fmt.Println(IsIP("192a168b1c1"))
	// fmt.Println(IsIP(`192\d168\a31\d3`))
}
