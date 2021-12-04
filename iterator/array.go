package iterator

import "fmt"

type Iterator interface {
  Next() int
  HasNext() bool
}

type IterableCollection interface {
  CreateIterator(direction bool) Iterator
}

type Array []int

func (a Array) CreateIterator(direction bool) Iterator {
  if direction {
    return &ForwardIterator{Arr: a}
  } else {
    return &BackwardIterator{Arr: a, CurrIdx: len(a)-1}
  }
}

type ForwardIterator struct {
  CurrIdx int
  Arr Array
}

func (f *ForwardIterator) Next() int {
  val := f.Arr[f.CurrIdx]
  f.CurrIdx++
  return val
}

func (f ForwardIterator) HasNext() bool {
  fmt.Println(f.Arr)
  length := len(f.Arr)
  return f.CurrIdx <= length-1
}

type BackwardIterator struct {
  CurrIdx int
  Arr Array
}

func (f *BackwardIterator) Next() int {
  val := f.Arr[f.CurrIdx]
  f.CurrIdx--
  return val
}

func (f BackwardIterator) HasNext() bool {
  return f.CurrIdx >= 0
}
