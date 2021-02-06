package utils

import "math"

type RemainingTicket struct {

}


func (*RemainingTicket)Creat(col int, value uint32, remain *[8][64]uint32) {
	for j := 0; j < 64; j++ {
		remain[col][j] = value
	}
}

func (*RemainingTicket)Add(start, end, col int, matrix *[8][64]uint32) bool {
	for i := start; i <= end; i++ {
		matrix[col][i] += 1
	}

	min := matrix[col][0]
	for j := 1; j < 63; j++ {
		if matrix[col][j] < min {
			min = matrix[col][j]
		}
	}

	matrix[col][63] = min
	return true
}

func (*RemainingTicket)Sub(start, end, col int, matrix [8][64]uint32) bool {
	for i := start; i <= end; i++ {
		if matrix[col][i] == 0 {
			return false
		} else {
			matrix[col][i] -= 1
			if matrix[col][i] < matrix[col][63] {
				matrix[col][63] = matrix[col][i]
			}
		}
	}
	return true
}

func (*RemainingTicket)Find(start, end, col uint32, matrix [8][64]uint32) uint32 {
	if matrix[col][63] > 20 {
		return math.MaxUint32
	} else {
		min := uint32(math.MaxUint32)
		for i := start; i <= end; i++ {
			if min > matrix[col][i] {
				min = matrix[col][i]
			}
		}
		return min
	}
}
