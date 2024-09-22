package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type BlockInfo struct {
	Number    *big.Int
	Hash      string
	TxCount   int
	Timestamp time.Time
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "indexer",
		Short: "A mini Ethereum blockchain indexer",
		Run: func(cmd *cobra.Command, args []string) {
			runIndexer()
		},
	}

	rootCmd.Flags().String("rpc", "", "URL of the Ethereum RPC node")
	rootCmd.Flags().Int64("start", 1, "Block number to start indexing from")
	rootCmd.Flags().String("out", "blocks.log", "Output file for block data")

	viper.BindPFlag("rpc", rootCmd.Flags().Lookup("rpc"))
	viper.BindPFlag("start", rootCmd.Flags().Lookup("start"))
	viper.BindPFlag("out", rootCmd.Flags().Lookup("out"))

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}

func runIndexer() {
	fmt.Println("Press Ctrl+C to stop")
	rpcURL := viper.GetString("rpc")
	startBlock := viper.GetInt64("start")
	outFile := viper.GetString("out")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	log.Printf("Connected to the Ethereum client: %s", rpcURL)

	file, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()

	blockCh := make(chan *BlockInfo)

	defer close(blockCh)

	var wg sync.WaitGroup
	defer wg.Wait()

	// start writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for blockInfo := range blockCh {
			writeBlockToFile(file, blockInfo)
		}
	}()

	// start fetching blocks sequentially from the starting block
	for i := startBlock; ; i++ {
		blockNumber := big.NewInt(i)
		wg.Add(1)
		go func(blockNumber *big.Int) {
			defer wg.Done()
			blockInfo, err := fetchBlock(client, blockNumber)
			if err != nil {
				log.Printf("Failed to fetch block %d: %v", blockNumber.Int64(), err)
				return
			}
			log.Printf("Fetched block: %d", blockNumber)
			blockCh <- blockInfo
		}(blockNumber)
		time.Sleep(2 * time.Second) // delay to avoid overwhelming the network
	}
}

func fetchBlock(
	client *ethclient.Client,
	blockNumber *big.Int,
) (*BlockInfo, error) {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	return &BlockInfo{
		Number:    block.Number(),
		Hash:      block.Hash().Hex(),
		TxCount:   len(block.Transactions()),
		Timestamp: time.Unix(int64(block.Time()), 0),
	}, nil
}

func writeBlockToFile(file *os.File, block *BlockInfo) {
	output := fmt.Sprintf("Number: %d Hash: %s TxCount: %d Timestamp: %s\n",
		block.Number.Int64(),
		block.Hash,
		block.TxCount,
		block.Timestamp.Format(time.RFC3339)) // yyyy-MM-dd'T'HH:mm:ssXXX
	if _, err := file.WriteString(output); err != nil {
		log.Fatalf("Failed to write to output file: %v", err)
	}
}
