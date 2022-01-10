# ![rainbow](doc/rainbow-chancery.png)

![logo](doc/small.png) | Dashboard for DeFi options trading highlighting market opportunities to ease strategies decision. With Rainbow, users can scan CEX/DEX markets to fetch available options data. <br><br> This project has been initiated by [Teal.Finance](https://teal.finance/)  during the [Ethereum's EthGlobal Hackathon](https://showcase.ethglobal.com/ethonline2021/rainbow), [Solana's Ignition Hackathon](https://devpost.com/software/rainbow-ai5p7m) and [Encode Club's Hack DeFi](https://www.encode.club/hack-defi).
---|---

## Live Demo

Rainbow is live at <https://teal.finance/rainbow/>
Please gives some feedback via  GitHub issue or contact us at Teal.Finance[at]protonmail.com or using Twitter [@TealFinance](https://twitter.com/TealFinance).

## Motivations

**Rainbow** was inspired by the following trends:

* Crypto-assets and DeFi becoming mainstream,

* Advent of the Internet of Blockchains’ world  with more cross-chain applications and communications,

* Crypto options trading growing (hopefully) bigger  than the spot market, like in traditional Finance.

The crypto derivatives markets are expected to grow a lot more in the upcoming years.
Specifically, the options markets will see the biggest growth because in Finance, Option market is much bigger
than the underlying spot market. This is lagging in Crypto, when we look at Deribit's volume, the main (centralized) venue for Crypto Options, compared to Perpetuals future.

More info on our motivations: <https://cryptonuage.com/posts/internet-of-decentralized-options/>

## Current status

Supported Options exchanges:

* [Deribit](https://www.deribit.com): Centralized exchange. Main venue for crypto Options trading.

* [Opyn(Ethereum)](https://www.opyn.co/):  DeFi protocol using TheGraph & 0x.

* [PsyOptions(Solana)](https://www.psyoptions.io/): DeFi protocol build on Serum order books.

## Requirements

* Go v1.17 or later
* Node v14 or later
* Yarn

[Snap](https://en.wikipedia.org/wiki/Snap_(package_manager)) provides a simple way to install these requirements on many Linux distributions:

    snap install go   --classic
    snap install node --classic  # also installs yarn

    # check
    go   version
    yarn versions

On Debian/Ubuntu, the command `sudo apt install golang` may be for an older version.
You can check that with `apt list --all-versions golang`.
If this is your case, you may install Go v1.17 using a different package name:

    sudo apt remove  golang
    sudo apt install golang-1.17

## Build

    git clone https://github.com/teal-finance/rainbow
    cd rainbow

### CLI

You may just use the pretty nice table printed by the command `./cli`.

    go build ./cmd/cli && ./cli

![CLI screenshot](doc/cli.jpg)

### Back-end

The front-end requires the server API.

    go build ./cmd/server && ./server

### Front-end

To run the Vue3 front-end in dev mode:

    cd frontend
    yarn
    yarn dev --open

Else, in prod mode, the back-end serves the web static files from `frontend/dist`.

    cd frontend
    yarn
    yarn build

Finally open <http://localhost:8090>

## Container

This Git repo provides a [Dockerfile](Dockerfile) for a secured and light all-in-one container image: 30 MB.

The image contains the hardened sever executable (with dynamic library)
and the front-end static files.

The container enables by default the CORS, the export ports and a rate limiter.
Some of these features can be customized using environments variables.

The Dockerfile supports Docker-20 and Podman-3. The following build configures CORS with `http://localhost:1111` and backend listening on port 2222:

```
podman build --build-arg addr=http://localhost:1111 --build-arg port=2222 -t rainbow .
podman stop rainbow # if already running in background
podman run --rm -p 0.0.0.0:1111:2222 -d --name rainbow rainbow
podman logs --follow rainbow
```

Open <http://localhost:1111>.

See the comments within the [Dockerfile](Dockerfile) for more information.

## Command line flags

Rainbow provides a complete HTTP server,
including a rate limiter, an export port (Prometheus monitoring),
and more. For more details see the underlying project
[Teal.Finance/Garcon](https://github.com/teal-finance/garcon/).

```sh
$ go build ./cmd/server
$ ./server -help
Usage of ./server:
  -addr string
        Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR (default "http://localhost")
  -burst int
        Max requests during a burst, has precedence over REQ_BURST (default 10)
  -dev
        Run rainbow in dev. mode
  -exp int
        Export port for Prometheus, has precedence over EXP_PORT
  -opa string
        Policy files (comma-separated filenames) for the Open Policy Agent using the Datalog/Rego format
  -port int
        API port, has precedence over MAIN_PORT (default 8090)
  -rate int
        Max requests per minute, has precedence over REQ_PER_MINUTE (default 30)
  -www string
        Folder of the web static files, has precedence over WWW_DIR (default "frontend/dist")
```

## Makefile commands

### Build

- `make build`: build the backend
- `make buildfront`: build the frontend
- `make buildall`: build the backend and the frontend

### Run

- `make run`: run the backend
- `make runfront`: run the frontend in dev mode

## API

### /v0/options

List all the options and their order books: <http://localhost:8090/v0/options>

```js
{
  "Name": "ETH-2021-10-29 23:59:59-3200-PUT",
  "Expiry": "2021-10-29 23:59:59",
  "Type": "PUT",
  "Asset": "ETH",             // ETH, BTC, SOL
  "Strike": 3200,
  "ExchangeType": "DEX",      // CEX or DEX
  "Chain": "Solana",          // Ethereum, Solana...
  "Layer": "L1",
  "Provider": "PsyOptions",   // Opyn, Lyra, Thales, Deribit, Psyoptions
  "QuoteCurrency": "USDC" // ETH, BTC...
  "Bid": [
    {
      "Price": 13.3,
      "Quantity": 5,
    },
    {
      "Price": 13.1,
      "Quantity": 10,
    }
  ],
  "Ask": [
    {
      "Price": 15.12,
      "Quantity": 5,
    },
    {
      "Price": 15.25,
      "Quantity": 9,
    }
  ]
}
```

### /v0/options/cp

List the options in Call/Put format: <http://localhost:8090/v0/options/cp>

Rainbow API is currently only used by its web front-end and has been influenced by
the [BFF pattern](https://blog.bitsrc.io/e4fa965128bf) (Backend for Frontend pattern).
This endpoints aims simplifying the front-end processing.

```js
{ "rows":[
    { "asset": "ETH",
      "expiry": "Dec 31 08:00",
      "provider": "Deribit",
      "call": { "bid": {"px": 1805, "size":  24},
                "ask": {"px": 1830, "size": 459}},
      "strike": 4400,
      "put":  { "bid": {"px": 1305, "size":  26},
                "ask": {"px": 1330, "size":  37}}
    },

    { "asset": "ETH",
      "expiry": "Dec 31 08:00",
      "provider": "Deribit",
      "call": { "bid": {"px": 1140, "size": 258},
                "ask": {"px": 1160, "size":  33}},
      "strike": 5200,
      "put":  { "bid": {"px": 1235, "size":  80},
                "ask": {"px":    0, "size":   0}}
    },

    // ...

    { "asset": "BTC",
      "expiry": "Nov 26 08:00",
      "provider": "Deribit",
      "call": { "bid": {"px": 2045, "size":  7.5},
                "ask": {"px": 2105, "size":  5.3}},
      "strike":50000,
      "put":  { "bid": {"px":  125, "size": 27.2},
                "ask": {"px":  135, "size": 70.2}}
    }
]}
```
