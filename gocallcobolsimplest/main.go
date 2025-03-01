package main

/*
#cgo LDFLAGS: XDDLL.x
int MYPROG(void);
*/
import "C"

import (
	"fmt"
	"runtime"

	"github.com/ibmruntimes/go-recordio/v2/utils"
)

func main() {
	runtime.LockOSThread()   // lock thread
	utils.ThreadEbcdicMode() // COBOL must run in ebcdic
	ret := C.MYPROG()
	utils.ThreadAsciiMode()  // restore
	runtime.UnlockOSThread() // unlock thread
	fmt.Printf("Call COBOL MYPROG returns %d\n", ret)
}
