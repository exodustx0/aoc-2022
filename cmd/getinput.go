package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const goDayTmpl = `package cmd

import "bufio"

func day%02d(input *bufio.Reader) error {
	return nil
}
`

var (
	session string

	getInputCmd = &cobra.Command{
		Short: "Get input prompt for a given day.",
		Use:   "getinput",
		RunE: func(_ *cobra.Command, _ []string) error {
			if time.Now().Before(time.Date(2022, 12, 2, 0, 0, 0, 0, time.FixedZone("UTC", -5*60*60))) {
				return errors.New("puzzle not available yet")
			}

			p := fmt.Sprintf("testdata/%02d.input", day)
			if _, err := os.Stat(p); !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("%q already exists", p)
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
			if err != nil {
				return err
			}

			req.AddCookie(&http.Cookie{Name: "session", Value: session})
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				return errors.New(res.Status)
			}

			inf, err := os.Create(p)
			if err != nil {
				return err
			}
			defer inf.Close()

			if _, err = io.Copy(inf, res.Body); err != nil {
				return err
			}

			gof, err := os.Create(fmt.Sprintf("cmd/%02d.go", day))
			if err != nil {
				return err
			}
			defer gof.Close()

			_, err = gof.WriteString(fmt.Sprintf(goDayTmpl, day))
			return err
		},
	}
)

func init() {
	f := getInputCmd.Flags()
	f.IntVarP(&day, "day", "d", -1, "the AoC day")
	f.StringVarP(&session, "session", "s", "", "the session cookie")
	getInputCmd.MarkFlagRequired("day")
	getInputCmd.MarkFlagRequired("session")
	rootCmd.AddCommand(getInputCmd)
}
