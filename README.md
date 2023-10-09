# Go Pointer And Value Data Types Benchmarks
The fundamental purpose of these benchmarks is to examine how Go manages memory and to evaluate the effects of using pointer and value data types.
## Run
```sh
$ make all
```
## Benchmark Results
```sh
$ go build -gcflags '-m -l' ./main.go

# command-line-arguments
./main.go:10:15: []int{...} escapes to heap
./main.go:16:3: moved to heap: ret
./main.go:16:15: []int{...} escapes to heap
./main.go:20:16: argSlice does not escape
./main.go:26:23: argSlice does not escape
./main.go:32:7: ds does not escape
./main.go:37:7: ds does not escape
./main.go:48:4: moved to heap: res
./main.go:52:30: y does not escape
./main.go:57:37: y does not escape
./main.go:58:4: moved to heap: res
./main.go:70:31: leaking param: ds to result ~r0 level=1
./main.go:78:27: leaking param: ds to result ~r0 level=0
./main.go:87:11: &DemoStruct{...} escapes to heap
./main.go:94:34: leaking param content: ds
./main.go:95:11: &DemoStruct{...} escapes to heap
./main.go:102:30: leaking param: ds
./main.go:103:11: &DemoStruct{...} escapes to heap

$ go test -bench . -benchmem -benchtime=30000000x

goos: linux
goarch: amd64
pkg: demo
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkReturnVar-8                            30000000                 2.313 ns/op         0 B/op          0 allocs/op
BenchmarkReturnVarPointer-8                     30000000                23.44 ns/op          8 B/op          1 allocs/op
BenchmarkReturnVarFromArgPointer-8              30000000                 2.322 ns/op         0 B/op          0 allocs/op
BenchmarkReturnVarPointerFromArgPointer-8       30000000                23.10 ns/op          8 B/op          1 allocs/op
BenchmarkReturnStructCopy-8                     30000000                 7.253 ns/op         0 B/op          0 allocs/op
BenchmarkReturnStructCopyFromArgPointer-8       30000000                13.78 ns/op          0 B/op          0 allocs/op
BenchmarkReturnStructCopyFromArgVal-8           30000000                15.87 ns/op          0 B/op          0 allocs/op
BenchmarkReturnStructPointer-8                  30000000                87.26 ns/op         80 B/op          1 allocs/op
BenchmarkReturnStructPointerFromArgPointer-8    30000000                92.62 ns/op         80 B/op          1 allocs/op
BenchmarkReturnStructPointerFromArgVal-8        30000000                96.08 ns/op         80 B/op          1 allocs/op
BenchmarkStructValueMethod-8                    30000000                 8.915 ns/op         0 B/op          0 allocs/op
BenchmarkStructPointerMethod-8                  30000000                 4.763 ns/op         0 B/op          0 allocs/op
BenchmarkReturnSlice-8                          30000000                53.72 ns/op         48 B/op          1 allocs/op
BenchmarkReturnSlicePointer-8                   30000000               117.8 ns/op          72 B/op          2 allocs/op
BenchmarkArgSlicePointerPassSlice-8             30000000                14.79 ns/op          0 B/op          0 allocs/op
BenchmarkArgSlicePassSlice1-8                   30000000                13.16 ns/op          0 B/op          0 allocs/op
BenchmarkArgSlicePassSlice2-8                   30000000                13.10 ns/op          0 B/op          0 allocs/op
PASS
ok      demo    17.729s
```
## Benchmark Tests
### main
```go
package main

type DemoStruct struct {
   A, B, C int
   D, E, F string
   G, H, I bool
}
//go:noinline
func  ReturnSlice() []int {
  ret := []int{1, 2, 3, 4, 5, 6}
  return ret
}

//go:noinline
func  ReturnSlicePointer() *[]int {
  ret := []int{1, 2, 3, 4, 5, 6}
  return &ret
}
//go:noinline
func  ArgSlice(argSlice []int)  {
   for i := 0; i < len(argSlice); i++{
      argSlice[i] = 54
   }
}
//go:noinline
func  ArgSlicePointer(argSlice *[]int)  {
   for i := 0; i < len(*argSlice); i++{
      (*argSlice)[i] = 54
   }
}
//go:noinline
func (ds *DemoStruct) StructPointerMethod(){
   ds.A = 3

}
//go:noinline
func (ds DemoStruct) StructValueMethod(){
   ds.A = 3
}
//go:noinline
func ReturnVar() int {
   y := 54
   return y * 54
}
//go:noinline
func ReturnVarPointer() *int {
   y := 54
   res := y * 54
   return &res
}
//go:noinline
func ReturnVarFromArgPointer(y *int) int {
   res := *y * 54
   return res
}
//go:noinline
func ReturnVarPointerFromArgPointer(y *int) *int {
   res := *y * 54
   return &res
}
//go:noinline
func CreateCopy() DemoStruct {
   return DemoStruct{
      A: 123, B: 456, C: 789,
      D: "ABC", E: "DEF", F: "HIJ",
      G: true, H: true, I: true,
   }
}
//go:noinline
func CreateCopyFromArgPointer(ds *DemoStruct) DemoStruct {
   return DemoStruct{
      A: ds.A, B: ds.B, C: ds.C,
      D: ds.D, E: ds.E, F: ds.F,
      G: ds.G, H: ds.H, I: ds.I,
   }
}
//go:noinline
func CreateCopyFromArgVal(ds DemoStruct) DemoStruct {
   return DemoStruct{
      A: ds.A, B: ds.B, C: ds.C,
      D: ds.D, E: ds.E, F: ds.F,
      G: ds.G, H: ds.H, I: ds.I,
   }
}
//go:noinline
func CreatePointer() *DemoStruct {
   return &DemoStruct{
      A: 123, B: 456, C: 789,
      D: "ABC", E: "DEF", F: "HIJ",
      G: true, H: true, I: true,
   }
}
//go:noinline
func CreatePointerFromArgPointer(ds *DemoStruct) *DemoStruct {
   return &DemoStruct{
      A: ds.A, B: ds.B, C: ds.C,
      D: ds.D, E: ds.E, F: ds.F,
      G: ds.G, H: ds.H, I: ds.I,
   }
}
//go:noinline
func CreatePointerFromArgVal(ds DemoStruct) *DemoStruct {
   return &DemoStruct{
      A: ds.A, B: ds.B, C: ds.C,
      D: ds.D, E: ds.E, F: ds.F,
      G: ds.G, H: ds.H, I: ds.I,
   }
}
func main(){}
```
### main_test
```go
package main

import "testing"

func BenchmarkReturnVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReturnVar()
	}
}
func BenchmarkReturnVarPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReturnVarPointer()
	}
}
func BenchmarkReturnVarFromArgPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y := 54
		_ = ReturnVarFromArgPointer(&y)
	}
}
func BenchmarkReturnVarPointerFromArgPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y := 54
		_ = ReturnVarPointerFromArgPointer(&y)
	}
}
func BenchmarkReturnStructCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateCopy()
	}
}
func BenchmarkReturnStructCopyFromArgPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		_ = CreateCopyFromArgPointer(&ds)
	}
}
func BenchmarkReturnStructCopyFromArgVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		_ = CreateCopyFromArgVal(ds)
	}
}
func BenchmarkReturnStructPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreatePointer()
	}
}
func BenchmarkReturnStructPointerFromArgPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		_ = CreatePointerFromArgPointer(&ds)
	}
}
func BenchmarkReturnStructPointerFromArgVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		_ = CreatePointerFromArgVal(ds)
	}
}
func BenchmarkStructValueMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		ds.StructValueMethod()
	}
}
func BenchmarkStructPointerMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ds := DemoStruct{
			A: 123, B: 456, C: 789,
			D: "ABC", E: "DEF", F: "HIJ",
			G: true, H: true, I: true,
		}
		ds.StructPointerMethod()
	}
}
func BenchmarkReturnSlice(b *testing.B){
   for i := 0; i < b.N; i++ {
      _ = ReturnSlice()
   }
}
func BenchmarkReturnSlicePointer(b *testing.B){
   for i := 0; i < b.N; i++ {
      _ = ReturnSlicePointer()
   }
}
func BenchmarkArgSlicePointerPassSlice(b *testing.B){
   for i := 0; i < b.N; i++ {
      s := []int{1,2,3,4,5,6,7,8,9,10}
      ArgSlicePointer(&s)
   }
}
func BenchmarkArgSlicePassSlice1(b *testing.B){
   for i := 0; i < b.N; i++ {
      s := []int{1,2,3,4,5,6,7,8,9,10}
      ArgSlice(s)
   }
}
func BenchmarkArgSlicePassSlice2(b *testing.B){
   for i := 0; i < b.N; i++ {
      s := []int{1,2,3,4,5,6,7,8,9,10}
      ArgSlice(s[:])
   }
}
```
## References
* https://go.dev/doc/faq#methods_on_values_or_pointers
