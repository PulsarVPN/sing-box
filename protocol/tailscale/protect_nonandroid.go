//go:build !android

package tailscale

import "github.com/pulsarvpn/sing-box/experimental/libbox/platform"

func setAndroidProtectFunc(platformInterface platform.Interface) {
}
