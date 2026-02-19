
package main
import (
    "flag"
    "fmt"
    "time"
)
func main() {
    target := flag.String("target", "", "Target IP")
    port := flag.String("port", "80", "Port")
    duration := flag.Int("duration", 60, "Duration")
    flag.Parse()
    fmt.Printf("Attack Start: %s:%s for %ds\n", *target, *port, *duration)
    time.Sleep(time.Duration(*duration) * time.Second)
}
