package geometry


type Vector struct {
	X float64
	Y float64
}



func AddVector(a, b Vector) Vector {
	v := Vector{a.X + b.X, a.Y + b.Y}
	return v

}




func SubVector(a, b Vector) Vector {
	v := Vector{a.X - b.X, a.Y - b.Y}
	return v

}




func MagnitudeVector(v Vector) float64 {
	res := Pow(v.X, 2) + Pow(v.Y, 2)
	res = Sqrt(res)

	return res
}



func DotProduct(a, b Vector) float64 {
	res_x := a.X * b.X
	res_y := a.Y * b.Y

	return res_x + res_y

}

func MultipleVector(a, b Vector) Vector {
	res_x := a.X * b.X
	res_y := a.Y * b.Y

	return Vector{res_x, res_y}

}




func AreaCircle(r float64) float64 {
	area := 3.14 * r * r
	return area
}




func AreaRectangle(x, y float64) float64 {
	return x * y
}



func Pow(a, b float64) float64 {
	res := 1.0
	i := 0.0

	for i = 0; i < b; i++ {
		res = res * a
	}

	return res
}



func CalculateAngle(a, b Vector) float64 {
	numerator := DotProduct(a, b)
	denominator := MagnitudeVector(a) * MagnitudeVector(b)

	return numerator / denominator

}



func Sqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for ; i < 10; i++ {
		z = z - (((z * z) - x) / (z + z))
	}

	return z

}


