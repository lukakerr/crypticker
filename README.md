# crypticker-cli

A simple command line program to get current cryptocurrency data.

<div style="text-align:center">
  <img src="https://i.imgur.com/lU9w7Ba.png" alt="crypticker">
</div>

### Install

To install automatically run:

```
curl https://raw.githubusercontent.com/lukakerr/crypticker-cli/master/install.sh | sh
```

This will execute `install.sh` and place the `crypticker` binary in `/usr/local/bin`.

After this, you are able to run the crypticker-cli by executing `crypticker`.

### Usage

```
crypticker      // Get top 10 coins
crypticker 20   // Pass an integer to get more coin data
```

### Running
 
```
git clone https://github.com/lukakerr/crypticker-cli.git
cd crypticker-cli
go run crypticker.go
```

### Building

```
go build crypticker.go
```

### To Do

- Add support for multiple currencies
- Add support for sorting by column
