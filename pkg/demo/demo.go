package demo

import (
	"fmt"

	"go.uber.org/zap"
)

// Demo function handles demo CLI command
func Demo(conf string) {
	var log = zap.S()
	log.Infow("Demo command called", "config", conf)
	fmt.Println(conf)
}
