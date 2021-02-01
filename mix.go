package wav_mixer

import (
	"bytes"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

//Recieves 2 .wav file bytes, makes each one a mono, and mixes them into a stereo,
//if no offset is desired, send 0
func MixWavsWithOffset(leftPath, rightPath []byte, outputPath string, leftOffsetSec, rightOffsetSec int) {

	leftCh, format := readWavBytesToChannel(leftPath, 1, 0, leftOffsetSec)
	rightCh, _ := readWavBytesToChannel(rightPath, 0, 1, rightOffsetSec)

	mixedStream := beep.Mix(leftCh, rightCh)

	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wav.Encode(f, mixedStream, format)
}

func readWavBytesToChannel(b []byte, left, right float64, secondsOffset int) (beep.Streamer, beep.Format) {
	f := bytes.NewReader(b)
	s, format, _ := wav.Decode(f)

	s.Seek(int(format.SampleRate) * secondsOffset)
	channel := multiplyChannels(left, right, s)
	return channel, format
}

func multiplyChannels(left, right float64, s beep.Streamer) beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		n, ok = s.Stream(samples)
		for i := range samples[:n] {
			samples[i][0] *= left
			samples[i][1] *= right
		}
		return n, ok
	})
}
