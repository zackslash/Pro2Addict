# Pro 2 Addict
A command-line tool written in Go; that converts MIDI drum tracks from Guitar Pro 6 to Addictive Drums 2.

### Pre-built Downloads

[Windows](https://drive.google.com/file/d/0BxVgyPt7oHs5azFuSG5uNk5Hc1U/view?usp=sharing&resourcekey=0-owKNkKZAdKfn88tBgpNbRg)

[Mac](https://drive.google.com/file/d/0BxVgyPt7oHs5ajRPeE1qWklGRFE/view?usp=sharing&resourcekey=0-PCbSROzoDpRvGQstrpJE7g)

[Linux](https://drive.google.com/file/d/0BxVgyPt7oHs5MTJjWVAxTXVMcDA/view?usp=sharing&resourcekey=0-2Qe0DpCjvSVVj8oIjm1OSA)

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

## Build
Source will compile under Go 1.x. [Glide](https://glide.sh/) has been used for package management.

## License
Pro2addict is released under the MIT License.
See the bundled [LICENSE](https://github.com/zackslash/pro2addict/blob/master/LICENCE) file for details.
