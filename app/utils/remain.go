package utils

import "math"

type name struct {

}

func Creat(value uint32) [][64]uint32 {
	remain := make([][64]uint32, 8)
	for i := 0; i < 8; i++ {
		for j := 0; j < 64; j++ {
			remain[i][j] = value
		}
	}
	return remain
}

func Add(start, end, col int, matrix [][64]uint32) bool {
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

func Sub(start, end, col int, matrix [][64]uint32) bool {
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

func Find(start, end, col int, matrix [][64]uint32) uint32 {
	if matrix[col][63] > 20 {
		// uint 无法返回-1
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
