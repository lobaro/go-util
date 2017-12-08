package progress

type OnChangedHandler func(prog Progress)

type NestedProgress struct {
	Current float64 // Ignored when having nested progress
	Total   float64 // Ignored when having nested progress

	name     string
	weight   float64
	nested   []*NestedProgress
	onChange OnChangedHandler
}

type Progresser interface {
	Advance(progress float64)
	IncreaseTotal(total float64)
}

type Progress interface {
	Percent() float64
}

func NewProgress(name string) *NestedProgress {
	return &NestedProgress{name: name}
}

func (p *NestedProgress) SetOnChanged(cb OnChangedHandler) {
	p.onChange = cb
}

func (p *NestedProgress) Percent() float64 {
	if p.nested != nil && len(p.nested) != 0 {
		sumWeight := float64(0)
		sumProgress := float64(0)
		for _, prog := range p.nested {
			sumProgress += prog.Percent() * prog.weight
			sumWeight += prog.weight
		}
		return sumProgress / sumWeight
	}

	if p.Total == 0 {
		return 0
	}
	res := p.Current / p.Total
	if res > 100 {
		res = 100
	}
	return res
}

func (p *NestedProgress) Reset() {
	p.Current = 0
	if p.nested != nil {
		for _, n := range p.nested {
			n.Reset()
		}
	}
}

func (p *NestedProgress) IncreaseTotal(total float64) {
	p.Total += total
}

func (p *NestedProgress) Advance(progress float64) {
	p.Current += progress

	p.notifyOnChanged()
	//logrus.WithField("progress", fmt.Sprintf("%s: %f/%f", p.name, p.Current, p.Total)).Info("Advance Progress")
}

func (p *NestedProgress) notifyOnChanged() {
	if p.onChange != nil {
		p.onChange(p)
	}
}

func (p *NestedProgress) onNestedChanged(prog Progress) {
	p.notifyOnChanged()
}

// CreateNested progress with a certain weight
func (p *NestedProgress) CreateNested(name string, weight float64) *NestedProgress {
	if p.nested == nil {
		p.nested = make([]*NestedProgress, 0)
	}

	prog := &NestedProgress{
		name:     name,
		weight:   weight,
		onChange: p.onNestedChanged,
	}

	p.nested = append(p.nested, prog)
	return prog
}
