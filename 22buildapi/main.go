package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake database
var courses []Course

//middleware/helper
func (c *Course) isEmpty() bool {
	return  c.CourseName==""
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API series in Golang</h1>"))
}


func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All course with API")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course with API")
	w.Header().Set("Content-Type", "application/json")

	//get that specific id
	params := mux.Vars(r)

	//looping through the course, matching the id and returning back the value
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No id given!")
	
}
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")

	//what if there is no json data
	if r.Body == nil{
		json.NewEncoder(w).Encode("There is no data")
		return
	} 

	//what if some one send this empty json {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	//params 
	params := mux.Vars(r)

	//loop, match id, remove that id, inject new id
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index + 1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("deleting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove that index
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index + 1:]...)
			break;
		}
	}
}

func main() {
  fmt.Println("hello and welcome to api")
}