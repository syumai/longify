# longify

* A command to output longified ascii art.
  - Inspired by Tweet from [@sheepla](https://github.com/sheepla): https://twitter.com/Sheeeeepla/status/1522199846870196225

## Installation

```console
go install github.com/syumai/longify@latest
```

## Usage

```console
$ cat cow.txt
           __n__n__
    .------`-\00/-'
   /  ##  ## (oo)
  /|    ##    |
 /  \## __   ./
     |//YY \|/
     |||   |||

$ cat cow.txt | longify
           __n__n__
    .------`-\00/-'
   /  ##  ## (oo)
  /|    ##    |
  /|    ##    |
  /|    ##    |
  /|    ##    |
  /|    ##    |
  /|    ##    |
 /  \## __   ./
     |//YY \|/
     |||   |||

$ cat cow.txt | longify -l 3 # specify length (default: random number in 3-13)
           __n__n__
    .------`-\00/-'
   /  ##  ## (oo)
  /|    ##    |
  /|    ##    |
  /|    ##    |
 /  \## __   ./
     |//YY \|/
     |||   |||

# cat cow.txt | longify -p 7 # speficy longified position (default: center)
           __n__n__
    .------`-\00/-'
   /  ##  ## (oo)
  /|    ##    |
 /  \## __   ./
     |//YY \|/
     |||   |||
     |||   |||
     |||   |||
     |||   |||
     |||   |||

```

* ASCII Art is created by [Shanaka Dias](https://www.asciiart.eu/animals/cows)

## License

MIT

## Author

syumai
