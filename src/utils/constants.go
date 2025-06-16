package utils

import (
	"strings"
)

const (
	MineRate                      = 3000
	MiningReward          float64 = 3000
	AdjustDifficultyEvery         = 2048
	MemPoolSize                   = 2048
)

var GenesisLastHash string = strings.Repeat("0", 64)
