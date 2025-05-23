/*
 * util_fast.S
 *
 * Portions are taken from the "copies and fills" library by Simon Hall
 * which is licensed under the GNU Lesser General Public License version 2.1
 *
 * Circle - A C++ bare metal environment for Raspberry Pi
 * Copyright (C) 2016-2021  R. Stange <rsta2@o2online.de>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

	.text

#if AARCH == 32

	.globl	memset
	.type   memset, %function
memset:
	tst	r0, #3
	bne	2f
	cmp	r2, #16
	blo	2f

	orr	r3, r1, r1, lsl #8
	orr	r12, r3, r3, lsl #16

	mov	r3, r0
1:	str	r12, [r3]
	str	r12, [r3, #4]
	str	r12, [r3, #8]
	str	r12, [r3, #12]
	add	r3, r3, #16
	sub	r2, r2, #16
	cmp	r2, #15
	bhi	1b
	b	3f

2:	mov	r3, r0
3:	cmp	r2, #0
	bxeq	lr

4:	strb	r1, [r3], #1
	subs	r2, r2, #1
	bne	4b
	bx	lr

	.globl	memcpy
	.type   memcpy, %function
memcpy:
	push	{r0}

	cmp	r2, #127
	bls	2f
	tst	r1, #3
	bne	2f
	tst	r0, #3
	bne	2f

	push	{r4-r10}
1:	ldmia	r1!, {r3-r10}
	sub	r2, #8*4
	stmia	r0!, {r3-r10}
	pld	[r1, #8*4*2]
	cmp	r2, #8*4-1
	bhi	1b
	pop	{r4-r10}

2:	cmp	r2, #0
	beq	4f

3:	ldrb	r3, [r1], #1
	subs	r2, #1
	strb	r3, [r0], #1
	bne	3b

4:	pop	{r0}
	bx	lr

#else

	.globl	memset
	.type   memset, %function
memset:
	cbz	x2, 3f
	mov	x10, x0

1:	tst	x10, #0x1f
	b.eq	4f

2:	strb	w1, [x10], #1
	sub  	x2, x2, #1
	cbnz	x2, 1b

3:	ret

4:	lsr	x9, x2, #5
	and	x2, x2, #0x1f
	cbz	x9, 6f
	and	x1, x1, #0xff
	mov	x8, #0x0101010101010101
	mul	x8, x8, x1

5:	sub	x9, x9, #1
	stp	x8, x8, [x10], #0x10
	stp	x8, x8, [x10], #0x10
	cbnz	x9, 5b
6:	cbz	x2, 3b
	b	2b

	.globl	memcpy
	.type   memcpy, %function
memcpy:
	mov	x8, x0

	cmp	x2, #127
	b.ls	2f
	tst	x1, #7
	b.ne	2f
	tst	x1, #7
	b.ne	2f

	mov	x3, #64
1:	ldp	x4, x5, [x1], #16
	ldp	x6, x7, [x1], #16
	sub	x2, x2, #32
	stp	x4, x5, [x0], #16
	stp	x6, x7, [x0], #16
	prfm	pldl1strm, [x1, x3]
	cmp	x2, #32-1
	b.hi	1b

2:	cmp	x2, #0
	b.eq	4f

3:	ldrb	w3, [x1], #1
	subs	x2, x2, #1
	strb	w3, [x0], #1
	b.ne	3b

4:	mov	x0, x8
	ret

#endif

/* End */
