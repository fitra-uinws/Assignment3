package main

import (
	"assignment3/router"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type WaterAndWindStatus struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var waterAndWindStatus WaterAndWindStatus

func main() {
	go GenerateRandomValuedJson()

	r := router.StartApp()
	r.Run(":8080")
}

func GenerateRandomValuedJson() {
	waterAndWindStatus.Water = rand.Intn(100) + 1
	waterAndWindStatus.Wind = rand.Intn(100) + 1

	waterAndWindStatusJson, err := json.Marshal(map[string]interface{}{"status": &waterAndWindStatus})
	if err != nil {
		fmt.Println("error :", err.Error())
	}

	_ = ioutil.WriteFile("json_file.json", waterAndWindStatusJson, 0644)

	time.Sleep(15 * time.Second)
	GenerateRandomValuedJson()
}
