package main

/*
#cgo CFLAGS: -I /usr/include/python3.5m
#cgo LDFLAGS: -lpython3.5m
#include <stdio.h>
#include <Python.h>
extern void start_response(PyObject*, PyObject*);
extern PyObject* goStartResponse(char*, PyObject*, PyObject*);

void (*start_fp)(PyObject*, PyObject*);

PyObject* MyBuildValue(const char *format, PyObject *env) {
	printf("Hello");
	return Py_BuildValue(format, env, &start_fp);
}

void start_response(PyObject *self, PyObject *args) {
	int ok = 0;
	char *status;
	PyObject *exc_info, *response_headers;
	ok = PyArg_ParseTuple(args, "sO|O", status, response_headers, exc_info);
	if (ok != 0){
		PyErr_SetString(PyExc_RuntimeError, "Failed to parse args");
	}
	goStartResponse(status, response_headers, exc_info);
}

int init() {
  start_fp = &start_response;
}
*/
import "C"

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {

	appPath := flag.String("app", "", "Need import path to the application")
	flag.Parse()

	if *appPath == "" {
		log.Fatal("app import path is required")
	}

	C.Py_Initialize()
	if C.Py_IsInitialized() == 0 {
		log.Fatal("Failed to initilize python interpreter")
	}

	p := strings.Split(*appPath, ".")
	if len(p) < 2 {
		log.Fatal("app name should be of the form <module>.<app>")
	}

	modName := strings.Join(p[:len(p)-1], ".")
	a := loadPyModule(modName)
	app := C.PyObject_GetAttrString(a, C.CString(p[len(p)-1]))

	env := buildEnviron()
	C.init()
	argsObject := C.MyBuildValue(C.CString("OO"), env)
	r := C.PyObject_CallObject(app, argsObject)
	if r == nil {
		errTraceback()
		log.Fatal(nil)
	}

	fmt.Println(r)
}

func loadPyModule(module string) *C.PyObject {
	mod := C.PyImport_ImportModule(C.CString(module))
	if mod == nil {
		errTraceback()
		log.Fatalf("failed to load %s", module)
	}

	return mod
}

func errTraceback() {
	if C.PyErr_Occurred() != nil {
		C.PyErr_PrintEx(1)
	}
}
