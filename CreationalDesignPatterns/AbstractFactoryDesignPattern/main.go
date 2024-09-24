package main

import "fmt"

func main() {

	tri := GetShapeFactory(3)
	quad := GetShapeFactory(4)

	s1 := tri.GetShape(EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s1.GetName(), " PARAMETER:", s1.GetParameter())
	s1.GetFeatures()

	s2 := tri.GetShape(EQUAL_NOT, []int{4, 3, 2})
	fmt.Println("SHAPE:", s2.GetName(), " PARAMETER:", s2.GetParameter())
	s2.GetFeatures()

	s3 := tri.GetShape(EQUAL_TWO_SIDES, []int{4, 5})
	fmt.Println("SHAPE:", s3.GetName(), " PARAMETER:", s3.GetParameter())
	s3.GetFeatures()

	s4 := tri.GetShape(ANGLE_90, []int{4, 3, 5})
	fmt.Println("SHAPE:", s4.GetName(), " PARAMETER:", s4.GetParameter())
	s4.GetFeatures()

	s5 := quad.GetShape(EQUAL_ALL_Angle_90, []int{4})
	fmt.Println("SHAPE:", s5.GetName(), " PARAMETER:", s5.GetParameter())
	s5.GetFeatures()

	s6 := quad.GetShape(EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s6.GetName(), " PARAMETER:", s6.GetParameter())
	s6.GetFeatures()

	s7 := quad.GetShape(EQUAL_TWO_SIDES_ANGLE_90, []int{4, 5})
	fmt.Println("SHAPE:", s7.GetName(), " PARAMETER:", s7.GetParameter())
	s7.GetFeatures()

	s8 := quad.GetShape(EQUAL_TWO_SIDES, []int{4, 5})
	fmt.Println("SHAPE:", s8.GetName(), " PARAMETER:", s8.GetParameter())
	s8.GetFeatures()

	s9 := quad.GetShape(EQUAL_NOT, []int{1, 2, 3, 4})
	fmt.Println("SHAPE:", s9.GetName(), " PARAMETER:", s9.GetParameter())
	s9.GetFeatures()

}
