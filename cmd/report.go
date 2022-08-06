/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/pezza/advent-of-code/puzzles"
	"github.com/spf13/cobra"
)

func noEntry() string {
	return tm.Color(". ", tm.CYAN)
}

func otherSolution() string {
	return tm.Color("* ", tm.GREEN)
}

func otherFirst() string {
	return tm.Color("* ", tm.BLUE)
}

func solvedFirst() string {
	return tm.Color("* ", tm.WHITE)
}

func solvedBoth() string {
	return tm.Bold(tm.Color("* ", tm.YELLOW))
}

func unreliable() string {
	return tm.Bold(tm.Color("* ", tm.RED))
}

func toughCookie() string {
	return tm.Bold(tm.Color("X ", tm.RED))
}

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Gives details of which puzzles are solved.",
	Long:  `Will report for all years which puzzles are solved, should extend to say how many stars are complete`,
	Run: func(cmd *cobra.Command, args []string) {
		minYear, maxYear := puzzles.GetManualYearRange()
		fmt.Println("Legend:")
		fmt.Printf("This project       : [ %v] First [ %v] Second\n", solvedFirst(), solvedBoth())
		fmt.Printf("Other projects     : [ %v] First [ %v] Second\n", otherFirst(), otherSolution())
		fmt.Printf("Unreliable         : [ %v]\n", unreliable())
		fmt.Printf("Oof size ... large : [ %v]\n", toughCookie())
		fmt.Println()
		fmt.Println("                       1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2")
		fmt.Println("     1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5")
		for year := minYear; year <= maxYear; year++ {
			fmt.Print(tm.Color(fmt.Sprint(year), tm.GREEN))
			fmt.Print(" ")

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
				case 2:
					fmt.Print(solvedBoth())
				case 3:
					fmt.Print(otherSolution())
				case 4:
					fmt.Print(unreliable())
				case 5:
					fmt.Print(toughCookie())
				case 6:
					fmt.Print(otherFirst())
				default:
					fmt.Print(noEntry())
				}

			}
			fmt.Println("")
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
