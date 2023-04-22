package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; ; i++ {
		fmt.Print(i, " ")
		callApi()
		time.Sleep(1 * time.Second)
	}
}

func callApi() {
	values := map[string]string{
		"wallet_id":  "0xed6a2dd9a563a3a682638fa543f63fb59bc9800a",
		"r":          "0x410cff6bfcf682d273db8bb957ab19dace641912212140e002996349ed4f5ead",
		"s":          "0x61e8c3e31d5423c39c50c8986144cf794b7df49f6915419c9f8ed0a039289503",
		"v":          "0x1b",
		"numberbonk": "29",
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.arbmario.io/hit", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
