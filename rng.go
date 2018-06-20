package rng

import (
	"math/rand"
	"time"
)

func New() *RNG {
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	return &RNG{
		r: *r,
	}
}

type RNG struct {
	r rand.Rand
}

func (r *RNG) Int63() int64 {
	return r.r.Int63()
}

func (r *RNG) Uint32() uint32 {
	return r.r.Uint32()
}

func (r *RNG) Uint64() uint64 {
	return r.r.Uint64()
}

func (r *RNG) Int31() int32 {
	return r.r.Int31()
}

func (r *RNG) Int() int {
	return r.r.Int()
}

func (r *RNG) Int63n(n int64) int64 {
	return r.r.Int63n(n)
}

func (r *RNG) Int63nn(min, max int64) int64 {
	return r.r.Int63n(max-min) + min
}

func (r *RNG) Int31n(n int32) int32 {
	return r.r.Int31n(n)
}

func (r *RNG) Int31nn(min, max int32) int32 {
	return r.r.Int31n(max-min) + min
}

func (r *RNG) Intn(n int) int {
	return r.r.Intn(n)
}

func (r *RNG) Intnn(min, max int) int {
	return r.r.Intn(max-min) + min
}

func (r *RNG) Float64() float64 {
	return r.r.Float64()
}

func (r *RNG) Float32() float32 {
	return r.r.Float32()
}

func (r *RNG) Perm(n int) []int {
	return r.r.Perm(n)
}

func (r *RNG) Shuffle(arr []interface{}) {
	for len(arr) > 0 {
		n := len(arr)
		randIndex := r.Intn(n)
		arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]
		arr = arr[:n-1]
	}
}
