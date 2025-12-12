//go:build !with_quic

package include

import (
	"context"
	"io"
	"net/http"

	"github.com/pulsarvpn/sing-box/adapter"
	"github.com/pulsarvpn/sing-box/adapter/inbound"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	"github.com/pulsarvpn/sing-box/common/listener"
	"github.com/pulsarvpn/sing-box/common/tls"
	C "github.com/pulsarvpn/sing-box/constant"
	"github.com/pulsarvpn/sing-box/dns"
	"github.com/pulsarvpn/sing-box/log"
	"github.com/pulsarvpn/sing-box/option"
	"github.com/pulsarvpn/sing-box/protocol/naive"
	"github.com/pulsarvpn/sing-box/transport/v2ray"
	"github.com/sagernet/sing/common/logger"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

func init() {
	v2ray.RegisterQUICConstructor(
		func(ctx context.Context, logger logger.ContextLogger, options option.V2RayQUICOptions, tlsConfig tls.ServerConfig, handler adapter.V2RayServerTransportHandler) (adapter.V2RayServerTransport, error) {
			return nil, C.ErrQUICNotIncluded
		},
		func(ctx context.Context, dialer N.Dialer, serverAddr M.Socksaddr, options option.V2RayQUICOptions, tlsConfig tls.Config) (adapter.V2RayClientTransport, error) {
			return nil, C.ErrQUICNotIncluded
		},
	)
}

func registerQUICInbounds(registry *inbound.Registry) {
	inbound.Register[option.HysteriaInboundOptions](registry, C.TypeHysteria, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.HysteriaInboundOptions) (adapter.Inbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
	inbound.Register[option.TUICInboundOptions](registry, C.TypeTUIC, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.TUICInboundOptions) (adapter.Inbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
	inbound.Register[option.Hysteria2InboundOptions](registry, C.TypeHysteria2, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Hysteria2InboundOptions) (adapter.Inbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
	naive.ConfigureHTTP3ListenerFunc = func(listener *listener.Listener, handler http.Handler, tlsConfig tls.ServerConfig, logger logger.Logger) (io.Closer, error) {
		return nil, C.ErrQUICNotIncluded
	}
}

func registerQUICOutbounds(registry *outbound.Registry) {
	outbound.Register[option.HysteriaOutboundOptions](registry, C.TypeHysteria, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.HysteriaOutboundOptions) (adapter.Outbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
	outbound.Register[option.TUICOutboundOptions](registry, C.TypeTUIC, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.TUICOutboundOptions) (adapter.Outbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
	outbound.Register[option.Hysteria2OutboundOptions](registry, C.TypeHysteria2, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Hysteria2OutboundOptions) (adapter.Outbound, error) {
		return nil, C.ErrQUICNotIncluded
	})
}

func registerQUICTransports(registry *dns.TransportRegistry) {
	dns.RegisterTransport[option.RemoteTLSDNSServerOptions](registry, C.DNSTypeQUIC, func(ctx context.Context, logger log.ContextLogger, tag string, options option.RemoteTLSDNSServerOptions) (adapter.DNSTransport, error) {
		return nil, C.ErrQUICNotIncluded
	})
	dns.RegisterTransport[option.RemoteHTTPSDNSServerOptions](registry, C.DNSTypeHTTP3, func(ctx context.Context, logger log.ContextLogger, tag string, options option.RemoteHTTPSDNSServerOptions) (adapter.DNSTransport, error) {
		return nil, C.ErrQUICNotIncluded
	})
}
