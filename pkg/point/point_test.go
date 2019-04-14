package point

import "fmt"

func ExamplePoint_Dim() {
	p := &PointBase{
		Vec: []float64{
			1.0, 2.0, 3.0,
		},
	}

	fmt.Println(p.Dim())

	// Output:
	// 3
}

func ExamplePoint_GetValue() {
	p := &PointBase{
		Vec: []float64{
			1.0, 2.0, 3.0,
		},
	}

	fmt.Println(p.GetValue(0))
	fmt.Println(p.GetValue(1))
	fmt.Println(p.GetValue(2))

	// Output:
	// 1
	// 2
	// 3
}

func ExamplePoint_Distance() {
	p1 := &PointBase{
		Vec: []float64{
			1.0, 2.0, 3.0,
		},
	}

	p2 := &PointBase{
		Vec: []float64{
			2.0, 3.0, 4.0,
		},
	}

	p3 := &PointBase{
		Vec: []float64{
			5.0, 4.0, 3.0,
		},
	}

	fmt.Println(p1.Distance(p1))
	fmt.Println(p1.Distance(p2))
	fmt.Println(p1.Distance(p3))
	fmt.Println(p2.Distance(p2))
	fmt.Println(p2.Distance(p3))
	fmt.Println(p3.Distance(p3))

	// Output:
	// 0
	// 1.7320508075688772
	// 4.47213595499958
	// 0
	// 3.3166247903554
	// 0
}

func ExamplePoint_Equals() {
	p1 := &PointBase{
		Vec: []float64{
			1.0, 2.0, 3.0,
		},
	}

	p2 := &PointBase{
		Vec: []float64{
			2.0, 3.0, 4.0,
		},
	}

	p3 := &PointBase{
		Vec: []float64{
			5.0, 4.0, 3.0,
		},
	}

	fmt.Println(p1.Equals(p1))
	fmt.Println(p1.Equals(p2))
	fmt.Println(p1.Equals(p3))
	fmt.Println(p2.Equals(p2))
	fmt.Println(p2.Equals(p3))
	fmt.Println(p3.Equals(p3))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
}
