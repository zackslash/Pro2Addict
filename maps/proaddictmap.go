// pro2addict
// Copyright 2017 Luke Hines
// Released under the MIT License
// https://tldrlegal.com/license/mit-license

package maps

// noteMap is the mapping of GP6 MIDI node to the equivalent AD2 note
var noteMap = map[int]int{
  35: 36,   // Kick (hit)

  41 : 65,  // Tom very low (hit) 
  43 : 67,  // Tom low (hit)   
  45 : 69,  // Tom medium (hit)  
  47 : 71,  // Tom high (hit)  
  48 : 72,  // Tom very high (hit)

  38 : 38,  //  Snare (hit) 
  //38 : 37,// Snare (rim shot) - Note is duplicate
  37 : 42,  // Snare (side stick)

  51 : 60,  // Ride (middle) 
  59 : 62,  // Ride (edge)  
  53 : 61,  // Ride (bell)  

  44 : 59,  // Pedal hihat (hit)   
  42 : 48,  // Hihat (closed)     
  46 : 51,  // Hihat (half)     
  //46 : 54,// Hihat (open) - Note is duplicate

  49 : 79,  // Crash medium (hit)
  57 : 77,  // Crash high (hit)

  //56 : x, // Cowbell high (hit) - kit excluded from AD2
  //56 : x, // Cowbell medium (hit) - kit excluded from AD2
  //56 : x, // Cowbell low (hit) - kit excluded from AD2

  55 : 89,  // Splash (hit)
  52 : 91,  // China (hit)      
}

// GetMappedNote returns AD2 note for givem GP6
func GetMappedNote(key int) int {
    return noteMap[key]
}