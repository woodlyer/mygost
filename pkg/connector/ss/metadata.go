package ss

import (
	"time"

	"github.com/go-gost/gost/pkg/common/util/ss"
	md "github.com/go-gost/gost/pkg/metadata"
	"github.com/shadowsocks/go-shadowsocks2/core"
)

type metadata struct {
	cipher         core.Cipher
	connectTimeout time.Duration
	noDelay        bool
}

func (c *ssConnector) parseMetadata(md md.Metadata) (err error) {
	const (
		method         = "method"
		password       = "password"
		key            = "key"
		connectTimeout = "timeout"
		noDelay        = "nodelay"
	)

	c.md.cipher, err = ss.ShadowCipher(
		md.GetString(method),
		md.GetString(password),
		md.GetString(key),
	)
	if err != nil {
		return
	}

	c.md.connectTimeout = md.GetDuration(connectTimeout)
	c.md.noDelay = md.GetBool(noDelay)

	return
}
