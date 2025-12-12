//go:build !android

package tailscale

import "github.com/PulsarVPN/sing-box/experimental/libbox/platform"

func setAndroidProtectFunc(platformInterface platform.Interface) {
}
