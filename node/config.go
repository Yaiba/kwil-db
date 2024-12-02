package node

import (
	"github.com/kwilteam/kwil-db/config"
	"github.com/kwilteam/kwil-db/core/crypto"
	"github.com/kwilteam/kwil-db/core/log"
	"github.com/kwilteam/kwil-db/node/types"
)

// Config is the configuration for a [Node] instance.
type Config struct {
	RootDir string
	PrivKey crypto.PrivateKey

	P2P *config.PeerConfig

	Mempool    types.MemPool
	BlockStore types.BlockStore
	Consensus  ConsensusEngine
	Logger     log.Logger
}