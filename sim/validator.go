package main 
import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const totalBlocks =10
	var slashed bool 
	var rewards, penalities int 
	for i:=1; i <= 10;i++ {
      if slashed {
		fmt.Println("block slashed halting at %d",i)
	  }
	switch {

	case rand.Float64() < 0.2: 
			fmt.Println("⚠️ Block %d missed sining \n",i)
			penalities += 10
	case rand.Float64() < 0.1: 
            fmt.Println("❌ Block %d double signed - slashing triggered \n",i)
			slashed = true
			penalities += 100
	default:		
		fmt.Printf("Validator successfuly signed blocked %d \n",i)
		rewards += 20
	}
    time.Sleep(1*time.second)
    }
	fmt.Println("\n 📊 final validator summary")
	fmt.printf("🔒 Slashed: %v\n",slashed)
	fmt.Printf(" Total Rewards %d\n",rewards)
	fmt.Printf("Total penalities %d",penalities)
}
