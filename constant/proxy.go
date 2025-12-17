package constant

const (
	TypeTun       = "tun"
	TypeRedirect  = "redirect"
	TypeTProxy    = "tproxy"
	TypeDirect    = "direct"
	TypeBlock     = "block"
	TypeDNS       = "dns"
	TypeHTTP      = "http"
	TypeMixed     = "mixed"
	TypeTrojan    = "trojan"
	TypeNaive     = "naive"
	TypeWireGuard = "wireguard"
	TypeAmneziaWG = "amneziawg"
	TypeHysteria  = "hysteria"
	TypeAnyTLS    = "anytls"
	TypeVLESS     = "vless"
	TypeTUIC      = "tuic"
	TypeHysteria2 = "hysteria2"
	TypeDERP      = "derp"
	TypeResolved  = "resolved"
	TypeSSMAPI    = "ssm-api"
)

const (
	TypeSelector = "selector"
	TypeURLTest  = "urltest"
)

func ProxyDisplayName(proxyType string) string {
	switch proxyType {
	case TypeTun:
		return "TUN"
	case TypeRedirect:
		return "Redirect"
	case TypeTProxy:
		return "TProxy"
	case TypeDirect:
		return "Direct"
	case TypeBlock:
		return "Block"
	case TypeDNS:
		return "DNS"
	case TypeHTTP:
		return "HTTP"
	case TypeMixed:
		return "Mixed"
	case TypeTrojan:
		return "Trojan"
	case TypeNaive:
		return "Naive"
	case TypeWireGuard:
		return "WireGuard"
	case TypeAmneziaWG:
		return "AmneziaWG"
	case TypeHysteria:
		return "Hysteria"
	case TypeVLESS:
		return "VLESS"
	case TypeTUIC:
		return "TUIC"
	case TypeHysteria2:
		return "Hysteria2"
	case TypeAnyTLS:
		return "AnyTLS"
	case TypeSelector:
		return "Selector"
	case TypeURLTest:
		return "URLTest"
	default:
		return "Unknown"
	}
}
