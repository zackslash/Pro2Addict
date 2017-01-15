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
	Start int32 // Relative to parent
	End   int32 // Relative to parent
}

// Header is the MIDI header chunk
type Header struct {
	Chunk  Chunk
	Format int16
	Tracks int16
}

// Note is a MIDI note event (on channel 10)
type Note struct {
	Chunk          Chunk
}

// Track is a MIDI track chunk
type Track struct {
	Chunk          Chunk
	Number         int64
	InstrumentName string
	TrackName      string
	Channel        int
	Notes		   []Note
}
