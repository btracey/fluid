package fluid2d

import (
	"github.com/btracey/fluid"
	"github.com/btracey/matrix/twod"
)

// TurbKinVisc predicts the turbulent kinematic viscosity using least squares
func TurbKinVisc(tau ReynoldsStress, S StrainRate, eps float64) fluid.TurbKinVisc {
	norm := S.Norm(twod.Frobenius)

	// Both symmetric
	sum := tau.UU()*S.Symmetric[0] +
		2*tau.UV()*S.Symmetric[1] +
		tau.VV()*S.Symmetric[2]

	return fluid.TurbKinVisc((-1.0 / 2.0) * (sum / (norm*norm + eps)))
}
