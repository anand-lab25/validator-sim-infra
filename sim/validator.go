package main 
import (
	"fmt"
	"time"
)

func main() {
	for i:=1; i <= 10;i++ {
		fmt.Printf("Validator signed blocked %d \n",i)
		time.Sleep(1*time.Second)
	}
}
