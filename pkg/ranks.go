package pkg

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var Grades = []string{
	"vB",
	"v0",
	"v1",
	"v2",
	"v3",
	"v4",
	"v5",
	"v6",
	"v7",
	"v8",
	"v9",
	"v10+",
}

type Input struct {
	Seconds int  `json:"time"`
	Grades  []string `json:"grades"`
}

type Output struct {
	Seconds    int    `json:"time"`
	AvgGrade   string `json:"average_grade"`
	NumRoutes  int    `json:"num_routes"`
	PointValue int    `json:"point_value"`
}

func (i Input) Compute() (Output, error) {
	// this actually needs to take a list of grades that were accomplished
	// and the number of routes done
	// TODO: modify to output AvgGrade and NumRoutes
	if valid := stringInSlice(i.Grade, Grades); valid {
		// now we know grade is valid
		ParseGrade
		pointVal := ParseGrade(i.Grade)
		o := Output{Seconds: i.Seconds, Grade: i.Grade, PointValue: pointVal}
		return o, nil
	} else {
		return Output{}, errors.New("Invalid Input field `Grade`.")
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ParseGrade(grade string) int {
	if grade == "v10+" {
		return 10
	} else if grade == "vB" {
		return 0
	} else {
		targetDigit := string(grade[len(grade)-1])
		fmt.Println(targetDigit, reflect.TypeOf(targetDigit))
		number, err := strconv.Atoi(targetDigit)
		Check(err)
		return number + 1
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
