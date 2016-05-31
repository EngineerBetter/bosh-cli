package cmd

import (
	boshuaa "github.com/cloudfoundry/bosh-init/uaa"
)

type ClientTokenSession struct {
	uaa       boshuaa.UAA
	lastToken boshuaa.Token
}

func NewClientTokenSession(uaa boshuaa.UAA) *ClientTokenSession {
	return &ClientTokenSession{uaa: uaa}
}

func (c *ClientTokenSession) TokenFunc(retried bool) (string, error) {
	if c.lastToken == nil || retried {
		token, err := c.uaa.ClientCredentialsGrant()
		if err != nil {
			return "", err
		}

		c.lastToken = token
	}

	return c.lastToken.Type() + " " + c.lastToken.Value(), nil
}
