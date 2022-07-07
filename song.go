package nbs

// Song is a decoded NBS song. A Song may be re-used safely by calling Play for every individual instance of a song
// being played.
type Song struct {
	// Title is the title of the NBS song.
	Title string
	// Description is the description of the NBS song.
	Description string
	// Author is the creator of the song.
	Author string
	// Layers contains each layer.
	Layers map[int16]*Layer
	// Length is the length of the song in ticks.
	Length int64
	// SongHeight is the amount of different layers.
	SongHeight int16
	// Speed is the speed of the song.
	Speed float32
}

// Play creates a new Player for the Song and starts it.
func (s Song) Play() *Player {
	c := make(chan Note)
	p := &Player{
		C:      c,
		layers: s.Layers,
		length: s.Length,
		speed:  s.Speed,
	}
	go p.run(c)
	return p
}
