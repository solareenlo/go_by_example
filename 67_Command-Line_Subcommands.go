package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmb := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmb.Bool("enable", false, "enable")
	fooName := fooCmb.String("name", "", "name")

	barCmb := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmb.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmb.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmb.Args())
	case "bar":
		barCmb.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", fooCmb.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
