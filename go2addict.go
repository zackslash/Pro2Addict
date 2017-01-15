// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	cli "gopkg.in/alecthomas/kingpin.v2"

	"github.com/zackslash/pro2addict/midiParse"
)

// release is the current version number
const release = "1.0.0"

var (
	midiInLocation  = cli.Arg("in", "location of your MIDI file").Required().String()
	midiOutLocation = cli.Arg("out", "location for output MIDI file").String()
	debugMode       = false
)

func main() {
	cli.Version(release)
	cli.Parse()

	b, err := ioutil.ReadFile(*midiInLocation)
	if err != nil {
		log.Fatalf(err.Error())
	}

	midi, err := midiParse.ParseMidi(b, debugMode)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("GOT MIDI with %d tracks.\n", len(midi.Tracks))

	for _, t := range midi.Tracks {
		fmt.Printf("Track:%d/ch%d - %s - %s\n", t.Number, t.Channel, t.InstrumentName, t.TrackName)
	}

	// Get GP6 track order

	// Flip GP6 track order to AD2 track order

	// Save new midi

}
