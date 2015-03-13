package splitter

import "fmt"
import "github.com/mkb218/gosndfile/sndfile"

func Sub_main() {
 
    fmt.Println("OK")
    
    s, _ := sndfile.GetLibVersion()
    
    fmt.Println(s)
    
    var f * sndfile.File
    var inf sndfile.Info
    
    f, err := sndfile.Open("data/test.wav", sndfile.Write, &inf)
}