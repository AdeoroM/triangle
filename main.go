package main

import (
	"flag"
	"fmt"
)

type Triangle struct {
	CatOP int
	CatAD int
	Hyp   int
	Name  string
}

func triangleInequality(tringle Triangle) bool {
	a := tringle.CatAD
	b := tringle.CatOP
	h := tringle.Hyp
	return a < (b+h) && b < (a+h) && h < (a+b)
}

func typeTriangle(tringle Triangle) string {
	a := tringle.CatAD
	b := tringle.CatOP
	h := tringle.Hyp
	name := tringle.Name

	if a == b && b == h {
		return fmt.Sprintf("Your %v triangle is Equlat", name)
	}
	if a == b || a == h || b == h {
		return fmt.Sprintf("Your %v triangle is Isosceles", name)
	}
	return fmt.Sprintf("Your %v triangle  is Scalene", name)
}

func main() {

	catOP := flag.Int("catOP", 0, "Save")
	catAD := flag.Int("catAD", 0, "Save")
	hyp := flag.Int("hyp", 0, "Save")
	name := flag.String("name", "", "Save")
	flag.Parse()
	triangle := Triangle{
		*catAD, *catOP, *hyp, *name,
	}
	if triangleInequality(triangle) {
		if *catAD != 0 && *catOP != 0 && *hyp != 0 {
			if *name != "" {
				fmt.Println(typeTriangle(triangle))
			} else {
				fmt.Printf("Hey, you should set the triangle name\n")
			}
		} else {
			fmt.Printf("Hey, you should set the triangle parameters: catOP, catAD and hyp\n")
		}
	} else {
		fmt.Printf("Your triangle is invalid\n")
	}
}
