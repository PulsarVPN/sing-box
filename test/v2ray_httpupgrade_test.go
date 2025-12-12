package main

import (
	"testing"

	C "github.com/PulsarVPN/sing-box/constant"
	"github.com/PulsarVPN/sing-box/option"
)

func TestV2RayHTTPUpgrade(t *testing.T) {
	t.Run("self", func(t *testing.T) {
		testV2RayTransportSelf(t, &option.V2RayTransportOptions{
			Type: C.V2RayTransportTypeHTTPUpgrade,
		})
	})
}
