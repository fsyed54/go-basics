package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct { //Struct
	Title string `json:"title"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos") //Sending GET request
	if err != nil{
		fmt.Println("Error with GET request:", err)
		return //Exit function
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) //Reading response body into a byte slice
	if err != nil{
		fmt.Println("Error reading response body:", err)
		return
	}

	var todos []Todo
	if err := json.Unmarshal(body, &todos); err != nil { //Unmarshal JSON response into a slice of Todo structs
		fmt.Println("Error JSON", err)
		return
	}

	file, err := os.Create("titles.txt") //Create or open a text file
	if err != nil{
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	for _, todo := range todos { //Iterate over todos slice
		_, err := file.WriteString(todo.Title + "\n") //Write titles to the text file

		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}
	fmt.Println("Titles have been written to the text file") //Successful case

}