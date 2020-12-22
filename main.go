package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	flags := SetFlags()
	if flag.Arg(0) == "" {
		fmt.Println("Please enter something or try -h for flags")
		os.Exit(1)
	}

	switch flags.EffectVar {
	case "caps":
		for i := range flag.Args() {
			string := strings.ToUpper(flag.Arg(i))
			fmt.Print(string + " ")
		}
	case "colour":
		coloured(flags)
	case "alt":
		for i := range flag.Args() {
			string := AlternateCase(flag.Arg(i))
			fmt.Print(string + " ")
		}

	default:
		for i := range flag.Args() {
			string := AlternateCase(flag.Arg(i))
			fmt.Print(string + " ")
		}
	}

}

func coloured(flags Flags) {

	for j := range flag.Args() {
		rs := strings.Split(flag.Arg(j), "")
		for i := range rs {
			switch rs[i] {
			case "@":
				rs[i] = "hash"
			case "!":
				rs[i] = "exclamation"
			case "?":
				rs[i] = "question"
			case "#":
				rs[i] = "hash"
			}
			print(flags, getColour(flags, i)+rs[i]+":")
		}
		if flags.PlaneVar == "horizontal" {
			fmt.Print(" ")
		}
	}
}

func getColour(flags Flags, index int) string {
	yellowMod := ":alphabet-yellow-"
	whiteMod := ":alphabet-white-"

	switch flags.ColourVar {
	case "yellow":
		return yellowMod
	case "white":
		return whiteMod
	case "alt":
		if index%2 == 0 {
			return yellowMod
		}
		return whiteMod
	}
	return ""
}

func AlternateCase(s string) string {
	rs, upper := []rune(s), false
	for i, r := range rs {
		if unicode.IsLetter(r) {
			if upper = !upper; upper {
				rs[i] = unicode.ToUpper(r)
			}
		}
	}
	return string(rs)
}

func print(flags Flags, value string) {
	switch flags.PlaneVar {
	case "vertical":
		fmt.Println(value)
	default:
		fmt.Print(value)
	}
}

//Flags struct
type Flags struct {
	ChangeVar                      int
	EffectVar, PlaneVar, ColourVar string
}

func SetFlags() Flags {
	var flags Flags
	flag.StringVar(&flags.PlaneVar, "plane", "horizontal", "Direction in which the text appears")
	flag.StringVar(&flags.ColourVar, "colour", "alt", "colours in which the text appears")
	flag.IntVar(&flags.ChangeVar, "change", 1, "how often an effect is changed")
	flag.StringVar(&flags.EffectVar, "effect", "alt", "text effect")
	flag.Parse()
	return flags
}
