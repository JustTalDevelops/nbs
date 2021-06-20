package nbs

// Note is a note in a note block song.
type Note struct {
	// Instrument is the instrument intended to be used.
	Instrument uint8
	// Key is the key the note should be played on.
	Key uint8
}
