// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package midiParse

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// http://cs.fit.edu/~ryan/cse4051/projects/midi/midi.html used as MIDI spec reference
var (
	channel10NoteEventPrefix = byte(0x99)
	instrumentMetaID         = []byte{0xFF, 0x04}
	sequenceTrackNameID      = []byte{0xFF, 0x03}
	channelIDPrefix          = []byte{0xFF, 0x20, 0x01}
)

// ParseMidi parses a midi byte array into midi struct
func ParseMidi(raw []byte, debug bool) (*MIDI, error) {
	m := MIDI{
		Raw: raw,
	}

	header, err := getHeaderFromBytes(m.Raw)
	if err != nil {
		return nil, err
	}
	m.Header = header

	if debug {
		fmt.Printf("HEADER: offset: %d, Format:%d ,Tracks:%d \n", m.Header.Chunk.End, m.Header.Format, m.Header.Tracks)
	}

	os := m.Header.Chunk.End
	var tn int64 = 1
	for {
		if int(os) >= len(m.Raw) {
			break
		}
		tr, err := getTrackFromBytes(m.Raw, os, tn)
		if err != nil {
			return nil, err
		}
		os = tr.Chunk.End
		m.Tracks = append(m.Tracks, tr)
		tn = tn + 1

		if debug {
			fmt.Printf("TRACK: offset: %d, number:%d \n", tr.Chunk.End, tr.Number)
		}

		// bytes '2' 99'2c'488360892c40 is note code
		fmt.Printf("Found %d drum notes\n", len(tr.Notes))

		for _, d := range tr.Notes {
			fmt.Printf("%x\n", d.Chunk.data)
		}
	}

	return &m, nil
}

func getTrackFromBytes(midi []byte, offset int32, number int64) (Track, error) {
	t := string(midi[offset:(offset + typeByteLen)])
	track := Track{}
	if !isValidTrackType(t) {
		return track, fmt.Errorf("Track specified invalid type")
	}

	l := midi[(offset + typeByteLen):(offset + typeByteLen + lenByteLen)]
	li := binary.BigEndian.Uint32(l)
	track.Chunk.Start = offset
	track.Chunk.End = offset + int32(li) + typeByteLen + lenByteLen
	track.Chunk.mType = t
	track.Chunk.mLen = li
	track.Chunk.data = midi[(offset + typeByteLen + lenByteLen):track.Chunk.End]
	track.InstrumentName = readInstrumentForTrack(track.Chunk.data)
	track.TrackName = readNameForTrack(track.Chunk.data)
	track.Channel = readChannelForTrack(track.Chunk.data)
	track.Notes = readDrumNoteEventsForTrack(track.Chunk.data)
	track.Number = number
	return track, nil
}

// readDrumNoteEventsForTrack reads track event data
func readDrumNoteEventsForTrack(trackData []byte) []Note {
	res := []Note{}
	offset := 0
	for _, b := range trackData {
		if b == channel10NoteEventPrefix {
			ev := Note{
				Chunk: Chunk{
					mLen:  8,
					data:  trackData[offset : offset+8],
					Start: int32(offset),
					End:   int32(offset + 8),
				},
			}

			res = append(res, ev)
		}
		offset++
	}
	return res
}

// readInstrumentForTrack reads instrument name from a track using Meta Event ID
func readInstrumentForTrack(trackData []byte) string {
	n := ""
	res := bytes.SplitN(trackData, instrumentMetaID, 1)
	if len(res) >= 1 {
		leA := res[0]
		if len(leA) > 0 {
			lenq := leA[3:4]
			st := fmt.Sprintf("%d", lenq[0])
			i, _ := strconv.Atoi(st)
			return string(leA[4 : 4+i])
		}
	}
	return n
}

// readNameForTrack reads instrument name from a track using Meta Event ID
func readNameForTrack(trackData []byte) string {
	n := ""
	res := bytes.SplitN(trackData, sequenceTrackNameID, 1)
	if len(res) >= 1 {
		leA := res[0]
		if len(leA) > 0 {
			lenq := leA[3:4]
			st := fmt.Sprintf("%d", lenq[0])
			i, _ := strconv.Atoi(st)
			return string(leA[4 : 4+i])
		}
	}
	return n
}

func readChannelForTrack(trackData []byte) int {
	res := bytes.SplitN(trackData, channelIDPrefix, 1)
	if len(res) >= 1 {
		leA := res[0]
		if len(leA) > 0 {
			lenq := leA[3:4]
			chs := fmt.Sprintf("%d", lenq[0])
			chi, _ := strconv.Atoi(chs)
			return chi
		}
	}
	return -1
}

func isValidTrackType(t string) bool {
	return t == typeTrack
}

func getHeaderFromBytes(midi []byte) (Header, error) {
	var offset int32 = (typeByteLen + lenByteLen)
	l := midi[typeByteLen:offset]
	li := binary.BigEndian.Uint32(l)
	head := Header{Chunk: Chunk{
		mType: string(midi[:typeByteLen]),
		mLen:  li,
		data:  midi[offset:(offset + int32(li))],
	}}
	head.Format = int16(binary.BigEndian.Uint16(head.Chunk.data[:2]))
	head.Tracks = int16(binary.BigEndian.Uint16(head.Chunk.data[2:4]))

	if !isValidHeaderType(head.Chunk.mType) {
		return head, fmt.Errorf("Header specified invalid header type")
	}

	if !isValidMidiFormat(head.Format) {
		return head, fmt.Errorf("Header specified invalid format")
	}
	head.Chunk.Start = 0
	head.Chunk.End = (offset + int32(li))
	return head, nil
}

func isValidHeaderType(t string) bool {
	return t == typeHeader
}

func isValidMidiFormat(format int16) bool {
	if format == 0 {
		// contains a single track
		return true
	}

	if format == 1 {
		// Contains one or more simultaneous tracks
		// The first track of a Format 1 file is special, and is also known as the 'Tempo Map'. It should contain all meta-events of the types Time Signature, and Set Tempo.
		// The meta-events Sequence/Track Name, Sequence Number, Marker, and SMTPE Offset. should also be on the first track of a Format 1 file.
		return true
	}

	if format == 2 {
		// Contains one or more independant tracks
		return true
	}

	return false
}
