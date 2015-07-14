// The MIT License (MIT)

// Copyright (c) 2015 Christopher Kinniburgh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cjk_math

import "testing"

// ---------------------------------------------------------------
// abs
func TestAbsInt(t *testing.T) {
	cases := []struct {
		in, want int
	}{
			{0, 0},
			{3, 3},
			{-3, 3},
			{4324, 4324},
			{-5839, 5839},
	}
	for _, c := range cases {
		got := AbsInt(c.in)
		if got != c.want {
			t.Errorf("AbsInt(%i) == %i, want %i", c.in, got, c.want)
		}
	}
}

func TestAbsInt8(t *testing.T) {
	cases := []struct {
		in, want int8
	}{
			{0, 0},
			{3, 3},
			{-3, 3},
			{127, 127},
			{-127, 127},
	}
	for _, c := range cases {
		got := AbsInt8(c.in)
		if got != c.want {
			t.Errorf("AbsInt8(%i) == %i, want %i", c.in, got, c.want)
		}
	}
}

func TestAbsInt16(t *testing.T) {
	cases := []struct {
		in, want int16
	}{
			{0, 0},
			{3, 3},
			{-3, 3},
			{127, 127},
			{-127, 127},
			{32767, 32767},
			{-32767, 32767},
	}
	for _, c := range cases {
		got := AbsInt16(c.in)
		if got != c.want {
			t.Errorf("AbsInt16(%i) == %i, want %i", c.in, got, c.want)
		}
	}
}

func TestAbsInt32(t *testing.T) {
	cases := []struct {
		in, want int32
	}{
			{0, 0},
			{3, 3},
			{-3, 3},
			{127, 127},
			{-127, 127},
			{32767, 32767},
			{-32767, 32767},
			{2147483647, 2147483647},
			{-2147483647, 2147483647},
	}
	for _, c := range cases {
		got := AbsInt32(c.in)
		if got != c.want {
			t.Errorf("AbsInt32(%i) == %i, want %i", c.in, got, c.want)
		}
	}
}

func TestAbsInt64(t *testing.T) {
	cases := []struct {
		in, want int64
	}{
			{0, 0},
			{3, 3},
			{-3, 3},
			{127, 127},
			{-127, 127},
			{32767, 32767},
			{-32767, 32767},
			{2147483647, 2147483647},
			{-2147483647, 2147483647},
			{9223372036854775807, 9223372036854775807},
			{-9223372036854775807, 9223372036854775807},
	}
	for _, c := range cases {
		got := AbsInt64(c.in)
		if got != c.want {
			t.Errorf("AbsInt64(%i) == %i, want %i", c.in, got, c.want)
		}
	}
}

