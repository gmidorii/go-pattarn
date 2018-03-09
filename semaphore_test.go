package main

// ExamplesyncWait is syncWait sysout test
func ExampleSemaphore() {
	Semaphore(2, "first", "second", "third")

	// Output:
	// first
	// second
	// third
}
