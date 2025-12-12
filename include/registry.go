package include

import (
	"context"

	box "github.com/pulsarvpn/sing-box"
	"github.com/pulsarvpn/sing-box/adapter/endpoint"
	"github.com/pulsarvpn/sing-box/adapter/inbound"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	"github.com/pulsarvpn/sing-box/adapter/service"
	"github.com/pulsarvpn/sing-box/dns"
	"github.com/pulsarvpn/sing-box/dns/transport"
	"github.com/pulsarvpn/sing-box/dns/transport/fakeip"
	"github.com/pulsarvpn/sing-box/dns/transport/hosts"
	"github.com/pulsarvpn/sing-box/dns/transport/local"
	"github.com/pulsarvpn/sing-box/protocol/anytls"
	"github.com/pulsarvpn/sing-box/protocol/block"
	"github.com/pulsarvpn/sing-box/protocol/direct"
	protocolDNS "github.com/pulsarvpn/sing-box/protocol/dns"
	"github.com/pulsarvpn/sing-box/protocol/group"
	"github.com/pulsarvpn/sing-box/protocol/http"
	"github.com/pulsarvpn/sing-box/protocol/mixed"
	"github.com/pulsarvpn/sing-box/protocol/naive"
	"github.com/pulsarvpn/sing-box/protocol/redirect"
	"github.com/pulsarvpn/sing-box/protocol/trojan"
	"github.com/pulsarvpn/sing-box/protocol/tun"
	"github.com/pulsarvpn/sing-box/protocol/vless"
	"github.com/pulsarvpn/sing-box/service/resolved"
	"github.com/pulsarvpn/sing-box/service/ssmapi"
)

func Context(ctx context.Context) context.Context {
	return box.Context(ctx, InboundRegistry(), OutboundRegistry(), EndpointRegistry(), DNSTransportRegistry(), ServiceRegistry())
}

func InboundRegistry() *inbound.Registry {
	registry := inbound.NewRegistry()

	tun.RegisterInbound(registry)
	redirect.RegisterRedirect(registry)
	redirect.RegisterTProxy(registry)
	direct.RegisterInbound(registry)

	http.RegisterInbound(registry)
	mixed.RegisterInbound(registry)

	trojan.RegisterInbound(registry)
	naive.RegisterInbound(registry)
	vless.RegisterInbound(registry)
	anytls.RegisterInbound(registry)

	registerQUICInbounds(registry)

	return registry
}

func OutboundRegistry() *outbound.Registry {
	registry := outbound.NewRegistry()

	direct.RegisterOutbound(registry)

	block.RegisterOutbound(registry)
	protocolDNS.RegisterOutbound(registry)

	group.RegisterSelector(registry)
	group.RegisterURLTest(registry)

	http.RegisterOutbound(registry)
	trojan.RegisterOutbound(registry)
	vless.RegisterOutbound(registry)
	anytls.RegisterOutbound(registry)

	registerQUICOutbounds(registry)
	registerWireGuardOutbound(registry)

	return registry
}

func EndpointRegistry() *endpoint.Registry {
	registry := endpoint.NewRegistry()

	registerWireGuardEndpoint(registry)

	return registry
}

func DNSTransportRegistry() *dns.TransportRegistry {
	registry := dns.NewTransportRegistry()

	transport.RegisterTCP(registry)
	transport.RegisterUDP(registry)
	transport.RegisterTLS(registry)
	transport.RegisterHTTPS(registry)
	hosts.RegisterTransport(registry)
	local.RegisterTransport(registry)
	fakeip.RegisterTransport(registry)
	resolved.RegisterTransport(registry)

	registerQUICTransports(registry)
	registerDHCPTransport(registry)

	return registry
}

func ServiceRegistry() *service.Registry {
	registry := service.NewRegistry()

	resolved.RegisterService(registry)
	ssmapi.RegisterService(registry)

	return registry
}
