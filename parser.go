package nbs

import (
	"bytes"
	"os"
)

// Parse parses a NBS file.
func Parse(file string) (*Song, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ParseBuf(bytes.NewBuffer(data))
}

// ParseBuf reads NBS data from a buffer.
func ParseBuf(buf *bytes.Buffer) (*Song, error) {
	layers := make(map[int16]*Layer)

	length, err := readShort(buf)
	if err != nil {
		panic(err)
	}

	var nbsVersion uint8
	if length == 0 {
		nbsVersion, err = buf.ReadByte()
		if err != nil {
			panic(err)
		}
		_, err = buf.ReadByte()
		if err != nil {
			panic(err)
		}
		if nbsVersion >= 3 {
			length, err = readShort(buf)
			if err != nil {
				panic(err)
			}
		}
	}
	songHeight, err := readShort(buf)
	if err != nil {
		panic(err)
	}
	title, err := readString(buf)
	if err != nil {
		panic(err)
	}
	author, err := readString(buf)
	if err != nil {
		panic(err)
	}
	_, err = readString(buf)
	if err != nil {
		panic(err)
	}
	description, err := readString(buf)
	if err != nil {
		panic(err)
	}
	rawSpeed, err := readShort(buf)
	if err != nil {
		panic(err)
	}
	_, err = buf.ReadByte()
	if err != nil {
		panic(err)
	}
	_, err = buf.ReadByte()
	if err != nil {
		panic(err)
	}
	_, err = buf.ReadByte()
	if err != nil {
		panic(err)
	}
	_, err = readInt(buf)
	if err != nil {
		panic(err)
	}
	_, err = readInt(buf)
	if err != nil {
		panic(err)
	}
	_, err = readInt(buf)
	if err != nil {
		panic(err)
	}
	_, err = readInt(buf)
	if err != nil {
		panic(err)
	}
	_, err = readInt(buf)
	if err != nil {
		panic(err)
	}
	_, err = readString(buf)
	if err != nil {
		panic(err)
	}
	if nbsVersion >= 4 {
		_, err = buf.ReadByte()
		if err != nil {
			panic(err)
		}
		_, err = buf.ReadByte()
		if err != nil {
			panic(err)
		}
		_, err = readShort(buf)
		if err != nil {
			panic(err)
		}
	}

	tick := int16(-1)
	for {
		jumpTicks, err := readShort(buf)
		if err != nil {
			panic(err)
		}
		if jumpTicks == 0 {
			break
		}
		tick += jumpTicks
		layer := int16(-1)
		for {
			jumpLayers, err := readShort(buf)
			if err != nil {
				panic(err)
			}
			if jumpLayers == 0 {
				break
			}
			layer += jumpLayers

			l, ok := layers[layer]
			if !ok {
				l = NewLayer()
				layers[layer] = l
			}

			instrument, err := buf.ReadByte()
			if err != nil {
				panic(err)
			}
			key, err := buf.ReadByte()
			if err != nil {
				panic(err)
			}

			// We don't support the new features, but we want to be able to parse new files still so we just ignore the new data.
			if nbsVersion >= 4 {
				_, err = buf.ReadByte()
				if err != nil {
					panic(err)
				}
				_, err = buf.ReadByte()
				if err != nil {
					panic(err)
				}
				_, err = readShort(buf)
				if err != nil {
					panic(err)
				}
			}

			l.SetNote(int64(tick), Note{
				Instrument: instrument,
				Key: key,
			})
		}
	}

	if nbsVersion > 0 && nbsVersion < 3 {
		length = tick
	}

	for i := int16(0); i < songHeight; i++ {
		l, ok := layers[i]
		if ok {
			l.Name, err = readString(buf)
			if err != nil {
				panic(err)
			}
			if nbsVersion >= 4 {
				buf.ReadByte()
			}

			l.Volume, err = buf.ReadByte()
			if err != nil {
				panic(err)
			}
			if nbsVersion >= 2 {
				buf.ReadByte()
			}
		}
	}

	speed := float32(rawSpeed) / 100

	return &Song{
		Title:       title,
		Description: description,
		Author:      author,
		Layers:      layers,
		Length:      int64(length),
		SongHeight:  songHeight,
		Speed:       speed,
	}, nil
}
