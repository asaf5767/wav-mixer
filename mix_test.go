package wav_mixer

import "testing"

func TestMixWavsWithOffset(t *testing.T) {
	MixWavsWithOffset("sample1.wav", "sample3.wav",
		"out.wav", 0, 0)
}
