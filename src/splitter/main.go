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
    /*
    
    fmt.Println(inf)
    fmt.Println(err)
    fmt.Println(" --- ")
    fmt.Println(*f)
    
    var buff []int32
    var total int64
    var i int64
    var l int64
    
    var part int64 = int64(inf.Samplerate) / int64(10)
    var parts int64 = inf.Frames / part
    var pbuff []int64
    var u int = 0
    
    fmt.Printf("part = %d, parts = %d\n", part, parts)
    
    buff = make([]int32, part)
    pbuff = make([]int64, parts + 2)
      
    i, _ = f.ReadFrames(buff)  
    for ; i != 0; total += i {
    
        
        sum := int64(0)
        
        for l = 0; l < i ; l++ {
            sum += int64(buff[l])
               
        }
        fmt.Printf("%d u read\n", u)
        pbuff[u] = sum / i
        u++
        i, _ = f.ReadFrames(buff)
    }
    
    fmt.Printf("%d bytes read\n", total)
    //fmt.Println(buff)
    */
}