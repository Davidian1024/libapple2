#include <stdio.h>
#include "libretro.h"

extern bool set_environment_callback();
void retro_set_environment(retro_environment_t callback) {
	printf("libapple2_cgo retro_set_environment() start\n");
	//set_environment_callback = callback;
	//callback = set_environment_callback;
	set_environment_callback();
	printf("libapple2_cgo retro_set_environment() end\n");
}

extern void set_video_refresh_callback();
void retro_set_video_refresh(retro_video_refresh_t callback) {
	printf("libapple2_cgo retro_set_video_refresh() start\n");
	//callback = set_video_refresh_callback;
	set_video_refresh_callback(callback);
	//v = callback; //Need to figure out how to make it so we can call the video_refresh_call back from Go space
	printf("libapple2_cgo retro_set_video_refresh() end\n");
}

extern void set_audio_sample_callback();
void retro_set_audio_sample(retro_audio_sample_t callback) {
	printf("libapple2_cgo retro_set_audio_sample() start\n");
	set_audio_sample_callback();
	printf("libapple2_cgo retro_set_audio_sample() end\n");
}

extern void set_audio_sample_batch_callback();
void retro_set_audio_sample_batch(retro_audio_sample_batch_t callback) {
	printf("libapple2_cgo retro_set_audio_sample_batch() start\n");
	set_audio_sample_batch_callback();
	printf("libapple2_cgo retro_set_audio_sample_batch() end\n");
}

extern void set_input_poll_callback();
void retro_set_input_poll(retro_input_poll_t callback) {
	printf("libapple2_cgo retro_set_input_poll() start\n");
	set_input_poll_callback();
	printf("libapple2_cgo retro_set_input_poll() end\n");
}

extern void set_input_state_callback();
void retro_set_input_state(retro_input_state_t callback) {
	printf("libapple2_cgo retro_set_input_state() start\n");
	set_input_state_callback();
	printf("libapple2_cgo retro_set_input_state() end\n");
}
