package utils

type ProgressionPrinter struct {
	spinner  []string
	rate     uint32
	spinRate uint32
	isActive bool
	message  string
}

// []string{"| ", "/ ", "- ", "\\ "}
//

func NewSpinnerPrinter(seq []string) *ProgressionPrinter {
	return &ProgressionPrinter{spinner: seq}
}

func (p *ProgressionPrinter) Start(updateRate uint32) error {
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

				if p.message != "" {
					_, _ = VaPrint("\r", p.spinner[idx], p.message)
				}

				nextSpinTick = tickCounter + uint64(updateRate)
			}

			tickCounter++
		}
	}()

	return nil
}

func (p *ProgressionPrinter) Write(text string) {
	if p.isActive {
		p.message = text
	}
}

func (p *ProgressionPrinter) Stop() {
	if !p.isActive {
		return
	}

	p.isActive = false
	VaPrintf("\r")
	return
}
