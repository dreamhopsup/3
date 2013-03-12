package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var reducemaxdiff_code cu.Function

type reducemaxdiff_args struct {
	arg_src1    unsafe.Pointer
	arg_src2    unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [5]unsafe.Pointer
}

// Wrapper for reducemaxdiff CUDA kernel, asynchronous.
func k_reducemaxdiff_async(src1 unsafe.Pointer, src2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config, str cu.Stream) {
	if reducemaxdiff_code == 0 {
		reducemaxdiff_code = fatbinLoad(reducemaxdiff_map, "reducemaxdiff")
	}

	var a reducemaxdiff_args

	a.arg_src1 = src1
	a.argptr[0] = unsafe.Pointer(&a.arg_src1)
	a.arg_src2 = src2
	a.argptr[1] = unsafe.Pointer(&a.arg_src2)
	a.arg_dst = dst
	a.argptr[2] = unsafe.Pointer(&a.arg_dst)
	a.arg_initVal = initVal
	a.argptr[3] = unsafe.Pointer(&a.arg_initVal)
	a.arg_n = n
	a.argptr[4] = unsafe.Pointer(&a.arg_n)

	args := a.argptr[:]
	cu.LaunchKernel(reducemaxdiff_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for reducemaxdiff CUDA kernel, synchronized.
func k_reducemaxdiff(src1 unsafe.Pointer, src2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	str := stream()
	k_reducemaxdiff_async(src1, src2, dst, initVal, n, cfg, str)
	syncAndRecycle(str)
}

var reducemaxdiff_map = map[int]string{0: ""}
