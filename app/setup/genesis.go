package setup

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/kwilteam/kwil-db/app/shared/display"
	"github.com/kwilteam/kwil-db/config"
	"github.com/kwilteam/kwil-db/core/crypto"
	"github.com/kwilteam/kwil-db/core/types"
	"github.com/kwilteam/kwil-db/node"
	"github.com/spf13/cobra"
)

var (
	genesisLong = `` + "`" + `genesis` + "`" + ` creates a new genesis.json file.

This command creates a new genesis file with optionally specified modifications.

Validators, balance allocations should have the format "pubkey:power", "address:balance" respectively.`

	genesisExample = `# Create a new genesis.json file in the current directory
kwil-admin setup genesis

# Create a new genesis.json file in a specific directory with a specific chain ID and a validator with 1 power
kwil-admin setup genesis --out /path/to/directory --chain-id mychainid --validator 890fe7ae9cb1fa6177555d5651e1b8451b4a9c64021c876236c700bc2690ff1d:1

# Create a new genesis.json with the specified allocation
kwil-admin setup genesis --alloc 0x7f5f4552091a69125d5dfcb7b8c2659029395bdf:100`
)

type genesisFlagConfig struct {
	chainID       string
	validators    []string
	allocs        []string
	withGas       bool
	leader        string
	dbOwner       string
	maxBlockSize  int64
	joinExpiry    int64
	voteExpiry    int64
	maxVotesPerTx int64
	genesisState  string
}

func GenesisCmd() *cobra.Command {
	var flagCfg genesisFlagConfig
	var output string

	cmd := &cobra.Command{
		Use:               "genesis",
		Short:             "`genesis` creates a new genesis.json file",
		Long:              genesisLong,
		Example:           genesisExample,
		DisableAutoGenTag: true,
		SilenceUsage:      true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			outDir, err := node.ExpandPath(output)
			if err != nil {
				return display.PrintErr(cmd, fmt.Errorf("failed to expand output path: %w", err))
			}

			err = os.MkdirAll(outDir, nodeDirPerm)
			if err != nil {
				return display.PrintErr(cmd, fmt.Errorf("failed to create output directory: %w", err))
			}

			genesisFile := config.GenesisFilePath(outDir)

			conf := config.DefaultGenesisConfig()
			conf, err = mergeGenesisFlags(conf, cmd, &flagCfg)
			if err != nil {
				return display.PrintErr(cmd, err)
			}

			existingFile, err := os.Stat(genesisFile)
			if err == nil && existingFile.IsDir() {
				return display.PrintErr(cmd, fmt.Errorf("a directory already exists at %s, please remove it first", genesisFile))
			} else if err == nil {
				return display.PrintErr(cmd, fmt.Errorf("file already exists at %s, please remove it first", genesisFile))
			}

			err = conf.SaveAs(genesisFile)
			if err != nil {
				return display.PrintErr(cmd, fmt.Errorf("failed to save genesis file: %w", err))
			}

			return display.PrintCmd(cmd, display.RespString("Created genesis.json file at "+genesisFile))
		},
	}

	bindGenesisFlags(cmd, &flagCfg)
	cmd.Flags().StringVar(&output, "out", "", "Output directory for the genesis.json file")

	return cmd
}

// bindGenesisFlags binds the genesis configuration flags to the given command.
func bindGenesisFlags(cmd *cobra.Command, cfg *genesisFlagConfig) {
	cmd.Flags().StringVar(&cfg.chainID, "chain-id", "", "chainID for the genesis.json file")
	cmd.Flags().StringSliceVar(&cfg.validators, "validators", nil, "public key, keyType and power of initial validator(s)") // accept: [pubkey1#keyType1:power1]
	cmd.Flags().StringSliceVar(&cfg.allocs, "allocs", nil, "address and initial balance allocation(s)")
	cmd.Flags().BoolVar(&cfg.withGas, "with-gas", false, "include gas costs in the genesis.json file")
	cmd.Flags().StringVar(&cfg.leader, "leader", "", "public key of the block proposer")
	cmd.Flags().StringVar(&cfg.dbOwner, "db-owner", "", "owner of the database")
	cmd.Flags().Int64Var(&cfg.maxBlockSize, "max-block-size", 0, "maximum block size")
	cmd.Flags().Int64Var(&cfg.joinExpiry, "join-expiry", 0, "Number of blocks before a join proposal expires")
	cmd.Flags().Int64Var(&cfg.voteExpiry, "vote-expiry", 0, "Number of blocks before a vote proposal expires")
	cmd.Flags().Int64Var(&cfg.maxVotesPerTx, "max-votes-per-tx", 0, "Maximum votes per transaction")
	cmd.Flags().StringVar(&cfg.genesisState, "genesis-snapshot", "", "path to genesis state snapshot file")
}

