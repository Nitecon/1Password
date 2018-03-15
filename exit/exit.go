package exit

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Printf("Fatal error occurred, exiting:\n==> %s", err)

	os.Exit(1)
}

func Fatalm(m string, err error) {
	fmt.Printf("%s\n==> %s", err)

	os.Exit(1)
}

func Ok(m string) {
	fmt.Println(m)
	os.Exit(0)
}
