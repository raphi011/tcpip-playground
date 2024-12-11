package network

import "testing"

func TestNetwork(t *testing.T) {
	host1 := &Host{name: "host-1"}
	host2 := &Host{name: "host-2"}

	swit := NewSwitch(16)

	rout := &Router{}

	swit.Connect(host1, 0)
	swit.Connect(host2, 1)
	swit.Connect(rout, 2)
}
