# interrupted

A simple wrapper to handle interrupts in Go with a Context and Waitgroup


## Usage

```go
package main

func main() {
	WaitForInterrupt(func(ctx context.Context, wg *sync.WaitGroup) {
		go asyncFunction1(ctx, wg)
        asyncFunction2(ctx, wg) // It is important to call this without a goroutine to allow the waitgroup to get the correct count, else the program will exit before the asyncFunction1 has started
	})
}
```
