// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package midiParse

import (
	"io/ioutil"
	"testing"
)

const (
	// testParamTrackCount is the known number of tracks in the test midi. Note: excludes header chunk
	testParamTrackCount = 6
)

func TestParse(t *testing.T) {
	b, err := ioutil.ReadFile("./test2.mid")
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

