package splitter

import "fmt"
//import "encoding/json"
import "github.com/mkb218/gosndfile/sndfile"


func printInfo(inf * sndfile.Info) {
    
    fmt.Printf("Frames/Samples : %d \n", inf.Frames)
    fmt.Printf("Sample rate : %d hz \n", inf.Samplerate)
    fmt.Printf("Number of channel : %d \n", inf.Channels)
    fmt.Printf("Calculated duration : %f \n", float64(inf.Frames) / float64(inf.Samplerate))
    fmt.Println("")
}

func printBegin() {
 
    var f * sndfile.File
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
}