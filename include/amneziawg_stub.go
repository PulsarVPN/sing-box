//go:build !with_amneziawg

package include

import (
	"context"

	"github.com/pulsarvpn/sing-box/adapter"
	"github.com/pulsarvpn/sing-box/adapter/endpoint"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	C "github.com/pulsarvpn/sing-box/constant"
	"github.com/pulsarvpn/sing-box/log"
	"github.com/pulsarvpn/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func registerAmneziaWGOutbound(registry *outbound.Registry) {
	outbound.Register[option.LegacyAmneziaWGOutboundOptions](registry, C.TypeAmneziaWG, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.LegacyAmneziaWGOutboundOptions) (adapter.Outbound, error) {
		return nil, E.New(`AmneziaWG is not included in this build, rebuild with -tags with_amneziawg`)
	})
}

func registerAmneziaWGEndpoint(registry *endpoint.Registry) {
	endpoint.Register[option.AmneziaWGEndpointOptions](registry, C.TypeAmneziaWG, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.AmneziaWGEndpointOptions) (adapter.Endpoint, error) {
		return nil, E.New(`AmneziaWG is not included in this build, rebuild with -tags with_amneziawg`)
	})
}

