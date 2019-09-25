package main

import (
	"fmt"
  "sort"
)

type student struct {
	name string
	sem1, sem2 int
	tot int
}


func (s student) univ() {
   x:=0
	 fmt.Println("Enter number of students:")
	 m := make(map[string]int)
	 var inputs int
	 fmt.Scan(&inputs)
	 for i:=0;i<inputs;i++{
	 	fmt.Println("Enter name of the student:")
	 	fmt.Scan(&s.name)
	 	fmt.Println("Enter sem1 and sem2 gpa of the student:")
	 	fmt.Scan(&s.sem1)
		fmt.Scan(&s.sem2)
		s.tot= (s.sem1+s.sem2)/2
		m[s.name] = s.tot
	 }
	fmt.Println(m)
  fmt.Println("do u want rank of the students: if yes, enter 1, else enter any number")
  fmt.Scan(&x)
  if (x==1) {
	n := map[int][]string{}
  var p []int
  for k, v := range m {
  n[v] = append(n[v], k)
  }
  for k := range n {
  p = append(p, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(p)))
  for i, k := range p {
   for _, s := range n[k] {
    fmt.Printf("%d.  %s : %d\n",i+1, s, k)
    }
    }
}
}

func main(){
	s:=student{}
	s.univ()
	}


