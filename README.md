# Clever Cloud's Materia KV Go CLI

This project shows how to connect to create a CLI for Materia serverless key-value store from [Clever Cloud](https://www.clever-cloud.com). It authenticates you with a [biscuit-based](https://www.biscuitsec.org/) token, sends commands and prints responses using the Redis protocol. It's a demonstration application using [Go-Redis client](https://redis.uptrace.dev/).

## How to use it

You need [Go language](https://go.dev/), to create [a Materia KV add-on](https://developers.clever-cloud.com/doc/addons/materia-kv/) on Clever Cloud and set the `KV_TOKEN` environment variable to connect. Then, to get/set a key:

```sh
mkv-cli get <key>
mkv-cli set <key> <value>
```

To increment/decrement a key:

```sh
mkv-cli incr <key>
mkv-cli decr <key>
```

To list all keys:

```sh
mkv-cli keys "*"
```

To delete a key:

```sh
mkv-cli del <key>
```

To list supported commands:

```sh
mkv-cli commands
```

You can also execute any supported command with the `raw` feature:

```sh
mkv-cli raw "<command>"
```

For more information, you can use the help command:

```sh
mkv-cli --help
```

## Build

```sh
make
```

## Build & Install

```sh
make install
```

or

```sh
go install github.com/davlgd/mkv-cli/
```

## Remove

```sh
make clean
```

or

```sh
make uninstall
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
