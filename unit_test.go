// Public domain

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

func ExampleFromSexaSec() {
	fmt.Println(unit.FromSexaSec(' ', 1, 0, 0))
	// Output:
	// 3600
}

func ExamplePMod() {
	// Example numbers from the go language spec:
	fmt.Println(" x      y     x % y   PMod(x, y)")
	{
		y := 3
		for _, x := range []int{5, -5} {
			fmt.Printf("%2d %6d %7d %12.1f\n",
				x, y, x%y, unit.PMod(float64(x), float64(y)))
		}
	}
	// A more typical use:
	x := -2.5
	y := 360.
	fmt.Printf("%.1f %4.0f %20.1f\n", x, y, unit.PMod(x, y))
	// Output:
	//  x      y     x % y   PMod(x, y)
	//  5      3       2          2.0
	// -5      3      -2          1.0
	// -2.5  360                357.5
}

func ExampleAngleFromDeg() {
	a := unit.AngleFromDeg(180)
	fmt.Println(a)
	fmt.Println(sexa.FmtAngle(a))
	// Output:
	// 3.141592653589793
	// 180°0′0″
}

func ExampleAngleFromMin() {
	a := unit.AngleFromMin(83)
	fmt.Println(sexa.FmtAngle(a))
	// Output:
	// 1°23′0″
}

