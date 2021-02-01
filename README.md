# Wav Mixer
A small go module that accepts 2 wav files and merges them, splitting to 2 different channels


# How to use:

Import this module using this command:
``` go get github.com/asaf5767/wav-mixer ```

then you'll be able to call the function ``` MixWavsWithOffset ``` with the following arguments: 

1. 2 Byte arrays which holds the wav files data (left side and then right side)
2. Merged output file path
3. Offset to begin each of the wavs from (left side and then right side)
