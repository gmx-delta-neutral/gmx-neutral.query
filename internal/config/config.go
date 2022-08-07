package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SnowtraceApiKey string `envconfig:"SNOWTRACE_API_KEY"`
	ArbiscanApiKey  string `envconfig:"ARBISCAN_API_KEY"`
	Addresses       struct {
		Account         string `yaml:"account", envconfig:"AVAX_ACCOUNT_ADDRESS"`
		RewardRouter    string `yaml:"rewardRouter", envconfig:"REWARD_ROUTER_ADDRESS"`
		Glp             string `yaml:"glp", envconfig:"GLP_ADDRESS"`
		PoolCommitter   string `yaml:"poolCommitter", envconfig:"POOL_COMMITTER_ADDRESS"`
		PoolStateHelper string `yaml:"poolStateHelper", envconfig:"POOL_STATE_HELPER_ADDRESS"`
		LeveragedPool   string `yaml:"leveragedPool", envconfig:"LEVERAGED_POOL_ADDRESS"`
		ShortBtcToken   string `yaml:"shortBtcToken", envconfig:"SHORT_BTC_TOKEN_ADDRESS"`
		PoolKeeper      string `yaml:"poolKeeper", envconfig:"POOL_KEEPER_ADDRESS"`
		BalancerVault   string `yaml:"balancerVault", envconfig:"BALANCER_VAULT_ADDRESS"`
	} `yaml:"addresses"`
	Hashes struct {
		AddLiquidity string `yaml:"addLiquidity", envconfig:"ADD_LIQUIDITY_HASH"`
		CreateCommit string `yaml:"createCommit", envconfig:"CREATE_COMMIT_HASH"`
		Swap         string `yaml:"swap", envconfig:"SWAP_HASH"`
	} `yaml:"hashes"`
	Decimals struct {
		Glp            int64 `yaml:"glp"`
		Usdc           int64 `yaml:"usdc"`
		ShortBtc3X     int64 `yaml:"shortBtc3X"`
		PercentDivisor int64 `yaml:"percentDivisor"`
	} `yaml:"decimals"`
	Rpc struct {
		Avalanche string `yaml:"avalanche", envconfig:"AVALANCHE_RPC"`
		Arbitrum  string `yaml:"arbitrum", envconfig:"ARBITRUM_RPC"`
	} `yaml:"rpc"`
	Id struct {
		TracerThreeLeverageShortBtc int64 `yaml:"tracerThreeLeverageShortBtc"`
	} `yaml:"id"`
	ThirdPartyApi struct {
		SnowtraceApiUrl string `yaml:"snowtraceApiUrl", envconfig:"SNOWTRACE_API_URL"`
		ArbiscanApiUrl  string `yaml:"arbiscanApiUrl", envconfig:"ARBISCAN_API_URL"`
		TracerApiUrl    string `yaml:"tracerApiUrl", envconfig:"TRACER_API_URL"`
	} `yaml:"thirdPartyApi"`
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
