package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	days = [...]func(*bufio.Reader) (partOne, partTwo any){
		day01,
		day02,
		day03,
		day04,
		day05,
		day06,
		day07,
		day08,
	}
	examples bool
	day      int

	rootCmd = &cobra.Command{
		Short:             "Compute solutions for the Advent of Code 2022.",
		Use:               "aoc-2022",
		Args:              cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		SilenceErrors:     true,
		SilenceUsage:      true,
		RunE: func(_ *cobra.Command, _ []string) error {
			var name string
			if examples {
				name = fmt.Sprintf("testdata/%02d.example", day)
			} else {
				name = fmt.Sprintf("testdata/%02d.input", day)
			}

			day--
			if day < 0 || day >= len(days) {
				return fmt.Errorf("input day must be 1 <= d <= %d", len(days))
			}

			f, err := os.Open(name)
			if err != nil {
				return err
			}
			defer f.Close()

			partOne, partTwo := days[day](bufio.NewReader(f))
			fmt.Println("Part one:", partOne)
			fmt.Println("Part two:", partTwo)
			return nil
		},
	}
)

func init() {
	f := rootCmd.Flags()
	f.IntVarP(&day, "day", "d", -1, "the AoC day")
	f.BoolVarP(&examples, "test-examples", "e", false, "test on the example input")
	rootCmd.MarkFlagRequired("day")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
