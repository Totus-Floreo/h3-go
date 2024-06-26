package h3_test

import (
	"fmt"

	"github.com/Totus-Floreo/h3-go"
)

func ExampleLatLngToCell() {
	latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
	resolution := 9
	c := h3.LatLngToCell(latLng, resolution)
	fmt.Printf("%s", c)
	// Output:
	// 8928308280fffff
}
