# libapple2
A libretro core for Apple ][ emulation


# foundational resources
github.com/libretro/RetroArch
github.com/ivanizag/apple2

# building
# optimizations disabled
go build -o libapple2.so -buildmode=c-shared .

# notes
Initial gccgo builds always segfaulted.  After switching to golang.org's compiler the core loads and details passed from retro_get_system_info() appear in the GUI.

# log
- 2020-09-20: Current builds failing with latest changes revolving around bridge_video_refresh_callback()
