package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/treeder/bump"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bump"
	app.Usage = "bump it dawg! See https://github.com/treeder/bump for more info."
	app.Action = bumper
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename",
			Value: "VERSION",
			Usage: "filename to look for version in",
		},
		cli.StringFlag{
			Name:  "input",
			Usage: "use this if you want to pass in a string to pass, rather than read it from a file. Cannot be used with --filename.",
		},
		cli.BoolFlag{
			Name:  "extract",
			Usage: "this will just find the version and return it, does not modify anything. Safe operation.",
		},
		cli.StringFlag{
			Name:  "format",
			Usage: "either M for major, M-m for major-minor or M-m-p",
		},
		cli.StringFlag{
			Name:  "replace",
			Usage: "overwrites the version with what you pass in here",
		},
		cli.IntFlag{
			Name:  "index",
			Usage: "if zero (default), uses first match. If greater than zero, uses nth match. If less than zero, starts at last match and goes backwards, ie: last match is -1.",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func bumper(c *cli.Context) error {
	arg := "patch"
	// fmt.Println("ARGS:", c.Args())
	if len(c.Args()) < 1 {
		// log.Fatal("Invalid arg")
	} else {
		arg = c.Args().First()
		arg = strings.ToLower(arg)
	}

	// check for `[bump X]` in input, user can pass in git commit messages to auto bump different versions
	if strings.Contains(arg, "[bump minor]") {
		arg = "minor"
	} else if strings.Contains(arg, "[bump major]") {
		arg = "major"
	}

	var err error
	var vbytes []byte
	filename := c.String("filename")
	if c.IsSet("input") {
		vbytes = []byte(c.String("input"))
	} else {
		vbytes, err = ioutil.ReadFile(filename)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("%v not found. Use either --filename or --input to change where to look for version.", filename)
			}
			return err
		}
	}

	index := c.Int("index")

	var old, new string
	var newcontent []byte
	if c.IsSet("replace") {
		// this will just write the passed in version directly
		replace := c.String("replace")
		old, new, _, newcontent, err = bump.ReplaceInContent(vbytes, replace, index)
	} else {
		old, new, _, newcontent, err = bump.BumpInContent(vbytes, arg, index)
	}
	if err != nil {
		return err
	}
	if c.Bool("extract") {
		print(c, old)
		return nil
	}

	//	fmt.Fprintln(os.Stderr, "Old version:", old)
	//	fmt.Fprintln(os.Stderr, "New version:", new)
	if !c.IsSet("input") {
		err = ioutil.WriteFile(filename, newcontent, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	print(c, new) // write it to stdout so scripts can use it
	return nil
}

func print(c *cli.Context, version string) {
	format := c.String("format")
	if format == "" {
		fmt.Print(version)
		return
	}

	v := semver.New(version)
	// else, we format it
	var b bytes.Buffer
	for _, char := range format {
		if char == 'M' {
			b.WriteString(strconv.FormatInt(v.Major, 10))
		} else if char == 'm' {
			b.WriteString(strconv.FormatInt(v.Minor, 10))
		} else if char == 'p' {
			b.WriteString(strconv.FormatInt(v.Patch, 10))
		} else {
			b.WriteRune(char)
		}
	}
	fmt.Print(b.String())
}
