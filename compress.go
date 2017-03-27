package tiger

import (
	"encoding/binary"
	"unsafe"
)

var littleEndian bool

func init() {
	x := uint32(0x04030201)
	y := [4]byte{0x1, 0x2, 0x3, 0x4}
	littleEndian = *(*[4]byte)(unsafe.Pointer(&x)) == y
}

func pass(a *uint64, b *uint64, c *uint64, x []uint64, mul uint64) {
	round(a, b, c, x[0], mul)
	round(b, c, a, x[1], mul)
	round(c, a, b, x[2], mul)
	round(a, b, c, x[3], mul)
	round(b, c, a, x[4], mul)
	round(c, a, b, x[5], mul)
	round(a, b, c, x[6], mul)
	round(b, c, a, x[7], mul)
}

func round(a *uint64, b *uint64, c *uint64, x uint64, mul uint64) {
	*c ^= x
	*a -= t1[*c&0xff] ^ t2[(*c>>16)&0xff] ^ t3[(*c>>32)&0xff] ^ t4[(*c>>48)&0xff]
	*b += t4[(*c>>8)&0xff] ^ t3[(*c>>24)&0xff] ^ t2[(*c>>40)&0xff] ^ t1[(*c>>56)&0xff]
	*b *= mul
}

func keySchedule(x []uint64) {
	x[0] -= x[7] ^ 0xa5a5a5a5a5a5a5a5
	x[1] ^= x[0]
	x[2] += x[1]
	x[3] -= x[2] ^ ((^x[1]) << 19)
	x[4] ^= x[3]
	x[5] += x[4]
	x[6] -= x[5] ^ ((^x[4]) >> 23)
	x[7] ^= x[6]
	x[0] += x[7]
	x[1] -= x[0] ^ ((^x[7]) << 19)
	x[2] ^= x[1]
	x[3] += x[2]
	x[4] -= x[3] ^ ((^x[2]) >> 23)
	x[5] ^= x[4]
	x[6] += x[5]
	x[7] -= x[6] ^ 0x0123456789abcdef
}

func compress(d *digest, data []byte) {
	// save_abc
	aa := d.a
	bb := d.b
	cc := d.c

	var x []uint64
	if littleEndian {
		x = []uint64{
			binary.LittleEndian.Uint64(data[0:8]),
			binary.LittleEndian.Uint64(data[8:16]),
			binary.LittleEndian.Uint64(data[16:24]),
			binary.LittleEndian.Uint64(data[24:32]),
			binary.LittleEndian.Uint64(data[32:40]),
			binary.LittleEndian.Uint64(data[40:48]),
			binary.LittleEndian.Uint64(data[48:56]),
			binary.LittleEndian.Uint64(data[56:64]),
		}
	} else {
		x = []uint64{
			binary.BigEndian.Uint64(data[0:8]),
			binary.BigEndian.Uint64(data[8:16]),
			binary.BigEndian.Uint64(data[16:24]),
			binary.BigEndian.Uint64(data[24:32]),
			binary.BigEndian.Uint64(data[32:40]),
			binary.BigEndian.Uint64(data[40:48]),
			binary.BigEndian.Uint64(data[48:56]),
			binary.BigEndian.Uint64(data[56:64]),
		}
	}

	pass(&d.a, &d.b, &d.c, x, 5)
	keySchedule(x)
	pass(&d.c, &d.a, &d.b, x, 7)
	keySchedule(x)
	pass(&d.b, &d.c, &d.a, x, 9)

	// feedfoward
	d.a ^= aa
	d.b -= bb
	d.c += cc
}
