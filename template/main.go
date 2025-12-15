package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Daily optimization template
func main() {
	fmt.Printf("Day X: [Topic]\n")
	fmt.Printf("Date: %s\n\n", time.Now().Format("2006-01-02"))

	fmt.Println("🚀 Starting optimization challenge...")

	// Run benchmarks
	benchmarkBefore()

	// Show optimization
	fmt.Println("\n" + getDivider())
	explainOptimization()

	// Run after optimization
	benchmarkAfter()

	// Calculate cost impact
	fmt.Println("\n" + getDivider())
	calculateCostImpact()

	fmt.Println("\n✅ Challenge completed!")
}

func getDivider() string {
	return strings.Repeat("=", 50)
}

func benchmarkBefore() {
	fmt.Println("📊 BEFORE OPTIMIZATION")
	fmt.Println(getDivider())

	start := time.Now()

	// TODO: Implement original (unoptimized) code here
	// Example:
	// result := unoptimizedFunction()

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %v\n", elapsed)

	// Measure memory
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory allocated: %.2f MB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Total allocations: %d\n", m.Mallocs)

	fmt.Printf("Database queries: [if applicable]\n")
}

func explainOptimization() {
	fmt.Println("🔧 OPTIMIZATION STRATEGY")
	fmt.Println(getDivider())

	fmt.Println("1. Problem identified:")
	fmt.Println("   - [Describe the inefficiency]")
	fmt.Println("2. Solution:")
	fmt.Println("   - [Describe the optimization]")
	fmt.Println("3. Expected improvements:")
	fmt.Println("   - [List expected gains]")
}

func benchmarkAfter() {
	fmt.Println("\n📈 AFTER OPTIMIZATION")
	fmt.Println(getDivider())

	start := time.Now()

	// TODO: Implement optimized code here
	// Example:
	// result := optimizedFunction()

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %v\n", elapsed)

	// Measure memory
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory allocated: %.2f MB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Total allocations: %d\n", m.Mallocs)

	fmt.Printf("Database queries: [if applicable]\n")

	// Calculate improvement percentage
	// Note: You'll need to store the before time to compare
}

func calculateCostImpact() {
	fmt.Println("💰 COST IMPACT ANALYSIS")
	fmt.Println(getDivider())

	fmt.Println("Assumptions:")
	fmt.Println("- 100,000 requests per day")
	fmt.Println("- AWS t3.medium: $0.0416/hour (~$30/month)")
	fmt.Println("- Data transfer: $0.09/GB")

	fmt.Println("\nCalculations:")
	fmt.Println("1. CPU cost savings:")
	fmt.Println("   - Before: $X/month")
	fmt.Println("   - After:  $Y/month")
	fmt.Println("   - Savings: $Z/month")

	fmt.Println("\n2. Memory cost savings:")
	fmt.Println("   - Before: $A/month")
	fmt.Println("   - After:  $B/month")
	fmt.Println("   - Savings: $C/month")

	fmt.Println("\n3. Database cost savings:")
	fmt.Println("   - Query reduction: P%")
	fmt.Println("   - Savings: $D/month")

	fmt.Println("\n📈 Total estimated savings:")
	fmt.Println("   Monthly:  $T/month")
	fmt.Println("   Annual:   $(T * 12)/year")
}

// Helper functions for common measurements
func printMemoryStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s:\n", label)
	fmt.Printf("  Alloc:      %.2f MB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("  TotalAlloc: %.2f MB\n", float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("  Mallocs:    %d\n", m.Mallocs)
	fmt.Printf("  Frees:      %d\n", m.Frees)
}

func runAndMeasure(name string, fn func()) time.Duration {
	fmt.Printf("\n⏱️  Running: %s\n", name)
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("Time: %v\n", elapsed)
	return elapsed
}
