package fluid2d

import "github.com/btracey/matrix/twod"

// ReyStress is the reynolds stress tensor (tau). It is symmetric.
type ReynoldsStress struct {
	twod.Symmetric
}

func (r *ReynoldsStress) SetAll(uu, uv, vv float64) {
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
	return r.Symmetric[2]
}
