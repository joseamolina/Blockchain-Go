package main

import (
	"fmt"
	"strings"
)

type Block struct {
	Data         map[string]string
	Hash         string
	PreviousHash string
	Timestamp    string
	Pow          int
}

// New Constructor
func New(data map[string]string, hash string, previoushash string, timestamp string, pow int) Block {
	e := Block{data, hash, previoushash, timestamp, pow}

	return e
}

func (e Block) mine(difficulty int) {
	var xhash = e.Hash
	var zeroes_before = strings.Repeat("0", difficulty)
	for !strings.HasPrefix(xhash, zeroes_before) {
		e.Pow = e.Pow + 1
	}

	fmt.Printf("Hola chavale %s\n", "Como estáis?", zeroes_before, xhash)
}
