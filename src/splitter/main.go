package splitter

import "fmt"
import "github.com/mkb218/gosndfile/sndfile"

func Sub_main() {
 
    fmt.Println("OK")
    
    s, _ := sndfile.GetLibVersion()
    
    fmt.Println(s)
    
    var f * sndfile.File
    var inf sndfile.Info
    
    inf.Format = 0
    f, err := sndfile.Open("data/test2.aiff", sndfile.Read, &inf)
    
    fmt.Println(*f)
    fmt.Println(err)
    
    
    
}