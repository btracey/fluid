package fluid2d

import "github.com/btracey/matrix/twod"

// TODO: make all of them matrices

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

// VelGrad is the velocity gradient tensor
type VelGrad struct {
	twod.Dense
}

func (v *VelGrad) Set(dudx, dudy, dvdx, dvdy float64) {
	v.SetDUDX(dudx)
	v.SetDUDY(dudy)
	v.SetDVDX(dvdx)
	v.SetDVDY(dudy)
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

// ReyStress is the reynolds stress tensor (tau). It is symmetric.
type ReynoldsStress struct {
	twod.Symmetric
}

func (r *ReynoldsStress) Set(uu, uv, vv float64) {
	r.SetUU(uu)
	r.SetUV(uv)
	r.SetVV(vv)
}

func (r *ReynoldsStress) SetUU(uu float64) {
	r.Symmetric[0] = uu
}

func (r *ReynoldsStress) SetUV(uv float64) {
	r.Symmetric[1] = uv
}

// Sets the VUth component. Note that it is symmetric
func (r *ReynoldsStress) SetVU(uv float64) {
	r.Symmetric[1] = uv
}

func (r *ReynoldsStress) SetVV(vv float64) {
	r.Symmetric[2] = vv
}

func (r ReynoldsStress) UU() float64 {
	return r.Symmetric[0]
}

func (r ReynoldsStress) UV() float64 {
	return r.Symmetric[1]
}

func (r ReynoldsStress) VU() float64 {
	return r.Symmetric[1]
}

func (r ReynoldsStress) VV() float64 {
	return r.Symmetric[1]
}
