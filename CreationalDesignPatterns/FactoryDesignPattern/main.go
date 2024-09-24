package main

import "fmt"

func main() {

	s0 := shapeFactory(0, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s0.GetName(), " PARAMETER:", s0.GetParameter())
	s0.GetFeatures()

	s1 := shapeFactory(3, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s1.GetName(), " PARAMETER:", s1.GetParameter())
	s1.GetFeatures()

	s2 := shapeFactory(3, EQUAL_NOT, []int{4, 3, 2})
	fmt.Println("SHAPE:", s2.GetName(), " PARAMETER:", s2.GetParameter())
	s2.GetFeatures()

	s3 := shapeFactory(3, EQUAL_TWO_SIDES, []int{4, 5})
	fmt.Println("SHAPE:", s3.GetName(), " PARAMETER:", s3.GetParameter())
	s3.GetFeatures()

	s4 := shapeFactory(3, ANGLE_90, []int{4, 3, 5})
	fmt.Println("SHAPE:", s4.GetName(), " PARAMETER:", s4.GetParameter())
	s4.GetFeatures()

	s5 := shapeFactory(4, EQUAL_ALL_Angle_90, []int{4})
	fmt.Println("SHAPE:", s5.GetName(), " PARAMETER:", s5.GetParameter())
	s5.GetFeatures()

	s6 := shapeFactory(4, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s6.GetName(), " PARAMETER:", s6.GetParameter())
	s6.GetFeatures()

	s7 := shapeFactory(4, EQUAL_TWO_SIDES_ANGLE_90, []int{4, 5})
	fmt.Println("SHAPE:", s7.GetName(), " PARAMETER:", s7.GetParameter())
	s7.GetFeatures()

	s8 := shapeFactory(4, EQUAL_TWO_SIDES, []int{4, 5})
	fmt.Println("SHAPE:", s8.GetName(), " PARAMETER:", s8.GetParameter())
	s8.GetFeatures()

	s9 := shapeFactory(4, EQUAL_NOT, []int{1, 2, 3, 4})
	fmt.Println("SHAPE:", s9.GetName(), " PARAMETER:", s9.GetParameter())
	s9.GetFeatures()

	s10 := shapeFactory(5, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s10.GetName(), " PARAMETER:", s10.GetParameter())
	s10.GetFeatures()

	s11 := shapeFactory(6, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s11.GetName(), " PARAMETER:", s11.GetParameter())
	s11.GetFeatures()

	s12 := shapeFactory(8, EQUAL_ALL, []int{4})
	fmt.Println("SHAPE:", s12.GetName(), " PARAMETER:", s12.GetParameter())
	s12.GetFeatures()
}
