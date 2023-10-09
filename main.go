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
