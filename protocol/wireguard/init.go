package wireguard

import (
	"github.com/pulsarvpn/sing-box/common/dialer"
	"github.com/pulsarvpn/wireguard-go/conn"
)

func init() {
	dialer.WgControlFns = conn.ControlFns
}
