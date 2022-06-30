package initfunc

import "fmt"

func init() {
	fmt.Println("[initfunc][init] initializing initmock")
}

func Placeholder() {
	fmt.Println("[initfunc][placeholder] I'm just a simple placeholder func")
}
