package models

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

//
// constants
//

type model string

// Model
const (
	CYCLES  model = "Model:Cycles"
	SAMPLES model = "Model:Samples"
)

type voice int8

// Voices/Tracks
const (
	T1 voice = iota
	T2
	T3
	T4
	T5
	T6
)

type notes int8

// Keys/letter notes
const (
	A0 notes = iota + 21
	As0
	B0
	C1
	Cs1
	D1
	Ds1
	E1
	F1
	Fs1
	G1
	Gs1
	A1
	As1
	B1
	C2
	Cs2
	D2
	Ds2
	E2
	F2
	Fs2
	G2
	Gs2
	A2
	As2
	B2
	C3
	Cs3
	D3
	Ds3
	E3
	F3
	Fs3
	G3
	Gs3
	A3
	As3
	B3
	C4
	Cs4
	D4
	Ds4
	E4
	F4
	Fs4
	G4
	Gs4
	A4
	As4
	B4
	C5
	Cs5
	D5
	Ds5
	E5
	F5
	Fs5
	G5
	Gs5
	A5
	As5
	B5
	C6
	Cs6
	D6
	Ds6
	E6
	F6
	Fs6
	G6
	Gs6
	A6
	As6
	B6
	C7
	Cs7
	D7
	Ds7
	E7
	F7
	Fs7
	G7
	Gs7
	A7
	As7
	B7
	C8
	Cs8
	D8
	Ds8
	E8
	F8
	Fs8
	G8
	Gs8
	A8
	As8
	B8

	Bf0 notes = As0
	Df1 notes = Cs1
	Ef1 notes = Ds1
	Gf1 notes = Fs1
	Af1 notes = Gs1
	Bf1 notes = As1
	Df2 notes = Cs2
	Ef2 notes = Ds2
	Gf2 notes = Fs2
	Af2 notes = Gs2
	Bf2 notes = As2
	Df3 notes = Cs3
	Ef3 notes = Ds3
	Gf3 notes = Fs3
	Af3 notes = Gs3
	Bf3 notes = As3
	Df4 notes = Cs4
	Ef4 notes = Ds4
	Gf4 notes = Fs4
	Af4 notes = Gs4
	Bf4 notes = As4
	Df5 notes = Cs5
	Ef5 notes = Ds5
	Gf5 notes = Fs5
	Af5 notes = Gs5
	Bf5 notes = As5
	Df6 notes = Cs6
	Ef6 notes = Ds6
	Gf6 notes = Fs6
	Af6 notes = Gs6
	Bf6 notes = As6
	Df7 notes = Cs7
	Ef7 notes = Ds7
	Gf7 notes = Fs7
	Af7 notes = Gs7
	Bf7 notes = As7
	Df8 notes = Cs8
	Ef8 notes = Ds8
	Gf8 notes = Fs8
	Af8 notes = Gs8
	Bf8 notes = As8
)

type chords int8

// Chords
const (
	Unisonx2 chords = iota
	Unisonx3
	Unisonx4
	Minor
	Major
	Sus2
	Sus4
	MinorMinor7
	MajorMinor7
	MinorMajor7
	MajorMajor7
	MinorMinor7Sus4
	Dim7
	MinorAdd9
	MajorAdd9
	Minor6
	Major6
	Minorb5
	Majorb5
	MinorMinor7b5
	MajorMinor7b5
	MajorAug5
	MinorMinor7Aug5
	MajorMinor7Aug5
	Minorb6
	MinorMinor9no5
	MajorMinor9no5
	MajorAdd9b5
	MajorMajor7b5
	MajorMinor7b9no5
	Sus4Aug5b9
	Sus4AddAug5
	MajorAddb5
	Major6Add4no5
	MajorMajor76no5
	MajorMajor9no5
	Fourths
	Fifths
)

type Parameter int8

