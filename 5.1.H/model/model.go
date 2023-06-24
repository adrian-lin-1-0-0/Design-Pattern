package model

type Model interface {
	Mul([]float64) []float64
}

type RealModel struct {
	matrix [][]float64
}

func NewRealModel(matrix [][]float64) *RealModel {
	return &RealModel{
		matrix: matrix,
	}
}

func (m *RealModel) Mul(v []float64) []float64 {
	result := make([]float64, len(v))
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v); j++ {
			result[i] += m.matrix[i][j] * v[j]
		}
	}
	return result
}
