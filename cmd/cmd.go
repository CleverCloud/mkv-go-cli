package cmd

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func HandleCommand(ctx context.Context, rdb *redis.Client, command string, args []string) {
	var key, value string

	if len(args) > 0 {
		key = args[0]
	}
	if len(args) > 1 {
		value = args[1]
	}

	switch command {
	case "append":
		executeAppend(ctx, rdb, key, value)
	case "commands":
		executeCommands(ctx, rdb)
	case "decr":
		executeDecr(ctx, rdb, key)
	case "dbsize":
		executeDbSize(ctx, rdb)
	case "del":
		executeDel(ctx, rdb, key)
	case "exists":
		executeExists(ctx, rdb, key)
	case "flushdb":
		executeFlushDb(ctx, rdb)
	case "get":
		executeGet(ctx, rdb, key)
	case "incr":
		executeIncr(ctx, rdb, key)
	case "keys":
		executeKeys(ctx, rdb, key)
	case "ping":
		executePing(ctx, rdb)
	case "raw":
		executeRawCommand(ctx, rdb, key)
	case "set":
		executeSet(ctx, rdb, key, value)
	case "strlen":
		executeStrLen(ctx, rdb, key)
	case "type":
		executeType(ctx, rdb, key)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Use -h or --help for usage information.")
	}
}
