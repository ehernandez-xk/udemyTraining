package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//JSONObject it is the json object
type JSONObject struct {
	Port              string
	DebugMode         bool
	ReplicationAgents []URL
	Dispatchers       []URL
}

//URL the server url
type URL struct {
	URL  string
	Port string
}

func main() {

	myJSON := JSONObject{}

	file, er := ioutil.ReadFile("./conf.json")
	if er != nil {
		fmt.Println("Error json: ", er)
	}

	json.Unmarshal(file, &myJSON)

	//The object contains all json info.
	fmt.Println(myJSON.Port)
	fmt.Println(myJSON.Dispatchers)
}
