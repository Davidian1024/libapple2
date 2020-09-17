//package libapple2
package main

import "fmt"
//import "unsafe"

/*
#include <stdlib.h>
#include "libretro.h"
*/
import (
	"C"
)

//export retro_get_system_info
func retro_get_system_info(info *C.struct_retro_system_info) {
	fmt.Println("libapple2 retro_get_system_info() start")
	info.need_fullpath    = false;
	info.valid_extensions = C.CString("dsk")
	info.library_version = C.CString("0.1")
	info.library_name = C.CString("libapple2")
	info.block_extract    = false;

	fmt.Println(info.need_fullpath)
	fmt.Println(*info.valid_extensions)
	fmt.Println(*info.library_version)
	fmt.Println(*info.library_name)
	fmt.Println(info.block_extract)
	fmt.Println("libapple2 retro_get_system_info() end")
	return
}

//func main() {
//	fmt.Println("libapple2")
//}

//export retro_init
func retro_init() {
	fmt.Println("libapple2 retro_init()")
}

func main() {}
