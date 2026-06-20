package main
import "fmt"
func main() {
	var w, log, adm int

	fmt.Scan(&w, &log, &adm)

	if w == 0 {
		fmt.Println("you must connect to wifi")
		return
	}
	if log == 0 {
		fmt.Println("you need to login first")
		return
	}
	if adm == 0 {
		fmt.Println("you must to login as admin")
		return
	}
	fmt.Println("done")
}