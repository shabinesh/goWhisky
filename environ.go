package main

/*
#cgo CFLAGS: -I /usr/include/python3.5m
#cgo LDFLAGS: -lpython3.5m
#include <Python.h>
*/
import "C"

func buildEnviron() *C.PyObject {
	d := C.PyDict_New()
	C.PyDict_SetItemString(d, C.CString("SERVER_NAME"), C.PyUnicode_FromString(C.CString("combat_server")))
	C.PyDict_SetItemString(d, C.CString("SERVER_PORT"), C.PyUnicode_FromString(C.CString("8000")))
	C.PyDict_SetItemString(d, C.CString("REQUEST_METHOD"), C.PyUnicode_FromString(C.CString("GET")))
	C.PyDict_SetItemString(d, C.CString("PATH_INFO"), C.PyUnicode_FromString(C.CString("/")))
	C.PyDict_SetItemString(d, C.CString("QUERY_STRING"), C.PyUnicode_FromString(C.CString("")))

	C.PyDict_SetItemString(d, C.CString("wsgi.url_scheme"), C.PyUnicode_FromString(C.CString("http")))
	return d
}