// mergeGenesisFlags merges the genesis configuration flags with the given configuration.
func mergeGenesisFlags(conf *config.GenesisConfig, cmd *cobra.Command, flagCfg *genesisFlagConfig) (*config.GenesisConfig, error) {
	makeErr := func(e error) error {
		return display.PrintErr(cmd, fmt.Errorf("failed to create genesis file: %w", e))
	}
	if cmd.Flags().Changed("chain-id") {
		conf.ChainID = flagCfg.chainID
	}

	if cmd.Flags().Changed("validators") {
		conf.Validators = nil
		for _, v := range flagCfg.validators {
			parts := strings.Split(v, ":")
			if len(parts) != 2 {
				return nil, makeErr(fmt.Errorf("invalid format for validator, expected key:power, received: %s", v))
			}

			keyParts := strings.Split(parts[0], "#")
			hexPub, err := hex.DecodeString(keyParts[0])
			if err != nil {
				return nil, makeErr(fmt.Errorf("invalid public key for validator: %s", parts[0]))
			}

			power, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, makeErr(fmt.Errorf("invalid power for validator: %s", parts[1]))
			}

			// string to Int
			keyType, err := strconv.ParseInt(keyParts[1], 10, 64)
			if err != nil {
				return nil, makeErr(fmt.Errorf("invalid power for validator: %s", keyParts[1]))
			}

			conf.Validators = append(conf.Validators, &types.Validator{
				PubKey:     hexPub,
				PubKeyType: crypto.KeyType(keyType),
				Power:      power,
			})
		}
	}

	if cmd.Flags().Changed("allocs") {
		conf.Allocs = nil
		for _, a := range flagCfg.allocs {
			parts := strings.Split(a, ":")
			if len(parts) != 2 {
				return nil, makeErr(fmt.Errorf("invalid format for alloc, expected address:balance, received: %s", a))
			}

			balance, ok := new(big.Int).SetString(parts[1], 10)
			if !ok {
				return nil, makeErr(fmt.Errorf("invalid balance for alloc: %s", parts[1]))
			}

			conf.Allocs[parts[0]] = balance
		}
	}

	if cmd.Flags().Changed("with-gas") {
		conf.DisabledGasCosts = !flagCfg.withGas
	}

	if cmd.Flags().Changed("leader") {
		conf.Leader = flagCfg.leader
	}

	if cmd.Flags().Changed("db-owner") {
		conf.DBOwner = flagCfg.dbOwner
	}

	if cmd.Flags().Changed("max-block-size") {
		conf.MaxBlockSize = flagCfg.maxBlockSize
	}

	if cmd.Flags().Changed("join-expiry") {
		conf.JoinExpiry = flagCfg.joinExpiry
	}

	if cmd.Flags().Changed("vote-expiry") {
		conf.VoteExpiry = flagCfg.voteExpiry
	}

	if cmd.Flags().Changed("max-votes-per-tx") {
		conf.MaxVotesPerTx = flagCfg.maxVotesPerTx
	}

	if cmd.Flags().Changed("genesis-state") {
		hash, err := appHashFromSnapshotFile(flagCfg.genesisState)
		if err != nil {
			return nil, makeErr(err)
		}
		conf.StateHash = hash
	}

	return conf, nil
}