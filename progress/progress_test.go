package progress

import "testing"

func TestProgress_Advance(t *testing.T) {
	p := NewProgress("dummy")

	p.Total = 100
	p.Reset()
	p.Advance(50)

	if p.Percent() != 0.5 {
		t.Errorf("Progress is %f expected 0.5", p.Percent())
	}
}

func TestProgress_Nested(t *testing.T) {
	p := NewProgress("parent")

	n1 := p.CreateNested("child1", 20)
	n1.Total = 100
	n2 := p.CreateNested("child2", 80)
	n2.Total = 100

	n1.Advance(20)
	n2.Advance(50)

	if p.Percent() != 0.44 {
		t.Errorf("Progress is %f expected 0.44", p.Percent())
	}
}
