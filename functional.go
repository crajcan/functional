package functional

import (
	//"time"
)

func dumbFib(x int) int {
	if x == 0 {
		return 0
	}
	if x < 2 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func fibTh(x int, res chan<- int) {
	if x == 0 {
		res <- 0
    return
	}
	if x < 2 {
		res <- 1
    return
	}

	n1, n2 := make(chan int), make(chan int)

	go fibTh(x-1, n1)
	go fibTh(x-2, n2)

  i := <-n1
  j := <-n2
	res <- (i + j)
  return
}

func fib(x int) int {
	ret := make(chan int)
	go fibTh(x, ret)
  res := <-ret
	return res
}

// func sum(x int) int {


/*
type Collection interface{
  split() interface{}
  app()   interface{}
  join()  interface{}
}

func cons(c Collection) element {
  return split(c,1)
}

func linearMap(a []int, b[]int, f int => ){
  if (a == []) { return []}
  return append(linearMap(), f(a[0]))
}

func leftMap(){
}

*/

func pMap() {

}

/*
func foldL(){
}

func foldR(){
}

func reduce(){
}

func mapReduce(){
}

func filter(){

}
*/
