# crypticker-cli

A simple command line program to get current cryptocurrency data.

<div style="text-align:center">
  <img src="https://i.imgur.com/FYMT4Vm.png" alt="crypticker">
</div>

### Running

To run first clone this repo `git clone https://github.com/lukakerr/crypticker-cli.git`.

Then `cd` into `crypticker-cli` and run `go run crypticker.go`.

### Building

To build `crypticker.go` run `go build crypticker.go`.

### Install

To install crypticker-cli, simply copy the `crypticker` binary to `/usr/local/bin`, or run `./install.sh` inside the repository which does this automatically.

### Usage

Run `crypticker` to get the top 10 coins.

Run `crypticker 20` to get the top 20 coins. This value can be changed to get more or less coin data.

### To Do

- Add support for multiple currencies
- Add support for sorting by column