const (
	// NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10
	SWEEP      Parameter = 18
	CONTOUR    Parameter = 19
	DELAY      Parameter = 12
	REVERB     Parameter = 13
	VOLUMEDIST Parameter = 7
	// SWING      Parameter = 15
	// CHANCE     Parameter = 14

	// model:cycles
	MACHINE     Parameter = 64
	CYCLESPITCH Parameter = 65
	DECAY       Parameter = 80
	COLOR       Parameter = 16
	SHAPE       Parameter = 17
	PUNCH       Parameter = 66
	GATE        Parameter = 67

	// model:samples
	PITCH        Parameter = 16
	SAMPLESTART  Parameter = 19
	SAMPLELENGTH Parameter = 20
	CUTOFF       Parameter = 74
	RESONANCE    Parameter = 71
	LOOP         Parameter = 17
	REVERSE      Parameter = 18
)

// Reverb & Delay settings
const (
	DELAYTIME Parameter = iota + 85
	DELAYFEEDBACK
	REVERBSIZE
	REVERBTONE
)

// LFO settings
const (
	LFOSPEED Parameter = iota + 102
	LFOMULTIPIER
	LFOFADE
	LFODEST
	LFOWAVEFORM
	LFOSTARTPHASE
	LFORESET
	LFODEPTH
)

type machine int8

// Machines
const (
	KICK machine = iota
	SNARE
	METAL
	PERC
	TONE
	CHORD
)

type scaleMode bool

const (
	PTN scaleMode = true
	TRK scaleMode = false
)

//
// data structures
//

// Project long description of the data structure, methods, behaviors and useage.
type Project struct {
	model
	*Sequencer
	// Free allows to bypass the sequencer and send triggers in real time.
	Free *Free
}

// Sequencer methods allows for Play, Pause, Resume, Stop.
type Sequencer struct {
	mu *sync.Mutex

	// midi fields
	drv midi.Driver
	in  midi.In
	out midi.Out
	wr  *writer.Writer

	// pattern map[int]*Pattern

	// playtime sequencer fields
	// pause           chan bool
	// resume          chan bool
	// stop            chan bool
	// next            chan bool
	// currentPattern  int
	// changeCountdown int

	// currentChain int
	// chains       []int
	// fillMode chan bool
	// chances chan float64
	// swing chan float64

	// scaleLock bool
	// trigLock  bool
}

// Free mathods allows to by-pass the sequencer.
type Free struct {
	midi *Sequencer
}

// Pattern holds all relevant tracks information for given pattern.
// type Pattern struct {
// 	track map[voice]*Track
// 	scale *scale
// 	tempo float64

// 	// sequencer playtime filed helper
// 	changingPattern bool
// }

// Track holds all relevant information for track.
// type Track struct {
// 	preset
// 	scale *scale
// 	trig  map[int]*Trig
// }

// type scale struct {
// 	mode   scaleMode
// 	length int
// 	scale  float64
// 	change int8
// }

type preset map[Parameter]int8

// Trig holds all relevant information for trig.
// type Trig struct {
// 	note *Note
// 	lock preset

// 	scale *scale
// 	nudge float64
// 	// condition float64
// }

// Note holds all relevant information for given trig.
// type Note struct {
// 	key      notes
// 	length   float64
// 	velocity int8
// }

//
// Project
//

// NewProject initiates and returns a *Project struct.
// TODO: better documentation
func NewProject(m model) (*Project, error) {
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	sequencer := &Sequencer{
		mu:  new(sync.Mutex),
		drv: drv,
		// pause:   make(chan bool),
		// resume:  make(chan bool),
		// stop:    make(chan bool),
		// next:    make(chan bool),
		// pattern: make(map[int]*Pattern),
	}

	// find elektron and assign it to in/out
	var helperIn, helperOut bool

	sequencer.mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), string(m)) {
			sequencer.in = in
			helperIn = true
		}
	}
	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), string(m)) {
			sequencer.out = out
			helperOut = true
		}
	}
	// check if nothing found
	if !helperIn && !helperOut {
		return nil, fmt.Errorf("device %s not found", m)
	}

	err = sequencer.in.Open()
	if err != nil {
		return nil, err
	}

	err = sequencer.out.Open()
	if err != nil {
		return nil, err
	}

	wr := writer.New(sequencer.out)
	sequencer.wr = wr
	sequencer.mu.Unlock()

	return &Project{
		model:     m,
		Sequencer: sequencer,
		Free: &Free{
			midi: sequencer,
		}}, nil
}

