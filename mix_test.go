package wav_mixer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestMixWavsWithOffset(t *testing.T) {
	MixWavsWithOffset("sample1.wav", "sample3.wav",
		"out.wav", 0, 0)
}

func TestMixWavsWithOffsetFromBytes(t *testing.T) {

	bytes1, err := ioutil.ReadFile("sample1.wav")
	assert.NoError(t, err)

	bytes2, err := ioutil.ReadFile("sample3.wav")
	assert.NoError(t, err)

	MixWavesWithOffsetFromBytes(bytes1, bytes2, "out.wav", 0, 0)
}
