#!/usr/bin/env bash

geth --datadir "../chainDB" --rpcapi net,eth,web3,personal,miner, --rpc --rpcport 8545 --nodiscover console 2>>"../logs/eth_output.log"
