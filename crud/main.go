package main

import (
	"encoding/json" // Package for encoding and decoding JSON
	"fmt"           // Package for formatted I/O operations
	"io"            // Package for I/O primitives
	"net/http"      // Package for HTTP client and server implementations
	"strings"       // Package for string manipulation
)

// Todo represents a task with various properties
// The `json:"xyz"` tags tell the JSON encoder/decoder how to map JSON fields to struct fields
type Todo struct {
	UserID    int    `json:"userId"`    // ID of the user who owns this todo
	Id        int    `json:"id"`        // Unique identifier for this todo
	Title     string `json:"title"`     // Title/description of the todo
	Completed bool   `json:"completed"` // Whether the todo is completed or not
}

// performGetRequest demonstrates how to make an HTTP GET request to fetch a todo item
func performGetRequest() {
	// Make a GET request to the JSONPlaceholder API to get a todo with ID 1
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	// Check if there was an error making the request
	if err != nil {
		fmt.Println("An Error occured ", err)
		return
	}

	// Ensure we close the response body when this function exits
	// defer makes this happen at the end of the function
	defer res.Body.Close()

	// Check if the response status code is what we expect (200 OK)
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", res.StatusCode)
		return
	}

	// Read all data from the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Some error occured ", err)
		return
	}

	// Print the raw response data as a string
	fmt.Println("The raw response data: ", string(data))

	// Create a variable to hold our Todo
	var todo Todo

	// Unmarshal (parse) the JSON data into our Todo struct
	// We pass &todo to give the function access to our variable's memory address
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the Todo struct
	fmt.Println(todo)
}

// performPostRequest demonstrates how to make an HTTP POST request to create a new todo
func performPostRequest() {
	// Create a new Todo that we want to send to the server
	todo := Todo{
		UserID:    23,       // Set the user ID
		Title:     "Roudra", // Set the title
		Completed: true,     // Mark it as completed
	}
	// Note: We don't set ID because the server will assign one for us

	// Convert our Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		println("Some error happened ", err)
		return
	}

	// Convert the JSON byte array to a string (optional step for demonstration)
	jsonString := string(jsonData)

	// Create a reader from our JSON string, which is required for the http.Post function
	jsonReader := strings.NewReader(jsonString)

	// URL we want to send our POST request to
	url := "https://jsonplaceholder.typicode.com/todos"

	// Send the POST request with our JSON data
	// The second parameter specifies that we're sending JSON data
	response, err := http.Post(url, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Some Error happened", err)
		return
	}

	// Ensure we close the response body when done
	defer response.Body.Close()

	// Read the response from the server
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response we got from the server
	// The server typically returns the created Todo with an assigned ID
	fmt.Println("Response:", string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23,             // Set the user ID
		Title:     "Roudra Codes", // Set the title
		Completed: true,           // Mark it as completed
	}

	// Convert our Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		println("Some error happened ", err)
		return
	}

	// Convert the JSON byte array to a string (optional step for demonstration)
	jsonString := string(jsonData)

	// Create a reader from our JSON string, which is required for the http.Post function
	jsonReader := strings.NewReader(jsonString)

	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, url, jsonReader)

	if err != nil {
		fmt.Println("Some error happened ", err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))
}

func performDeleteRequest() {

	// For DELETE requests, we typically don't need to send a body
	// We just need the URL of the resource to delete
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Create a new DELETE request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))
}

// The main function - program execution starts here
func main() {
	fmt.Println("Learning web services")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
