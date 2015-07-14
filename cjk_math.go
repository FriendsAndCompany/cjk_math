package cjk_math

// ---------------------------------------------------------------
// abs
func AbsInt(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
func AbsInt8(a int8) int8 {
	if a < 0 {
		a = -a
	}
	return a
}
func AbsInt16(a int16) int16 {
	if a < 0 {
		a = -a
	}
	return a
}
func AbsInt32(a int32) int32 {
	if a < 0 {
		a = -a
	}
	return a
}
func AbsInt64(a int64) int64 {
	if a < 0 {
		a = -a
	}
	return a
}

// ---------------------------------------------------------------
// min
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}
func MinInt16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
func MinUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// ---------------------------------------------------------------
// max
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func MaxInt8(a, b int8) int8 {
	if a < b {
		return b
	}
	return a
}
func MaxInt16(a, b int16) int16 {
	if a < b {
		return b
	}
	return a
}
func MaxInt32(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}
func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
func MaxUint(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}
func MaxUint8(a, b uint8) uint8 {
	if a < b {
		return b
	}
	return a
}
func MaxUint16(a, b uint16) uint16 {
	if a < b {
		return b
	}
	return a
}
func MaxUint32(a, b uint32) uint32 {
	if a < b {
		return b
	}
	return a
}
func MaxUint64(a, b uint64) uint64 {
	if a < b {
		return b
	}
	return a
}

// ---------------------------------------------------------------
// pow
func PowInt(b, p int) int {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 0
	}
	var out int = 1

	for p > 0 {
		out *= b
		if out < 0 { // Catch overflows
			return -1
		}
		p--
	}
	return out
}
func PowInt8(b, p int8) int8 {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 0
	}
	var out int8 = 1

	for p > 0 {
		out *= b
		if out < 0 { // Catch overflows
			return -1
		}
		p--
	}
	return out
}
func PowInt16(b, p int16) int16 {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 0
	}
	var out int16 = 1

	for p > 0 {
		out *= b
		if out < 0 { // Catch overflows
			return -1
		}
		p--
	}
	return out
}
func PowInt32(b, p int32) int32 {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 0
	}
	var out int32 = 1

	for p > 0 {
		out *= b
		if out < 0 { // Catch overflows
			return -1
		}
		p--
	}
	return out
}
func PowInt64(b, p int64) int64 {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 0
	}
	var out int64 = 1

	for p > 0 {
		out *= b
		if out < 0 { // Catch overflows
			return -1
		}
		p--
	}
	return out
}

func PowUint(b, p uint) uint {
	if p == 0 {
		return 1
	}
	var out uint = 1

	for p > 0 {
		if MulOverflowUint(out, b) { // Catch overflows
			return 0
		}
		out *= b
		p--
	}
	return out
}
func PowUint8(b, p uint8) uint8 {
	if p == 0 {
		return 1
	}
	var out uint8 = 1

	for p > 0 {
		if MulOverflowUint8(out, b) { // Catch overflows
			return 0
		}
		out *= b
		p--
	}
	return out
}
func PowUint16(b, p uint16) uint16 {
	if p == 0 {
		return 1
	}
	var out uint16 = 1

	for p > 0 {
		if MulOverflowUint16(out, b) { // Catch overflows
			return 0
		}
		out *= b
		p--
	}
	return out
}
func PowUint32(b, p uint32) uint32 {
	if p == 0 {
		return 1
	}
	var out uint32 = 1

	for p > 0 {
		if MulOverflowUint32(out, b) { // Catch overflows
			return 0
		}
		out *= b
		p--
	}
	return out
}
func PowUint64(b, p uint64) uint64 {
	if p == 0 {
		return 1
	}
	var out uint64 = 1

	for p > 0 {
		if MulOverflowUint64(out, b) { // Catch overflows
			return 0
		}
		out *= b
		p--
	}
	return out
}

// ---------------------------------------------------------------
// overflows

func MulOverflowUint(a, b uint) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func MulOverflowUint8(a, b uint8) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func MulOverflowUint16(a, b uint16) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func MulOverflowUint32(a, b uint32) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func MulOverflowUint64(a, b uint64) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}



// func rountFloat64(val float64, roundOn float64, places int ) (newVal float64) {
// 	var round float64
// 	pow := math.Pow(10, float64(places))
// 	digit := pow * val
// 	_, div := math.Modf(digit)
// 	if div >= roundOn {
// 		round = math.Ceil(digit)
// 	} else {
// 		round = math.Floor(digit)
// 	}
// 	newVal = round / pow
// 	return
// }
