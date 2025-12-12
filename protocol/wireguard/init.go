package wireguard

import (
	"github.com/PulsarVPN/sing-box/common/dialer"
	"github.com/sagernet/wireguard-go/conn"
)

func init() {
	dialer.WgControlFns = conn.ControlFns
}
