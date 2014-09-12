package fluid2d

import "github.com/btracey/matrix/twod"

// StrainRate is the strain-rate tensor (DUi/Dxj + Duj/Dxi) / 2
type StrainRate struct {
	twod.Symmetric
}

func (s *StrainRate) Set(diag1, offDiag, diag2 float64) {
	s.Symmetric[0] = diag1
	s.Symmetric[1] = offDiag
	s.Symmetric[2] = diag2
}
