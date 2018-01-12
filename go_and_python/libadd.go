// run it with:
// 1. go build -buildmode=c-archive libadd.go
//    result: libadd.a, libadd.h
cat libadd.def
EXPORTS
        Add
// 2. gcc -shared -pthread -o libadd.dll libadd.a -lWinMM -lntdll -lWS2_32
//    gcc -shared -o libadd.dll libadd.def libadd.a -Wl,--allow-multiple-definition -static -lstdc++ -lwinmm -lntdll -lWs2_32
//    result: libadd.dll

//libadd.go
package main

import "C"

//export Add
func Add(left, right int) int {
	return left + right
}

func main() {
}
