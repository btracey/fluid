package fluid2d

import "github.com/btracey/matrix/twod"

// StrainRate is the strain-rate tensor (DUi/Dxj + Duj/Dxi) / 2
type StrainRate struct {
	twod.Symmetric
}
