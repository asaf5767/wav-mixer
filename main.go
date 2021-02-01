package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

func main() {

	leftCh, format := readWavToChannel("sample3.wav", 1, 0)
	rightCh, format := readWavToChannel("sample1.wav", 0, 1)

	mixedStream := beep.Mix(leftCh, rightCh)

	f, err := os.Create("mixed.wav")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	wav.Encode(f, mixedStream, format)
}

func readWavToChannel(path string, left, right float64) (beep.Streamer, beep.Format) {
	f, _ := os.Open(path)
	s1, format, _ := wav.Decode(f)
	channel := multiplyChannels(left, right, s1)
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
