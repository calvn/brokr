# TODO

## General

- [ ] Rate limiting tracking on the client
- [ ] Determine best type (probably not float64) for holding currency
- [ ] Handle errors gracefully
  - [x] `json: Invalid access token`
  - [ ] `json: cannot unmarshal number 0E-8 into Go value of type int` on `LastFillQuantity` and such
  - [ ] Better error handling on `MarshalJSON` and `UnmarshalJSON`
- [ ] Change slice of objects to slice of pointers on structs
- [ ] Godocs
- [ ] Fix time zone. Get:"2016-08-05T17:24:34Z" | Want:"2016-08-05T17:24:34.000Z"
- [ ] Fix quotes to markets/quotes
- [ ] Fix OrderService to OrdersService

## Error handling

There are case where the endpoint returns 200, but there is an actual error.

For example on `/v1/accounts/{account_id}/orders`:
```
{
  "errors": {
    "error": [
      "Backoffice rejected override of the order.",
      "InitialMargin"
    ]
  }
}
```

## Potential enhancements

Structs that are basically wrappers around slices (i.e. `Orders`) could be turned directly into slices. There would have to be changes in MarshalJSON() and UnmarshalJSON() for this to work. This way the underlying slice object can be accessed directly. In other words, `Orders.Order[0]` could be turned into `Orders[0]`.

## User endpoints
|                | endpoint | test coverage | docs |
|----------------|----------|---------------|------|
| user/profile   | ✓        | ✓             |      |
| user/balances  | ✓        |               |      |
| user/positions | ✓        |               |      |
| user/history   | ✓        |               |      |
| user/gainloss  | ✓        |               |      |
| user/orders    | ✓        |               |      |

- [x] ~~Fix `OrdersAccountEntry` to dynamically map to object~~
- [x] `Order` should support indexing if it is a slice
- [x] Add a generic userRequest method that gets called by all user-related methods


## Account endpoints

|                       | endpoint | test coverage | docs |
|-----------------------|----------|---------------|------|
| account/balances      | ✓        |               |      |
| account/positions     | ✓        |               |      |
| account/history       | ✓        |               |      |
| account/gainloss      | ✓        |               |      |
| account/orders        | ✓        |               |      |
| account/orders/status | ✓        |               |      |

## Trading endpoints

|                       | endpoint | test coverage | docs |
|-----------------------|----------|---------------|------|
| order/create          | ✓        |               |      |
| order/create_multileg | ✓        |               |      |
| order/preview         | ✓        |               |      |
| order/update          | ✓        |               |      |
| order/cancel          | ✓        |               |      |

## Market Data endpoints

|                                 | endpoint | test coverage | docs |
|---------------------------------|----------|---------------|------|
| GET markets/quotes              | ✓        |               |      |
| GET markets/timesales           |          |               |      |
| GET markets/options/chains      |          |               |      |
| GET markets/options/strikes     |          |               |      |
| GET markets/options/expirations |          |               |      |
| GET markets/history             |          |               |      |
| GET markets/clock               |          |               |      |
| GET markets/calendar            |          |               |      |
| GET markets/search              |          |               |      |
| GET markets/lookup              |          |               |      |

## Fundamentals endpoints

*The fundamental/ endpoint is still in beta, so most of the values in the JSON response has not been normalized. Thus, its equivalent methods will probably not be implemented in the near future until there is stable release from Tradier, or unless there is a desire from the community to do so.*

## Watchlists endpoints

|                                | endpoint | test coverage | docs |
|--------------------------------|----------|---------------|------|
| GET watchlists                 | ✓        |               |      |
| GET watchlist/{id}             | ✓        |               |      |
| POST watchlist                 | ✓        |               |      |
| PUT watchlist                  | ✓        |               |      |
| DELETE watchlist               | ✓        |               |      |
| POST watchlists/{id}/symbols   | ✓        |               |      |
| DELETE watchlists/{id}/symbols | ✓        |               |      |

## Streaming endpoints
