package fluid2d

import "github.com/btracey/matrix/twod"

type Velocity struct {
	twod.Vector
}

func (v Velocity) U() float64 {
	return v.Vector[0]
}

func (v *Velocity) SetU(val float64) {
	v.Vector[0] = val
}

func (v Velocity) V() float64 {
	return v.Vector[1]
}

func (v *Velocity) SetV(val float64) {
	v.Vector[1] = val
}
