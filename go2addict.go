// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	cli "gopkg.in/alecthomas/kingpin.v2"

	"github.com/zackslash/pro2addict/maps"
	"github.com/zackslash/pro2addict/midiParse"
)

const (
	release = "1.0.0"
	drumKey = "drum"
)

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

	fmt.Printf("Got MIDI with %d tracks.\n", len(midi.Tracks))

	res := make([]byte, len(b))
	copy(res, b)

	for _, t := range midi.Tracks {
		if debugMode == true {
			fmt.Printf("Track:%d/ch%d - %s - %s\n", t.Number, t.Channel, t.InstrumentName, t.TrackName)
		}
		if strings.Contains(strings.ToLower(t.InstrumentName), drumKey) {
			for _, n := range t.Notes {
				s := fmt.Sprintf("%d", n.GetNoteData()[1])
				i, _ := strconv.Atoi(s)
				rep := maps.GetMappedNote(i)
				new := make([]byte, len(n.GetNoteData()))
				copy(new, n.GetNoteData())
				new[1] = byte(rep)

				if debugMode == true {
					fmt.Printf("From:%x To:%x\n", n.GetNoteData(), new)
				}

				res = bytes.Replace(res, n.GetNoteData(), new, 1)
			}
		}
	}

	output := ""
	if *midiOutLocation != "" {
		output = *midiOutLocation
	} else {
		output = *midiInLocation+"conv.mid"
	}

	ioutil.WriteFile(output, res, 0644)
	fmt.Printf("Done - %s\n", output)
}
