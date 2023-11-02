package helper

import "fmt"

func AddWeightedMultiplier(PointChangePreWeight int, WeightPoints int) int {
	var FinalPointChange float64
	FinalPointChange = float64((PointChangePreWeight * 10)) / float64((WeightPoints + PointChangePreWeight))
	fmt.Println(FinalPointChange, "Winenrs final point change+")
	FinalPointChangeInt := int(FinalPointChange)
	return FinalPointChangeInt
}
