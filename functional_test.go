package functional

import (
  "testing"
)

func BenchmarkDumbFib19(b *testing.B) {
    for n := 0; n < b.N; n++ {
        dumbFib(19)
    }
}

func BenchmarkFib19(b *testing.B) {
    for n := 0; n < b.N; n++ {
        fib(19)
    }
}

func TestFib(t *testing.T) {
  if fib(0) != 0 {
    t.Errorf("got wrong value")
  }
  if fib(1) != 1 {
    t.Errorf("got wrong value")
  }
   if fib(2) != 1 {
    t.Errorf("got wrong value")
  }
   if fib(3) != 2 {
    t.Errorf("got wrong value")
  }
  if fib(19) != 4181 {
    t.Errorf("got wrong value")
  }
  if dumbFib(19) != 4181 {
    t.Errorf("got wrong value")
  }
}
