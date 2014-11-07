package fluid2d

import (
	"math"

	"github.com/btracey/fluid"
	"github.com/btracey/matrix/twod"
)

// VelGrad is the velocity gradient tensor
type VelGrad struct {
	twod.Dense
}

func (v *VelGrad) SetAll(dudx, dudy, dvdx, dvdy float64) {
	v.SetDUDX(dudx)
	v.SetDUDY(dudy)
	v.SetDVDX(dvdx)
	v.SetDVDY(dvdy)
}

func (v VelGrad) DUDX() float64 {
	return v.Dense[0][0]
}

func (v *VelGrad) SetDUDX(val float64) {
	v.Dense[0][0] = val
}

func (v VelGrad) DUDY() float64 {
	return v.Dense[1][0]
}

func (v *VelGrad) SetDUDY(val float64) {
	v.Dense[1][0] = val
}

func (v VelGrad) DVDX() float64 {
	return v.Dense[0][1]
}

func (v *VelGrad) SetDVDX(val float64) {
	v.Dense[0][1] = val
}

func (v VelGrad) DVDY() float64 {
	return v.Dense[1][1]
}

func (v *VelGrad) SetDVDY(val float64) {
	v.Dense[1][1] = val
}

// Symmetric returs the symmetric part of the velocity gradient (stress tensor)
func (v *VelGrad) SymmetricPart() StressTensor {
	return StressTensor{v.Dense2d().SymmetricPart()}
}

func (v *VelGrad) Vorticity() fluid.Vorticity {
	a := v.Dense[1][0] - v.Dense[0][1]
	a *= a
	a *= 2
	return fluid.Vorticity(math.Sqrt(a))
}
