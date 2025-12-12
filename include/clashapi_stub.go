//go:build !with_clash_api

package include

import (
	"context"

	"github.com/pulsarvpn/sing-box/adapter"
	"github.com/pulsarvpn/sing-box/experimental"
	"github.com/pulsarvpn/sing-box/log"
	"github.com/pulsarvpn/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func init() {
	experimental.RegisterClashServerConstructor(func(ctx context.Context, logFactory log.ObservableFactory, options option.ClashAPIOptions) (adapter.ClashServer, error) {
		return nil, E.New(`clash api is not included in this build, rebuild with -tags with_clash_api`)
	})
}
