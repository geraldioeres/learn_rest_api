package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name    string
	Email   string
	Address string
}

type PeopleSwapi struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
}

func main() {
	http.HandleFunc("/v1/users", GetUserController)
	fmt.Println("Server API telah berjalan di port 8000")
	http.ListenAndServe(":8000", nil)
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response, _ := http.Get("https://swapi.dev/api/people/1/")
		responseBody, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		peopleSwapi := PeopleSwapi{}
		json.Unmarshal(responseBody, &peopleSwapi)

		user := User{peopleSwapi.Name, peopleSwapi.Mass, peopleSwapi.Height}
		resultJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Gagal Convert", http.StatusInternalServerError)
			return
		}
		w.Write(resultJSON)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
