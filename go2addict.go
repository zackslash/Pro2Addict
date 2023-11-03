// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/zackslash/pro2addict/maps"
	"github.com/zackslash/pro2addict/midiParse"

	cli "gopkg.in/alecthomas/kingpin.v2"
)

const (
	release = "1.0.1"
	drumKey = "drum"
)

var (
	midiInLocation  = cli.Arg("in", "location of your Guitar Pro MIDI file").Required().String()
	midiOutLocation = cli.Arg("out", "location to output converted (AD2) MIDI file").String()
	debugMode       = false
)

func main() {
	cli.Version(release)
	cli.Parse()

	// Read and parse the MIDI file.
	originalBytes, midi, err := readMIDIFile(*midiInLocation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read MIDI with %d tracks.\n", len(midi.Tracks))

	// Convert drum notes in the MIDI data.
	convertedBytes := convertDrumNotes(midi, originalBytes)

	// Determine the output file location.
	output := *midiOutLocation
	if output == "" {
		output = *midiInLocation + "-ad2.mid"
	}

	// Write the converted MIDI data to the output file.
	if err := writeMIDIFile(output, convertedBytes); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done - %s\n", output)
}

// readMIDIFile reads a MIDI file from the given filePath and parses it.
// It returns the file contents, the parsed MIDI data, and an error if any.
func readMIDIFile(filePath string) ([]byte, *midiParse.MIDI, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	midi, err := midiParse.ParseMidi(b, debugMode)
	if err != nil {
		return nil, nil, err
	}

	return b, midi, nil
}

// convertDrumNotes takes a MIDI object and the original file bytes,
// and converts the drum notes using a mapping function. It returns the modified bytes.
func convertDrumNotes(midi *midiParse.MIDI, originalBytes []byte) []byte {
	res := make([]byte, len(originalBytes))
	copy(res, originalBytes)

	for _, t := range midi.Tracks {
		if debugMode {
			fmt.Printf("Track:%d/ch%d - %s - %s\n", t.Number, t.Channel, t.InstrumentName, t.TrackName)
		}

		if strings.Contains(strings.ToLower(t.InstrumentName), drumKey) {
			res = replaceNotes(t, res)
		}
	}

	return res
}

// replaceNotes takes a single track and the current result bytes,
// replaces the note data for drum notes, and returns the updated result bytes.
func replaceNotes(track midiParse.Track, res []byte) []byte {
	for _, n := range track.Notes {
		oldData := n.GetNoteData()
		s := fmt.Sprintf("%d", oldData[1])
		i, _ := strconv.Atoi(s)
		rep := maps.GetMappedNote(i)
		newData := make([]byte, len(oldData))
		copy(newData, oldData)
		newData[1] = byte(rep)

		if debugMode {
			fmt.Printf("From:%x To:%x\n", oldData, newData)
		}

		res = bytes.Replace(res, oldData, newData, 1)
	}
	return res
}

// writeMIDIFile writes the modified MIDI data to a file at the specified output location.
// It returns an error if the write fails.
func writeMIDIFile(output string, data []byte) error {
	return os.WriteFile(output, data, 0644)
}
