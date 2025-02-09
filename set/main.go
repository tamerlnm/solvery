package main

import (
	"fmt"
	"set/pkg/model"
)

func main() {
	set1 := model.NewSet()
	set2 := model.NewSet()

	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2.Add("b")
	set2.Add("c")
	set2.Add("e")

	unionSet := set1.Union(set2)
	fmt.Println("Объединение:", unionSet.ToSlice())

	subtractionSet := set1.Subtraction(set2)
	fmt.Println("Разность:", subtractionSet.ToSlice())

	intersectionSet := set1.Intersection(set2)
	fmt.Println("Пересечение:", intersectionSet.ToSlice())
}
