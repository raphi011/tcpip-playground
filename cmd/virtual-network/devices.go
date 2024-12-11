package network

import (
	"fmt"
	"io"
	"net"
	"net/netip"
)

type NetworkDevice interface {
	NIC() *NIC
}

type Router struct {
	nic NIC
}

func (r *Router) NIC() *NIC {
	return &r.nic
}

func (s *Switch) Connect(device NetworkDevice, port int) error {
	if port < 0 || port >= len(s.ports) {
		return fmt.Errorf("attempt to connect to non-existant port %d", port)
	}

	if s.ports[port] != nil {
		return fmt.Errorf("port %d is already used, disconnect first", port)
	}

	s.ports[port] = &Port{device.NIC().Connect()}

	return nil
}

type Switch struct {
	ports []*Port
}

func NewSwitch(size int) *Switch {
	return &Switch{ports: make([]*Port, 16)}
}

type Host struct {
	nic  NIC
	name string
}

func (h *Host) Start() {

}

func (h *Host) NIC() *NIC {
	return &h.nic
}

type Port struct {
	con io.ReadWriter
}

type NIC struct {
	address netip.Addr
	mac     net.HardwareAddr
	con     io.ReadWriter
}

func (n *NIC) Connect() io.ReadWriter {
	return n.con
}
