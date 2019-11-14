## Sokoban [![GoDoc](https://godoc.org/github.com/tristangoossens/sokoban-go?status.svg)](https://godoc.org/github.com/tristangoossens/sokoban-go) [![Go Report Card](https://goreportcard.com/badge/github.com/tristangoossens/sokoban-go)](https://goreportcard.com/report/github.com/tristangoossens/sokoban-go) [![Build Status](https://travis-ci.com/tristangoossens/sokoban-go.svg?branch=master)](https://travis-ci.com/tristangoossens/sokoban-go)

![Logo](https://github.com/tristangoossens/sokoban-go/blob/master/images/logo.png)

Retro game sokoban clone made in Go! Created by tristangoossens. ![GitHub followers](https://img.shields.io/github/followers/tristangoossens?style=social)  
Please star this repository to help my project grow! ![GitHub stars](https://img.shields.io/github/stars/tristangoossens/sokoban-go?style=social)

## The game

Here is a screenshot of sokoban's titlescreen! 

![Titlescreen](https://github.com/tristangoossens/sokoban-go/blob/master/images/game-titlescreen.png)

Screenshot of sokoban's first level!

![Level1](https://github.com/tristangoossens/sokoban-go/blob/master/images/game-gamescreen.png)

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

- [ ] **Read levels from 1 file(instead of 20)**
- [ ] **Implement support for multiple crates and goals.**
- [ ] Make the player, crate and goal a square(equal width and height)
- [ ] Add more levels(goal is around 40-50)
- [ ] Implement save current level button
- [ ] Add a timer to the gamescreen to time your best times.
    - [ ] Save your best time to a markdown file(table).
- [ ] Add test files for all game files
- [ ] Start from last saved level.

*Feel free to help me add the functions above with a pull request!*

## Links

- [Sokoban on Godoc.](https://godoc.org/github.com/tristangoossens/sokoban-go/game)
- [Sokoban on repl.it.](https://repl.it/@tristangoossens/sokoban-go)

