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


#### A vim Journal
So, some notes on how using vim has progressed.  This is using go as the language for the 2022 event.
* Day 1: Mainly getting used to NerdTree and `Vex` commands to have multiple windows open.  `ctrl-w =` nice shortcut to make all windows same size.
* Day 2: Hating Life!  Why did I think this would be a good idea.  Spent more time today accidentlly leaving insert mode and losing text.  Generally fighting with Vim.
* Day 3: so `dd` is actually cut and not delete is it? and this overwrites the `yy` does it?  #Registers#, `reg` .. nice.  So whilst `dd` overwrites the default reg, `yy` also copies to register 0 and can paste from there.
* Day 4: Really, really, need to change that cursor from black.  Using GruvBox dark theme, can't see it.  Turns out this is controlled in the WSL colour theme. Now site to a nice bright Cyan.  Also installed `vim-airline`, edit mode and changes, active window now much more apparant. 
* Day 5: #Debugging#  Had to roll a stack implementation for this, and deal with all the pointer/receiver issues.  Lots of step-through debugging today, as can't find a way of getting log output in vim-go test output.  However, vim-go has a good integration with delve, and althrough `GoDebugTestFunc` wasnt working out of the box `GoDebugTest` was (it might just be a working folder issue).  Debugging was really nice, but probably need to setup some default layouts for the windows, for AoC debugging goroutine probably isnt too important.  One thing this did throw up is the ability to need to move around windows quickly, so cue first #Key Mapping#, after some googling, remapped the default windows navigate keys to ctrl-h/j/k/l rathen that the ctrl-w h/j/k/l chord.  Much faster    
* Day 6: Nothing new, but going to need to figure out how to capture stdout/err for debugging as well.



