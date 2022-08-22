/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	tm "github.com/buger/goterm"
	"github.com/pezza/advent-of-code/puzzles"
	"github.com/spf13/cobra"
)

func noEntry() string {
	return tm.Color(". ", tm.CYAN)
}

func solvedFirst() string {
	return tm.Color("░ ", tm.YELLOW)
}

func solvedBoth() string {
	return tm.Bold(tm.Color("█ ", tm.YELLOW))
}

func otherFirst() string {
	return tm.Color("░ ", tm.GREEN)
}

func otherSolution() string {
	return tm.Color("█ ", tm.GREEN)
}

func unreliable() string {
	return tm.Bold(tm.Color("▓ ", tm.MAGENTA))
}

func toughCookie() string {
	return tm.Bold(tm.Color("‼ ", tm.RED))
}

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Gives details of which puzzles are solved.",
	Long:  `Will report for all years which puzzles are solved, should extend to say how many stars are complete`,
	Run: func(cmd *cobra.Command, args []string) {
		minYear, maxYear := puzzles.GetManualYearRange()
		fmt.Println()
		fmt.Println(tm.Bold(tm.Color("Advent of Code summary report:", tm.WHITE)))
		fmt.Println("")
		fmt.Printf("This project       : [ %v] First [ %v] Second\n", solvedFirst(), solvedBoth())
		fmt.Printf("Other projects     : [ %v] First [ %v] Second\n", otherFirst(), otherSolution())
		fmt.Printf("Problematic        : [ %v] Slow! [ %v] Oof size ... large\n", unreliable(), toughCookie())
		fmt.Println()
		fmt.Println("                       1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2")
		fmt.Println("     1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5")

		totalCount := 0
		for year := maxYear; year >= minYear; year-- {
			fmt.Print(tm.Color(fmt.Sprint(year), tm.GREEN))
			fmt.Print(" ")

			yearCount := 0

			for day := 1; day <= 25; day++ {
				puzz, err := puzzles.GetPuzzle(day, year)

				if err != nil {
					fmt.Print(noEntry())
					continue
				}

				_, _, _, state := puzz.Describe()

				switch state {
				case 1:
					fmt.Print(solvedFirst())
					yearCount = yearCount + 1
				case 2:
					fmt.Print(solvedBoth())
					yearCount = yearCount + 2
				case 3:
					fmt.Print(otherSolution())
					yearCount = yearCount + 2
				case 4:
					fmt.Print(unreliable())
					yearCount = yearCount + 2
				case 5:
					fmt.Print(toughCookie())
				case 6:
					fmt.Print(otherFirst())
					yearCount = yearCount + 1
				default:
					fmt.Print(noEntry())
				}
			}
			fmt.Print(tm.Color(strconv.Itoa(yearCount), tm.YELLOW))
			fmt.Println("")
			fmt.Println("")
			totalCount += yearCount
		}

		fmt.Println()
		fmt.Println(tm.Color("Total stars:", tm.WHITE), tm.Bold(tm.Color(fmt.Sprint(totalCount), tm.WHITE)))
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
