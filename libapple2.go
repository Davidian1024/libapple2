//package libapple2
package main

import (
	"fmt"
	"unsafe"
	//gopointer "github.com/mattn/go-pointer"
)

/*
#include <stdlib.h>
#include "libretro.h"

typedef const void cvoid_t;
typedef const char cchar_t;
typedef const struct retro_game_info cretro_game_info_t;
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

	fmt.Println("need_fullpath:", info.need_fullpath)
	fmt.Println("valid_extensions:", C.GoString(info.valid_extensions))
	fmt.Println("library_version:", C.GoString(info.library_version))
	fmt.Println("library_name:", C.GoString(info.library_name))
	fmt.Println("block_extract:", info.block_extract)
	fmt.Println("libapple2 retro_get_system_info() end")
	return
}

//export retro_init
func retro_init() {
	fmt.Println("libapple2 retro_init() start")
	fmt.Println("libapple2 retro_init() end")
}

//export retro_deinit
func retro_deinit() {
	fmt.Println("libapple2 retro_deinit() start")
	fmt.Println("libapple2 retro_deinit() end")
}

//export retro_api_version
func retro_api_version() C.unsigned {
	fmt.Println("libapple2 retro_api_version() start")
	fmt.Println("libapple2 retro_api_version() end")
	return C.unsigned(1)
}

//export retro_get_system_av_info
func retro_get_system_av_info(info *C.struct_retro_system_av_info) {
	fmt.Println("libapple2 retro_get_system_av_info() start")
	fmt.Println("libapple2 retro_get_system_av_info() end")
}

//export set_environment_callback
func set_environment_callback() C.bool{
	fmt.Println("libapple2 set_environment_callback() start")
	fmt.Println("libapple2 set_environment_callback() end")
	return false
}

//export set_video_refresh_callback
func set_video_refresh_callback() C.bool {
	fmt.Println("libapple2 set_video_refresh_callback() start")
	fmt.Println("libapple2 set_video_refresh_callback() end")
	return false
}

//export set_audio_sample_callback
func set_audio_sample_callback() {
	fmt.Println("libapple2 set_audio_sample_callback() start")
	fmt.Println("libapple2 set_audio_sample_callback() end")
}

//export set_audio_sample_batch_callback
func set_audio_sample_batch_callback() {
	fmt.Println("libapple2 set_audio_sample_batch_callback() start")
	fmt.Println("libapple2 set_audio_sample_batch_callback() end")
}

//export set_input_poll_callback
func set_input_poll_callback() {
	fmt.Println("libapple2 set_input_poll_callback() start")
	fmt.Println("libapple2 set_input_poll_callback() end")
}

//export set_input_state_callback
func set_input_state_callback() {
	fmt.Println("libapple2 set_input_state_callback() start")
	fmt.Println("libapple2 set_input_state_callback() end")
}

//export retro_set_controller_port_device
func retro_set_controller_port_device(port C.unsigned, device C.unsigned) {
	fmt.Println("libapple2 retro_set_controller_port_device() start")
	fmt.Println("libapple2 retro_set_controller_port_device() end")
}

//export retro_reset
func retro_reset() {
	fmt.Println("libapple2 retro_reset() start")
	fmt.Println("libapple2 retro_reset() end")
}

var running = false

//export retro_run
func retro_run() {
	if running {
	} else {
		fmt.Println("libapple2 retro_run() start")
		fmt.Println("libapple2 retro_run() setting running to true to supress these messages")
		running = true
		fmt.Println("libapple2 retro_run() end")
	}
}

//export retro_serialize_size
func retro_serialize_size() C.size_t {
	fmt.Println("libapple2 retro_serialize_size() start")
	fmt.Println("libapple2 retro_serialize_size() end")
	return C.size_t(0)
}

//export retro_serialize
func retro_serialize(data unsafe.Pointer, size C.size_t) C.bool {
	fmt.Println("libapple2 retro_serialize() start")
	fmt.Println("libapple2 retro_serialize() end")
	return false
}

//export retro_unserialize
func retro_unserialize(data *C.cvoid_t, size C.size_t) C.bool {
	fmt.Println("libapple2 retro_unserialize() start")
	fmt.Println("libapple2 retro_unserialize() end")
	return false
}

//export retro_cheat_reset
func retro_cheat_reset() {
	fmt.Println("libapple2 retro_cheat_reset() start")
	fmt.Println("libapple2 retro_cheat_reset() end")
}

//export retro_cheat_set
func retro_cheat_set(index C.unsigned, enabled C.bool, code *C.cchar_t) {
	fmt.Println("libapple2 retro_cheat_set() start")
	fmt.Println("libapple2 retro_cheat_set() end")
}

//export retro_load_game
func retro_load_game(game *C.cretro_game_info_t) C.bool {
	fmt.Println("libapple2 retro_load_game() start")
	fmt.Println("libapple2 retro_load_game() end")
	return true
}

//export retro_load_game_special
func retro_load_game_special(game_type C.unsigned, info *C.cretro_game_info_t, num_info C.size_t) C.bool {
	fmt.Println("libapple2 retro_load_game_special() start")
	fmt.Println("libapple2 retro_load_game_special() end")
	return false
}

//export retro_unload_game
func retro_unload_game() {
	fmt.Println("libapple2 retro_unload_game() start")
	fmt.Println("libapple2 retro_unload_game() end")
}

//export retro_get_region
func retro_get_region() C.unsigned {
	fmt.Println("libapple2 retro_get_region() start")
	fmt.Println("libapple2 retro_get_region() end")
	return C.unsigned(0)
}

//export retro_get_memory_data
func retro_get_memory_data(id C.unsigned) unsafe.Pointer {
	fmt.Println("libapple2 retro_get_memory_data() start")
	fmt.Println("libapple2 retro_get_memory_data() end")
	return nil
}

//export retro_get_memory_size
func retro_get_memory_size(id C.unsigned) C.size_t {
	fmt.Println("libapple2 retro_get_memory_size() start")
	fmt.Println("libapple2 retro_get_memory_size() end")
	return C.size_t(0)
}

func main() {
	fmt.Println("libapple2 main() start")
	fmt.Println("libapple2 main() end")
}
