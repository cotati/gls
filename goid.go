package gls

import (
	"runtime"
	"unsafe"

	"github.com/v2pro/plz/reflect2"
)

// offset
var goidOffset uintptr

func init() {
	offsetMap := map[string]uintptr{
		"go1.9":    152,
		"go1.9.1":  152,
		"go1.9.2":  152,
		"go1.9.3":  152,
		"go1.9.4":  152,
		"go1.9.5":  152,
		"go1.9.6":  152,
		"go1.9.7":  152,
		"go1.10":   152,
		"go1.10.1": 152,
		"go1.10.2": 152,
		"go1.10.3": 152,
		"go1.10.4": 152,
		"go1.11":   152,
		"go1.11.1": 152,
		"go1.11.2": 152,
		"go1.12.7": 152,
	}

	version := runtime.Version()
	o, exists := offsetMap[version]
	if exists {
		goidOffset = o
	} else {
		panic("not supported golang version:" + version)
	}
}

// GoID returns the goroutine id of current goroutine
func GoID() int64 {
	g := getg()
	p_goid := (*int64)(unsafe.Pointer(uintptr(g) + goidOffset))
	return *p_goid
}

// 获取offset
func getOffset() uintptr {
	gType := reflect2.TypeByName("runtime.g").(reflect2.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	goidField := gType.FieldByName("goid")
	offset := goidField.Offset()
	return offset
}

func getg() uintptr
