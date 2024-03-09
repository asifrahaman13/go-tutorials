package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `jsonn:"courseName"`
	CoursePrice int     `jsonn:"price"`
	Author      *Author `jsonn:"author"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers.
//serve home route.

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Learning code online.</h1>"))
}

// Function to get all the courses.
func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Function to get one course.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course...")
	w.Header().Set("Content-Type", "applications/json")

	// Grab the id from the requests
	params := mux.Vars(r)

	// Loop throgh the courses. Find matching id and eturn the response.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id.")
	return
}

// Function to insert data
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course...")
	w.Header().Set("Content-Type", "applications/json")

	// What if the body is empty.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data is found inside the json.")
	}

	// Generate a unique id, string.
	// Append course into the courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// Update the course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop, id, remove, add with my id.
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			fmt.Println("Course updated")
			return
		}
	}
}

// API to delete the courses.
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop, id, remove operations (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

// Middleware and helpers -file
func main() {
	fmt.Println("API - Learn go lang code online.")

	// Create a router
	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", CoursePrice: 299, Author: &Author{Fullname: "Cameron Green", Website: "green.dev"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "Go lang", CoursePrice: 2419, Author: &Author{Fullname: "Marshall D Teach", Website: "green.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// Listen your server into the PORT=4000
	log.Fatal(http.ListenAndServe(":4000", r))
}
