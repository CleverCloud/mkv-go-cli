package utils

import "fmt"

func PrintHelp() {
	fmt.Println(`mkv-cli is a simple CLI tool for Clever Cloud's Materia serverless KV store.

Usage:
  command [OPTIONS] <COMMAND> [ARGUMENTS]

You need to set the KV_TOKEN environment variable with your Materia KV token to connect properly.

Commands:
  ping       Ping the server
  raw        Send a raw command

  get        Get a value
  set        Set a value
  append     Append a value
  incr       Increment a value
  decr       Decrement a value
  del        Delete a key
  strlen     Get the length of a key
  type       Get the type of a key
  keys       Get all keys matching a pattern
  exists     Check if a key exists
  
  commands   Get the list of supported commands
  dbsize     Get the number of keys in the database
  flushdb    Flush the database

Options:
-h, –help Show this help message`)
}
