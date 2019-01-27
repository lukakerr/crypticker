# crypticker-cli

A simple command line program to get current cryptocurrency data.

<div style="text-align:center">
  <img src="https://i.imgur.com/oPsZZ2M.png" alt="crypticker">
</div>

### Install

To install automatically run:

```bash
$ curl https://raw.githubusercontent.com/lukakerr/crypticker/master/install.sh | sh
```

This will execute `install.sh` and place the `crypticker` binary in `/usr/local/bin`.

After this, you are able to run the crypticker-cli by executing `crypticker`.

### Usage

```bash
$ crypticker             # Get top 10 coins
$ crypticker -l 20       # Pass an integer to get more coin data
$ crypticker -c bitcoin  # Pass a coin name to get singular data
$ crypticker -h          # Get help
```

### Running

```bash
$ git clone https://github.com/lukakerr/crypticker.git

$ cd crypticker-cli

# Install tablewriter dependency
$ go get github.com/olekukonko/tablewriter

$ go run main.go
```

### Building

```bash
$ go build main.go
```

### To Do

- [x] Add support for individual coins
- [ ] Add support for multiple currencies
- [ ] Add support for sorting by column
