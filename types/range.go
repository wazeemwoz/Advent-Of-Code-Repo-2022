package types

import . "github.com/wazeemwoz/advent2022/utils"

type Range struct {
	Left, Right int
}

func (r Range) Size() int {
	return r.Right - (r.Left - 1)
}

func (r Range) Unions(other Range) int {
	span := Range{Min(r.Left, other.Left), Max(r.Right, other.Right)}

	return Max((r.Size()+other.Size())-span.Size(), 0)
}
