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
	"math"
)

func printImage(in string, out string) {

	// Number of seconds worth of buffer to allocate.
	const Seconds = 8
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

func printLine(file *image.RGBA, p int, h int) {
	
	for i := 0; i < h; i++ {
		file.Set(p, i, color.RGBA{0x00, 0x00, 0x00, 0xff})
	}
	
}

func printSpectr(file *image.RGBA, p int, h int, pow int, c color.RGBA) {
	
	h = h/2
	
	if pow < h {
		for i := h; i > pow; i-- {
			file.Set(p, i, c)
		}
	} else {
		for i := h; i < pow; i++ {
			file.Set(p, i, c)
		}
	}
}

func printImageShort(in string, out string, div int) {

	// Number of seconds worth of buffer to allocate.
	const Seconds = 8
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
	if err != nil {
		return
	}

	// Compression à 1/100. On ignore le channel 0
	const ch = 1
	var packFrame int64 = int64(info.Samplerate / int32(div))
	nbrPack := int(math.Ceil(float64(numSamples) / float64(packFrame)))
	
	//Buffers
	var subTotal int64 = 0
	var nbr int64 = 0
	var actualPack int = 0
	totalBuff := make([]int16, nbrPack)
	
	totalBuffMin := make([]int16, nbrPack)
	totalBuffMax := make([]int16, nbrPack)
	
	outimage := image.NewRGBA(image.Rect(0, 0, nbrPack, ImageHeight))
	
	fmt.Printf("Width=%d, Height=%d \n", nbrPack, ImageHeight)
	
	// GO
	for i := 0; i < numSamples; i++ {
		for channel := 0; channel < numChannels; channel++ {
			
			if channel == ch {
				val := buffer[i*numChannels+channel]
				subTotal += int64(val)
				nbr++
				
				if totalBuffMin[actualPack] > val {
					totalBuffMin[actualPack] = val
				}
				if totalBuffMax[actualPack] < val {
					totalBuffMax[actualPack] = val
				}
				
			}
			if nbr == packFrame {
				
				totalBuff[actualPack] = int16(subTotal / nbr)
				nbr = 0
				subTotal = 0
				actualPack++
			}
		}
	}
	
	// Both math.Abs and math.Max operate on float64. Hm.
	max := int16(0)
	min := int16(0)
	for _, v := range totalBuff {
		if v > max {
			max = v
		}
		if v < min {
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
	for i := 0; i < nbrPack; i++ {
		//y := int(float64(totalBuff[i])*mult+float64(ImageHeight)/2)
		min := int(float64(totalBuffMin[i])*mult+float64(ImageHeight)/2)
		max := int(float64(totalBuffMax[i])*mult+float64(ImageHeight)/2)

		var t color.RGBA
		var t2 color.RGBA
		
		if int(totalBuff[i]) < th {
			t = color.RGBA{0xff, 0x00, 0x00, 0xff}
			t2 = color.RGBA{0xff, 0xf0, 0x00, 0xff}
			
		} else {
			t = color.RGBA{0x00, 0x00, 0xff, 0xff}
			t2 = color.RGBA{0x00, 0xf0, 0xff, 0xff}
		}
		
		outimage.Set(i, min, t)
		outimage.Set(i, max, t)
		printSpectr(outimage, i, ImageHeight, min, t2)
		printSpectr(outimage, i, ImageHeight, max, t2)
		
		if i % div == 0 {
			printLine(outimage, i, ImageHeight)
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

	//printBegin()
	//printImageShort("data/test2.aiff", "data/out_5000.png", 5000)
	printImageShort("data/test2.aiff", "data/out_900.png", 900)
	printImageShort("data/test2.aiff", "data/out_500.png", 500)
	printImageShort("data/test2.aiff", "data/out_400.png", 400)
	printImageShort("data/test2.aiff", "data/out_300.png", 300)
	printImageShort("data/test2.aiff", "data/out_200.png", 200)
	printImageShort("data/test2.aiff", "data/out_150.png", 150)
	printImageShort("data/test2.aiff", "data/out_120.png", 120)
	//printImageShort("data/test2.aiff", "data/out_50.png", 50)
	//printImageShort("data/test2.aiff", "data/out_10.png", 10)
	algo()
}
