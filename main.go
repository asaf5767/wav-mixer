package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

func main() {
	//TODO:
	//1. accept args -> left file, right file, offsets and output
	//2. make this an exec / library
	leftCh, format := readWavToChannel("sample3.wav", 1, 0, 20)
	rightCh, _ := readWavToChannel("sample1.wav", 0, 1, 10)

	mixedStream := beep.Mix(leftCh, rightCh)

	f, err := os.Create("mixed.wav")
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
