package main

import (
	em "github.com/bh90210/elektronmodels"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

func main() {
	// start a new project
	p, err := em.NewProject(em.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	p.Pattern(INTRO).
		Scale(15, 1.0, 0).
		Tempo(125.0)

	// kick
	kick := []int{0, 4, 8, 12}
	for _, v := range kick {
		p.Pattern(INTRO).Track(em.T1).Trig(v)
	}

	// // snare
	snare := []int{4, 12}
	for _, v := range snare {
		p.Pattern(INTRO).Track(em.T2).Trig(v)
	}

	// // hi-hat
	for i := 0; i <= 15; i++ {
		p.Pattern(INTRO).Track(em.T3).Trig(i)
	}

	// tom
	tom := []int{2, 3, 6, 7, 10, 11, 14, 15}
	for _, v := range tom {
		p.Pattern(INTRO).Track(em.T4).Trig(v)
	}

	// tone
	tone := []int{5, 9, 10}
	for _, v := range tone {
		p.Pattern(INTRO).Track(em.T5).Trig(v)
	}

	// synth
	synth := []int{3, 7, 9}
	for _, v := range synth {
		p.Pattern(INTRO).Track(em.T6).Trig(v)
	}

	// preset := make(map[em.Parameter]int8)
	// preset := em.PT1()
	// preset[em.COLOR] = 100
	// p.Pattern(INTRO).Track(em.T1).Preset(preset)
	// p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset)

	// p.Chain(INTRO, INTRO, VERSE, CHORUS).Play()
	// p.Next()

	p.Play(INTRO)
	// p.Next(INTRO)

	// p.Stop()

	// free
	// var i int = 100
	// p.Free.CC(em.T1, em.MACHINE, 1)
	// p.Free.CC(em.T1, em.GATE, 0)
	// p.Free.CC(em.T1, em.COLOR, 100)
	// p.Free.CC(em.T1, em.CONTOUR, 0)
	// p.Free.CC(em.T1, em.SWEEP, 100)
	// p.Free.CC(em.T1, em.REVERB, 0)
	// p.Free.Note(em.T1, em.C4, 120, 200)
	// p.Free.CC(em.T1, em.MACHINE, 2)
	// time.Sleep(time.Duration(i) * time.Millisecond)
	// p.Free.CC(em.T1, em.GATE, 1)
	// p.Free.CC(em.T1, em.COLOR, 120)
	// p.Free.CC(em.T1, em.CONTOUR, 0)
	// p.Free.CC(em.T1, em.SWEEP, 120)
	// p.Free.CC(em.T1, em.REVERB, 120)
	// p.Free.CC(em.T1, em.REVERBSIZE, 120)
	// p.Free.Note(em.T1, em.A4, 120, 500)
	// p.Free.CC(em.T1, em.MACHINE, 6)
	// time.Sleep(time.Duration(i*4) * time.Millisecond)
	// p.Free.CC(em.T1, em.GATE, 0)
	// p.Free.CC(em.T1, em.COLOR, 100)
	// p.Free.CC(em.T1, em.CONTOUR, 0)
	// p.Free.CC(em.T1, em.SWEEP, 100)
	// p.Free.CC(em.T1, em.REVERB, 0)
	// p.Free.CC(em.T1, em.REVERBSIZE, 10)
	// p.Free.Note(em.T1, em.D6, 120, 1500)
	// go func() {
	// 	var color, contour, sweep em.Parameter
	// 	color = 100
	// 	sweep = 100
	// 	contour = 0
	// 	for {
	// 		color -= 1
	// 		sweep -= 1
	// 		contour += 1
	// 		time.Sleep(10 * time.Millisecond)
	// 		p.Free.CC(em.T1, em.COLOR, int8(color))
	// 		p.Free.CC(em.T1, em.CONTOUR, int8(contour))
	// 		p.Free.CC(em.T1, em.SWEEP, int8(sweep))
	// 		p.Free.CC(em.T1, em.SHAPE, int8(sweep))
	// 	}
	// }()
	// time.Sleep(2 * time.Second)
}
