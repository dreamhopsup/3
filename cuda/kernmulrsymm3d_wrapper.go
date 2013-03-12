package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var kernmulRSymm3D_code cu.Function

type kernmulRSymm3D_args struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftMz  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKzz unsafe.Pointer
	arg_fftKyz unsafe.Pointer
	arg_fftKxz unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_N0     int
	arg_N1     int
	arg_N2     int
	argptr     [12]unsafe.Pointer
}

// Wrapper for kernmulRSymm3D CUDA kernel, asynchronous.
func k_kernmulRSymm3D_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, fftKxz unsafe.Pointer, fftKxy unsafe.Pointer, N0 int, N1 int, N2 int, cfg *Config, str cu.Stream) {
	if kernmulRSymm3D_code == 0 {
		kernmulRSymm3D_code = cu.ModuleLoadData(kernmulRSymm3D_ptx).GetFunction("kernmulRSymm3D")
	}

	var a kernmulRSymm3D_args

	a.arg_fftMx = fftMx
	a.argptr[0] = unsafe.Pointer(&a.arg_fftMx)
	a.arg_fftMy = fftMy
	a.argptr[1] = unsafe.Pointer(&a.arg_fftMy)
	a.arg_fftMz = fftMz
	a.argptr[2] = unsafe.Pointer(&a.arg_fftMz)
	a.arg_fftKxx = fftKxx
	a.argptr[3] = unsafe.Pointer(&a.arg_fftKxx)
	a.arg_fftKyy = fftKyy
	a.argptr[4] = unsafe.Pointer(&a.arg_fftKyy)
	a.arg_fftKzz = fftKzz
	a.argptr[5] = unsafe.Pointer(&a.arg_fftKzz)
	a.arg_fftKyz = fftKyz
	a.argptr[6] = unsafe.Pointer(&a.arg_fftKyz)
	a.arg_fftKxz = fftKxz
	a.argptr[7] = unsafe.Pointer(&a.arg_fftKxz)
	a.arg_fftKxy = fftKxy
	a.argptr[8] = unsafe.Pointer(&a.arg_fftKxy)
	a.arg_N0 = N0
	a.argptr[9] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[10] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[11] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(kernmulRSymm3D_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for kernmulRSymm3D CUDA kernel, synchronized.
func k_kernmulRSymm3D(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, fftKxz unsafe.Pointer, fftKxy unsafe.Pointer, N0 int, N1 int, N2 int, cfg *Config) {
	str := stream()
	k_kernmulRSymm3D_async(fftMx, fftMy, fftMz, fftKxx, fftKyy, fftKzz, fftKyz, fftKxz, fftKxy, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

const kernmulRSymm3D_ptx = `
.version 3.0
.target sm_30
.address_size 64


.entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .f32 	%f<39>;
	.reg .pred 	%p<8>;
	.reg .s32 	%r<87>;
	.reg .s64 	%rl<54>;


	ld.param.u32 	%r1, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r2, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r3, [kernmulRSymm3D_param_11];
	.loc 2 35 1
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %tid.y;
	mad.lo.s32 	%r7, %r5, %r4, %r6;
	.loc 2 36 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r27, %r8, %r9, %r10;
	setp.lt.s32 	%p2, %r27, %r3;
	setp.lt.s32 	%p3, %r7, %r2;
	and.pred  	%p4, %p3, %p2;
	.loc 2 44 1
	setp.gt.s32 	%p5, %r1, 0;
	and.pred  	%p6, %p4, %p5;
	.loc 2 38 1
	@!%p6 bra 	BB0_6;

	ld.param.u32 	%r70, [kernmulRSymm3D_param_10];
	.loc 2 55 1
	shr.u32 	%r29, %r70, 31;
	add.s32 	%r30, %r70, %r29;
	shr.s32 	%r31, %r30, 1;
	add.s32 	%r32, %r31, 1;
	setp.lt.s32 	%p1, %r7, %r32;
	sub.s32 	%r35, %r70, %r7;
	ld.param.u32 	%r73, [kernmulRSymm3D_param_11];
	mad.lo.s32 	%r86, %r73, %r35, %r27;
	mad.lo.s32 	%r80, %r73, %r7, %r27;
	shl.b32 	%r81, %r80, 1;
	or.b32  	%r82, %r81, 1;
	mov.u32 	%r83, 0;

BB0_2:
	mov.u32 	%r84, %r86;
	mov.u32 	%r18, %r84;
	ld.param.u64 	%rl43, [kernmulRSymm3D_param_0];
	cvta.to.global.u64 	%rl16, %rl43;
	.loc 2 48 1
	mul.wide.s32 	%rl17, %r81, 4;
	add.s64 	%rl10, %rl16, %rl17;
	ld.global.f32 	%f1, [%rl10];
	.loc 2 49 1
	mul.wide.s32 	%rl18, %r82, 4;
	add.s64 	%rl19, %rl16, %rl18;
	add.s64 	%rl11, %rl19, -4;
	.loc 2 49 1
	ld.global.f32 	%f2, [%rl19];
	ld.param.u64 	%rl44, [kernmulRSymm3D_param_1];
	cvta.to.global.u64 	%rl20, %rl44;
	.loc 2 50 1
	add.s64 	%rl12, %rl20, %rl17;
	ld.global.f32 	%f3, [%rl12];
	.loc 2 51 1
	add.s64 	%rl21, %rl20, %rl18;
	add.s64 	%rl13, %rl21, -4;
	.loc 2 51 1
	ld.global.f32 	%f4, [%rl21];
	ld.param.u64 	%rl45, [kernmulRSymm3D_param_2];
	cvta.to.global.u64 	%rl22, %rl45;
	.loc 2 52 1
	add.s64 	%rl14, %rl22, %rl17;
	ld.global.f32 	%f5, [%rl14];
	.loc 2 53 1
	add.s64 	%rl23, %rl22, %rl18;
	add.s64 	%rl15, %rl23, -4;
	.loc 2 53 1
	ld.global.f32 	%f6, [%rl23];
	.loc 2 55 1
	@%p1 bra 	BB0_4;

	ld.param.u64 	%rl50, [kernmulRSymm3D_param_6];
	cvta.to.global.u64 	%rl24, %rl50;
	.loc 2 67 1
	mul.wide.s32 	%rl25, %r18, 4;
	add.s64 	%rl26, %rl24, %rl25;
	ld.global.f32 	%f13, [%rl26];
	neg.f32 	%f38, %f13;
	ld.param.u64 	%rl53, [kernmulRSymm3D_param_8];
	cvta.to.global.u64 	%rl27, %rl53;
	.loc 2 69 1
	add.s64 	%rl28, %rl27, %rl25;
	ld.global.f32 	%f14, [%rl28];
	neg.f32 	%f37, %f14;
	mov.u32 	%r85, %r18;
	bra.uni 	BB0_5;

BB0_4:
	ld.param.u64 	%rl49, [kernmulRSymm3D_param_6];
	cvta.to.global.u64 	%rl29, %rl49;
	.loc 2 59 1
	mul.wide.s32 	%rl30, %r80, 4;
	add.s64 	%rl31, %rl29, %rl30;
	ld.global.f32 	%f38, [%rl31];
	ld.param.u64 	%rl52, [kernmulRSymm3D_param_8];
	cvta.to.global.u64 	%rl32, %rl52;
	.loc 2 61 1
	add.s64 	%rl33, %rl32, %rl30;
	ld.global.f32 	%f37, [%rl33];
	ld.param.u32 	%r69, [kernmulRSymm3D_param_10];
	mad.lo.s32 	%r50, %r83, %r69, %r7;
	ld.param.u32 	%r72, [kernmulRSymm3D_param_11];
	.loc 2 46 1
	mad.lo.s32 	%r20, %r50, %r72, %r27;
	mov.u32 	%r85, %r20;

BB0_5:
	mov.u32 	%r21, %r85;
	ld.param.u64 	%rl51, [kernmulRSymm3D_param_7];
	cvta.to.global.u64 	%rl34, %rl51;
	mul.wide.s32 	%rl35, %r21, 4;
	add.s64 	%rl36, %rl34, %rl35;
	ld.param.u64 	%rl48, [kernmulRSymm3D_param_5];
	cvta.to.global.u64 	%rl37, %rl48;
	add.s64 	%rl38, %rl37, %rl35;
	ld.param.u64 	%rl47, [kernmulRSymm3D_param_4];
	cvta.to.global.u64 	%rl39, %rl47;
	add.s64 	%rl40, %rl39, %rl35;
	ld.param.u64 	%rl46, [kernmulRSymm3D_param_3];
	cvta.to.global.u64 	%rl41, %rl46;
	add.s64 	%rl42, %rl41, %rl35;
	ld.global.f32 	%f15, [%rl38];
	ld.global.f32 	%f16, [%rl40];
	ld.global.f32 	%f17, [%rl42];
	.loc 2 72 1
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.f32 	%f20, [%rl36];
	.loc 2 72 1
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	st.global.f32 	[%rl10], %f21;
	.loc 2 73 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rl11+4], %f24;
	.loc 2 74 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	st.global.f32 	[%rl12], %f27;
	.loc 2 75 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rl13+4], %f30;
	.loc 2 76 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	st.global.f32 	[%rl14], %f33;
	.loc 2 77 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rl15+4], %f36;
	ld.param.u32 	%r68, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r71, [kernmulRSymm3D_param_11];
	mul.lo.s32 	%r65, %r71, %r68;
	mad.lo.s32 	%r22, %r71, %r68, %r18;
	shl.b32 	%r66, %r65, 1;
	add.s32 	%r82, %r82, %r66;
	add.s32 	%r81, %r81, %r66;
	mad.lo.s32 	%r80, %r71, %r68, %r80;
	.loc 2 44 18
	add.s32 	%r83, %r83, 1;
	ld.param.u32 	%r67, [kernmulRSymm3D_param_9];
	.loc 2 44 1
	setp.lt.s32 	%p7, %r83, %r67;
	mov.u32 	%r86, %r22;
	.loc 2 44 1
	@%p7 bra 	BB0_2;

BB0_6:
	.loc 2 79 2
	ret;
}


`
