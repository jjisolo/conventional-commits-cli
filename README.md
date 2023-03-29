# Conventional commit CLI tool written in Go
I did not find any CLI interface for making conventional commits, so i decided to create my own. Please enjoy :)

![application screenshot](https://github.com/jjisolo/conventional-commits-cli/blob/main/screenshot.png)

# Installation

First clone the repository
```
$ git clone https://github.com/jjisolo/conventional-commits-cli
```

Then run the build and installation
```
$ cd conventional-commits-cli
$ go build .
$ go install
```
Now the program can be called via
```
$ ccommit [-a|-A|..filenames]
```

Optionally you can alias the ```ccommit``` to ``cc`` in your shell

# Todo

- [ ] Do not commnit when the process shut downs(for instance <C-c> is pressed by the user)
- [ ] Better git output(some colors maybe)
- [ ] Handle more git options
