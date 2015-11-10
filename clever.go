package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://api.clever.com/v1.1/sections?limit=10000", nil)
	req.Header.Set("Authorization", "Bearer DEMO_TOKEN")
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var f interface{}
	var students = 0
	err := json.Unmarshal(body, &f)

	if err == nil {
		a := f.(map[string]interface{})
		b := a["data"].([]interface{})
		var i = 0
		for i = 0; i < len(b); i++ {
			c := b[i].(map[string]interface{})
			d := c["data"].(map[string]interface{})
			e := d["students"].([]interface{})
			students += len(e)
			fmt.Printf("Adding %v more students. Total now %v\n", len(e), students)
		}
		var av = (students/i + 1)
		fmt.Printf("Average students per section is %v\n", av)
	}

}
