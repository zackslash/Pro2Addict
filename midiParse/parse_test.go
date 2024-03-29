// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package midiParse

import (
	"os"
	"testing"
)

const (
	// testParamTrackCount is the known number of tracks in the test midi. Note: excludes header chunk
	testParamTrackCount = 1
)

func TestParse(t *testing.T) {
	b, err := os.ReadFile("./test.mid")
	if err != nil {
		t.Error(err)
	}

	midi, err := ParseMidi(b, false)
	if err != nil {
		t.Error(err)
	}

	if len(midi.Tracks) != testParamTrackCount {
		t.Errorf("Expected track count %d does not match result track count %d", testParamTrackCount, len(midi.Tracks))
	}
}