func ExampleAngleFromSec() {
	a := unit.AngleFromSec(83)
	fmt.Println(sexa.FmtAngle(a))
	// Output:
	// 1′23″
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

func ExampleAngle_Cos() {
	a := unit.AngleFromDeg(60)
	fmt.Printf("%f\n", a.Cos())
	// Output:
	// 0.500000
}

func ExampleAngle_Deg() {
	a := unit.Angle(math.Pi / 2)
	fmt.Println(a.Deg())
	// Output:
	// 90
}

func ExampleAngle_Div() {
	a := unit.AngleFromDeg(90)
	fmt.Println(sexa.FmtAngle(a.Div(2)))
	// Output:
	// 45°0′0″
}

func ExampleAngle_HourAngle() {
	a := unit.AngleFromDeg(-30)
	fmt.Println(sexa.FmtAngle(a))
	fmt.Println(sexa.FmtHourAngle(a.HourAngle()))
	// Output:
	// -30°0′0″
	// -2ʰ0ᵐ0ˢ
}

func ExampleAngle_Min() {
	a := unit.NewAngle(' ', 1, 23, 0)
	fmt.Printf("%f\n", a.Min())
	// Output:
	// 83.000000
}

func ExampleAngle_Mod1() {
	fmt.Println(sexa.FmtAngle(unit.AngleFromDeg(23).Mod1()))
	fmt.Println(sexa.FmtAngle(unit.AngleFromDeg(383).Mod1()))
	fmt.Println(sexa.FmtAngle(unit.AngleFromDeg(-337).Mod1()))
	// Output:
	// 23°0′0″
	// 23°0′0″
	// 23°0′0″
}

func ExampleAngle_Mul() {
	a := unit.AngleFromDeg(45)
	fmt.Println(sexa.FmtAngle(a.Mul(2)))
	// Output:
	// 90°0′0″
}

func ExampleAngle_RA() {
	a := unit.AngleFromDeg(-30)
	fmt.Println(sexa.FmtAngle(a))
	fmt.Println(sexa.FmtRA(a.RA()))
	// Output:
	// -30°0′0″
	// 22ʰ0ᵐ0ˢ
}

func ExampleAngle_Rad() {
	a := unit.NewAngle(' ', 180, 0, 0) // conversion to radians happens here
	fmt.Println(a.Rad())               // no cost to access radians here
	// Output:
	// 3.141592653589793
}

func ExampleAngle_Sec() {
	a := unit.NewAngle(' ', 0, 1, 23)
	fmt.Printf("%f\n", a.Sec())
	// Output:
	// 83.000000
}

func ExampleAngle_Sin() {
	a := unit.AngleFromDeg(60)
	fmt.Printf("%f\n", a.Sin())
	// Output:
	// 0.866025
}

func ExampleAngle_Sincos() {
	a := unit.AngleFromDeg(60)
	s, c := a.Sincos()
	fmt.Printf("%f  %f\n", s, c)
	// Output:
	// 0.866025  0.500000
}

func ExampleAngle_Tan() {
	a := unit.AngleFromDeg(60)
	fmt.Printf("%f\n", a.Tan())
	// Output:
	// 1.732051
}

func ExampleAngle_Time() {
	a := unit.AngleFromDeg(-30)
	fmt.Println(sexa.FmtAngle(a))
	fmt.Println(sexa.FmtTime(a.Time()))
	// Output:
	// -30°0′0″
	// -2ʰ0ᵐ0ˢ
}

func ExampleHourAngleFromHour() {
	h := unit.HourAngleFromHour(12)
	fmt.Println(h)
	fmt.Println(sexa.FmtHourAngle(h))
	fmt.Println(sexa.FmtHourAngle(unit.HourAngleFromHour(-2)))
	// Output:
	// 3.141592653589793
	// 12ʰ0ᵐ0ˢ
	// -2ʰ0ᵐ0ˢ
}

func ExampleHourAngleFromMin() {
	fmt.Println(sexa.FmtHourAngle(unit.HourAngleFromMin(-83)))
	// Output:
	// -1ʰ23ᵐ0ˢ
}

func ExampleHourAngleFromSec() {
	fmt.Println(sexa.FmtHourAngle(unit.HourAngleFromSec(-83)))
	// Output:
	// -1ᵐ23ˢ
}

func ExampleNewHourAngle() {
	fmt.Println(sexa.FmtHourAngle(unit.NewHourAngle(' ', 0, 32, 41)))
	// Output:
	// 32ᵐ41ˢ
}

func ExampleHourAngle_Angle() {
	h := unit.HourAngleFromHour(-2)
	fmt.Println(sexa.FmtHourAngle(h))
	fmt.Println(sexa.FmtAngle(h.Angle()))
	// Output:
	// -2ʰ0ᵐ0ˢ
	// -30°0′0″
}

func ExampleHourAngle_Cos() {
	h := unit.HourAngleFromHour(4)
	fmt.Printf("%f\n", h.Cos())
	// Output:
	// 0.500000
}

func ExampleHourAngle_Hour() {
	h := unit.NewHourAngle(' ', 1, 30, 0)
	fmt.Println(h.Hour())
	// Output:
	// 1.5
}

func ExampleHourAngle_Div() {
	h := unit.HourAngleFromHour(6)
	fmt.Println(sexa.FmtHourAngle(h.Div(2)))
	// Output:
	// 3ʰ0ᵐ0ˢ
}

func ExampleHourAngle_Min() {
	h := unit.NewHourAngle(' ', 1, 23, 0)
	fmt.Printf("%f\n", h.Min())
	// Output:
	// 83.000000
}

func ExampleHourAngle_Mul() {
	h := unit.HourAngleFromHour(3)
	fmt.Println(sexa.FmtHourAngle(h.Mul(2)))
	// Output:
	// 6ʰ0ᵐ0ˢ
}

func ExampleHourAngle_RA() {
	h := unit.HourAngleFromHour(-2)
	fmt.Println(sexa.FmtHourAngle(h))
	fmt.Println(sexa.FmtRA(h.RA()))
	// Output:
	// -2ʰ0ᵐ0ˢ
	// 22ʰ0ᵐ0ˢ
}

func ExampleHourAngle_Rad() {
	a := unit.NewHourAngle(' ', 12, 0, 0) // conversion to radians happens here
	fmt.Println(a.Rad())                  // no cost to access radians here
	// Output:
	// 3.141592653589793
}

func ExampleHourAngle_Sec() {
	a := unit.NewHourAngle(' ', 0, 1, 23)
	fmt.Printf("%f\n", a.Sec())
	// Output:
	// 83.000000
}

func ExampleHourAngle_Sin() {
	h := unit.HourAngleFromHour(4)
	fmt.Printf("%f\n", h.Sin())
	// Output:
	// 0.866025
}

func ExampleHourAngle_Sincos() {
	h := unit.HourAngleFromHour(4)
	s, c := h.Sincos()
	fmt.Printf("%f  %f\n", s, c)
	// Output:
	// 0.866025  0.500000
}

func ExampleHourAngle_Tan() {
	h := unit.HourAngleFromHour(4)
	fmt.Printf("%f\n", h.Tan())
	// Output:
	// 1.732051
}

func ExampleHourAngle_Time() {
	h := unit.HourAngleFromHour(-2)
	fmt.Println(sexa.FmtHourAngle(h))
	fmt.Println(sexa.FmtTime(h.Time()))
	// Output:
	// -2ʰ0ᵐ0ˢ
	// -2ʰ0ᵐ0ˢ
}

func ExampleNewRA() {
	r := unit.NewRA(9, 14, 55.8) // converts to radians
	fmt.Println(r)               // radians, native representation
	fmt.Println(r.Hour())        // hours, just scaled from radians
	fmt.Println(sexa.FmtRA(r))   // sexagesimal
	// Output:
	// 2.4213389045230334
	// 9.248833333333334
	// 9ʰ14ᵐ56ˢ
}

func ExampleRAFromDeg() {
	fmt.Println(sexa.FmtRA(unit.RAFromDeg(-30)))
	// Output:
	// 22ʰ0ᵐ0ˢ
}

func ExampleRAFromHour() {
	fmt.Println(sexa.FmtRA(unit.RAFromHour(-2)))
	// Output:
	// 22ʰ0ᵐ0ˢ
}

func ExampleRAFromMin() {
	fmt.Println(sexa.FmtRA(unit.RAFromMin(83)))
	// Output:
	// 1ʰ23ᵐ0ˢ
}

func ExampleRAFromSec() {
	fmt.Println(sexa.FmtRA(unit.RAFromSec(83)))
	// Output:
	// 1ᵐ23ˢ
}

func ExampleRAFromRad() {
	fmt.Println(sexa.FmtRA(unit.RAFromRad(math.Pi)))
	// Output:
	// 12ʰ0ᵐ0ˢ
}

func ExampleRA_Add() {
	r := unit.RAFromHour(22)
	fmt.Println(sexa.FmtRA(r.Add(unit.HourAngleFromHour(4))))
	// Output:
	// 2ʰ0ᵐ0ˢ
}

func ExampleRA_Angle() {
	r := unit.RAFromHour(2)
	fmt.Println(sexa.FmtRA(r))
	fmt.Println(sexa.FmtAngle(r.Angle()))
	// Output:
	// 2ʰ0ᵐ0ˢ
	// 30°0′0″
}

func ExampleRA_Cos() {
	h := unit.RAFromHour(4)
	fmt.Printf("%f\n", h.Cos())
	// Output:
	// 0.500000
}

func ExampleRA_Deg() {
	r := unit.RAFromHour(2)
	fmt.Printf("%f\n", r.Deg())
	// Output:
	// 30.000000
}

func ExampleRA_Hour() {
	r := unit.NewRA(1, 30, 0)
	fmt.Println(r.Hour())
	// Output:
	// 1.5
}

func ExampleRA_Min() {
	r := unit.NewRA(1, 23, 0)
	fmt.Printf("%f\n", r.Min())
	// Output:
	// 83.000000
}

func ExampleRA_HourAngle() {
	r := unit.RAFromHour(2)
	fmt.Println(sexa.FmtRA(r))
	fmt.Println(sexa.FmtHourAngle(r.HourAngle()))
	// Output:
	// 2ʰ0ᵐ0ˢ
	// 2ʰ0ᵐ0ˢ
}

func ExampleRA_Rad() {
	r := unit.NewRA(12, 0, 0) // conversion to radians happens here
	fmt.Println(r.Rad())      // no cost to access radians here
	// Output:
	// 3.141592653589793
}

func ExampleRA_Sec() {
	r := unit.NewRA(0, 1, 23)
	fmt.Printf("%f\n", r.Sec())
	// Output:
	// 83.000000
}

func ExampleRA_Sin() {
	r := unit.RAFromHour(4)
	fmt.Printf("%f\n", r.Sin())
	// Output:
	// 0.866025
}

func ExampleRA_Sincos() {
	r := unit.RAFromHour(4)
	s, c := r.Sincos()
	fmt.Printf("%f  %f\n", s, c)
	// Output:
	// 0.866025  0.500000
}

func ExampleRA_Tan() {
	r := unit.RAFromHour(4)
	fmt.Printf("%f\n", r.Tan())
	// Output:
	// 1.732051
}

func ExampleRA_Time() {
	r := unit.RAFromHour(2)
	fmt.Println(sexa.FmtRA(r))
	fmt.Println(sexa.FmtTime(r.Time()))
	// Output:
	// 2ʰ0ᵐ0ˢ
	// 2ʰ0ᵐ0ˢ
}

func ExampleTime() {
	t := unit.Time(1.23)
	fmt.Println(t)
	fmt.Printf("%.2s\n", sexa.FmtTime(t))
	// Output:
	// 1.23
	// 1.23ˢ
}

func ExampleNewTime() {
	t := unit.NewTime('-', 12, 34, 45.6)
	fmt.Println(sexa.FmtTime(t))
	// Output:
	// -12ʰ34ᵐ46ˢ
}

func ExampleTimeFromDay() {
	t := unit.TimeFromDay(1)
	fmt.Println(t)
	fmt.Println(sexa.FmtTime(t))
	// Output:
	// 86400
	// 24ʰ0ᵐ0ˢ
}

func ExampleTimeFromHour() {
	t := unit.TimeFromHour(-1)
	fmt.Println(t)
	fmt.Println(sexa.FmtTime(t))
	// Output:
	// -3600
	// -1ʰ0ᵐ0ˢ
}

func ExampleTimeFromMin() {
	fmt.Println(sexa.FmtTime(unit.TimeFromMin(.5)))
	// Output:
	// 30ˢ
}

func ExampleTimeFromRad() {
	t := unit.TimeFromRad(math.Pi)
	fmt.Println(sexa.FmtTime(t))
	// Output:
	// 12ʰ0ᵐ0ˢ
}

func ExampleTime_Angle() {
	t := unit.TimeFromHour(-2)
	fmt.Println(sexa.FmtTime(t))
	fmt.Println(sexa.FmtAngle(t.Angle()))
	// Output:
	// -2ʰ0ᵐ0ˢ
	// -30°0′0″
}

func ExampleTime_Day() {
	t := unit.NewTime(0, 48, 36, 0)
	fmt.Println(t.Day())
	// Output:
	// 2.025
}

func ExampleTime_Div() {
	t := unit.TimeFromHour(6)
	fmt.Println(sexa.FmtTime(t.Div(2)))
	// Output:
	// 3ʰ0ᵐ0ˢ
}

func ExampleTime_Hour() {
	t := unit.NewTime(0, 2, 15, 0)
	fmt.Println(t.Hour())
	// Output:
	// 2.25
}

func ExampleTime_HourAngle() {
	t := unit.TimeFromHour(-1.5)
	fmt.Println(sexa.FmtHourAngle(t.HourAngle()))
	// Output:
	// -1ʰ30ᵐ0ˢ
}

func ExampleTime_Min() {
	t := unit.NewTime(0, 0, 1, 30)
	fmt.Println(t.Min())
	// Output:
	// 1.5
}

func ExampleTime_Mod1() {
	t := unit.TimeFromHour(25)
	fmt.Println(t)
	fmt.Println(t.Mod1())
	fmt.Println(sexa.FmtTime(t.Mod1()))
	// Output:
	// 90000
	// 3600
	// 1ʰ0ᵐ0ˢ
}

func ExampleTime_Mul() {
	t := unit.TimeFromHour(3)
	fmt.Println(sexa.FmtTime(t.Mul(2)))
	// Output:
	// 6ʰ0ᵐ0ˢ
}

func ExampleTime_RA() {
	t := unit.TimeFromHour(-2)
	fmt.Println(sexa.FmtTime(t))
	fmt.Println(sexa.FmtRA(t.RA()))
	// Output:
	// -2ʰ0ᵐ0ˢ
	// 22ʰ0ᵐ0ˢ
}

func ExampleTime_Rad() {
	t := unit.NewTime(0, 12, 0, 0)
	fmt.Println(t.Rad())
	// Output:
	// 3.141592653589793
}

func ExampleTime_Sec() {
	t := unit.NewTime(0, 0, 1, 30)
	fmt.Println(t.Sec())
	// Output:
	// 90
}
