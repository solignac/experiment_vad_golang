package splitter

import (
	"fmt"
	"github.com/mkb218/gosndfile/sndfile"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func printLine(file *image.RGBA, p int, h int) {
	for i := 0; i < h; i++ {
		file.Set(p, i, color.RGBA{0x00, 0x00, 0x00, 0xff})
	}
}

func printSpectr(file *image.RGBA, x int, height int, ampl int, col color.RGBA) {

	height = height / 2
	if ampl < height {
		for i := height; i > ampl; i-- {
			file.Set(x, i, col)
		}
	} else {
		for i := height; i < ampl; i++ {
			file.Set(x, i, col)
		}
	}
}

func printImageShort(in string, out string, div int) {

	// Number of seconds worth of buffer to allocate.
	const Seconds = 10
	// Height per channel.
	const ImageHeight = 300

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

	// Nouveau seuil
	middlePos := 0
	totalTmp := 0
	for _, v := range totalBuffMax {
		totalTmp += int(v)
	}
	middlePos = totalTmp / len(totalBuffMax)

	th := middlePos

	mult := float64(float64(ImageHeight)/float64(max)) / 2

	th2 := int((float64(th) * mult) + float64(ImageHeight/2))

	fmt.Printf("Th = %d, th2=%d \n", th, th2)
	fmt.Printf("Mult = %f \n", mult)

	// Signed float so add 1 to turn [-1, 1] into [0, 2].
	for i := 0; i < nbrPack; i++ {
		//y := int(float64(totalBuff[i])*mult+float64(ImageHeight)/2)
		min := int(float64(totalBuffMin[i])*mult + float64(ImageHeight)/2)
		max := int(float64(totalBuffMax[i])*mult + float64(ImageHeight)/2)

		var t color.RGBA
		var t2 color.RGBA

		if int(totalBuffMax[i]) > th {
			t = color.RGBA{0xff, 0x00, 0x00, 0xff}  // Red
			t2 = color.RGBA{0xff, 0xa0, 0x00, 0xff} // Yellow

		} else {
			t = color.RGBA{0x00, 0x00, 0xff, 0xff}  // Dark blue
			t2 = color.RGBA{0x00, 0xf0, 0xff, 0xff} // Light blue
		}

		outimage.Set(i, min, t)
		outimage.Set(i, max, t)
		printSpectr(outimage, i, ImageHeight, min, t2)
		printSpectr(outimage, i, ImageHeight, max, t2)

		outimage.Set(i, th2, color.RGBA{0xff, 0x00, 0x00, 0xff})

		if i%div == 0 {
			printLine(outimage, i, ImageHeight)
		}
	}
	png.Encode(imageFile, outimage)
}

func printInfo(fileName string) {

	var info sndfile.Info

	_, err := sndfile.Open(fileName, sndfile.Read, &info)

	if err != nil {
		log.Fatal("Error", err)
	}

	fmt.Printf("Frames/Samples : %d \n", info.Frames)
	fmt.Printf("Sample rate : %d hz \n", info.Samplerate)
	fmt.Printf("Number of channel : %d \n", info.Channels)
	fmt.Printf("Calculated duration : %f \n", float64(info.Frames)/float64(info.Samplerate))
	fmt.Println("")
}

func Sub_main() {

	fmt.Println("Start")
	s, _ := sndfile.GetLibVersion()
	fmt.Println(s)

	const audioFile = "data/test2.aiff"

	printInfo(audioFile)
	printImageShort(audioFile, "data/out_900.png", 900)
	printImageShort(audioFile, "data/out_500.png", 500)
	printImageShort(audioFile, "data/out_400.png", 400)
	printImageShort(audioFile, "data/out_300.png", 300)
	printImageShort(audioFile, "data/out_200.png", 200)
	printImageShort(audioFile, "data/out_150.png", 150)
	printImageShort(audioFile, "data/out_120.png", 120)
}
