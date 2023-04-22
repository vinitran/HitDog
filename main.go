package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	for i := 0; ; i++ {
		callApi()
		callApi2()
	}
}

type Data struct {
	User struct {
		ID             string `json:"_id"`
		WalletID       string `json:"wallet_id"`
		R              string `json:"r"`
		S              string `json:"s"`
		V              string `json:"v"`
		Score          int    `json:"score"`
		LastTimeUpdate int64  `json:"lastTimeUpdate"`
	} `json:"user"`
	Totalpoint int `json:"totalpoint"`
	PointHonk  int `json:"point_honk"`
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
		fmt.Println(err)
		return
	}

	resp, err := http.Post("https://api.arbmario.io/hit", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("err", err)
		return
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return
	}

	fmt.Println(data.User.Score)
}

func callApi2() {
	values := map[string]string{"wallet_id": "0x4bff03efff59e0782517d9fc3e3af196254f1dff", "r": "0x3ca1ab5241b8686ed64cd9871f6c36c42c470163acd4042a74b504eef06268b5", "s": "0x285cd859d5824c93d19280feacd850938b1d0954769545b9e2f36ea3e96b8e2e", "v": "0x1b", "numberbonk": "18"}

	json_data, err := json.Marshal(values)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("https://api.arbmario.io/hit", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("err", err)
		return
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return
	}

	fmt.Println(data.User.Score)
}
