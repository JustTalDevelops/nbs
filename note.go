package nbs

import (
	"github.com/df-mc/dragonfly/server/world/sound"
)

// Note is a note in a note block song.
type Note struct {
	// Instrument is the instrument intended to be used.
	Instrument Instrument
	// Key is the key the note should be played on.
	Key int
}

// Sound converts the Note to a sound.Note so that it may be played as a soon in Dragonfly.
func (n Note) Sound() sound.Note {
	return sound.Note{
		Instrument: n.Instrument.instrument(),
		Pitch:      n.Key - 33,
	}
}
