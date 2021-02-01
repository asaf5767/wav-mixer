package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"os"
)

func main() {
	f, _ := os.Open("sample1.wav")
	f2, _ := os.Open("sample3.wav")

	s1, _, _ := wav.Decode(f)
	s2, _, _ := wav.Decode(f2)

	mixedStream := beep.Mix(s1, s2)
	format := beep.Format{
		SampleRate:  44100,
		NumChannels: 2,
		Precision:   2,
	}

	f, err := os.Create("mixed.wav")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	wav.Encode(f, mixedStream, format)
}
