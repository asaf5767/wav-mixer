package wav_mixer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestMixWavsWithOffsetFromBytes(t *testing.T) {

	bytes1, err := ioutil.ReadFile("sample1.wav")
	assert.NoError(t, err)

	bytes2, err := ioutil.ReadFile("sample3.wav")
	assert.NoError(t, err)

	MixWavsWithOffset(bytes1, bytes2, "out.wav", 0, 0)
}
