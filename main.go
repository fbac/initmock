package initfunc

import "fmt"

func init() {
	fmt.Println("[initfunc][init] initializing!")
}

func main() {
	fmt.Println("[initfunc][main] initialized!")
}
