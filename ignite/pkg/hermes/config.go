package hermes

import (
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
		AccountPrefix  string         `json:"account_prefix"`
		ClockDrift     string         `json:"clock_drift"`
		EventSource    EventSource    `json:"event_source"`
		GasPrice       GasPrice       `json:"gas_price"`
		GrpcAddr       string         `json:"grpc_addr"`
		Id             string         `json:"id"`
		KeyName        string         `json:"key_name"`
		MaxGas         int            `json:"max_gas"`
		RpcAddr        string         `json:"rpc_addr"`
		RpcTimeout     string         `json:"rpc_timeout"`
		StorePrefix    string         `json:"store_prefix"`
		TrustThreshold TrustThreshold `json:"trust_threshold"`
		TrustingPeriod string         `json:"trusting_period"`
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
)

func Parse(path string) (cfg Config, err error) {
	err = toml.Unmarshal([]byte(path), &cfg)
	return
}

func DefaultConfig() Config {
	return Config{
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
