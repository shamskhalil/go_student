package routes

import "fmt"

func MyErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
