package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	days = [...]func(*bufio.Reader) error{
		day01,
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

			return days[day](bufio.NewReader(f))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "the AoC day")
	rootCmd.Flags().BoolVarP(&examples, "test-examples", "e", false, "test on the example input")
	rootCmd.MarkFlagRequired("day")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
