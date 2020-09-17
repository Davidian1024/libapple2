# libapple2
A libretro core for Apple ][ emulation


# foundational resources
github.com/libretro/RetroArch
github.com/ivanizag/apple2

# building
# optimizations disabled
go build -gcflags '-N -l' -o libgithub.com-Davidian1024-libapple2.so -buildmode=c-shared .
go install -gcflags '-N -l' -buildmode=shared .

