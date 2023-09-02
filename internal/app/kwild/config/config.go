package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cometCfg "github.com/cometbft/cometbft/config"
	"github.com/kwilteam/kwil-db/pkg/log"
	"github.com/spf13/viper"
)

type KwildConfig struct {
	RootDir string

	AppCfg   *AppConfig       `mapstructure:"app"`
	ChainCfg *cometCfg.Config `mapstructure:"chain"`
	Logging  *Logging         `mapstructure:"log"`
}

type Logging struct {
	LogLevel    string   `mapstructure:"log_level"`
	LogFormat   string   `mapstructure:"log_format"`
	OutputPaths []string `mapstructure:"output_paths"`
}

type AppConfig struct {
	GrpcListenAddress  string         `mapstructure:"grpc_listen_addr"`
	HttpListenAddress  string         `mapstructure:"http_listen_addr"`
	PrivateKey         string         `mapstructure:"private_key"`
	SqliteFilePath     string         `mapstructure:"sqlite_file_path"`
	ExtensionEndpoints []string       `mapstructure:"extension_endpoints"`
	WithoutGasCosts    bool           `mapstructure:"without_gas_costs"`
	WithoutNonces      bool           `mapstructure:"without_nonces"`
	SnapshotConfig     SnapshotConfig `mapstructure:"snapshots"`
	TLSCertFile        string         `mapstructure:"tls_cert_file"`
	TLSKeyFile         string         `mapstructure:"tls_key_file"`
	Hostname           string         `mapstructure:"hostname"`
	Log                log.Config
}

type SnapshotConfig struct {
	Enabled         bool   `mapstructure:"enabled"`
	RecurringHeight uint64 `mapstructure:"snapshot_heights"`
	MaxSnapshots    uint64 `mapstructure:"max_snapshots"`
	SnapshotDir     string `mapstructure:"snapshot_dir"`
}

func (cfg *KwildConfig) LoadKwildConfig(rootDir string) error {
	// Expand root dir path
	rootDir, err := ExpandPath(rootDir)
	if err != nil {
		return fmt.Errorf("failed to expand rootdir path: %v", err)
	}
	cfg.RootDir = rootDir

	cfgFile := filepath.Join(rootDir, "config.toml")
	err = cfg.ParseConfig(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	err = cfg.sanitizeCfgPaths()
	if err != nil {
		return fmt.Errorf("failed to sanitize config paths: %v", err)
	}

	cfg.configureLogging()
	cfg.configureCerts()

	if err := cfg.ChainCfg.ValidateBasic(); err != nil {
		return fmt.Errorf("invalid chain configuration data: %v", err)
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (cfg *KwildConfig) ParseConfig(cfgFile string) error {
	/*
		Lots of Viper magic here, but the gist is:
		We want to be able to set config values via
			-  flags
			-  environment variables
			-  config file
			-  default values

		for env variables support:
		Requirement is, we need to be able to config from env variables with a prefix "KWILD_"

		It can be done 2 ways:
		1. AutomaticEnv: off mode
			- This will not bind env variables to config values automatically
			- We need to manually bind env variables to config values (this is what we are doing currently)
			- As we bound flags to viper, viper is already aware of the config structure mapping,
				so we can explicitly call viper.BindEnv() on all the keys in viper.AllKeys()
			- else we would have to reflect on the config structure and bind env variables to config values

		2. AutomaticEnv: on mode
			- This is supposed to automatically bind env variables to config values
				(but it doesn't work without doing a bit more work from our side)
			- One way to make this work is add default values using either viper.SetDefault() for all the config values
			  or can do viper.MergeConfig(<serialized config>)
			- Serializing is really painful as cometbft has a field which is using map<interface><interface> though its deprecated.
				which prevents us from doing the AutomaticEnv binding
		Issues referencing the issues (or) correct usage of AutomaticEnv: https://github.com/spf13/viper/issues/188
		For now, we are going with the first approach

		Note:
		The order of preference of various modes of config supported by viper is:
		explicit call to Set > flags > env variables > config file > default values
	*/

	for _, key := range viper.AllKeys() {
		envKey := "KWILD_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		viper.BindEnv(key, envKey)
	}

	if fileExists(cfgFile) {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("reading config: %v", err)
		}
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("decoding config: %v", err)
	}

	return nil
}

func DefaultConfig() *KwildConfig {
	cfg := &KwildConfig{
		ChainCfg: cometCfg.DefaultConfig(),
		AppCfg: &AppConfig{
			GrpcListenAddress: "localhost:50051",
			HttpListenAddress: "localhost:8080",
			SqliteFilePath:    "data/kwil.db",
			WithoutGasCosts:   true,
			WithoutNonces:     false,
			SnapshotConfig: SnapshotConfig{
				Enabled:         false,
				RecurringHeight: uint64(10000),
				MaxSnapshots:    3,
				SnapshotDir:     "snapshots",
			},
		},
		Logging: &Logging{
			LogLevel:    "info",
			LogFormat:   "plain",
			OutputPaths: []string{"stdout"},
		},
	}

	// PEX is recommended to be disabled for validators: https://docs.cometbft.com/v0.37/core/validators#validator-node-configuration
	cfg.ChainCfg.P2P.PexReactor = false

	/*
	 As all we are validating are tx signatures, no need to go through Validation again
	 To be set to true when we have Validations based on gas, nonces, account balance, etc.
	*/
	cfg.ChainCfg.Mempool.Recheck = false
	return cfg
}

func (cfg *KwildConfig) configureLogging() {
	// App Logging
	cfg.AppCfg.Log.Level = cfg.Logging.LogLevel
	cfg.AppCfg.Log.OutputPaths = cfg.Logging.OutputPaths

	// Chain Logging
	cfg.ChainCfg.LogLevel = cfg.Logging.LogLevel
	cfg.ChainCfg.LogFormat = cfg.Logging.LogFormat
}

func (cfg *KwildConfig) configureCerts() {
	if cfg.AppCfg.TLSCertFile != "" {
		cfg.AppCfg.TLSCertFile = rootify(cfg.AppCfg.TLSCertFile, cfg.RootDir)
		cfg.ChainCfg.RPC.TLSCertFile = cfg.AppCfg.TLSCertFile
	}

	if cfg.AppCfg.TLSKeyFile != "" {
		cfg.AppCfg.TLSKeyFile = rootify(cfg.AppCfg.TLSKeyFile, cfg.RootDir)
		cfg.ChainCfg.RPC.TLSKeyFile = cfg.AppCfg.TLSKeyFile
	}
}

func rootify(path, rootDir string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(rootDir, path)
}

func (cfg *KwildConfig) sanitizeCfgPaths() error {
	rootDir := cfg.RootDir
	cfg.AppCfg.SqliteFilePath = rootify(cfg.AppCfg.SqliteFilePath, rootDir)
	cfg.AppCfg.SnapshotConfig.SnapshotDir = rootify(cfg.AppCfg.SnapshotConfig.SnapshotDir, rootDir)

	cfg.ChainCfg.SetRoot(filepath.Join(rootDir, "abci"))
	return nil
}

func ExpandPath(path string) (string, error) {
	var expandedPath string

	if tail, cut := strings.CutPrefix(path, "~/"); cut {
		// Expands ~ in the path
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		expandedPath = filepath.Join(homeDir, tail)
	} else {
		// Expands relative paths
		absPath, err := filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("failed to get absolute path of file: %v due to error: %v", path, err)
		}
		expandedPath = absPath
	}
	return expandedPath, nil
}
