package nbs

import (
	"go.uber.org/atomic"
	"time"
)

// Song is a decoded NBS song.
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
	// player is the player attached to the song.
	player Player
	// played is true if the song is being played.
	played atomic.Bool
	// paused is true if the song is currently paused.
	paused atomic.Bool
}

// Play starts playing the song. It will not do anything if the song is already being played.
func (s *Song) Play() {
	if s.played.Load() || s.player == nil {
		return
	}

	s.played.Store(true)

	var lastPlayed, tick int64
	for {
		if !s.played.Load() {
			break
		}

		notReadyForNextNote := time.Now().UnixNano()/int64(time.Millisecond)-lastPlayed < int64(50*s.Delay())
		if s.paused.Load() || notReadyForNextNote {
			continue
		}

		tick++
		if tick > s.Length {
			break
		}

		// Play each note in each layer.
		for _, l := range s.Layers {
			note, ok := l.Note(tick)
			if !ok {
				continue
			}

			// Send the note to the attached player.
			s.player.Play(s, note)
		}

		// Update the last played time.
		lastPlayed = time.Now().UnixNano() / int64(time.Millisecond)

		time.Sleep(20 * time.Millisecond)
	}

	s.played.Store(false)
	s.paused.Store(false)
}

// Stop stops playing the song that is currently playing.
func (s *Song) Stop() {
	s.played.Store(false)
}

// Pause toggles the pause boolean. If the song is already paused, then it will resume the song. Otherwise, it will
// pause the song and no new notes will be played.
func (s *Song) Pause() {
	s.paused.Toggle()
}

// Player attaches a new player to the song. This doesn't play the song automatically.
func (s *Song) Player(player Player) {
	s.player = player
}

// Delay is the tick delay which is calculated using the speed.
func (s *Song) Delay() float32 {
	return 20 / s.Speed
}
