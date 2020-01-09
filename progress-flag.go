package debug_panel

type ProgressFlag struct {
	Name           string
	OriginalStatus Status
	CurrentStatus  Status
}

func NewProgressFlag(name string, status Status) *ProgressFlag {
	pf := &ProgressFlag{
		Name:           name,
		OriginalStatus: status,
		CurrentStatus:  status,
	}
	panel.addProgressFlag(pf)
	return pf
}

func (p *ProgressFlag) SetStatus(s Status) {
	p.CurrentStatus = s
}

func (p *ProgressFlag) Reset() {
	p.CurrentStatus = p.OriginalStatus
}

