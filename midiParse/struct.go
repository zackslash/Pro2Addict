// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package midiParse

const (
	typeByteLen = 4
	lenByteLen  = 4
	typeHeader  = "MThd"
	typeTrack   = "MTrk"
)

// MIDI is the top level midi structure
type MIDI struct {
	Header Header
	Tracks []Track
	Raw    []byte
}

// Chunk represents a single MIDI raw chunk
type Chunk struct {
	mType string
	mLen  uint32
	data  []byte
	Start int32
	End   int32
}

// Header is the MIDI header chunk
type Header struct {
	Chunk  Chunk
	Format int16
	Tracks int16
}

// Track is a MIDI track chunk
type Track struct {
	Chunk          Chunk
	Number         int64
	InstrumentName string
	TrackName      string
	Channel        int
}

