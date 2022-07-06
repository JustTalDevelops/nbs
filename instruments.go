package nbs

import (
	"github.com/df-mc/dragonfly/server/block/instrument"
	"strconv"
)

// Instrument is one of the instruments supported through the NBS format.
type Instrument int

// instrument converts an Instrument to its representation in Dragonfly.
func (i Instrument) instrument() instrument.Instrument {
	switch i {
	case InstrumentHarp:
		return instrument.Piano()
	case InstrumentBass:
		return instrument.Bass()
	case InstrumentBassDrum:
		return instrument.BassDrum()
	case InstrumentSnare:
		return instrument.Snare()
	case InstrumentHat:
		return instrument.ClicksAndSticks()
	case InstrumentGuitar:
		return instrument.Guitar()
	case InstrumentFlute:
		return instrument.Flute()
	case InstrumentBell:
		return instrument.Bell()
	case InstrumentChime:
		return instrument.Chimes()
	case InstrumentXylophone:
		return instrument.Xylophone()
	case InstrumentIronXylophone:
		return instrument.IronXylophone()
	case InstrumentCowBell:
		return instrument.CowBell()
	case InstrumentDidgeridoo:
		return instrument.Didgeridoo()
	case InstrumentBit:
		return instrument.Bit()
	case InstrumentBanjo:
		return instrument.Banjo()
	case InstrumentPling:
		return instrument.Pling()
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
