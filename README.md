# brokr

brokr is a CLI application that allows you to manage your brokerage account, including the ability to place orders, view pending orders and existing positions, and get real-time quotes.

brokr currently only supports managing Tradier brokerage accounts, but it has been designed from the beginning to support implementing other brokerages.

![brokr Demo GIF](assets/static/brokr_demo.gif)

## Installing

```sh
$ go get github.com/calvn/brokr
$ cd $GOPATH/src/github.com/calvn/brokr
$ make build
```

Compiled binaries will be available once brokr is in a relatively stable release.

## Getting started

To get started, configure brokr using `brokr config tradier --account <ACCOUNT ID> --token <ACCESS TOKEN>`. Your `.brokr.yaml` will look something like this:

```yml
brokerage: tradier
preview_order: true
tradier:
  account: <ACCOUNT ID>
  access_token: <ACCESS TOKEN>
```

After tradier configuration has been set, you can start interacting with your account.

### Sane defaults

Since brokr currently only supports Tradier, it will be brokr's default brokerage.

brokr will also set `preview_order` to `true` to avoid accidental order placement.


## Limitations

brokr is currently limited to buy and selling stocks. The ability to trade options will be implemented in the near future.

brokr currently does not have short-sell and buy-to-cover implemented. This feature will be added in a future release.


## Contributing

Contributions are always welcomed. To contribute, fork the repository, make the necessary changes, and send in a pull request back to master. Opening an accompanied issue is strongly recommended.

### Development dependencies

- Go 1.8+
- Glide
