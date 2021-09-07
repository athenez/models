package main

import (
	em "github.com/athenez/models"
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
		Scale(em.TRK, 15, 4.0, 0).
		Tempo(250.0)

	p.Pattern(INTRO).Track(em.T2).Scale(15, 3.0)

	// // kick
	kick := []int{0, 4, 8, 12}
	for _, v := range kick {
		p.Pattern(INTRO).Track(em.T2).Trig(v).Note(em.A4, 200, 120)
	}

	// preset1 := em.PT3()
	// preset1[em.GATE] = 1
	// preset1[em.PUNCH] = 1
	// preset1[em.DECAY] = 120
	// // p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset1)
	// // p.Pattern(INTRO).Track(em.T1).Trig(4).Lock(preset1)

	// p.Pattern(INTRO).Track(em.T2).Trig(8).Lock(preset1)
	// p.Pattern(INTRO).Track(em.T2).Trig(0).Nudge(250)
	// // p.Pattern(INTRO).Track(em.T1).Trig(12).Lock(preset2)

	// // snare
	// snare := []int{0, 4, 8, 12}
	// for _, v := range snare {
	// 	p.Pattern(INTRO).Track(em.T1).Trig(v)
	// }

	// // hi-hat
	// for i := 0; i <= 15; i++ {
	// 	p.Pattern(INTRO).Track(em.T3).Trig(i).Scale(float64(i))
	// 	// p.Pattern(INTRO).Track(em.T3).Trig(i)
	// }

	// // tom
	// tom := []int{2, 3, 6, 7, 10, 11, 14, 15}
	// for _, v := range tom {
	// 	p.Pattern(INTRO).Track(em.T4).Trig(v)
	// }

	// // tone
	// tone := []int{5, 9, 10}
	// for _, v := range tone {
	// 	p.Pattern(INTRO).Track(em.T5).Trig(v)
	// }

	// // synth
	// synth := []int{3, 7, 9}
	// for _, v := range synth {
	// 	p.Pattern(INTRO).Track(em.T6).Trig(v)
	// }

	// preset := make(map[em.Parameter]int8)
	// preset := em.PT1()
	// preset[em.COLOR] = 100
	// p.Pattern(INTRO).Track(em.T1).Preset(preset)
	// p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset)
	// p.Pattern(VERSE).
	// 	Scale(em.TRK, 15, 4.0, 4).
	// 	Tempo(250.0)

	p.Pattern(VERSE).Scale(em.TRK, 15, 1.0, 7).Tempo(250.0).Track(em.T2).Scale(15, 3.0)

	// kick = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, v := range kick {
		p.Pattern(VERSE).Track(em.T2).Trig(v).Note(em.A4, 200, 120)
	}

	p.Chain(INTRO, VERSE).Play()

	// go p.Play(INTRO)

	// p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset1)
	// p.Pattern(INTRO).Track(em.T1).Trig(4).Lock(preset1)

	// p.Pattern(VERSE).Track(em.T6).Trig(8).Lock(preset1)
	// p.Pattern(VERSE).Track(em.T6).Trig(0).Nudge(10)

	// time.Sleep(1 * time.Second)
	// fmt.Println("dsafdsaf")
	// p.Change(VERSE)

	// time.Sleep(1 * time.Second)
	// p.Pause()
	// time.Sleep(1 * time.Second)
	// p.Resume()
	// time.Sleep(1 * time.Second)
	// p.Stop()

	// free
	// var i int = 500
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
