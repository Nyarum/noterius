package rand

import (
	"fmt"
	gorand "math/rand"
	"sync/atomic"
	"time"
)

const (
	poolSize = 4096
	max      = 1 << 63
	mask     = max - 1
)

// random number generator pool
var pool = make([]*Rand, poolSize)
var pos uint64

func init() {
	gorand.Seed(time.Now().UnixNano())
	for i := range pool {
		pool[i] = New().RandomSeed()
	}
}

func Next() *Rand {
	apos := int(atomic.AddUint64(&pos, 1) % poolSize)
	return pool[apos]
}

type Rand struct {
	src [2]uint64
}

func New() *Rand {
	return &Rand{}
}

func (r *Rand) Seed(a, b uint64) *Rand {
	r.src[0] = a
	r.src[1] = b
	return r
}

func (r *Rand) RandomSeed() *Rand {
	r.src[0] = uint64(gorand.Uint32())<<32 + uint64(gorand.Uint32())
	r.src[1] = uint64(gorand.Uint32())<<32 + uint64(gorand.Uint32())
	return r
}

// this is xorshift+ https://en.wikipedia.org/wiki/Xorshift
func (r *Rand) Uint64() uint64 {
	s1 := r.src[0]
	s0 := r.src[1]
	r.src[0] = s0
	s1 ^= s1 << 23
	r.src[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return r.src[1] + s0
}

func Uint64() uint64 {
	return Next().Uint64()
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 {
	return int64(r.Uint64()) & mask
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 {
	return Next().Int63()
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 {
	return uint32(r.Int63() >> 31)
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 {
	return Next().Uint32()
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 {
	return int32(r.Int63() >> 32)
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 {
	return Next().Int31()
}

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	u := uint(r.Int63())
	return int(u << 1 >> 1) // clear sign bit if int == int32
}

// Int returns a non-negative pseudo-random int.
func Int() int {
	return Next().Int()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 {
	return Next().Int63n(n)
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 {
	return Next().Int31n(n)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	return Next().Intn(n)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 {
	// A clearer, simpler implementation would be:
	//	return float64(Int63n(1<<53)) / (1<<53)
	// However, Go 1 shipped with
	//	return float64(Int63()) / (1 << 63)
	// and we want to preserve that value stream.
	//
	// There is one bug in the value stream: Int63() may be so close
	// to 1<<63 that the division rounds up to 1.0, and we've guaranteed
	// that the result is always less than 1.0. To fix that, we treat the
	// range as cyclic and map 1 back to 0. This is justified by observing
	// that while some of the values rounded down to 0, nothing was
	// rounding up to 0, so 0 was underrepresented in the results.
	// Mapping 1 back to zero restores some balance.
	// (The balance is not perfect because the implementation
	// returns denormalized numbers for very small Int63(),
	// and those steal from what would normally be 0 results.)
	// The remapping only happens 1/2⁵³ of the time, so most clients
	// will not observe it anyway.
	f := float64(r.Int63()) / (1 << 63)
	if f == 1 {
		f = 0
	}
	return f
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64 {
	return Next().Float64()
}

func (r *Rand) Float64Range(a, b float64) float64 {
	if !(a < b) {
		panic(fmt.Sprintf("Invalid range: %.2f ~ %.2f", a, b))
	}
	return a + r.Float64()*(b-a)
}

func Float64Range(a, b float64) float64 {
	return Next().Float64Range(a, b)
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 {
	// Same rationale as in Float64: we want to preserve the Go 1 value
	// stream except we want to fix it not to return 1.0
	// There is a double rounding going on here, but the argument for
	// mapping 1 to 0 still applies: 0 was underrepresented before,
	// so mapping 1 to 0 doesn't cause too many 0s.
	// This only happens 1/2²⁴ of the time (plus the 1/2⁵³ of the time in Float64).
	f := float32(Float64())
	if f == 1 {
		f = 0
	}
	return f
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32 {
	return Next().Float32()
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []int {
	return Next().Perm(n)
}
