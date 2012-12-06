package nimble

import (
	"code.google.com/p/nimble-cube/core"
)

// ChanN is a Chan that passes N-component data.
type ChanN struct {
	comp []Chan1
}

// TODO: re-use same *Info for all chan1's, tag for components?
func MakeChanN(nComp int, tag, unit string, m *Mesh, memType MemType, bufBlocks int) ChanN {
	c := make([]Chan1, nComp)
	for i := range c {
		c[i] = MakeChan1(tag, unit, m, memType, bufBlocks)
	}
	return ChanN{c}
}

func AsChan(buffer []Slice, tag, unit string, m *Mesh) ChanN {
	nComp := len(buffer)
	c := make([]Chan1, nComp)
	for i := range c {
		c[i] = asChan1(buffer[i], tag, unit, m)
	}
	return ChanN{c}
}

func (c ChanN) Mesh() *Mesh          { return c.comp[0].Mesh }
func (c ChanN) Size() [3]int         { return c.comp[0].Size() }
func (c ChanN) Unit() string         { return c.comp[0].Unit() }
func (c ChanN) Tag() string          { return c.comp[0].Tag() }
func (c ChanN) NComp() int           { return len(c.comp) }
func (c ChanN) Comp(i int) Chan1     { return c.comp[i] }
func (c ChanN) BufLen() int          { return c.comp[0].BufLen() }
func (c ChanN) NBufferedBlocks() int { return c.comp[0].NBufferedBlocks() }
func (c ChanN) MemType() MemType     { return c.comp[0].slice.MemType }
func (c ChanN) ChanN() ChanN         { return c } // implements Chan iface

func (c ChanN) UnsafeData() []Slice {
	s := make([]Slice, c.NComp())
	for i := range s {
		s[i] = c.comp[i].slice
	}
	return s
}

//func (c ChanN) Chan3() Chan3 {
//	core.Assert(c.NComp() == 3)
//	return Chan3{c.comp}
//}

func (c ChanN) Chan1() Chan1 {
	core.Assert(c.NComp() == 1)
	return c.comp[0]
}

// WriteNext locks and returns a slice of length n for
// writing the next n elements to the Chan3.
// When done, WriteDone() should be called to "send" the
// slice down the Chan3. After that, the slice is not valid any more.
func (c ChanN) WriteNext(n int) []Slice {
	next := make([]Slice, c.NComp())
	for i := range c.comp {
		c.comp[i].WriteNext(n)
		a, b := c.comp[i].mutex.WRange()
		next[i] = c.comp[i].slice.Slice(a, b)
	}
	return next
}

// WriteDone() signals a slice obtained by WriteNext() is fully
// written and can be sent down the Chan3.
func (c ChanN) WriteDone() {
	for i := range c.comp {
		c.comp[i].WriteDone()
	}
}

//func (c ChanN) WriteDelta(Δstart, Δstop int) [][]float32 {
//	next := make([][]float32, c.NComp())
//	for i := range c {
//		c[i].WriteDelta(Δstart, Δstop)
//		a, b := c[i].mutex.WRange()
//		next[i] = c[i].slice.Slice(a, b).list
//	}
//	return next
//}
