package unit_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleFromSexa() {
	// Typical usage:  non-negative d, m, s, and '-' to indicate negative:
	fmt.Println(unit.FromSexa('-', 20, 30, 0))
	// Putting the minus sign on d is not equivalent:
	fmt.Println(unit.FromSexa(' ', -20, 30, 0))
	fmt.Println()
	// Other combinations can give the same result though:
	fmt.Println(unit.FromSexa(' ', -20, -30, 0))
	fmt.Println(unit.FromSexa(' ', -21, 30, 0))
	fmt.Println(unit.FromSexa(' ', -22, 90, 0))
	fmt.Println(unit.FromSexa('-', 22, -90, 0))
	// Output:
	// -20.5
	// -19.5
	//
	// -20.5
	// -20.5
	// -20.5
	// -20.5
}

func ExampleNewAngle() {
	fmt.Println(sexa.FmtAngle(unit.NewAngle('-', 0, 32, 41)))
	fmt.Println(sexa.FmtAngle(unit.NewAngle(' ', 0, 32, 41)))
	fmt.Println(sexa.FmtAngle(unit.NewAngle('+', 0, 32, 41)))
	fmt.Println(sexa.FmtAngle(unit.NewAngle(0, 0, 32, 41)))
	// Output:
	// -32′41″
	// 32′41″
	// 32′41″
	// 32′41″
}

func ExampleAngle_Deg() {
	a := unit.Angle(math.Pi / 2)
	fmt.Println(a.Deg())
	// Output:
	// 90
}

func ExampleAngle_Rad() {
	a := unit.NewAngle(' ', 180, 0, 0) // conversion to radians happens here
	fmt.Println(a.Rad())               // no cost to access radians here
	// Output:
	// 3.141592653589793
}

func ExampleNewRA() {
	a := unit.NewRA(9, 14, 55.8) // converts to radians
	fmt.Println(a)               // radians, native representation
	fmt.Println(a.Hour())        // hours, just scaled from radians
	fmt.Println(sexa.FmtRA(a))   // sexagesimal
	// Output:
	// 2.4213389045230334
	// 9.248833333333334
	// 9ʰ14ᵐ56ˢ
}

func ExampleNewTime() {
	t := unit.NewTime('-', 12, 34, 45.6)
	fmt.Println(sexa.FmtTime(t))
	// Output:
	// -12ʰ34ᵐ46ˢ
}

func ExampleTime_Sec() {
	t := unit.NewTime(0, 0, 1, 30)
	fmt.Println(t.Sec())
	// Output:
	// 90
}

func ExampleTime_Min() {
	t := unit.NewTime(0, 0, 1, 30)
	fmt.Println(t.Min())
	// Output:
	// 1.5
}

func ExampleTime_Hour() {
	t := unit.NewTime(0, 2, 15, 0)
	fmt.Println(t.Hour())
	// Output:
	// 2.25
}

func ExampleTime_Day() {
	t := unit.NewTime(0, 48, 36, 0)
	fmt.Println(t.Day())
	// Output:
	// 2.025
}

func ExampleTime_Rad() {
	t := unit.NewTime(0, 12, 0, 0)
	fmt.Println(t.Rad())
	// Output:
	// 3.141592653589793
}
