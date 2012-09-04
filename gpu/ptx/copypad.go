package ptx

const COPYPAD = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Wed Aug  1 02:51:19 2012 (1343782279)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00000508_00000000-9_copypad.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/copypad.cu"

.visible .entry copypad(
	.param .u64 copypad_param_0,
	.param .u32 copypad_param_1,
	.param .u32 copypad_param_2,
	.param .u32 copypad_param_3,
	.param .u64 copypad_param_4,
	.param .u32 copypad_param_5,
	.param .u32 copypad_param_6,
	.param .u32 copypad_param_7,
	.param .u32 copypad_param_8,
	.param .u32 copypad_param_9,
	.param .u32 copypad_param_10
)
{
	.reg .pred 	%p<10>;
	.reg .s32 	%r<40>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copypad_param_0];
	ld.param.u32 	%r19, [copypad_param_2];
	ld.param.u32 	%r20, [copypad_param_3];
	ld.param.u64 	%rd4, [copypad_param_4];
	ld.param.u32 	%r21, [copypad_param_5];
	ld.param.u32 	%r22, [copypad_param_6];
	ld.param.u32 	%r23, [copypad_param_7];
	ld.param.u32 	%r24, [copypad_param_8];
	ld.param.u32 	%r25, [copypad_param_9];
	ld.param.u32 	%r26, [copypad_param_10];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 17 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 18 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 20 1
	setp.ge.s32 	%p1, %r4, %r23;
	setp.ge.s32 	%p2, %r8, %r22;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_4;

	.loc 2 29 1
	setp.lt.s32 	%p4, %r4, %r20;
	setp.lt.s32 	%p5, %r8, %r19;
	and.pred  	%p6, %p5, %p4;
	.loc 2 41 1
	setp.gt.s32 	%p7, %r21, 0;
	.loc 2 29 1
	and.pred  	%p8, %p6, %p7;
	@!%p8 bra 	BB0_4;
	bra.uni 	BB0_2;

BB0_2:
	.loc 2 41 1
	add.s32 	%r28, %r3, %r26;
	mad.lo.s32 	%r29, %r1, %r2, %r28;
	add.s32 	%r30, %r7, %r25;
	mad.lo.s32 	%r31, %r5, %r6, %r30;
	mad.lo.s32 	%r32, %r24, %r19, %r31;
	mad.lo.s32 	%r38, %r20, %r32, %r29;
	mul.lo.s32 	%r10, %r20, %r19;
	mad.lo.s32 	%r37, %r23, %r8, %r4;
	mul.lo.s32 	%r12, %r23, %r22;
	mov.u32 	%r39, 0;

BB0_3:
	.loc 2 51 1
	mul.wide.s32 	%rd5, %r37, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mul.wide.s32 	%rd7, %r38, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;
	.loc 2 41 1
	add.s32 	%r38, %r38, %r10;
	add.s32 	%r37, %r37, %r12;
	.loc 2 41 18
	add.s32 	%r39, %r39, 1;
	.loc 2 41 1
	setp.lt.s32 	%p9, %r39, %r21;
	@%p9 bra 	BB0_3;

BB0_4:
	.loc 2 53 2
	ret;
}


`
