/*
hello.go — a deliberately tiny Go application.

The point of this project isn't the code (it's a few functions), it's the
CI/CD pipeline wrapped around it: tests, multi-format container builds
(Docker + Apptainer/Singularity), vulnerability scanning (e.g. govulncheck),
registry publishing, remote deployment via SSH, and binary releasing.

Gonum is used here on purpose so the built container images / binaries have a
real third-party CGO-free dependency to fetch, compile, and scan.
*/
package main

import (
	"fmt"
	"os"
	"runtime"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/floats"
)

// DemoStats holds the results of numerical calculations.
type DemoStats struct {
	Size int
	Sum  float64
	Mean float64
	Std  float64
}

// Greeting returns a friendly greeting.
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// ComputeDemoStats performs trivial statistical work using Gonum.
func ComputeDemoStats(size int) DemoStats {
	if size <= 0 {
		size = 10
	}

	// Generate a slice equivalent to np.arange(1, size + 1)
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = float64(i + 1)
	}

	// Gonum functions to exercise statistical operations
	sum := floats.Sum(arr)
	mean := stat.Mean(arr, nil)
	std := stat.StdDev(arr, nil) // Population/Sample standard deviation

	return DemoStats{
		Size: size,
		Sum:  sum,
		Mean: mean,
		Std:  std,
	}
}

func main() {
	args := os.Args[1:]
	name := "World"
	if len(args) > 0 {
		name = args[0]
	}

	fmt.Println(Greeting(name))
	fmt.Printf("Go version   : %s\n", runtime.Version())
	fmt.Println("Gonum version: v0.15.0") // Or fetched dynamically via build tags/deps

	stats := ComputeDemoStats(10)
	fmt.Printf("Demo Gonum stats over 1..%d: sum=%.1f, mean=%.1f, std=%.4f\n",
		stats.Size, stats.Sum, stats.Mean, stats.Std)
}
