# :christmas_tree: Advent-of-code :star:
[![Go](https://github.com/PezzA/advent-of-code/actions/workflows/go.yml/badge.svg)](https://github.com/PezzA/advent-of-code/actions/workflows/go.yml)

This repo contains solutions to puzzles on the [advent of code](https://www.adventofcode.com) website.  Currently this 
covers solutions for years 2015-2020.  However not all years are complete. :sob:

To run a puzzle 

`./adventof-code run <year> <day>`

By default the puzzle input is embedded into the solution itself, but can you can specify your own input with the --file flag.

`./advent-of=code run <year> <day> --file <path to input>`

### Package Contents

#### main.go & /cmd
Cobra scaffolding for the cli application

#### cli
cli UI functions

#### puzzles
the main body of the package.  The root of this package defines interfaces to use the puzzles, then each sub package within puzzles contains the puzzles solutions for a particular year







