package utils

type ProgressionPrinter struct {
	spinner  []string
	isActive bool
}

// []string{"| ", "/ ", "- ", "\\ "}
//

func NewSpinnerPrinter(seq []string) *ProgressionPrinter {
	return &ProgressionPrinter{spinner: seq}
}

func (p *ProgressionPrinter) Start(updateRate uint32, msg string) {
	p.isActive = true

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

				_, _ = VaPrint("\r", p.spinner[idx], msg)

				nextSpinTick = tickCounter + uint64(updateRate)
			}

			tickCounter++
		}
	}()
}

func (p *ProgressionPrinter) Stop() {
	if !p.isActive {
		return
	}

	p.isActive = false
	VaPrintf("\r")
	return
}
