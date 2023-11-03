# Pro 2 Addict

A small command-line tool to convert MIDI drum tracks exported from Guitar Pro 6 to Addictive Drums.

### Usage
```bash
usage: pro2addict [<flags>] <in> [<out>]
Example : ./pro2addict great_drums.mid ./great_drums_ad2.mid

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Args:
  <in>     location of your Guitar Pro MIDI file
 [<out>]  location to output converted (AD2) MIDI file (If no output is specified the result will be output to the same directory and filename as the original, appended with "-ad2.mid"
```

## License
Pro2addict is released under the MIT License.
See the bundled [LICENSE](https://github.com/zackslash/pro2addict/blob/master/LICENCE) file for details.
