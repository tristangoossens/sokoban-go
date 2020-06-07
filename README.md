## Sokoban [![GoDoc](https://godoc.org/github.com/tristangoossens/sokoban-go?status.svg)](https://godoc.org/github.com/tristangoossens/sokoban-go) [![Go Report Card](https://goreportcard.com/badge/github.com/tristangoossens/sokoban-go)](https://goreportcard.com/report/github.com/tristangoossens/sokoban-go) [![Build Status](https://travis-ci.com/tristangoossens/sokoban-go.svg?branch=master)](https://travis-ci.com/tristangoossens/sokoban-go)

![Logo](https://github.com/tristangoossens/sokoban-go/blob/master/img/sokoban-logo.png)

Retro game sokoban clone made in Go! Created by tristangoossens. ![GitHub followers](https://img.shields.io/github/followers/tristangoossens?style=social)  
Please star this repository to help my project grow! ![GitHub stars](https://img.shields.io/github/stars/tristangoossens/sokoban-go?style=social)

[![Run on Repl.it](https://repl.it/badge/github/tristangoossens/sokoban)](https://repl.it/github/tristangoossens/sokoban) 

## The game

Here is a screenshot of sokoban's titlescreen! 

![Titlescreen](https://github.com/tristangoossens/sokoban-go/blob/master/img/game-titlescreen.png)

Screenshot of sokoban's first level!

![Level1](https://github.com/tristangoossens/sokoban-go/blob/master/img/game-gamescreen.png)

## How to play

First of all you will need Go, you can find more information [here](https://golang.org/).

When you have installed Go, you will need to install the game:

```shell
git clone https://github.com/tristangoossens/sokoban-go.git
```

Then play it using the following command:

```shell
go run run.go
```

Or build and run using the following:

```
go build run.go
./run
```

## Bucket list for future versions

- [x] **Implement support for multiple crates and goals.**
- [ ] (Could)Read levels from 1 file(instead of 20).
- [x] Get rid of unnecessary variables.
- [x] Clean up code.
- [ ] Add more levels(goal is around 30, more for v3).
- [x] Create a new UI to support bigger levels.
- [ ] (Could)Add soundeffects for events.
- [ ] Add new functionality. (v3)
    - [ ] Teleporter.
    - [ ] Player trap.
    - [ ] One Way Wall.
- [ ] Implement a level selection ui. (v3)
- [x] Implement save current level button.
- [x] (Should)Add a timer to the gamescreen to time your best times.
    - [x] (Should)Save your best time to a markdown file(table).
- [ ] Add test files for all game files (v3)
- [x] Start from last saved level.

*Feel free to help me add the functions above with a pull request!*

## Links

- [Sokoban on Godoc.](https://godoc.org/github.com/tristangoossens/sokoban-go/game)
- [Sokoban on repl.it.](https://repl.it/@tristangoossens/sokoban-go)

