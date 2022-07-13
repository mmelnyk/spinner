//go:build exclude

package main

import (
	"context"
	"fmt"
	"time"

	"go.melnyk.org/mansi"
	"go.melnyk.org/spinner"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Print(mansi.CursorHide)
	spin := spinner.NewSpinner(spinner.WithDelay(100*time.Millisecond), spinner.WithStyle(spinner.Style3Dots), spinner.WithElapsedTimer())
	spin.Message("Long process message...")
	spin.Process(ctx)
	// long process
	time.Sleep(2 * time.Second)
	spin.Message("Still working...")
	time.Sleep(2 * time.Second)
	spin.Message("Almost done...")
	time.Sleep(2 * time.Second)
	spin.Message("Ready")
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Print(mansi.SoftReset)
}
