package wav_mixer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

//Recieves 2 .wav file paths, makes each one a mono, and mixes them into a stereo,
//if no offset is desired, send 0
func MixWavsWithOffset(leftPath, rightPath, outputPath string, leftOffsetSec, rightOffsetSec int) {

	leftCh, format := readWavToChannel(leftPath, 1, 0, leftOffsetSec)
	rightCh, _ := readWavToChannel(rightPath, 0, 1, rightOffsetSec)

	mixedStream := beep.Mix(leftCh, rightCh)

	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wav.Encode(f, mixedStream, format)
}

func readWavToChannel(path string, left, right float64, secondsOffset int) (beep.Streamer, beep.Format) {
	f, _ := os.Open(path)
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

//func samplesFromTime(t time.Duration, ) int {
//	return int(t * time.Duration(sr) / time.Second)
//}
