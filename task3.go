package main

import (
	"encoding/json"
	"errors"
)

func multiplyByTwo(k *int) error {
	*k *= 2
	return nil
}

func printMoreTen(k int64) error {
	if k < 10 {
		return errors.New("k must be > 10")
	}
	return nil
}

func dejson() error {
	type jsStruct struct {
		Data int  `json:"data"`
		OK   bool `json:"ok"`
	}
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	err := json.Unmarshal(in, &out)
	if err != nil {
		return err
	}
	return json.Unmarshal(in, &out)
}

func main() {
	var r int = 11
	if err := printMoreTen(int64(r)); err != nil {
		panic(err)
	}
}