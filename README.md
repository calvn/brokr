# brokr

brokr is a CLI application that allows you to manage your Tradier brokerage accounts, including the ability to place orders, view pending orders and existing positions, and get real-time quotes.

```
$ brokr
brokr - bringing your trades into the console.
  It currently supports making trades against Tradier.

Made with ♥︎ in Go.

Usage:
  brokr [flags]
  brokr [command]

Available Commands:
  buy         Preview or place a buy order
  cancel      Cancel a pending order
  config      Configure .brokr.yaml
  info        Show information about brokr settings
  orders      Get pending orders for an account
  positions   Get positions for an account
  quote       Get quotes from a set of symbols
  sell        Preview or place a sell order
  version     Display detailed brokr version information

Flags:
  -v, --version   Print version and exit

Use "brokr [command] --help" for more information about a command.
```

## Features and functionality

All of the commands listed in `brokr --help` should be available for use.

brokr is currently limited to buy and selling stocks. The ability to trade options will be implemented in the near future.

## Development

- Go 1.8+
- Glide
