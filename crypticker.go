package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Coin struct {
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	PriceUSD     string `json:"price_usd"`
	DayVolume    string `json:"24h_volume_usd"`
	PctChange1H  string `json:"percent_change_1h"`
	PctChange24H string `json:"percent_change_24h"`
}

type Coins []Coin

const RED = "\x1b[31;1m"
const GREEN = "\x1b[32;1m"
const YELLOW = "\x1b[33;1m"
const RESET = "\x1b[0;1m"

func main() {
	var limit string
	var coin string

	flag.StringVar(&limit, "l", "10", "Set the limit of coins to get")
	flag.StringVar(&coin, "c", "", "Get data for an individual coin")

	flag.Parse()

	// Escape variables for URL
	safeLimit := url.QueryEscape(limit)
	safeCoin := url.QueryEscape(coin)

	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s?limit=%s", safeCoin, safeLimit)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Coins

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	// Table data
	data := [][]string{}
	for _, element := range record {

		pct1H, err := strconv.ParseFloat(element.PctChange1H, 64)
		pct24H, err := strconv.ParseFloat(element.PctChange24H, 64)

		if err != nil {
			fmt.Println(err)
		}

		if pct1H < 0 {
			element.PctChange1H = RED + element.PctChange1H + "%" + RESET
		} else {
			element.PctChange1H = GREEN + element.PctChange1H + "%" + RESET
		}

		if pct24H < 0 {
			element.PctChange24H = RED + element.PctChange24H + "%" + RESET
		} else {
			element.PctChange24H = GREEN + element.PctChange24H + "%" + RESET
		}

		data = append(data, []string{
			element.Rank,
			element.Symbol,
			element.PriceUSD,
			element.DayVolume,
			element.PctChange1H,
			element.PctChange24H,
		})
	}

	// Table header names
	names := []string{
		YELLOW + "Rank" + RESET,
		YELLOW + "Coin" + RESET,
		YELLOW + "Price (USD)" + RESET,
		YELLOW + "Volume (24H)" + RESET,
		YELLOW + "Change (1H)" + RESET,
		YELLOW + "Change (24H)" + RESET,
	}

	// Build table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(names)
	table.SetCenterSeparator("─")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range data {
		table.Append(v)
	}

	// Render table
	fmt.Print("\n")
	table.Render()
	fmt.Print("\n")
}
