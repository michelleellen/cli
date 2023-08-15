package hermes_test

import (
	"context"
	"testing"

	"github.com/ignite/cli/ignite/pkg/hermes"
)

// rm -rf ~/.hermes ~/.ignite ~/.venus ~/.mars
// mkdir -p /Users/danilopantani/.hermes/
// cp /Users/danilopantani/Desktop/go/src/github.com/ignite/cli/ignite/pkg/hermes/config.toml /Users/danilopantani/.hermes/config.toml
// echo "middle useful federal remove kind found state weekend hammer disagree summer nephew caution boost expect" > mnemonic.txt
// hermes keys add --chain mars-1 --mnemonic-file mnemonic.txt
// hermes keys add --chain venus-1 --mnemonic-file mnemonic.txt
// hermes create client --host-chain mars-1 --reference-chain venus-1
// hermes create client --host-chain venus-1 --reference-chain mars-1
// hermes create connection --a-chain mars-1 --a-client 07-tendermint-0 --b-client 07-tendermint-0
// hermes create channel --a-chain mars-1 --a-connection connection-0 --a-port dex --b-port dex
// hermes query channels --show-counterparty --chain mars-1
// hermes start
// hermes tx ft-transfer --timeout-seconds 1000 --dst-chain venus-1 --src-chain mars-1 --src-port transfer --src-channel channel-0 --amount 100000
// hermes tx ft-transfer --timeout-seconds 10000 --denom ibc/C1840BD16FCFA8F421DAA0DAAB08B9C323FC7685D0D7951DC37B3F9ECB08A199 --dst-chain mars-1 --src-chain venus-1 --src-port transfer --src-channel channel-0 --amount 100000

// ignite relayer configure -a
// --source-rpc "http://0.0.0.0:26657"
// --source-faucet "http://0.0.0.0:4500"
// --source-port "dex"
// --source-version "dex-1"
// --source-gasprice "0.0000025stake"
// --source-prefix "cosmos"
// --source-gaslimit 300000
// --target-rpc "http://0.0.0.0:26659"
// --target-faucet "http://0.0.0.0:4501"
// --target-port "dex"
// --target-version "dex-1"
// --target-gasprice "0.0000025stake"
// --target-prefix "cosmos"
// --target-gaslimit 300000

func TestHermes(t *testing.T) {
	ctx := context.Background()
	h, err := hermes.New()
	if err != nil {
		panic(err)
	}
	defer h.Cleanup()

	err = h.Create(ctx, hermes.CommandClient, hermes.WithFlags(
		hermes.Flags{
			hermes.FlagHostChain:      "ibc-1",
			hermes.FlagReferenceChain: "ibc-0",
		}),
	)
	if err != nil {
		panic(err)
	}

	err = h.Create(ctx, hermes.CommandClient, hermes.WithFlags(
		hermes.Flags{
			hermes.FlagHostChain:      "ibc-0",
			hermes.FlagReferenceChain: "ibc-1",
		}),
	)
	if err != nil {
		panic(err)
	}

	err = h.Create(ctx, hermes.CommandConnection, hermes.WithFlags(
		hermes.Flags{
			hermes.FlagChainA:  "ibc-0",
			hermes.FlagClientA: "07-tendermint-0",
			hermes.FlagClientB: "07-tendermint-0",
		}),
	)
	if err != nil {
		panic(err)
	}

	err = h.Create(ctx, hermes.CommandChannel, hermes.WithFlags(
		hermes.Flags{
			hermes.FlagChainA:      "ibc-0",
			hermes.FlagConnectionA: "connection-0",
			hermes.FlagPortA:       "transfer",
			hermes.FlagPortB:       "transfer",
		}),
	)
	if err != nil {
		panic(err)
	}

	err = h.Query(ctx, hermes.CommandChannels, hermes.WithFlags(
		hermes.Flags{
			hermes.FlagShowCounterparty: true,
			hermes.FlagChain:            "ibc-0",
		}),
	)
	if err != nil {
		panic(err)
	}

	err = h.Start(ctx)
	if err != nil {
		panic(err)
	}
}
