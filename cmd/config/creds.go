package config

import (
	"strings"

	boshuaa "github.com/cloudfoundry/bosh-init/uaa"
)

type Creds struct {
	Username string
	Password string

	// For UAA clients
	Client       string
	ClientSecret string

	// For UAA users
	RefreshToken string
}

func (c Creds) IsBasic() bool {
	return len(c.Username) > 0
}

func (c Creds) IsBasicComplete() bool {
	return len(c.Username) > 0 && len(c.Password) > 0
}

func (c Creds) IsUAA() bool {
	return len(c.Client) > 0 || len(c.RefreshToken) > 0
}

func (c Creds) IsUAAClient() bool {
	// Clients dont use refresh tokens
	return len(c.Client) > 0
}

func (c Creds) Description() string {
	if len(c.Username) > 0 {
		return credsDesc{name: c.Username, kind: "user"}.String()
	}

	if len(c.RefreshToken) > 0 {
		info, err := boshuaa.NewTokenInfoFromValue(c.RefreshToken)
		if err != nil {
			return credsDesc{name: "?"}.String()
		}

		desc := credsDesc{
			name: info.Username,
			kind: "user",
			misc: strings.Join(info.Scopes, ", "),
		}

		return desc.String()
	}

	if len(c.Client) > 0 {
		return credsDesc{name: c.Client, kind: "client"}.String()
	}

	return credsDesc{full: "anonymous user"}.String()
}

type credsDesc struct {
	name string
	kind string
	misc string
	full string
}

func (c credsDesc) String() string {
	if len(c.full) > 0 {
		return c.full
	}

	str := ""

	if len(c.kind) > 0 {
		str = c.kind + " "
	}

	str += "'" + c.name + "'"

	if len(c.misc) > 0 {
		str += " (" + c.misc + ")"
	}

	return str
}