//
// sequencer
//

// // Play starts playing the given pattern. It is a blocking function.
// func (s *Sequencer) Play(ids ...int) {
// 	block := make(chan bool)
// 	// if user did not specify a pattern neither Chain method used, print an error
// 	if len(ids) == 0 && len(s.chains) == 0 {
// 		fmt.Println("error: no pattern selected")
// 		return
// 	}

// 	var pattern *Pattern
// 	var id int

// 	if len(s.chains) != 0 {
// 		if s.currentChain == 0 {
// 			pattern = s.pattern[s.chains[0]]
// 			id = s.chains[s.currentChain]
// 		} else {
// 			pattern = s.pattern[s.chains[s.currentChain]]
// 			id = s.chains[s.currentChain]
// 		}
// 	} else {
// 		pattern = s.pattern[ids[0]]
// 		id = ids[0]
// 	}

// 	// check pattern exists
// 	if pattern == nil {
// 		fmt.Println("error: pattern does not exist")
// 		return
// 	}

// 	s.currentPattern = id
// 	s.changeCountdown = int(pattern.scale.change)

// 	for i := 0; i <= 5; i++ {
// 		voice := voice(i)
// 		if track, ok := s.pattern[id].track[voice]; ok {
// 			// block if changing pattern is on
// 			if pattern.changingPattern {
// 				pattern.changingPattern = false
// 				<-s.next
// 			}

// 			go func() {
// 				// check if track has preset
// 				// if not set default preset for track
// 				if len(track.preset) == 0 {
// 					track.preset = defaultPreset(voice)
// 				}

// 				// apply preset
// 				for parameter, value := range track.preset {
// 					s.cc(voice, parameter, value)
// 				}

// 				// special case for lock on trig 0
// 				if trig, ok := track.trig[0]; ok {
// 					if len(trig.lock) != 0 {
// 						for k, v := range trig.lock {
// 							s.cc(voice, k, v)
// 						}
// 						s.trigLock = true
// 					}
// 				}
// 			}()

// 			var scl float64
// 			var lng int
// 			switch pattern.scale.mode {
// 			case PTN:
// 				lng = pattern.scale.length
// 				scl = pattern.scale.scale
// 			case TRK:
// 				lng = track.scale.length
// 				scl = track.scale.scale
// 			}

// 			tick := time.NewTicker(
// 				time.Duration(60000/((pattern.tempo*2)*scl)) * time.Millisecond)

// 			go func() {
// 				var count int
// 				var counter = make(chan int)
// 				var fire = make(chan bool)

// 				// lock patcher
// 				go func(fire chan bool) {
// 					for {
// 						currentCount := <-counter
// 						if trig, ok := track.trig[currentCount+1]; ok {
// 							if len(trig.lock) != 0 {
// 								<-fire
// 								for k, v := range trig.lock {
// 									s.cc(voice, k, v)
// 								}
// 								s.trigLock = true
// 								continue
// 							}
// 						}

// 						if s.trigLock {
// 							<-fire
// 							for k, v := range track.preset {
// 								s.cc(voice, k, v)
// 							}
// 							s.trigLock = false
// 						}
// 					}
// 				}(fire)
// 			loop:
// 				for {
// 					select {
// 					case <-tick.C:
// 						// if count is bigger than pattern/track length reset to zero
// 						// and start looping over the start.
// 						if count > lng {
// 							count = 0
// 						}

// 						if (count+1) > lng && len(s.chains) != 0 && !pattern.changingPattern {
// 							s.currentChain++
// 							if len(s.chains) > s.currentChain {
// 								pattern.changingPattern = true
// 								s.chainChange(s.chains[s.currentChain])
// 							} else {
// 								pattern.changingPattern = true
// 								s.currentChain = 0
// 								s.chainChange(s.chains[0])
// 							}
// 							// s.next <- true
// 							// block <- true
// 							// break loop
// 						}

