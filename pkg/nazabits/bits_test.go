// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazabits_test

import (
	"testing"

	"github.com/q191201771/naza/pkg/assert"
	"github.com/q191201771/naza/pkg/nazabits"
)

func TestGetBit8(t *testing.T) {
	assert.Equal(t, uint8(0), nazabits.GetBit8(0, 0))
	assert.Equal(t, uint8(1), nazabits.GetBit8(1<<1, 1))
	assert.Equal(t, uint8(1), nazabits.GetBit8(1<<1+1<<2, 2))
	assert.Equal(t, uint8(1), nazabits.GetBit8(1<<3+1<<4+1<<5, 3))
	assert.Equal(t, uint8(0), nazabits.GetBit8(0, 4))
	assert.Equal(t, uint8(1), nazabits.GetBit8(1<<5, 5))
	assert.Equal(t, uint8(0), nazabits.GetBit8(1+1<<5+1<<7, 6))
	assert.Equal(t, uint8(0), nazabits.GetBit8(0, 7))
}

func TestGetBits8(t *testing.T) {
	// 0110 1001 = 1 + 8 + 32 + 64 = 105
	v := uint8(105)
	assert.Equal(t, uint8(1), nazabits.GetBits8(v, 0, 1))
	assert.Equal(t, uint8(0), nazabits.GetBits8(v, 7, 1))
	assert.Equal(t, uint8(1), nazabits.GetBits8(v, 0, 2))
	assert.Equal(t, uint8(1), nazabits.GetBits8(v, 6, 2))
	assert.Equal(t, uint8(1), nazabits.GetBits8(v, 0, 3))
	assert.Equal(t, uint8(5), nazabits.GetBits8(v, 3, 3))
	assert.Equal(t, uint8(10), nazabits.GetBits8(v, 2, 4))
	assert.Equal(t, uint8(26), nazabits.GetBits8(v, 2, 5))
	assert.Equal(t, uint8(26), nazabits.GetBits8(v, 2, 6))
	assert.Equal(t, uint8(105), nazabits.GetBits8(v, 0, 7))
	assert.Equal(t, uint8(105), nazabits.GetBits8(v, 0, 8))
}

func TestGetBit16(t *testing.T) {
	v := []byte{48, 57} // 12345 {0011 0000, 0011 1001}
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 0))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 1))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 2))
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 3))
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 4))
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 5))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 6))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 7))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 8))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 9))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 10))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 11))
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 12))
	assert.Equal(t, uint8(1), nazabits.GetBit16(v, 13))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 14))
	assert.Equal(t, uint8(0), nazabits.GetBit16(v, 15))
}

func TestGetBits16(t *testing.T) {
	v := []byte{48, 57} // 12345 {0011 0000, 0011 1001}
	assert.Equal(t, uint16(12345), nazabits.GetBits16(v, 0, 16))
	assert.Equal(t, uint16(1), nazabits.GetBits16(v, 0, 1))
	assert.Equal(t, uint16(0), nazabits.GetBits16(v, 1, 2))
	assert.Equal(t, uint16(6), nazabits.GetBits16(v, 2, 3))
	assert.Equal(t, uint16(7), nazabits.GetBits16(v, 3, 4))
	assert.Equal(t, uint16(3), nazabits.GetBits16(v, 4, 5))
	assert.Equal(t, uint16(1), nazabits.GetBits16(v, 5, 6))
	assert.Equal(t, uint16(64), nazabits.GetBits16(v, 6, 7))
	assert.Equal(t, uint16(96), nazabits.GetBits16(v, 7, 8))
	assert.Equal(t, uint16(0), nazabits.GetBits16(v, 8, 1))
	assert.Equal(t, uint16(0), nazabits.GetBits16(v, 9, 2))
	assert.Equal(t, uint16(4), nazabits.GetBits16(v, 10, 3))
	assert.Equal(t, uint16(6), nazabits.GetBits16(v, 11, 4))
	assert.Equal(t, uint16(3), nazabits.GetBits16(v, 12, 3))
	assert.Equal(t, uint16(1), nazabits.GetBits16(v, 13, 2))
	assert.Equal(t, uint16(0), nazabits.GetBits16(v, 14, 1))
	assert.Equal(t, uint16(0), nazabits.GetBits16(v, 15, 1))
}

func TestBitReader_ReadBit(t *testing.T) {
	res := []uint8{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1}
	v := []byte{48, 57} // 12345 {0011 0000, 0011 1001}
	br := nazabits.NewBitReader(v)
	for _, b := range res {
		assert.Equal(t, b, br.ReadBit())
	}
}

