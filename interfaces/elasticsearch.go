package interfaces

import (
	"fmt"

	configs "github.com/crowdeco/bima/v2/configs"
	"github.com/fatih/color"
)

type Elasticsearch struct {
}

func (e *Elasticsearch) Run(servers []configs.Server) {
	color.New(color.FgCyan, color.Bold).Printf("✓ ")
	fmt.Println("Refill Data on Elasticsearch Storage...")

	for _, server := range servers {
		go server.RepopulateData()
	}
}

func (e *Elasticsearch) IsBackground() bool {
	return true
}

func (e *Elasticsearch) Priority() int {
	return configs.HIGEST_PRIORITY + 1
}
