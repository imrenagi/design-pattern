package iterator

import "fmt"

func Execute() {
  arr  := Array{1, 2, 3, 5}
  it := arr.CreateIterator(true)

  val := it.Next()
  fmt.Println(val)

  arr = append(arr, 6)

  for it.HasNext() {
    fmt.Println(it.Next())
  }
  fmt.Println("------")


  //
  // backwardIt := arr.CreateIterator(false)
  // for backwardIt.HasNext() {
  //   fmt.Println(backwardIt.Next())
  // }
}
