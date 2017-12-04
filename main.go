package main

func main() {

	year, day, error := checkParams()

	if error != nil {
		outputUseage(error)
		return
	}

	dailyRunner, error := getPuzzle(day, year)

	if error != nil {
		outputUseage(error)
		return
	}

	runner(dailyRunner)
}
