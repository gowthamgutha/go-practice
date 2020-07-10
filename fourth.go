// structs example
package main

import (
	"fmt"
)

type Student struct {
	sno   int
	sname string
}

func (s Student) getSno() int {
	return s.sno
}

// this is value type receiver method
func (s Student) getSname() string {
	return s.sname
}

// this is reference type receiver method (modification allowed)
func (s *Student) setSno(sno int) {
	s.sno = sno
}

func main() {
	student := Student{sno: 101, sname: "Rama"}
	fmt.Println(student.sno, student.sname)
	fmt.Println(student.getSno(), student.getSname())
	student.setSno(102)
	fmt.Println(student.getSno())
	fmt.Println(Student{sno: 103, sname: "Lakshmana"}.getSno())
	fmt.Println(Student{sno: 104, sname: "Sita"}.sname)
}
