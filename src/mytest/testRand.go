package mytest

import (
	"fmt"
	"math/rand"
)

func TestRand() {
	fmt.Println("rand.Int()=", rand.Int())
	fmt.Println("rand.Intn(10)=", rand.Intn(10))
	fmt.Println("rand.Int31()=", rand.Int31())
	fmt.Println("rand.Int31n(10)=", rand.Int31n(10))
	fmt.Println("rand.Int63()=", rand.Int63())
	fmt.Println("rand.Int63n(10)=", rand.Int63n(10))
	fmt.Println("rand.Float32()=", rand.Float32())
	fmt.Println("rand.Float64()=", rand.Float64())
	fmt.Println("rand.Perm(10)=", rand.Perm(10))
	fmt.Println("rand.ExpFloat64()=", rand.ExpFloat64())
	fmt.Println("rand.Uint32()=", rand.Uint32())

	r := rand.New(rand.NewSource(12346789))
	fmt.Println("r.Int()=", r.Int())
	fmt.Println("r.Intn(10)=", r.Intn(10))
	fmt.Println("r.Int31()=", r.Int31())
	fmt.Println("r.Int31n(10)=", r.Int31n(10))
	fmt.Println("r.Int63()=", r.Int63())
	fmt.Println("r.Int63n(10)=", r.Int63n(10))
	fmt.Println("r.Float32()=", r.Float32())
	fmt.Println("r.Float64()=", r.Float64())
	fmt.Println("r.Perm(10)=", r.Perm(10))
	fmt.Println("r.ExpFloat64()=", r.ExpFloat64())
	fmt.Println("r.Uint32()=", r.Uint32())

}
