# crypticker-cli

A simple command line program to get current cryptocurrency data.

<div style="text-align:center">
  <img src="https://i.imgur.com/lU9w7Ba.png" alt="crypticker">
</div>

### Install

To install automatically run:

`curl https://raw.githubusercontent.com/lukakerr/crypticker-cli/master/install.sh | sh`

This will execute `install.sh` and place the `crypticker` binary in `/usr/local/bin`.

After this, you are able to run the crypticker-cli by executing `crypticker`.

### Usage

Run `crypticker` to get the top 10 coins.

Run `crypticker 20` to get the top 20 coins. This value can be changed to get more or less coin data.

### Running

To run first clone this repo `git clone https://github.com/lukakerr/crypticker-cli.git`.

Then `cd` into `crypticker-cli` and run `go run crypticker.go`.

### Building

To build `crypticker.go` run `go build crypticker.go`.

### To Do

- Add support for multiple currencies
- Add support for sorting by column