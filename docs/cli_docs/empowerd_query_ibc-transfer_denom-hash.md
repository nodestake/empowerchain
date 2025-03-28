## empowerd query ibc-transfer denom-hash

Query the denom hash info from a given denom trace

### Synopsis

Query the denom hash info from a given denom trace

```
empowerd query ibc-transfer denom-hash [trace] [flags]
```

### Examples

```
empowerd query ibc-transfer denom-hash transfer/channel-0/uatom
```

### Options

```
      --grpc-addr string   the gRPC endpoint to use for this chain
      --grpc-insecure      allow gRPC over insecure channels, if not TLS the server must use TLS
      --height int         Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help               help for denom-hash
      --node string        \<host\>:\<port\> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string      Output format (text|json) (default "text")
```

### SEE ALSO

* [empowerd query ibc-transfer](empowerd_query_ibc-transfer.md)	 - IBC fungible token transfer query subcommands

