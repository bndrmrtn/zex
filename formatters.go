package zex

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/goterm"
	"github.com/fatih/color"
)

// displayServeInfo displays the server information
func displayServeInfo(listenAddr string, dev bool) {
	var mode = "production"
	if dev {
		mode = "development"
	}

	c := color.New(color.FgMagenta)
	c.Println(" _________  __")
	c.Println("|_  / _ \\ \\/ /")
	c.Println(" / /  __/>  <  ", color.New(color.FgHiYellow).Sprint(mode))
	c.Println("/___\\___/_/\\_\\ ", color.New(color.FgHiGreen).Sprintf("v%s", Version))

	c = color.New(color.FgBlue, color.Bold)
	c.Printf("↳ Server listening on %s\n\n", listenAddr)

	if dev {
		c = color.New(color.FgRed, color.Bold)
		c.Println("Running in development mode. Do not use in production! 🚨")
	}
}

// methodSpaces returns the method and the spaces
func methodSpaces(method string) (string, int) {
	var (
		max = 6
		l   = max - len(method)
	)
	if len(method) > max {
		l = 0
	}
	return method, l
}

// colorMethodName returns the colorized method name
func colorMethodName(method string) string {
	switch strings.TrimSpace(method) {
	case "GET":
		return color.New(color.FgHiGreen).Sprint(method)
	case "POST":
		return color.New(color.FgHiBlue).Sprint(method)
	case "PUT":
		return color.New(color.FgHiCyan).Sprint(method)
	case "PATCH":
		return color.New(color.FgHiYellow).Sprint(method)
	case "DELETE":
		return color.New(color.FgHiRed).Sprint(method)
	default:
		return color.New(color.FgHiMagenta).Sprint(method)
	}
}

// logRoutes logs the routes
func logRoutes(routes []Route) {
	for _, route := range routes {
		for _, p := range route.NormalizedPaths() {
			method, l := methodSpaces(route.Method())
			colorMethod := colorMethodName(method)
			mDots := strings.Repeat(".", l)

			colorMethod = mDots + colorMethod

			name := route.GetName()
			if name == "" {
				name = "unnamed"
			}

			width := goterm.Width()
			width = width - len(mDots+method) - len(p) - len(name) - 5 /* 5 spaces */

			if width < 5 {
				width = 5
			}

			dots := strings.Repeat(".", width)

			fmt.Printf(" %s %s %s %s \n", colorMethod, p, dots, color.New(color.FgHiBlack).Sprint(name))
		}
	}
}

// serverLogger logs the server request information
func serverLogger(start time.Time, method string, path string) {
	method, l := methodSpaces(method)
	colorMethod := colorMethodName(method)
	mDots := strings.Repeat(".", l)

	colorMethod = mDots + colorMethod

	timeString := time.Since(start).String()
	colorTime := color.New(color.FgHiBlack).Sprint(timeString)

	width := goterm.Width()
	width = width - len(mDots+method) - len(path) - len(timeString) - 5 /* 5 spaces */

	if width < 5 {
		width = 5
	}

	dots := strings.Repeat(".", width)

	fmt.Printf(" %s %s %s %s \n", colorMethod, path, dots, colorTime)
}