// 						if s.changeCountdown == 0 && pattern.changingPattern {
// 							s.next <- true
// 							// block <- true
// 							break loop
// 						}

// 						if pattern.changingPattern {
// 							s.changeCountdown--
// 						}

// 						// send the current count to lock patcher
// 						counter <- count

// 						if trig, ok := track.trig[count]; ok {
// 							// scale check/reset
// 							switch {
// 							case trig.scale != nil:
// 								tick.Reset(time.Duration(60000/((pattern.tempo*2)*trig.scale.scale)) * time.Millisecond)
// 								s.scaleLock = true
// 								// continue

// 							case s.scaleLock:
// 								tick.Reset(time.Duration(60000/((pattern.tempo*2)*scl)) * time.Millisecond)
// 								s.scaleLock = false
// 							}

// 							if trig.nudge != 0 {
// 								time.Sleep(time.Duration(trig.nudge) * time.Millisecond)
// 							}

// 							s.noteon(voice,
// 								trig.note.key,
// 								trig.note.velocity)
// 							go func(fire chan bool) {
// 								time.Sleep(time.Millisecond * time.Duration(trig.note.length))
// 								s.noteoff(voice, trig.note.key)
// 								fire <- true
// 							}(fire)
// 						}

// 						count++

// 					case <-s.pause:
// 						<-s.resume
// 					case <-s.stop:
// 						break loop
// 					}
// 				}
// 			}()
// 		}
// 	}

// 	// block := make(chan bool)
// 	<-block
// }

// // Pause the sequencer.
// func (s *Sequencer) Pause() {
// 	s.pause <- true
// }

// // Resume the sequencer.
// func (s *Sequencer) Resume() {
// 	s.resume <- true
// }

// // Stop the sequencer.
// func (s *Sequencer) Stop() {
// 	s.stop <- true
// }

// func (s *Sequencer) Change(id int) {
// 	log.Println("current", s.currentPattern)
// 	log.Println("changing to", id)
// 	// s.pattern[s.currentPattern].changingPattern = true
// 	s.pattern[id].changingPattern = true
// 	go s.Play(id)
// }

// // Chain allows for chaining in a series loop multiple patterns at once.
// func (s *Sequencer) Chain(patterns ...int) *Sequencer {
// 	s.chains = append(s.chains, patterns...)
// 	return s
// }

// // Pattern returns the specified pattern out of project's pattern collection.
// // Allows to access pattern's methods.
// func (s *Sequencer) Pattern(id int) *Pattern {
// 	if _, ok := s.pattern[id]; !ok {
// 		s.pattern[id] = &Pattern{
// 			track: make(map[voice]*Track),
// 			scale: &scale{PTN, 15, 1.0, 15},
// 			tempo: 120 * 2,
// 		}
// 	}

// 	return s.pattern[id]
// }

// Close midi connection. Use it with defer after creating a new project.
func (s *Sequencer) Close() {
	s.in.Close()
	s.out.Close()
	s.drv.Close()
}

func (s *Sequencer) noteon(t voice, n notes, vel int8) {
	s.wr.SetChannel(uint8(t))
	writer.NoteOn(s.wr, uint8(n), uint8(vel))
}

func (s *Sequencer) noteoff(t voice, n notes) {
	s.wr.SetChannel(uint8(t))
	writer.NoteOff(s.wr, uint8(n))
}

func (s *Sequencer) cc(t voice, par Parameter, val int8) {
	s.wr.SetChannel(uint8(t))
	writer.ControlChange(s.wr, uint8(par), uint8(val))
}

func (s *Sequencer) pc(t voice, pc int8) {
	s.wr.SetChannel(uint8(t))
	writer.ProgramChange(s.wr, uint8(pc))
}

// func (s *Sequencer) chainChange(id int) {
// 	// s.pattern[id].changingPattern = true
// 	// go s.Play(id)
// 	s.Change(id)
// }

// //
// // pattern
// //

// // Scale sets the scale for the pattern.
// // If scaleMode is set to track TRK the provided scale settings are used as default to the rest of the tracks.
// // This mimics synth's own functionality.
// func (p *Pattern) Scale(mode scaleMode, length int, scale float64, change int8) *Pattern {
// 	p.scale.mode = mode
// 	p.scale.length = length
// 	p.scale.scale = scale
// 	p.scale.change = change
// 	return p
// }

