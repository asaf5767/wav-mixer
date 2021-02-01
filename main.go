package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

func main() {
	f, _ := os.Open("sample1.wav")
	f2, _ := os.Open("sample3.wav")

	s1, format, _ := wav.Decode(f)
	s2, _, _ := wav.Decode(f2)

	//s1.Seek(50000)

	leftCh := multiplyChannels(1, 0, s1)
	rightCh := multiplyChannels(0, 1, s2)

	mixedStream := beep.Mix(leftCh, rightCh)

	f, err := os.Create("mixed.wav")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	wav.Encode(f, mixedStream, format)
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
