/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pezza/advent-of-code/cli"
	"github.com/pezza/advent-of-code/puzzles"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [year] [day]",
	Short: "Runs a specific puzzle.",
	Long: `Will run a specific puzzle, requires 2 arguents year and day.  By default, the puzzle input is embedded (and is the authors puzzle input).  
	To pass in a different puzzle input use the --file flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Handle re-showing cursor on Ctrl+C
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-c
			cli.Interrupted()
			os.Exit(1)
		}()

		year, _ := strconv.Atoi(args[0])
		day, _ := strconv.Atoi(args[1])
		if puzzle, err := puzzles.GetPuzzle(day, year); err != nil {

		} else {
			cli.Runner(puzzle)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 2 {
			return errors.New("run command requires <year> and <day> arguments. ")
		}

		if year, err := strconv.Atoi(args[0]); err != nil || year < 2015 || year > 2021 {
			return errors.New("year must be an integer between 2015 and 2021")
		}

		if day, err := strconv.Atoi(args[1]); err != nil || day < 1 || day > 25 {
			return errors.New("day must be an integer between 1 and 25")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().StringP("file", "f", "", "Path to input file to use.")
}
