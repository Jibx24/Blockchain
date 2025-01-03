package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Block structure
type Block struct {
	Index        int       `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	Data         string    `json:"data"`
	Hash         string    `json:"hash"`
}

// Blockchain structure
type Blockchain struct {
	Chain []Block
	mu    sync.Mutex
}

// Create a new block
func (bc *Blockchain) CreateBlock(data string) Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:        len(bc.Chain),
		Timestamp:    time.Now(),
		PreviousHash: previousBlock.Hash,
		Data:         data,
		Hash:         "",
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

// Initialize Blockchain
func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now(),
		PreviousHash: "0",
		Data:         "Genesis Block",
		Hash:         "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return &Blockchain{
		Chain: []Block{genesisBlock},
	}
}

// Calculate Hash
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + block.PreviousHash + block.Data
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// API Handlers
func main() {
	r := gin.Default()
	blockchain := NewBlockchain()

	// Add appointment endpoint
	r.POST("/appointment", func(c *gin.Context) {
		var appointment struct {
			PatientID string `json:"patient_id"`
			DoctorID  string `json:"doctor_id"`
			Date      string `json:"date"`
			Time      string `json:"time"`
		}
		if err := c.ShouldBindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, _ := json.Marshal(appointment)
		newBlock := blockchain.CreateBlock(string(data))
		c.JSON(http.StatusCreated, newBlock)
	})

	// Get all blocks
	r.GET("/blocks", func(c *gin.Context) {
		c.JSON(http.StatusOK, blockchain.Chain)
	})

	r.Run(":8080") // Start server
}
