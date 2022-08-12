package ebs

import "fmt"

type ApiReturn struct {
	number int16
}

func PrintTest(s string) string {
	return fmt.Sprintf("Testing %s", s)
}
