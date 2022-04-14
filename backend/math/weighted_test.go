package math

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightedDistribution(t *testing.T) {
	p1 := NewWeightedPoint("normal", 20)
	c := NewWeightedDistribution(
		NewWeightedPoint("none", 50),
		p1,
		NewWeightedPoint("lucky", 30),
	)
	assert.Equal(t, c.points[0].ratio, uint(50))
	assert.Equal(t, c.points[1].ratio, uint(70))
	assert.Equal(t, c.points[2].ratio, uint(100))
	assert.Equal(t, p1.ratio, uint(20))
}

func TestWeightedDistributionSampleOutOfBound(t *testing.T) {
	source := []byte{90}

	c := NewWeightedDistribution(
		NewWeightedPoint("none", 50),
		NewWeightedPoint("lucky", 30),
	)
	_, err := c.Sample(bytes.NewReader(source))
	assert.ErrorIs(t, err, errOutOfBound)
}

func TestWeightedDistributionSample(t *testing.T) {

	c := NewWeightedDistribution(
		NewWeightedPoint("none", 50),
		NewWeightedPoint("lucky", 30),
		NewWeightedPoint("epic", 20),
	)

	source := []byte{90}
	sample, err := c.Sample(bytes.NewReader(source))
	assert.NoError(t, err)
	assert.Equal(t, "epic", sample)

	source = []byte{80}
	sample, err = c.Sample(bytes.NewReader(source))
	assert.NoError(t, err)
	assert.Equal(t, "epic", sample)

	source = []byte{79}
	sample, err = c.Sample(bytes.NewReader(source))
	assert.NoError(t, err)
	assert.Equal(t, "lucky", sample)

	source = []byte{50}
	sample, err = c.Sample(bytes.NewReader(source))
	assert.NoError(t, err)
	assert.Equal(t, "lucky", sample)

	source = []byte{49}
	sample, err = c.Sample(bytes.NewReader(source))
	assert.NoError(t, err)
	assert.Equal(t, "none", sample)
}
