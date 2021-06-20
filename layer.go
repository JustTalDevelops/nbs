package nbs

// Layer is a layer in the NBS data.
type Layer struct {
	// Volume is the volume of the layer.
	Volume uint8
	// Name is the name of the layer.
	Name string
	// notes is a map from tick to note.
	notes map[int64]Note
}

// Note gets a note at a tick.
func (l *Layer) Note(tick int64) (note Note, ok bool) {
	note, ok = l.notes[tick]
	return
}

// SetNote sets a note at a tick.
func (l *Layer) SetNote(tick int64, note Note) {
	l.notes[tick] = note
}

// NewLayer initializes a new layer.
func NewLayer() *Layer {
	return &Layer{
		Volume: 100,
		notes:  make(map[int64]Note),
	}
}
