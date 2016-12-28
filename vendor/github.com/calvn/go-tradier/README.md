# go-tradier

[![Go Report Card](https://goreportcard.com/badge/github.com/calvn/go-tradier?style=flat-square)](https://goreportcard.com/report/github.com/calvn/go-tradier)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)                                           ](https://godoc.org/github.com/calvn/go-tradier/tradier)

Golang library for interacting with the [Tradier API](https://developer.tradier.com/documentation/)

***Note:*** *This library is still under development - use with discretion!*

## Authentication
go-tradier does not directly handle authentication. However, it uses `http.Client`, so authentication can be done by passing an `http.Client` that can handle authentication. For instance, you can use the [oauth2](https://github.com/golang/oauth2) library to achieve proper authentication. For a full working example, refer to the `examples/` directory.

```go
import "golang.org/x/oauth2"

func main() {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := tradier.NewClient(tc)

  // Returns the profile of the user
  profile, _, err := client.User.Profile()
}
```

## License

This library is licensed under the MIT License as provided in [here](LICENSE.md).

*Made with <3 in Go. Heavily borrowed from and influenced by Google's go-github library*
