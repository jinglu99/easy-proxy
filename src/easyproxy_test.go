package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestNetworkSetUP(t *testing.T) {
	r := exec.Command("networksetup" , "-getsocksfirewallproxy", "Wi-Fi",)
	out, _ := r.CombinedOutput()
	fmt.Println(string(out))
}
