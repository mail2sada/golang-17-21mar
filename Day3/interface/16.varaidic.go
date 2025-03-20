package main

import "fmt"

func ADD(a, b interface{}, collection ...interface{}) int {
	sum := 0
	switch a.(type) {
	case int:
		sum += a.(int)
	case float32:
		sum += int(a.(float32))
	case float64:
		sum += int(a.(float64))
	default:
		panic("unhandled type")
	}

	switch b.(type) {
	case int:
		sum += b.(int)
	case float32:
		sum += int(b.(float32))
	case float64:
		sum += int(b.(float64))
	default:
		panic("unhandled type")
	}

	for _, ele := range collection {
		switch ele.(type) {
		case int:
			sum += ele.(int)
		case float32:
			sum += int(ele.(float32))
		case float64:
			sum += int(ele.(float64))
		default:
			panic("unhandled type")
		}
	}
	return sum
}

func main() {
	fmt.Println("Pain points...")

	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println(rec)
		}
	}()

	s := ADD(10, 20)

	fmt.Println(s)

	s = ADD(10, 10.5)

	fmt.Println(s)

	s = ADD(10, "hello")
	fmt.Println(s)
}
