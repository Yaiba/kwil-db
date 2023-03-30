package client2

import (
	"crypto/ecdsa"
	chainCodes "kwil/pkg/chain/types"
)

type ClientOpt func(*Client)

func WithPrivateKey(key *ecdsa.PrivateKey) ClientOpt {
	return func(c *Client) {
		c.PrivateKey = key
	}
}

func WithChainCode(chainCode int32) ClientOpt {
	return func(c *Client) {
		c.chainCode = chainCodes.ChainCode(chainCode)
	}
}

func WithProviderAddress(address string) ClientOpt {
	return func(c *Client) {
		c.providerAddress = address
	}
}

func WithPoolAddress(address string) ClientOpt {
	return func(c *Client) {
		c.poolAddress = address
	}
}

func WithChainRpcUrl(url string) ClientOpt {
	return func(c *Client) {
		c.chainRpcUrl = url
	}
}

func WithoutProvider() ClientOpt {
	return func(c *Client) {
		c.usingProvider = false
	}
}
