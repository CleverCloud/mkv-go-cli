package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

func executeAppend(ctx context.Context, rdb *redis.Client, key, value string) {
	if key == "" || value == "" {
		fmt.Println("You need to provide a key and a value to append")
		return
	}
	answer, err := rdb.Append(ctx, key, value).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeCommands(ctx context.Context, rdb *redis.Client) {
	answer, err := rdb.Command(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range answer {
		fmt.Println(c.Name)
	}
}

func executeDbSize(ctx context.Context, rdb *redis.Client) {
	answer, err := rdb.DBSize(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeDecr(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key to decrement")
		return
	}
	answer, err := rdb.Decr(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeDel(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key to delete")
		return
	}
	answer, err := rdb.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeExists(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key name to check")
		return
	}
	answer, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeFlushDb(ctx context.Context, rdb *redis.Client) {
	answer, err := rdb.FlushDB(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeGet(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key to get")
		return
	}
	answer, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeIncr(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key to increment")
		return
	}
	answer, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeKeys(ctx context.Context, rdb *redis.Client, pattern string) {
	if pattern == "" {
		pattern = "*"
	}
	answer, err := rdb.Keys(ctx, pattern).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executePing(ctx context.Context, rdb *redis.Client) {
	answer, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeRawCommand(ctx context.Context, rdb *redis.Client, command string) {
	if command == "" {
		fmt.Println("You need to provide a command to execute")
		return
	}
	parts := strings.Split(command, " ")
	args := make([]interface{}, len(parts))
	for i, v := range parts {
		args[i] = v
	}
	answer, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeSet(ctx context.Context, rdb *redis.Client, key, value string) {
	if key == "" || value == "" {
		fmt.Println("You need to provide a key and a value to set")
		return
	}
	answer, err := rdb.Set(ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeStrLen(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key name to check")
		return
	}
	answer, err := rdb.StrLen(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func executeType(ctx context.Context, rdb *redis.Client, key string) {
	if key == "" {
		fmt.Println("You need to provide a key name to check")
		return
	}
	answer, err := rdb.Type(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}
