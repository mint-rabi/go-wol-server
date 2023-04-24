package go_wol_server

import (
	"fmt"
	"github.com/sabhiram/go-wol/wol"
	"net"
)

func Send(address string) error {
	bcastAddr := fmt.Sprintf("%s:%v", "255.255.255.255", 9)
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		return err
	}
	mp, err := wol.New(address)
	if err != nil {
		return err
	}

	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		return err
	}
	return nil
}
