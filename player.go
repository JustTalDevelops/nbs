package nbs

import (
	"go.uber.org/atomic"
	"time"
)

// Player plays a Song. Its field C may be iterated over to obtain Notes from the song whenever a new Note should be
// played like so:
//
//   for note := range Player.C {
//     fmt.Println(note)
//   }
//
// The channel C will be closed when the song ends or after Player.Stop is called, meaning the loop will end.
type Player struct {
	// C receives a Note everytime the Song being played by the Player reaches a point where a note should be played.
	// C is closed as soon as the song ends or when Player.Stop() is called. It is therefore possible to iterate over
	// this channel.
	C <-chan Note

	paused, stopped atomic.Bool
	// layers contains each layer.
	layers map[int16]*Layer
	// length is the length of the song in ticks.
	length int64
	// speed is the speed of the song.
	speed float32
}

// Stop stops the current song from playing. Stop closes Player.C as soon as the Song reaches its next tick.
func (p *Player) Stop() {
	p.stopped.Store(true)
}

// Pause toggles the pause boolean. If the song is already paused, then it will resume the song. Otherwise, it will
// pause the song and no new notes will be played.
func (p *Player) Pause() {
	p.paused.Toggle()
}

// run starts running the Player, submitting new Notes to the channel passed.
func (p *Player) run(c chan Note) {
	tick := int64(0)
	ticker := time.NewTicker(time.Duration(float64(time.Second) / float64(p.speed)))
	defer ticker.Stop()
	defer close(c)

	for range ticker.C {
		if tick > p.length || p.stopped.Load() {
			return
		} else if p.paused.Load() {
			continue
		}

		// Play each note in each layer.
		for _, l := range p.layers {
			note, ok := l.Note(tick)
			if !ok {
				continue
			}
			c <- note
		}
		tick++
	}
}
