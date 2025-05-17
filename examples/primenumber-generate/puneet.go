package primenumbergenerate
 
import (
    "fmt"
    "math"
    "sort"
    "sync"
)
 
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
 
func findPrimesInRange(start, end int, wg *sync.WaitGroup, ch chan<- int) {
    defer wg.Done()
    for num := start; num <= end; num++ {
        if isPrime(num) {
            ch <- num
        }
    }
}
 
func DoPuneet() {
    N := 100
 
    ch := make(chan int, N)
 
    var wg sync.WaitGroup
 
    // Number of workers (tune this based on your CPU cores)
    numWorkers := 8
    chunkSize := N / numWorkers
 
    for i := 0; i < numWorkers; i++ {
        start := i*chunkSize + 1
        end := (i + 1) * chunkSize
 
        if i == numWorkers-1 {
            end = N
        }
 
        wg.Add(1)
        go findPrimesInRange(start, end, &wg, ch)
    }
    go func() {
        wg.Wait()
        close(ch)
    }()
 
    var primes []int
    for prime := range ch {
        primes = append(primes, prime)
    }
    sort.Ints(primes)
    fmt.Println("Prime numbers up to", N, "are:", primes)
}