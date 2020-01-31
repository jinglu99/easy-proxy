package rewriter

import (
	"errors"
	"fmt"
	"github.com/armon/go-socks5"
	"github.com/jingleWang/easy-proxy/src/rule"
	"golang.org/x/net/context"
	"net"
)

type EasyRewriter struct {
	Rules *[]rule.Rule
}

func (rewriter *EasyRewriter) Rewrite(ctx context.Context, request *socks5.Request) (context.Context, *socks5.AddrSpec) {

	dest := request.DestAddr
	if dest.FQDN == "" {
		return ctx, request.DestAddr
	}

	for _, r := range *rewriter.Rules {
		if matched, addr, err := doRewriteIfMatch(dest, r); matched && err == nil {
			return ctx, addr
		}
	}
	return ctx, request.DestAddr
}

func doRewriteIfMatch(dest *socks5.AddrSpec, r rule.Rule) (bool, *socks5.AddrSpec, error) {
	if r.IsMatch(dest.FQDN) {
		dr := r.(*rule.DefaultRule)
		fmt.Printf("host: %s matched rule: %s\n", dest.FQDN, dr.Pattern)
		addr, err := net.ResolveIPAddr("ip", dr.Host)
		if err != nil {
			fmt.Printf("can't resolve IP addr of %s\n", dr.Host)
			return true, nil, errors.New("can't resolve IP addr of " + dr.Host)
		}

		port := dest.Port
		if dr.Port > 0 && dr.Port <= 65535 {
			port = dr.Port
		}

		return true, &socks5.AddrSpec{
			FQDN: dest.FQDN,
			IP:   addr.IP,
			Port: port,
		}, nil
	}
	return false, nil, nil
}
