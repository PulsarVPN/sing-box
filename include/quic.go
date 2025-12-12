//go:build with_quic

package include

import (
	"github.com/pulsarvpn/sing-box/adapter/inbound"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	"github.com/pulsarvpn/sing-box/dns"
	"github.com/pulsarvpn/sing-box/dns/transport/quic"
	"github.com/pulsarvpn/sing-box/protocol/hysteria"
	"github.com/pulsarvpn/sing-box/protocol/hysteria2"
	_ "github.com/pulsarvpn/sing-box/protocol/naive/quic"
	"github.com/pulsarvpn/sing-box/protocol/tuic"
	_ "github.com/pulsarvpn/sing-box/transport/v2rayquic"
)

func registerQUICInbounds(registry *inbound.Registry) {
	hysteria.RegisterInbound(registry)
	tuic.RegisterInbound(registry)
	hysteria2.RegisterInbound(registry)
}

func registerQUICOutbounds(registry *outbound.Registry) {
	hysteria.RegisterOutbound(registry)
	tuic.RegisterOutbound(registry)
	hysteria2.RegisterOutbound(registry)
}

func registerQUICTransports(registry *dns.TransportRegistry) {
	quic.RegisterTransport(registry)
	quic.RegisterHTTP3Transport(registry)
}
