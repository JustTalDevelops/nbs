package nbs

// Player is an interface that music players can embed and hook onto songs to receive new notes to play as they come in.
type Player interface {
	// Play is called when a new note should be played.
	Play(song *Song, note Note)
}
