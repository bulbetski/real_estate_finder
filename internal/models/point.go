package models

import (
	"fmt"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func NewPoint(lat float64, lng float64) *Point {
	return &Point{
		Lat: lat,
		Lng: lng,
	}
}

func (p *Point) GetLatitude() float64 {
	return p.Lat
}

func (p *Point) GetLongitude() float64 {
	return p.Lng
}

func Midpoint(points []Point) (*Point, error) {
	if points == nil || len(points) != 2 {
		return &Point{}, fmt.Errorf("incorrect number of points")
	}
	return calculateMidpoint(points[0], points[1]), nil
}

func calculateMidpoint(p1, p2 Point) *Point {
	return NewPoint((p1.Lat+p2.Lat)/2, (p1.Lng+p2.Lng)/2)
}
