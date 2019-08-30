package main
import "fmt"
func main() {
	a, b := "hello", "world"
	a,b=b,a
	fmt.Println(a, b)
}
