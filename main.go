package main

import (
	"fmt"
	"os"
	"path"

	"github.com/atotto/clipboard"
	"github.com/skip2/go-qrcode"
	"github.com/skratchdot/open-golang/open"
)

func main() {

	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	temp := os.TempDir()
	path := path.Join(temp, "qrcode.png")

	err = qrcode.WriteFile(text, qrcode.Medium, 256, path)
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	}

	open.Run(path)
}