func TestBitReader_ReadBits8(t *testing.T) {
	v := []byte{48, 57, 48, 57}
	br := nazabits.NewBitReader(v)
	assert.Equal(t, uint8(0), br.ReadBits8(1))
	assert.Equal(t, uint8(1), br.ReadBits8(2))
	assert.Equal(t, uint8(4), br.ReadBits8(3))
	assert.Equal(t, uint8(0), br.ReadBits8(4))
	assert.Equal(t, uint8(28), br.ReadBits8(5))
	assert.Equal(t, uint8(38), br.ReadBits8(6))
	assert.Equal(t, uint8(3), br.ReadBits8(7))

	br = nazabits.NewBitReader(v)
	// {0011 0000, 0011 1001, 0 011 0000, 0011 1001}
	assert.Equal(t, uint8(48), br.ReadBits8(8))
}

func TestBitReader_ReadBits16(t *testing.T) {
	v := []byte{48, 57, 48, 57}
	br := nazabits.NewBitReader(v)
	assert.Equal(t, uint16(48), br.ReadBits16(8))
	assert.Equal(t, uint16(1), br.ReadBits16(3))
	assert.Equal(t, uint16(25), br.ReadBits16(5))
	assert.Equal(t, uint16(12345), br.ReadBits16(16))
}

func TestBitReader_ReadBits32(t *testing.T) {
	v := []byte{48, 57, 48, 57}
	br := nazabits.NewBitReader(v)
	assert.Equal(t, uint32(809054265), br.ReadBits32(32))
}

func TestBitReader_ReadBytes(t *testing.T) {
	v := []byte{48, 57}
	br := nazabits.NewBitReader(v)
	assert.Equal(t, v, br.ReadBytes(2))
}

func TestBitWriter_WriteBit(t *testing.T) {
	v := make([]byte, 2)
	bw := nazabits.NewBitWriter(v)
	bs := []uint8{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1}
	for _, b := range bs {
		bw.WriteBit(b)
	}
	assert.Equal(t, uint8(48), v[0])
	assert.Equal(t, uint8(57), v[1])

	v = make([]byte, 2)
	bw = nazabits.NewBitWriter(v)
	bs = []uint8{2, 4, 3, 5, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1}
	for _, b := range bs {
		bw.WriteBit(b)
	}
	assert.Equal(t, uint8(48), v[0])
	assert.Equal(t, uint8(57), v[1])

	// 对非0原值进行位写入
	v = []uint8{0xF0}
	bw = nazabits.NewBitWriter(v)
	bw.WriteBit(0)
	bw.WriteBit(0)
	bw.WriteBit(0)
	bw.WriteBit(0)
	bw.WriteBit(1)
	bw.WriteBit(1)
	bw.WriteBit(1)
	bw.WriteBit(1)
	assert.Equal(t, uint8(0x0F), v[0])
}

func TestBitWriter_WriteBits8(t *testing.T) {
	v := make([]byte, 2)
	bw := nazabits.NewBitWriter(v)
	bw.WriteBits8(1, 0)
	bw.WriteBits8(2, 1)
	bw.WriteBits8(3, 4)
	bw.WriteBits8(4, 0)
	bw.WriteBits8(5, 28)
	bw.WriteBits8(1, 1)
	assert.Equal(t, uint8(48), v[0])
	assert.Equal(t, uint8(57), v[1])

	v = make([]byte, 1)
	bw = nazabits.NewBitWriter(v)
	bw.WriteBits8(3, 1+8+32+128)
	assert.Equal(t, uint8(1<<5), v[0])
}

func TestBitWriter_WriteBits16(t *testing.T) {
	v := make([]byte, 2)
	bw := nazabits.NewBitWriter(v)
	bw.WriteBits16(1, 0)
	bw.WriteBits16(2, 1)
	bw.WriteBits16(3, 4)
	bw.WriteBits16(4, 0)
	bw.WriteBits16(5, 28)
	bw.WriteBits16(1, 1)
	assert.Equal(t, uint8(48), v[0])
	assert.Equal(t, uint8(57), v[1])

	v = make([]byte, 1)
	bw = nazabits.NewBitWriter(v)
	bw.WriteBits16(3, 1+8+32+128)
	assert.Equal(t, uint8(1<<5), v[0])

	v = make([]byte, 2)
	bw = nazabits.NewBitWriter(v)
	bw.WriteBits16(16, 0xFFFF)
	assert.Equal(t, uint8(0xFF), v[0])
	assert.Equal(t, uint8(0xFF), v[1])
}

func BenchmarkGetBits16(b *testing.B) {
	v := []byte{48, 57}
	for i := 0; i < b.N; i++ {
		nazabits.GetBits16(v, 0, 16)
	}
}

func BenchmarkBitReader_ReadBits8(b *testing.B) {
	v := []byte{48, 57}
	for i := 0; i < b.N; i++ {
		br := nazabits.NewBitReader(v)
		br.ReadBits8(9)
	}
}

func BenchmarkBitWriter_WriteBits16(b *testing.B) {
	v := make([]byte, 2)
	for i := 0; i < b.N; i++ {
		bw := nazabits.NewBitWriter(v)
		bw.WriteBits16(9, 28)
	}
}