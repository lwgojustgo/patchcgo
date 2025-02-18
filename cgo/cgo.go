package cgo

// #include <stdlib.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func AToByteHelp(v string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&v))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func GoStringToCString(logPath string) *C.char {
	return C.CString(logPath)
}

func ByteArrayToCString(ba []byte) *C.char {
	ba = append(ba, 0) // 防止转成C++的string后，末尾没有结束符，导致C++判断失误
	s := bytes2String(ba)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(p.Data))
}

func UIntCToGo(ptr unsafe.Pointer) uint32 {
	bodySize := uint32(*(*C.uint)(ptr))
	return bodySize
}

func FreeCStr(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}
