package core

import (
	"errors"
)

var Grades = map[string]int{
	"vB":   1,
	"v0":   2,
	"v1":   3,
	"v2":   4,
	"v3":   5,
	"v4":   6,
	"v5":   7,
	"v6":   8,
	"v7":   9,
	"v8":   10,
	"v9":   11,
	"v10+": 12,
}

type Route struct {
	Grade    string `json:"grade"`
	Finished bool   `json:"finished"`
}

func (r Route) PointValue() (int, error) {
	// TODO: implement logic for if failed route
	val, ok := Grades[r.Grade]
	if ok {
		return val, nil
	}
	return -1, errors.New("invalid route grade")
}

type Input struct {
	Seconds int     `json:"seconds" form:"seconds"`
	Routes  []Route `json:"routes" form:"routes"`
}

func (i Input) TotalPoints() (int, error) {
	var routePoints []int
	for _, route := range i.Routes {
		val, err := route.PointValue()
		if err != nil {
			return -1, err
		}
		routePoints = append(routePoints, val)
	}
	return sum(routePoints), nil
}

func (i Input) AverageDifficulty() (float64, error) {
	var routePoints []int
	for _, route := range i.Routes {
		val, err := route.PointValue()
		if err != nil {
			return -1, err
		}
		routePoints = append(routePoints, val)
	}
	return mean(routePoints), nil
}

type Output struct {
	TotalTime         int     `json:"total_time"`
	TotalPoints       int     `json:"total_points"`
	NumRoutes         int     `json:"number_of_routes"`
	AverageDifficulty float64 `json:"average_difficulty"`
	Routes            []Route `json:"routes"`
	Score             float64 `json:"score"`
}

func (i Input) Compute() (Output, error) {
	points, err := i.TotalPoints()
	if err != nil {
		return Output{}, err
	}
	difficulty, err := i.AverageDifficulty()
	if err != nil {
		return Output{}, err
	}
	o := Output{
		TotalTime:         i.Seconds,
		TotalPoints:       points,
		NumRoutes:         len(i.Routes),
		AverageDifficulty: difficulty,
		Routes:            i.Routes,
	}
	return o, nil
}

func NewGame() (Input, error) {
	return Input{}, nil
}
