package main

func main() {

	day, runTest, error := checkParams()

	if error != nil {
		outputUseage(error)
		return
	}

	dailyData, dailyRunner, error := getPuzzle(day)

	if error != nil {
		outputUseage(error)
		return
	}

	runner(dailyData, dailyRunner, runTest)
}
