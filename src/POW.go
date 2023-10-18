package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

// Block represents a block in the blockchain.
type Block struct {
	Index      int
	Data       string
	Hash       string
	PrevHash   string
	Difficulty int
}

// ProofOfWork represents the Proof of Work mechanism.
type ProofOfWork struct{}

// Mine attempts to find a valid hash for a given block and difficulty level.
func (pow ProofOfWork) Mine(block *Block) {
	targetPrefix := strings.Repeat("0", block.Difficulty)

	for nonce := 0; ; nonce++ {
		data := strconv.Itoa(block.Index) + block.Data + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(data))
		block.Hash = fmt.Sprintf("%x", hash)

		if strings.HasPrefix(block.Hash, targetPrefix) {
			fmt.Printf("Block mined! Nonce: %d, Hash: %s\n", nonce, block.Hash)
			break
		}
	}
}

func main() {
	// Create a new block with some data and set the difficulty level.
	block := Block{
		Index:      1,
		Data:       "Hello, Blockchain!",
		PrevHash:   "0000000000000000",
		Difficulty: 4, // Adjust the difficulty level to control mining speed.
	}

	pow := ProofOfWork{}
	pow.Mine(&block)

	// Print the mined block.
	fmt.Printf("Mined Block: %+v\n", block)
}

def calculate_new_capacity(capacity):
    new_cap = capacity[0]
    new_capacity_array = [new_cap]

    for i in range(1, len(capacity)):
        new_cap += capacity[i] / new_cap
        new_capacity_array.append(new_cap)

    return new_capacity_array

# Example usage:
capacity = [10, 5, 20, 15]
result = calculate_new_capacity(capacity)
print(result)