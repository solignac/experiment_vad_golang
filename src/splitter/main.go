package splitter

import "fmt"
import "github.com/mkb218/gosndfile/sndfile"

func Sub_main() {
 
    fmt.Println("OK")
    
    s, _ := sndfile.GetLibVersion()
    
    fmt.Println(s)
}