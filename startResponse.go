package main

/*
#cgo CFLAGS: -I /usr/include/python3.5m
#cgo LDFLAGS: -lpython3.5m
#include <Python.h>
*/
import "C"

import "fmt"

//export goStartResponse
func goStartResponse(status *C.char, responseHeaders, excInfo *C.PyObject) *C.PyObject {
	/*
			   This func is passed as argument to the application callable,
		       arguments are
		       - status a string like "200 OK"
		       - response_headers: list of tuples (header_name, header_value)
		       - exc_info: sys.exc_info() tuple
	*/

	// parse the args to respective type
	fmt.Print("Done!!\n")
	// Extract Headers
	// if ok = C.PyList_Check(responseHeaders); !ok {
	// 	C.PyErr_SetString(C.PyExc_TypeError, "response headers should a list type")
	// 	os.Exit(2)
	// }
	return nil
}
