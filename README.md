# :christmas_tree: Advent-of-code :star:
[![Build](https://github.com/PezzA/advent-of-code/actions/workflows/build-only.yml/badge.svg)](https://github.com/PezzA/advent-of-code/actions/workflows/build-only.yml)<br/>
This repo contains solutions to puzzles on the [advent of code](https://www.adventofcode.com) website.  Currently this 
covers solutions for years 2015-2020.  However not all years are complete. :sob:


|  |  |
|--|--|
| 2022 | Go, but not allowed to leave neovim editor |
| 2021 | Typescript |
| 2020 | Go |
| 2019 | Go |
| 2018 | Go |
| 2017 | Go |
| 2016 | Elm, Go  |
| 2015 | C#, Go |

## Usage
To run a puzzle 

`./adventof-code run <year> <day>`

By default the puzzle input is embedded into the solution itself, but can you can specify your own input with the --file flag.

`./advent-of=code run <year> <day> --file <path to input>`

## Package Contents

### main.go & /cmd
Cobra scaffolding for the cli application

#### cli
cli UI functions

#### puzzles
the main body of the package.  The root of this package defines interfaces to use the puzzles, then each sub package within puzzles contains the puzzles solutions for a particular year







