package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davlgd/mkv-cli/cmd"
	"github.com/davlgd/mkv-cli/connect"
	"github.com/davlgd/mkv-cli/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use -h or --help for usage information.")
		return
	}

	command := os.Args[1]
	if command == "-h" || command == "--help" {
		utils.PrintHelp()
		return
	}

	auth := os.Getenv("KV_TOKEN")
	if auth == "" {
		fmt.Println("KV_TOKEN environment variable is not set.")
		return
	}

	ctx := context.Background()
	rdb := connect.GetClient()

	if len(os.Args) > 2 {
		cmd.HandleCommand(ctx, rdb, command, os.Args[2:])
	} else {
		cmd.HandleCommand(ctx, rdb, command, []string{})
	}
}
