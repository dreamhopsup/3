package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var reducemaxvecnorm2_code cu.Function

type reducemaxvecnorm2_args struct {
	arg_x       unsafe.Pointer
	arg_y       unsafe.Pointer
	arg_z       unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [6]unsafe.Pointer
}

// Wrapper for reducemaxvecnorm2 CUDA kernel, asynchronous.
func k_reducemaxvecnorm2_async(x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config, str cu.Stream) {
	if reducemaxvecnorm2_code == 0 {
		reducemaxvecnorm2_code = fatbinLoad(reducemaxvecnorm2_map, "reducemaxvecnorm2")
	}

	var a reducemaxvecnorm2_args

	a.arg_x = x
	a.argptr[0] = unsafe.Pointer(&a.arg_x)
	a.arg_y = y
	a.argptr[1] = unsafe.Pointer(&a.arg_y)
	a.arg_z = z
	a.argptr[2] = unsafe.Pointer(&a.arg_z)
	a.arg_dst = dst
	a.argptr[3] = unsafe.Pointer(&a.arg_dst)
	a.arg_initVal = initVal
	a.argptr[4] = unsafe.Pointer(&a.arg_initVal)
	a.arg_n = n
	a.argptr[5] = unsafe.Pointer(&a.arg_n)

	args := a.argptr[:]
	cu.LaunchKernel(reducemaxvecnorm2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for reducemaxvecnorm2 CUDA kernel, synchronized.
func k_reducemaxvecnorm2(x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	str := stream()
	k_reducemaxvecnorm2_async(x, y, z, dst, initVal, n, cfg, str)
	syncAndRecycle(str)
}

var reducemaxvecnorm2_map = map[int]string{0: ""}
