package fluid2d

import "github.com/btracey/fluid"

// TurbKinVisc predicts the turbulent kinematic viscosity using least squares
func TurbKinVisc(tau ReynoldsStress, S StrainRate, eps float64) fluid.KinematicViscosity {
	/*
		norm := S.Norm(twod.Frobenius)

		// Both symmetric
		sum := tau.UU()*S.Symmetric[0] +
			2*tau.UV()*S.Symmetric[1] +
			tau.VV()*S.Symmetric[2]

		return fluid.KinematicViscosity((-1.0 / 2.0) * (sum / (norm*norm + eps)))
	*/
	// 2 * S_ij * S_ij
	denom := 2 * (S.Symmetric[0]*S.Symmetric[0] +
		2*S.Symmetric[1]*S.Symmetric[1] +
		S.Symmetric[2]*S.Symmetric[2])

	// Tau_ij Sij
	num := tau.UU()*S.Symmetric[0] +
		2*tau.UV()*S.Symmetric[1] +
		tau.VV()*S.Symmetric[2]

	return fluid.KinematicViscosity(-num / (denom + eps))
}
