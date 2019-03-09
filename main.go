package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type User struct {
	ID        string   `json:"id,omitempty`
	Username  string   `json:"username,omitempty`
	Password  string   `json:"password,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}


// users endpoint
func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

				

	// db, err := sql.Open("postgres", "host=172.19.0.3 user=drios dbname=drios password=123 sslmode=disable")
	db, err := sqlx.Open("postgres", "host=172.19.0.3 user=drios dbname=drios password=123 sslmode=disable")
// db = sqlx.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
		log.Fatal(err)
	}
	defer db.Close()

	rows, err2 := db.Queryx("SELECT username, id, password FROM users")

	if err2 != nil {
		log.Fatal(err2)
	}

	for rows.Next() {
		var username string
		var id string
		var user User
		if e := rows.Scan(&username, &id); e != nil {

		}
		rows.StructScan(&user)
		// users = append(users, user)

		users = append(users, user)

		users = append(users, User{
				ID: user.ID, 
				Username: user.Username, 
				Firstname: "Dave", 
				Lastname: "Rios", 
				Address: &Address{City: "City Miami", State: "State FL"}})
				
	}




	json.NewEncoder(w).Encode(users)
}
func GetIndexEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to my api in golang, great.")
}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	router.HandleFunc("/api", GetIndexEndpoint).Methods("GET")
	
	router.HandleFunc("/api/users", GetUsersEndpoint).Methods("GET")

	router.HandleFunc("/api/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/api/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1528", router))
}












func main1() {
	// db, err := sql.Open("postgres", "host=172.19.0.3 user=drios password=123 sslmode=disable")
	db, err := sql.Open("postgres", "host=172.19.0.3 user=drios dbname=drios password=123 sslmode=disable")

	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
		log.Fatal(err)
	}
	defer db.Close()

	rows, err2 := db.Query("SELECT username FROM users")

	if err2 != nil {
		log.Fatal(err2)
	}

	for rows.Next() {
		var username string
		if e := rows.Scan(&username); e != nil {

		}
		fmt.Printf("username is %s\n", username)
	}

}








/*
https://gist.github.com/itang/8090953

*/
