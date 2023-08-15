#!/bin/bash

rm -rf ~/.hermes ~/.ignite ~/.venus ~/.mars

mkdir -p /Users/danilopantani/.hermes/
cp /Users/danilopantani/Desktop/go/src/github.com/ignite/cli/ignite/pkg/hermes/config.toml /Users/danilopantani/.hermes/config.toml

echo "letter column benefit acoustic evidence false trim cave jump pluck awesome lion" > mnemonic1.txt
echo "jeans payment lock client result enemy bullet rug crush deny month salad" > mnemonic2.txt
hermes keys add --chain mars-1 --mnemonic-file mnemonic1.txt
hermes keys add --chain venus-1 --mnemonic-file mnemonic2.txt

hermes create client --host-chain mars-1 --reference-chain venus-1
hermes create client --host-chain venus-1 --reference-chain mars-1

hermes create connection --a-chain mars-1 --a-client 07-tendermint-0 --b-client 07-tendermint-0

hermes create channel --a-chain mars-1 --a-connection connection-0 --a-port transfer --b-port transfer

hermes query channels --show-counterparty --chain mars-1

hermes start

hermes tx ft-transfer --timeout-seconds 1000 --dst-chain venus-1 --src-chain mars-1 --src-port transfer --src-channel channel-0 --amount 100000
hermes tx ft-transfer --timeout-seconds 10000 --denom ibc/C1840BD16FCFA8F421DAA0DAAB08B9C323FC7685D0D7951DC37B3F9ECB08A199 --dst-chain mars-1 --src-chain venus-1 --src-port transfer --src-channel channel-0 --amount 100000
