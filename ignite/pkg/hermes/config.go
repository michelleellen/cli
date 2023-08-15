package hermes

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type (
	Config struct {
		Chains    []Chain   `json:"chains"`
		Global    Global    `json:"global"`
		Telemetry Telemetry `json:"telemetry"`
		Mode      Mode      `json:"mode"`
	}

	Chain struct {
		Id             string         `json:"id"`
		RpcAddr        string         `json:"rpc_addr"`
		GrpcAddr       string         `json:"grpc_addr"`
		EventSource    EventSource    `json:"event_source"`
		RpcTimeout     string         `json:"rpc_timeout"`
		AccountPrefix  string         `json:"account_prefix"`
		KeyName        string         `json:"key_name"`
		StorePrefix    string         `json:"store_prefix"`
		DefaultGas     int            `json:"default_gas"`
		MaxGas         int            `json:"max_gas"`
		GasPrice       GasPrice       `json:"gas_price"`
		GasMultiplier  float64        `json:"gas_multiplier"`
		MaxMsgNum      int            `json:"max_msg_num"`
		MaxTxSize      int            `json:"max_tx_size"`
		ClockDrift     string         `json:"clock_drift"`
		MaxBlockTime   string         `json:"max_block_time"`
		TrustingPeriod string         `json:"trusting_period"`
		TrustThreshold TrustThreshold `json:"trust_threshold"`
		AddressType    AddressType    `json:"address_type"`
	}

	EventSource struct {
		BatchDelay string `json:"batch_delay"`
		Mode       string `json:"mode"`
		Url        string `json:"url"`
	}

	GasPrice struct {
		Denom string  `json:"denom"`
		Price float64 `json:"price"`
	}

	TrustThreshold struct {
		Denominator string `json:"denominator"`
		Numerator   string `json:"numerator"`
	}

	AddressType struct {
		Derivation string `json:"derivation"`
	}

	Global struct {
		LogLevel string `json:"log_level"`
	}

	Telemetry struct {
		Enabled bool   `json:"enabled"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
	}

	Mode struct {
		Channels    Channels    `json:"channels"`
		Clients     Clients     `json:"clients"`
		Connections Connections `json:"connections"`
		Packets     Packets     `json:"packets"`
	}

	Channels struct {
		Enabled bool `json:"enabled"`
	}

	Clients struct {
		Enabled      bool `json:"enabled"`
		Misbehaviour bool `json:"misbehaviour"`
		Refresh      bool `json:"refresh"`
	}

	Connections struct {
		Enabled bool `json:"enabled"`
	}

	Packets struct {
		ClearInterval  int  `json:"clear_interval"`
		ClearOnStart   bool `json:"clear_on_start"`
		Enabled        bool `json:"enabled"`
		TxConfirmation bool `json:"tx_confirmation"`
	}

	// ChainOption configures chain hermes configs.
	ChainOption func(*Chain)
)

func WithEventSource(mode, url, batchDelay string) ChainOption {
	return func(c *Chain) {
		c.EventSource = EventSource{
			BatchDelay: batchDelay,
			Mode:       mode,
			Url:        url,
		}
	}
}

func WithRPCTimeout(timeout string) ChainOption {
	return func(c *Chain) {
		c.RpcTimeout = timeout
	}
}

func WithAccountPrefix(prefix string) ChainOption {
	return func(c *Chain) {
		c.AccountPrefix = prefix
	}
}

func WithKeyName(key string) ChainOption {
	return func(c *Chain) {
		c.KeyName = key
	}
}

func WithStorePrefix(prefix string) ChainOption {
	return func(c *Chain) {
		c.StorePrefix = prefix
	}
}

func WithDefaultGas(defaultGas int) ChainOption {
	return func(c *Chain) {
		c.DefaultGas = defaultGas
	}
}

func WithMaxGas(maxGas int) ChainOption {
	return func(c *Chain) {
		c.MaxGas = maxGas
	}
}

func WithGasPrice(price float64, denom string) ChainOption {
	return func(c *Chain) {
		c.GasPrice = GasPrice{
			Denom: denom,
			Price: price,
		}
	}
}

func WithGasMultiplier(gasMultipler float64) ChainOption {
	return func(c *Chain) {
		c.GasMultiplier = gasMultipler
	}
}

func WithMaxMsgNum(maxMsg int) ChainOption {
	return func(c *Chain) {
		c.MaxMsgNum = maxMsg
	}
}

func WithMaxTxSize(size int) ChainOption {
	return func(c *Chain) {
		c.MaxTxSize = size
	}
}

func WithClockDrift(clock string) ChainOption {
	return func(c *Chain) {
		c.ClockDrift = clock
	}
}

func WithMaxBlockTime(maxBlockTime string) ChainOption {
	return func(c *Chain) {
		c.MaxBlockTime = maxBlockTime
	}
}

func WithTrustingPeriod(trustingPeriod string) ChainOption {
	return func(c *Chain) {
		c.TrustingPeriod = trustingPeriod
	}
}

func WithTrustThreshold(numerator, denominator string) ChainOption {
	return func(c *Chain) {
		c.TrustThreshold = TrustThreshold{
			Denominator: denominator,
			Numerator:   numerator,
		}
	}
}

func WithAddressType(derivation string) ChainOption {
	return func(c *Chain) {
		c.AddressType = AddressType{Derivation: derivation}
	}
}

func (c *Config) AddChain(chainID, rpcAddr, grpcAddr string, options ...ChainOption) error {
	rpcUrl, err := url.Parse(rpcAddr)
	if err != nil {
		return err
	}

	chain := Chain{
		Id:       chainID,
		RpcAddr:  rpcAddr,
		GrpcAddr: grpcAddr,
		EventSource: EventSource{
			BatchDelay: "500ms",
			Mode:       "push",
			Url:        fmt.Sprintf("ws://%s:%s", rpcUrl.Host, rpcUrl.Port()),
		},
		RpcTimeout:    "15s",
		AccountPrefix: "cosmos",
		KeyName:       "wallet",
		StorePrefix:   "ibc",
		DefaultGas:    100000,
		MaxGas:        10000000,
		GasPrice: GasPrice{
			Denom: "stake",
			Price: 0.01,
		},
		GasMultiplier:  1.1,
		MaxMsgNum:      30,
		MaxTxSize:      2097152,
		ClockDrift:     "5s",
		MaxBlockTime:   "10s",
		TrustingPeriod: "14days",
		TrustThreshold: TrustThreshold{
			Denominator: "3",
			Numerator:   "1",
		},
		AddressType: AddressType{
			Derivation: "cosmos",
		},
	}
	for _, o := range options {
		o(&chain)
	}

	c.Chains = append(c.Chains, chain)
	return nil
}

func (c *Config) Save() error {
	configPath, err := ConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	return toml.NewEncoder(file).Encode(c)
}

func Parse(path string) (cfg Config, err error) {
	err = toml.Unmarshal([]byte(path), &cfg)
	return
}

func DefaultConfig() *Config {
	return &Config{
		Chains: []Chain{},
		Global: Global{
			LogLevel: "info",
		},
		Mode: Mode{
			Channels: Channels{
				Enabled: true,
			},
			Clients: Clients{
				Enabled:      true,
				Misbehaviour: true,
				Refresh:      true,
			},
			Connections: Connections{
				Enabled: true,
			},
			Packets: Packets{
				ClearInterval:  100,
				ClearOnStart:   true,
				Enabled:        true,
				TxConfirmation: true,
			},
		},
		Telemetry: Telemetry{
			Enabled: true,
			Host:    "127.0.0.1",
			Port:    3001,
		},
	}
}

func ConfigPath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHomeDir, ".hermes", "config.toml"), nil
}
