package config

import (
	"errors"
	"flag"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	cfgOnce       sync.Once
	defaultConfig Config
)

type configImpl struct {
	root   map[string]string
	source map[interface{}]interface{}
	prefix string
}

func (c *configImpl) As(out interface{}) error {
	b, err := yaml.Marshal(c.source)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, out)
}

func (c *configImpl) Exists(key string) bool {
	if _, ok := c.root[c.normalize(key)]; ok {
		return true
	}
	return false
}

func (c *configImpl) ToStringMap() map[string]string {
	result := make(map[string]string)

	for k, v := range c.root {
		if strings.HasPrefix(k, c.prefix) {
			k2 := strings.TrimPrefix(k, c.prefix)
			result[k2] = v
		}
	}

	return result
}

func (c *configImpl) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range c.source {
		key := k.(string)
		if strings.HasPrefix(key, c.prefix) {
			k2 := strings.TrimPrefix(key, c.prefix)
			result[k2] = v
		}
	}

	return result
}

func (c *configImpl) Select(key string) Config {
	if key == "" {
		return c
	}

	m, ok := c.source[key]
	if !ok {
		return emptyConfig
	}

	var k = c.normalize(key) + "."
	if ok && reflect.TypeOf(m).Kind() == reflect.Map {
		return &configImpl{c.root, m.(map[interface{}]interface{}), k}
	}

	return &configImpl{c.root, make(map[interface{}]interface{}), k}
}

func (c *configImpl) String(key string) string {
	return c.GetString(key, "")
}

func (c *configImpl) GetString(key string, defaultValue string) string {
	if v, ok := c.root[c.normalize(key)]; ok {
		s := reflect.ValueOf(v).String()
		return s
	}

	return defaultValue
}

func (c *configImpl) Int32(key string, defaultValue int32) int32 {
	v, err := c.GetInt32(key, defaultValue)
	if err == nil {
		return v
	}

	log.Default().Println("Failed to parse num as int32", err)
	return defaultValue
}

func (c *configImpl) Int64(key string, defaultValue int64) int64 {
	v, err := c.GetInt64(key, defaultValue)
	if err == nil {
		return v
	}

	log.Default().Println("Failed to parse num as int64", err)
	return defaultValue
}

func (c *configImpl) GetInt32(key string, defaultValue int32) (int32, error) {
	s := c.String(key)
	if s == "" {
		return defaultValue, nil
	}

	result, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}

	return int32(result), nil
}

func (c *configImpl) GetInt64(key string, defaultValue int64) (int64, error) {
	s := c.String(key)
	if s == "" {
		return defaultValue, nil
	}

	result, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}

	return int64(result), nil
}

func (c *configImpl) normalize(key string) string {
	return c.prefix + key
}

var loadedConfigSources []ConfigSource

func getConfigSourcesInternal() []ConfigSource {
	getConfigInteral() //ensure it is loaded

	var local []ConfigSource
	return append(local, loadedConfigSources...)
}

func getConfigInteral() Config {
	//should look for a metaConfig.etc to specify things like useEnv, various files, etc
	cfgOnce.Do(func() {
		configFile := *flag.String("meta-config", "", "Path to configuration file")
		flag.Parse()
		if configFile == "" {
			configFile = getConfigFile("./meta-config.yaml")
			if configFile == "" {
				configFile = getConfigFile("./meta-config.yml")
				if configFile == "" {
					configFile = getConfigFile("./meta-config.json")
					if configFile == "" {
						flag.Usage()
						os.Exit(2)
					}
				}
			}
		}

		root_builder := &configBuilderImpl{}

		cfg, err := root_builder.UseFile("meta-config", configFile).Build()
		if err != nil {
			panic(err)
		}

		if env_settings := cfg.GetString("env-settings", ""); env_settings != "" {
			env, err := builder().UseFile("env", env_settings).Build()
			if err != nil {
				panic(err)
			}

			for k, v := range env.ToStringMap() {
				os.Setenv(k, os.ExpandEnv(v))
			}
		}

		b := builder()
		for k, v := range cfg.ToStringMap() {
			values := strings.Split(v, ",")
			if len(values) == 1 {
				b = b.UseFile(k, v)
			} else if len(values) == 2 {
				b = b.UseFileSelection(k, strings.TrimSpace(values[0]), strings.TrimSpace(values[1]))
			} else {
				panic("invalid config value: " + v)
			}
		}

		cfg, err = b.Build()
		if err != nil {
			panic(err)
		}

		loadedConfigSources = root_builder.sources
		defaultConfig = cfg
	})

	return defaultConfig
}

func getConfigFile(path string) string {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return ""
	}

	return path
}