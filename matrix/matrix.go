// Package matrix implements functions to deal with matrixes
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
	rows int
	cols int
}

func (m Matrix) Rows() [][]int {
	dst := make([][]int, m.rows)
	for i := range dst {
		dst[i] = make([]int, m.cols)
		copy(dst[i], m.data[i])
	}
	return dst
}

func (m Matrix) Cols() [][]int {
	rows := m.Rows()
	dst := make([][]int, m.cols)
	for i := 0; i < m.cols; i++ {
		dst[i] = make([]int, m.rows)
		for j := 0; j < m.rows; j++ {
			dst[i][j] = rows[j][i]
		}
	}
	return dst
}

func (m *Matrix) Set(r int, c int, val int) bool {
	if r < 0 || m.rows <= r || c < 0 || m.cols <= c {
		return false
	}
	m.data[r][c] = val
	return true
}

// New takes a matrix String and returns a 2D int numbers slice
func New(s string) (*Matrix, error) {
	m := new(Matrix)
	rows := strings.Split(s, "\n")
	m.rows = len(rows)
	m.data = make([][]int, m.rows)

	for i, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if m.cols != 0 && m.cols != len(cols) {
			return nil, errors.New("invalid")
		}
		m.cols = len(cols)
		m.data[i] = make([]int, m.cols)
		for j, col := range cols {
			n, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			m.data[i][j] = n
		}
	}
	return m, nil
}
