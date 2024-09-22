# Ethereum Mini Blockchain Indexer

## Overview

This CLI application fetches block information from the Ethereum blockchain and saves it to a file.

The information collected includes:
- Block number
- Block hash
- Number of transactions
- Block timestamp

The app starts fetching from a block number provided via a cmd flag (`--start`) and continues sequentially.

## Prerequisites

- Go 1.16+
- An Ethereum RPC endpoint (e.g., Infura, local Geth node). I used this `https://mainnet.infura.io/v3/8dcac4c10513450baedf76ff28e48bf0`

## Install

To install dependencies, run:

```bash
go mod tidy
```

## Building the binary

```bash
go build main.go
```

## Running the Indexer

```bash
go run main.go run --rpc=<your-rpc-url> --start=<block-number> --out=<output-file>
```

### Example

```bash
go run main.go run --rpc=https://mainnet.infura.io/v3/YOUR-PROJECT-ID --start=1 --out=blocks.log
```

This will connect to an Ethereum RPC node and start fetching block data from block 1, saving the information to `blocks.log`.

## Output Example

```
Number: 1
Hash: 0xabc...
TxCount: 10
Timestamp: 2023-07-10T10:30:00Z

Number: 2
Hash: 0xdef...
TxCount: 15
Timestamp: 2023-07-10T10:30:12Z
```
