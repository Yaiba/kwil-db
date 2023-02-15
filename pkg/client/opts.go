package client

type ClientOption func(*client)

// if a user configures something that is also set by the service config,
// it will use the user's config over the service config

// WithProviderAddress sets the provider address to use
func WithProviderAddress(address string) ClientOption {
	return func(c *client) {
		c.providerAddress = address
	}
}

// WithChainRpcUrl sets the chain rpc url to use
func WithChainRpcUrl(url string) ClientOption {
	return func(c *client) {
		c.chainRpcUrl = &url
	}
}

// WithEscrowAddress sets the escrow contract address to use
func WithEscrowAddress(address string) ClientOption {
	return func(c *client) {
		c.escrowContractAddress = address
	}
}

// WithChainCode sets the chain code to use
func WithChainCode(code int64) ClientOption {
	return func(c *client) {
		c.chainCode = code
	}
}

// WithoutServiceConfig disables the use of the config
// provided by the kwil provider
func WithoutServiceConfig() ClientOption {
	return func(c *client) {
		c.usingServiceCfg = false
	}
}
