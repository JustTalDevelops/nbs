package nbs

import (
	"github.com/df-mc/dragonfly/server/world/sound"
	"strconv"
)

// Instrument is one of the instruments supported through the NBS format.
type Instrument int

// instrument converts an Instrument to its representation in Dragonfly.
func (i Instrument) instrument() sound.Instrument {
	switch i {
	case InstrumentHarp:
		return sound.Piano()
	case InstrumentBass:
		return sound.Bass()
	case InstrumentBassDrum:
		return sound.BassDrum()
	case InstrumentSnare:
		return sound.Snare()
	case InstrumentHat:
		return sound.ClicksAndSticks()
	case InstrumentGuitar:
		return sound.Guitar()
	case InstrumentFlute:
		return sound.Flute()
	case InstrumentBell:
		return sound.Bell()
	case InstrumentChime:
		return sound.Chimes()
	case InstrumentXylophone:
		return sound.Xylophone()
	case InstrumentIronXylophone:
		return sound.IronXylophone()
	case InstrumentCowBell:
		return sound.CowBell()
	case InstrumentDidgeridoo:
		return sound.Didgeridoo()
	case InstrumentBit:
		return sound.Bit()
	case InstrumentBanjo:
		return sound.Banjo()
	case InstrumentPling:
		return sound.Pling()
	}
	panic("unsupported instrument type " + strconv.Itoa(int(i)))
}

// The following are all supported NBS instruments.
const (
	InstrumentHarp Instrument = iota
	InstrumentBass
	InstrumentBassDrum
	InstrumentSnare
	InstrumentHat
	InstrumentGuitar
	InstrumentFlute
	InstrumentBell
	InstrumentChime
	InstrumentXylophone
	InstrumentIronXylophone
	InstrumentCowBell
	InstrumentDidgeridoo
	InstrumentBit
	InstrumentBanjo
	InstrumentPling
)
