# TODO
Laundry list

## commands
- [x] config
- [x] info
- [x] quote
- [x] buy
- [x] sell
- [x] cancel
- [x] orders
- [x] positions
- [ ] buy open
- [ ] sell open
- [ ] buy close
- [ ] sell close


### config
- [x] Initialize config file
- [x] Merge existing config file
- [x] Merge flag, env, and config file settings in priority
- [x] Properly use and load environment variables
- [ ] ~~BROKR_ACCESS_TOKEN should be globally available~~
    - Used for all commands that require auth
- [x] fix brokerage interface to take in appropriate config

#### defaults
- [x] Subcommands should print help if no other values provided
- [x] Ability to toggle preview-order (default should be set to true)
- [x] Attempt to set default account if there is only one account number
- [ ] Set sane defaults for placing orders
  - [ ] Day expiry by default
  - [x] Preview order *enabled* by default
- [ ] Errors from commands should not return back `--help`

### quotes
- [x] Basic functionality
- [x] Render from template
- [x] Correctly display time as non-epoch values
- [x] Correctly set precision to two decimal values
- [x] Better output view
  - [x] Possibly render from template
- [ ] Colorize/symbolize delta changes
- [ ] Simple quote by default (symbol, price, change, date)
  - [ ] Verbose option to output more info

### info
- [ ] Display current remaining rate limits
- [x] Display current configuration, shortcut to `cat ~/.brokr.yaml`

### account
- [ ] Display all accounts under the user
  - [ ] `$ brokr accounts`
  - [ ] Mark account that is currently in use
- [ ] Display account information
  - [ ] `$ brokr account 123456`
  - [ ] Account number (account type)
  - [ ] Account value
  - [ ] Buying power (Stock/Option)
  - [ ] Cash
  - [ ] Stock value (Long/Short)
  - [ ] Option value (Long/Short)
- [ ] Change to use `account #`
  - [ ] `$ brokr account use 123456`

### orders/order/cancel
- [x] Display all orders
  - `$ brokr orders`
- [ ] Display pending orders
  - `$ brokr orders pending`
- [ ] Display order 812
  - `$ brokr order 812`
- [ ] Cancel order 812
  - `$ brokr order cancel 812`

### positions
`$ brokr positions`

### version
- [x] Version for more verbose info (e.g. go version, built time, os/arch)

## placing orders
- [ ] Support for equities
- [ ] Support for single-leg options
- [ ] Support for multi-leg options
- [ ] Provide subcommand flags for overwriting defaults (global)
  - [ ] `--duration={day,gtc}`, `-d`
  - [ ] `--preview={true,false}`, `-p`
- [ ] Ability to convert a short to sell if there is a position
- [ ] Ability to close a position regardless of side
- [ ] Ability to place limit and stop based on '>' '<'

- [ ] Order commands
  - [ ] equities
    - [ ] buy - b
    - [ ] sell - s
    - [ ] short -- ss
    - [ ] cover - c
  - [ ] options
    - [ ] open buy - ob
    - [ ] close sell - cs
    - [ ] open sell - os
    - [ ] close buy - cb
  - [ ] close

#### Example:

`$ brokr buy 100 appl limit 120` or `$ brokr buy 100 aapl < 120`
> buy 100 shares of AAPL with limit set at $120, day duration

`$ brokr b 100 aapl l 120`
> potentially provide the ability to shorten command

`$ brokr sell 100 appl limit 120`
> buy 100 shares of AAPL with limit set at $120, day duration

`$ brokr s 100 aapl l 120`
> potentially provide the ability to shorten command

`$ brokr short aapl stop 120 limit 100`
> short 100 shares of aapl with stop at $120 and limit at $100

`$ brokr buy open 100 AAPL161223C00115000`
> Single-leg order: buy-to-open 100 AAPL 115@Dec-23-2016 call contracts

`$ brokr bo AAPL161223C00115000`
> shorter command

`$ brokr sell close aapl@115 161223`
> potentially make it easier to place options orders without knowing options symbol

`$ brokr buy open 100 CSCO150117C00035000 sell open 100 CSCO140118C00008000`
> Multi-leg order: buy-to-open 100 CSCO 35@Jan-17-2015 call contracts and sell-to-open 100 CSCO 35@Jan-18-2014 call contracts

`$ brokr close aapl` or `$brokr c aapl` Limit/Stop: `$ brokr close aapl limit 100`
> Close all AAPL positions (excluding options)


***Note:*** *`open/close`comes before `buy/sell` to explicitly differentiate an option order from an equity order*

Should there be only buy/sell on orders? For instance, *sell short* turns into *sell* if there is a long position in the account. Same with *buy-to-cover* into *buy*. It will need to make an additional call in this case to make sure that there is an existing position in the account (possibly through preview).

## cli behavior
- [ ] Interactive/continuous command
- [ ] better error handling other than using `fmt.Println()`

## docs
- [ ] Write up docs for usage

## Tradier API issues

### POST ``/v1/accounts/{account_id}/orders`
Requests on a non-existent `account_id` returns application/xml regardless of whether -H "Accept: application/json" is provided.
