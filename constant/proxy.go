package constant

const (
	TypeTun          = "tun"
	TypeRedirect     = "redirect"
	TypeDirect       = "direct"
	TypeBlock        = "block"
	TypeDNS          = "dns"
	TypeSOCKS        = "socks"
	TypeHTTP         = "http"
	TypeMixed        = "mixed"
	TypeVMess        = "vmess"
	TypeNaive        = "naive"
	TypeVLESS        = "vless"
	TypeTUIC         = "tuic"
	TypeHysteria2    = "hysteria2"
	TypeTailscale    = "tailscale"
	TypeDERP         = "derp"
	TypeResolved     = "resolved"
	TypeSSMAPI       = "ssm-api"
	TypeCCM          = "ccm"
	TypeOCM          = "ocm"
	TypeOOMKiller    = "oom-killer"
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
	case TypeDirect:
		return "Direct"
	case TypeBlock:
		return "Block"
	case TypeDNS:
		return "DNS"
	case TypeSOCKS:
		return "SOCKS"
	case TypeHTTP:
		return "HTTP"
	case TypeMixed:
		return "Mixed"
	case TypeVMess:
		return "VMess"
	case TypeNaive:
		return "Naive"
	case TypeVLESS:
		return "VLESS"
	case TypeTUIC:
		return "TUIC"
	case TypeHysteria2:
		return "Hysteria2"
	case TypeTailscale:
		return "Tailscale"
	case TypeSelector:
		return "Selector"
	case TypeURLTest:
		return "URLTest"
	default:
		return "Unknown"
	}
}