// // Tempo set's pattern tempo.
// func (p *Pattern) Tempo(tempo float64) *Pattern {
// 	p.tempo = tempo * 2
// 	return p
// }

// // Track return a new track.
// func (p *Pattern) Track(id voice) *Track {
// 	if _, ok := p.track[id]; !ok {
// 		p.track[id] = &Track{
// 			preset: defaultPreset(id),
// 			trig:   make(map[int]*Trig),
// 			scale: &scale{
// 				length: 15,
// 				scale:  0.1,
// 			},
// 		}
// 	}

// 	return p.track[id]
// }

// //
// // Track
// //

// // SetScale sets a new scale for the track.
// // If not set a default one is used.
// func (t *Track) Scale(length int, factor float64) *Track {
// 	t.scale.length = length
// 	t.scale.scale = factor
// 	return t
// }

// // Preset applies preset to track.
// func (t *Track) Preset(p preset) *Track {
// 	t.preset = p
// 	return t
// }

// // SetParameter assigned a parameter to the preset.
// // First argument is a Parameter type and second value an int8.
// func (t *Track) Parameter(parameter Parameter, value int8) *Track {
// 	t.preset[parameter] = value
// 	return t
// }

// // Trig returns a new trig.
// func (t *Track) Trig(id int) *Trig {
// 	if _, ok := t.trig[id]; !ok {
// 		t.trig[id] = &Trig{note: &Note{C4, 200, 110}}
// 	}

// 	return t.trig[id]
// }

// //
// // trig
// //

// // Lock allows for setting trigger/preset locks.
// func (t *Trig) Lock(preset preset) *Trig {
// 	t.lock = preset
// 	return t
// }

// // Note sets note's key, length and velocity to trig.
// func (t *Trig) Note(key notes, length float64, velocity int8) *Trig {
// 	// t.note = &note{key, length, velocity}
// 	t.note.key = key
// 	t.note.length = length
// 	t.note.velocity = velocity
// 	return t
// }

// // Scale allows for setting individual trig scale.
// func (t *Trig) Scale(factor float64) *Trig {
// 	t.scale = &scale{scale: factor}
// 	return t
// }

// // Nudge sets a delay prior to firing the trig.
// func (t *Trig) Nudge(amount float64) *Trig {
// 	t.nudge = amount
// 	return t
// }

// //
// // note
// //

// // Key set note's key.
// func (n *Note) Key(key notes) {
// 	n.key = key
// }

// // Length Trig Length sets the duration of the notes. When a note has finished playing a NOTE OFF command
// // is sent. The INF setting equals infinite note length. This parameter only applies if GATE is set to ON or
// // when sending trig length data over MIDI. (0.125–128, INF)
// func (n *Note) Length(length float64) {
// 	n.length = length
// }

// // Velocity sets note's velocity.
// func (n *Note) Velocity(velocity int8) {
// 	n.velocity = velocity
// }

//
// free
//

// Preset immediately sets (CC) provided parameters.
func (f *Free) Preset(track voice, preset preset) {
	for parameter, value := range preset {
		f.midi.cc(track, parameter, value)
	}
}

// Note fires immediately a midi note on signal followed by a note off specified duration in milliseconds (ms).
// Optionally user can pass a preset too for convenience.
func (f *Free) Note(track voice, note notes, velocity int8, duration float64, pre ...preset) {
	if len(pre) != 0 {
		for i, _ := range pre {
			f.Preset(track, pre[i])
		}
	}

	f.midi.noteon(track, note, velocity)
	go func() {
		time.Sleep(time.Millisecond * time.Duration(duration))
		f.midi.noteoff(track, note)
	}()
}

// CC control change.
func (f *Free) CC(track voice, parameter Parameter, value int8) {
	f.midi.cc(track, parameter, value)
}

// PC program control change.
func (f *Free) PC(t voice, pc int8) {
	f.midi.pc(t, pc)
}
