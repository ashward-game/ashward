package math

import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"
)

var errOutOfBound = errors.New("out of bound: sum of weights might be less then 1")

type WeightedPoint struct {
	value interface{}
	ratio uint
}

func NewWeightedPoint(value interface{}, ratio uint) WeightedPoint {
	return WeightedPoint{
		value: value,
		ratio: ratio,
	}
}

type WeightedDistribution struct {
	points []WeightedPoint
}

// NewWeightedDistribution returns a collection (or distribution) of a set of
// pre-defined points with specific weight.
// NewWeightedDistribution only supports points with non-negative integer weight.
// Sum of weights should be 1.
func NewWeightedDistribution(points ...WeightedPoint) *WeightedDistribution {
	for i := 1; i < len(points); i++ {
		points[i].ratio += points[i-1].ratio
	}
	return &WeightedDistribution{
		points: points,
	}
}

// Sample returns a sample from the defined distribution, using
// `reader`` as the source of randomness.
// If sum of weights is less than 1, return an error.
// If sum of weights is larger than 1, no error is returned
// (though it might not work as expected).
func (wd *WeightedDistribution) Sample(reader io.Reader) (interface{}, error) {
	// since weights are integer, resolution can just be 100.
	// increase resolution if weights are float for more precision.
	resolution := int64(100)

	random, err := rand.Int(reader, big.NewInt(int64(resolution)))
	if err != nil {
		return nil, err
	}
	rand := uint(random.Uint64())
	for _, point := range wd.points {
		if rand < point.ratio {
			return point.value, nil
		}
	}
	return nil, errOutOfBound
}
