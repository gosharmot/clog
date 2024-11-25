package clog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Fatalf ...
func Fatalf(format string, v ...interface{}) {
	fmt.Printf(color.RedString("Fatal error: ")+format+"\n", v...)
	os.Exit(1)
}

// Fatal ...
func Fatal(v interface{}) {
	fmt.Printf(color.RedString("Fatal error: ")+"\n", v)
	os.Exit(1)
}

// Errorf ...
func Errorf(format string, v ...interface{}) {
	fmt.Printf(color.RedString("Error: ")+format+"\n", v...)
}

// Error ...
func Error(v interface{}) {
	fmt.Printf(color.RedString("Error: ")+"%v\n", v)
}

// Infof ...
func Infof(format string, v ...interface{}) {
	fmt.Printf(color.GreenString("Info: ")+format+"\n", v...)
}

// Info ...
func Info(v interface{}) {
	fmt.Printf(color.GreenString("Info: ")+"%v\n", v)
}

// Warningf ...
func Warningf(format string, v ...interface{}) {
	fmt.Printf(color.YellowString("Warn: ")+format+"\n", v...)
}

// Warning ...
func Warning(v interface{}) {
	fmt.Printf(color.YellowString("Warn: ")+"%v\n", v)
}
