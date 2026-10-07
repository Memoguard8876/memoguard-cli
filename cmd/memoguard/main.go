package main

import (
	"context"
	"os"

	cli "github.com/memoguard8876/memoguard-cli"
)

func main() {
	os.Exit(cli.Run(context.Background(), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
