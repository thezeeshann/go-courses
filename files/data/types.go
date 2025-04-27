package data

import "fmt"

type distance float64
type distanceKm float64

func (miles distance) ToKm() distanceKm {
	return distanceKm(1.666 * miles)
}

func Test() {
	// d := Distance(32.1) // distance constructor
	// fmt.Println(d)
	d := distance(4.5)
	mk := d.ToKm()
	fmt.Println(mk)
}
