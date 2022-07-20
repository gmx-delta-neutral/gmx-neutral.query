package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SnowtraceApiKey string `envconfig:"SNOWTRACE_API_KEY"`
	Addresses       struct {
		Account      string `yaml:"account", envconfig:"Account"`
		RewardRouter string `yaml:"rewardRouter", envconfig:"REWARD_ROUTER_ADDRESS"`
		Glp          string `yaml:"glp", envconfig:"GLP_ADDRESS"`
	} `yaml:"addresses"`
	Hashes struct {
		AddLiquidity string `yaml:"addLiquidity", envconfig:"ADD_LIQUIDITY_HASH"`
	} `yaml:"hashes"`
	Decimals struct {
		Glp  int64 `yaml:"glp"`
		Usdc int64 `yaml:"usdc"`
	} `yaml:"decimals"`
}

func NewConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	f, err := os.Open("../../internal/config/config.yml")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("", &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}
