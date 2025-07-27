package utils

import (
	"fmt"
	"os"
	"time"
)

type ProgressionPrinter struct {
	spinner  []string
	isActive bool
	start    time.Time
}

// []string{"| ", "/ ", "- ", "\\ "}
//

func NewSpinnerPrinter(seq []string) *ProgressionPrinter {
	return &ProgressionPrinter{spinner: seq}
}

func (p *ProgressionPrinter) Start(updateRate uint32, msg string) {
	p.isActive = true
	p.start = time.Now()

	go func() {
		tickCounter := uint64(0)
		nextSpinTick := uint64(0)

		var idx = 0
		for p.isActive {
			if nextSpinTick <= tickCounter {
				idx++
				if idx >= len(p.spinner) {
					idx = 0
				}

				_, _ = fmt.Print("\r", p.spinner[idx], VaSprintf("%s (%s)", msg, time.Since(p.start)))

				nextSpinTick = tickCounter + uint64(updateRate)
			}

			tickCounter++
		}
	}()
}

func (p *ProgressionPrinter) Stop() time.Duration {
	if !p.isActive {
		return 0
	}

	p.isActive = false
	_, _ = fmt.Fprint(os.Stdout, "\r")
	return time.Since(p.start)
}
