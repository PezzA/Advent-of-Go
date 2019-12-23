package Day201923

import (
	"fmt"
	"sync"
	"time"

	"github.com/pezza/advent-of-code/common"

	"github.com/pezza/advent-of-code/2019/intcode"
)

type computer struct {
	in  chan int64
	out chan int64
}

func runProgram(input string) {
	var handler [50]chan common.Point
	var ins [50]chan int64
	var outs [50]chan int64

	for i := 0; i < 50; i++ {
		ins[i] = make(chan int64, 10)
		outs[i] = make(chan int64, 10)
		handler[i] = make(chan common.Point, 10)
	}

	natIn := make(chan common.Point, 0)

	var wg sync.WaitGroup
	wg.Add(100)

	for index := range ins {
		go func(i int) {
			defer wg.Done()

			vm := intcode.New(input)
			ins[i] <- int64(i)
			vm.RunProgram(nil, nil, ins[i], outs[i], nil)

		}(index)

		go func(i int) {
			defer wg.Done()
			add, x, y := int64(0), int64(0), int64(0)

			step := 0

			for {
				select {
				case val := <-outs[i]:
					switch step {
					case 0:
						add = val
						step = 1
					case 1:
						x = val
						step = 2
					case 2:
						y = val
						step = 0

						if add == 255 {
							natIn <- common.Point{int(x), int(y)}
						} else {
							fmt.Print(".")
							handler[add] <- common.Point{int(x), int(y)}
						}
					default:
					}
				}
			}
		}(index)

		go func(i int) {
			for {
				select {
				case val := <-handler[i]:
					ins[i] <- int64(val.X)
					ins[i] <- int64(val.Y)
				default:
				}
			}
		}(index)
	}

	// NAT
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)

		point := common.Point{}

		outputs := make(map[int]bool, 0)

		for {
			select {
			case val := <-natIn:
				point = val
			case t := <-ticker.C:

				count := 0
				for index := range ins {
					count += len(ins[index])
					count += len(outs[index])
				}

				if count == 0 && point.X != 0 && point.Y != 0 {
					fmt.Println("NAT RESTART", t, "=>", count, "IDLE", point)
					ins[0] <- int64(point.X)
					ins[0] <- int64(point.Y)

					if _, ok := outputs[point.Y]; ok {
						fmt.Println("#################################", point.Y)
					} else {
						outputs[point.Y] = true
					}
					point = common.Point{0, 0}
				}
			default:
			}
		}
	}()

	wg.Wait()
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
