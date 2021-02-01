package wav_mixer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"io/ioutil"
	"math/rand"
	"os"
)

const (
	tempName1Left  = "temp1left"
	tempName2Right = "temp2right"
)

//Recieves 2 .wav byte arrays, makes each one a mono, and mixes them into a stereo,
//if no offset is desired, send 0
func MixWavesWithOffsetFromBytes(left, right []byte, outputPath string, leftOffsetSec, rightOffsetSec int) {
	fileName1 := tempName1Left + randStringBytes(5)
	fileName2 := tempName2Right + randStringBytes(5)

	createTempFileFromBytes(fileName1, left)
	defer os.Remove(fileName1)

	createTempFileFromBytes(fileName2, right)
	defer os.Remove(fileName2)

	MixWavsWithOffset(fileName1, fileName2, outputPath, leftOffsetSec, rightOffsetSec)
}

func createTempFileFromBytes(fileName string, left []byte) {
	permissions := 0666
	err := ioutil.WriteFile(fileName, left, os.FileMode(permissions))
	if err != nil {
		panic("failed to write left side file")
	}
}

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

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
