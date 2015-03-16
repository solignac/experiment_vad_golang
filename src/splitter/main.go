package splitter

//import "encoding/json"
//import "github.com/mkb218/gosndfile/sndfile"

import (
	"fmt"
	"github.com/mkb218/gosndfile/sndfile"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func printImage(in string, out string) {

	// Number of seconds worth of buffer to allocate.
	const Seconds = 5
	// Height per channel.
	const ImageHeight = 200

	var info sndfile.Info
	soundFile, err := sndfile.Open(in, sndfile.Read, &info)

	if err != nil {
		log.Fatal("Error", err)
	}
	defer soundFile.Close()

	imageFile, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	buffer := make([]int16, Seconds*info.Samplerate*info.Channels)
	numRead, err := soundFile.ReadItems(buffer)
	numSamples := int(numRead / int64(info.Channels))
	numChannels := int(info.Channels)
	outimage := image.NewRGBA(image.Rect(0, 0, numSamples, ImageHeight*numChannels))
	if err != nil {
		return
	}
	// Both math.Abs and math.Max operate on float64. Hm.
	max := int16(0)
	min := int16(0)
	for i, v := range buffer {
		if (i%2) == 1 && v > max {
			max = v
		}
		if (i%2) == 1 && v < min {
			min = v
		}
	}
	fmt.Printf("Max = %d, min = %d \n", max, min)

	var th int = ((int(max) - int(min)) / 2) + int(min)

	th = 0

	fmt.Printf("Th = %d \n", th)

	// Work out scaling factor to normalise signaland get best use of space.
	mult := float64(float64(ImageHeight)/float64(max)) / 2

	fmt.Printf("Mult = %f \n", mult)

	// Signed float so add 1 to turn [-1, 1] into [0, 2].
	for i := 0; i < numSamples; i++ {
		for channel := 0; channel < numChannels; channel++ {
			y := int(float64(buffer[i*numChannels+channel])*mult+float64(ImageHeight)/2) + ImageHeight*channel

			if channel == 0 { // Left Ref
				//outimage.Set(i, y, color.RGBA{0xff, 0x00, 0x00, 0xff})
			} else { // Right Green

				if int(buffer[i*numChannels+channel]) < th {
					outimage.Set(i, y, color.RGBA{0xff, 0x00, 0x00, 0xff})
				} else {
					outimage.Set(i, y, color.RGBA{0x00, 0x00, 0xff, 0xff})
				}
				//outimage.Set(i, y+1, color.RGBA64{0xFF, 0x00, 0x00, 0xFF})
			}
			//outimage.Set(i, y+1, color.Gray16{0xf000})
		}
	}
	png.Encode(imageFile, outimage)
}

func printInfo(inf *sndfile.Info) {

	fmt.Printf("Frames/Samples : %d \n", inf.Frames)
	fmt.Printf("Sample rate : %d hz \n", inf.Samplerate)
	fmt.Printf("Number of channel : %d \n", inf.Channels)
	fmt.Printf("Calculated duration : %f \n", float64(inf.Frames)/float64(inf.Samplerate))
	fmt.Println("")
}

func algo() {

	var inf sndfile.Info
	//var f * sndfile.File

	_, _ = sndfile.Open("data/test2.aiff", sndfile.Read, &inf)

	// En secondes
	const tempo float64 = 0.5

	var fpersec = inf.Samplerate
	var ftempo int = int(float64(fpersec) * tempo)

	fmt.Printf("Pack of %d frames\n", ftempo)
}

func printBegin() {

	var f *sndfile.File
	var inf sndfile.Info

	var abuff []int16 = make([]int16, 12)
	var bbuff []int16 = make([]int16, 12)

	f, _ = sndfile.Open("data/test2.aiff", sndfile.Read, &inf)

	printInfo(&inf)

	fmt.Println("Sample")
	f.ReadFrames(abuff)
	fmt.Println(abuff)

	f.Close()
	f, _ = sndfile.Open("data/test2.aiff", sndfile.Read, &inf)

	fmt.Println("Item")
	f.ReadItems(bbuff)
	fmt.Println(bbuff)

	f.Close()
}

func Sub_main() {

	fmt.Println("OK")

	s, _ := sndfile.GetLibVersion()

	fmt.Println(s)

	printBegin()
	//printImage("data/test2.aiff", "data/out.png")
	algo()
}