// ---------------------------------------------------------------
// min
func TestMinInt(t *testing.T) {
		cases := []struct {
		ina, inb, want int
	}{
			{0, 0, 0},
			{10, 1, 1},
			{1, 10, 1},
			{-100, 1, -100},
			{1, -100, -100},
			{-100, 1000, -100},
			{1000, -100, -100},
	}
	for _, c := range cases {
		got := MinInt(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MinInt(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMinInt8(t *testing.T) {
		cases := []struct {
		ina, inb, want int8
	}{
			{0, 0, 0},
			{10, 1, 1},
			{1, 10, 1},
			{-100, 1, -100},
			{1, -100, -100},
			{-128, 127, -128},
			{127, -128, -128},
	}
	for _, c := range cases {
		got := MinInt8(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MinInt8(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMinInt16(t *testing.T) {
		cases := []struct {
		ina, inb, want int16
	}{
			{0, 0, 0},
			{10, 1, 1},
			{1, 10, 1},
			{-100, 1, -100},
			{1, -100, -100},
			{-128, 127, -128},
			{127, -128, -128},
			{-32768, 32767, -32768},
			{32767, -32768, -32768},
	}
	for _, c := range cases {
		got := MinInt16(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MinInt16(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMinInt32(t *testing.T) {
		cases := []struct {
		ina, inb, want int32
	}{
			{0, 0, 0},
			{10, 1, 1},
			{1, 10, 1},
			{-100, 1, -100},
			{1, -100, -100},
			{-128, 127, -128},
			{127, -128, -128},
			{-32768, 32767, -32768},
			{32767, -32768, -32768},
			{-2147483648, 2147483647, -2147483648},
			{2147483647, -2147483648, -2147483648},

	}
	for _, c := range cases {
		got := MinInt32(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MinInt32(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMinInt64(t *testing.T) {
		cases := []struct {
		ina, inb, want int64
	}{
			{0, 0, 0},
			{10, 1, 1},
			{1, 10, 1},
			{-100, 1, -100},
			{1, -100, -100},
			{-128, 127, -128},
			{127, -128, -128},
			{-32768, 32767, -32768},
			{32767, -32768, -32768},
			{-2147483648, 2147483647, -2147483648},
			{2147483647, -2147483648, -2147483648},
			{-9223372036854775808, 9223372036854775807, -9223372036854775808},
			{9223372036854775807, -9223372036854775808, -9223372036854775808},
	}
	for _, c := range cases {
		got := MinInt64(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MinInt64(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

// ---------------------------------------------------------------
// max
func TestMaxInt(t *testing.T) {
		cases := []struct {
		ina, inb, want int
	}{
			{0, 0, 0},
			{10, 1, 10},
			{1, 10, 10},
			{-100, 1, 1},
			{1, -100, 1},
			{-100, 1000, 1000},
			{1000, -100, 1000},
	}
	for _, c := range cases {
		got := MaxInt(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MaxInt(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMaxInt8(t *testing.T) {
		cases := []struct {
		ina, inb, want int8
	}{
		{0, 0, 0},
		{10, 1, 10},
		{1, 10, 10},
		{-100, 1, 1},
		{1, -100, 1},
		{-128, 127, 127},
		{127, -128, 127},
	}
	for _, c := range cases {
		got := MaxInt8(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MaxInt8(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMaxInt16(t *testing.T) {
		cases := []struct {
		ina, inb, want int16
	}{
		{0, 0, 0},
		{10, 1, 10},
		{1, 10, 10},
		{-100, 1, 1},
		{1, -100, 1},
		{-128, 127, 127},
		{127, -128, 127},
		{-32768, 32767, 32767},
		{32767, -32768, 32767},
	}
	for _, c := range cases {
		got := MaxInt16(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MaxInt16(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMaxInt32(t *testing.T) {
		cases := []struct {
		ina, inb, want int32
	}{
		{0, 0, 0},
		{10, 1, 10},
		{1, 10, 10},
		{-100, 1, 1},
		{1, -100, 1},
		{-128, 127, 127},
		{127, -128, 127},
		{-32768, 32767, 32767},
		{32767, -32768, 32767},
		{-2147483648, 2147483647, 2147483647},
		{2147483647, -2147483648, 2147483647},
	}
	for _, c := range cases {
		got := MaxInt32(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MaxInt32(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestMaxInt64(t *testing.T) {
		cases := []struct {
		ina, inb, want int64
	}{
			{0, 0, 0},
			{10, 1, 10},
			{1, 10, 10},
			{-100, 1, 1},
			{1, -100, 1},
			{-128, 127, 127},
			{127, -128, 127},
			{-32768, 32767, 32767},
			{32767, -32768, 32767},
			{-2147483648, 2147483647, 2147483647},
			{2147483647, -2147483648, 2147483647},
			{-9223372036854775808, 9223372036854775807, 9223372036854775807},
			{9223372036854775807, -9223372036854775808, 9223372036854775807},
	}
	for _, c := range cases {
		got := MaxInt64(c.ina, c.inb)
		if got != c.want {
			t.Errorf("MaxInt64(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

// ---------------------------------------------------------------
// pow

func TestPowInt(t *testing.T) {
	cases := []struct {
		ina, inb, want int
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 823543},
		// TODO: Calculate close overflow situations for each.
		{32767, 32768, -1}, // Overflow checking.
		{2147483647, 2147483648, -1},

	}
	for _, c := range cases {
		got := PowInt(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowInt(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowInt8(t *testing.T) {
	cases := []struct {
		ina, inb, want int8
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, -1},
		// TODO: Calculate close overflow situations for each.

	}
	for _, c := range cases {
		got := PowInt8(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowInt8(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowInt16(t *testing.T) {
	cases := []struct {
		ina, inb, want int16
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, -1},
		// TODO: Calculate close overflow situations for each.

	}
	for _, c := range cases {
		got := PowInt16(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowInt16(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowInt32(t *testing.T) {
	cases := []struct {
		ina, inb, want int32
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 823543},
		{7, 14, -1},
		// TODO: Calculate close overflow situations for each.

	}
	for _, c := range cases {
		got := PowInt32(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowInt32(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
func TestPowInt64(t *testing.T) {
	cases := []struct {
		ina, inb, want int64
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 29, -1},
		// TODO: Calculate close overflow situations for each.

	}
	for _, c := range cases {
		got := PowInt64(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowInt64(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}



func TestPowUint(t *testing.T) {
	cases := []struct {
		ina, inb, want uint
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 823543},
		{823543, 823543, 0},
		// TODO: Calculate close overflow situations for each.
	}
	for _, c := range cases {
		got := PowUint(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowUint(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowUint8(t *testing.T) {
	cases := []struct {
		ina, inb, want uint8
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 0},
		// TODO: Calculate close overflow situations for each.
	}
	for _, c := range cases {
		got := PowUint8(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowUint8(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowUint16(t *testing.T) {
	cases := []struct {
		ina, inb, want uint16
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 0},
		// TODO: Calculate close overflow situations for each.
	}
	for _, c := range cases {
		got := PowUint16(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowUint16(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowUint32(t *testing.T) {
	cases := []struct {
		ina, inb, want uint32
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 823543},
		{823543, 823543, 0},
		// TODO: Calculate close overflow situations for each.
	}
	for _, c := range cases {
		got := PowUint32(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowUint32(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}

func TestPowUint64(t *testing.T) {
	cases := []struct {
		ina, inb, want uint64
	}{
		{1, 0, 1},
		{1, 1, 1},
		{7, 1, 7},
		{0, 1, 0},
		{7, 7, 823543},
		{823543, 823543, 0},
		// TODO: Calculate close overflow situations for each.
	}
	for _, c := range cases {
		got := PowUint64(c.ina, c.inb)
		if got != c.want {
			t.Errorf("PowUint64(%i, %i) == %i, want %i", c.ina, c.inb, got, c.want)
		}
	}
}
