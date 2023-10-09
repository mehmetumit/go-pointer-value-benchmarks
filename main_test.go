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
