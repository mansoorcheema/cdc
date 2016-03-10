/* Package geometry implements simple geomery operations
*/

package geometry

//Vector represents a two dimensional vector
type Vector struct {
	X float64
	Y float64
}


//AddVector adds two vectors
func AddVector(a, b Vector) Vector {
	v := Vector{a.X + b.X, a.Y + b.Y}
	return v

}



//SubVector subtracts two vectors
func SubVector(a, b Vector) Vector {
	v := Vector{a.X - b.X, a.Y - b.Y}
	return v

}



//MagnitudeVector finds magnitude of vector
func MagnitudeVector(v Vector) float64 {
	res := Pow(v.X, 2) + Pow(v.Y, 2)
	res = Sqrt(res)

	return res
}


//DotProduct finds dot product of two vectors
func DotProduct(a, b Vector) float64 {
	res_x := a.X * b.X
	res_y := a.Y * b.Y

	return res_x + res_y

}

//MultiplyVector multiplies two vectors
func MultiplyVector(a, b Vector) Vector {
	res_x := a.X * b.X
	res_y := a.Y * b.Y

	return Vector{res_x, res_y}

}



//AreaCircle finds area of a circle
func AreaCircle(r float64) float64 {
	area := 3.14 * r * r
	return area
}



//AreaRectangle finds area of rectangle
func AreaRectangle(x, y float64) float64 {
	return x * y
}


//Pow finds exponent i.e a raised to b
func Pow(a, b float64) float64 {
	res := 1.0
	i := 0.0

	for i = 0; i < b; i++ {
		res = res * a
	}

	return res
}


//CalculateAngle calculates cosine between two vectors
func CalculateAngle(a, b Vector) float64 {
	numerator := DotProduct(a, b)
	denominator := MagnitudeVector(a) * MagnitudeVector(b)

	return numerator / denominator

}


//Sqrt finds squareroot of a float64
func Sqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for ; i < 10; i++ {
		z = z - (((z * z) - x) / (z + z))
	}

	return z

}


