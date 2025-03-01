package main

/*
#cgo LDFLAGS: XDDLL.x

typedef int (*callback_t)(const char *astring, int *intp, int mycount);

int SUBPROG(callback_t cb);

extern int TESTAPI(char *astring, int *intp, int mycount);

*/
import "C"

import (
	"fmt"
	"runtime"

	"github.com/ibmruntimes/go-recordio/v2/utils"
	"github.com/indece-official/go-ebcdic"
)

//export TESTAPI
func TESTAPI(astring *C.char, intp *C.int, mycount C.int) C.int {
	utils.ThreadAsciiMode()
	str, _ := ebcdic.Decode([]byte(C.GoString(astring)), ebcdic.EBCDIC037)
	fmt.Println("A STRING: ", str)
	fmt.Println("AN INT: ", *intp)
	fmt.Println("MYCOUNT: ", mycount)
	utils.ThreadEbcdicMode()
	return C.int(400)
}

func main() {
	runtime.LockOSThread()   // lock thread
	utils.ThreadEbcdicMode() // COBOL must run in ebcdic
	ret := C.SUBPROG((C.callback_t)(C.TESTAPI))
	utils.ThreadAsciiMode()  // restore
	runtime.UnlockOSThread() // unlock thread
	fmt.Printf("Call COBOL SUBPROG returns %d\n", ret)
}
