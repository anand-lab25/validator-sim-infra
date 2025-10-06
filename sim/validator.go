package main 
import (
	"fmt"
	"time"
	"math/rand"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	rewardGuage = prometheuse.NewGauge(prometheuse.GaugeOpts{
		Name:"validator_reward_total",
		Help: "total reward earned by Validators",
	})
	penaltyGauge = prometheuse.NewGauge(prometheuse.GaugeOpts{
		Name:"validator_penalty_total",
		Help:"total penalty levied on a block",
	})
	slashedGauge = prometheuse.NewGauge(prometheuse.GaugeOpts{
		Name:"validator_slashed",
		Help:"Slashing status 1 if slashed otherwise 0",
	})
)


func init() {
    prometheuse.MustRegister(rewardGauge,penaltyGauge,slashedGauge)
}



func main() {
	rand.Seed(time.Now().UnixNano())
	const totalBlocks =10
	var slashed bool 
	var rewards, penalities int 
	for i:=1; i <= 10;i++ {
      if slashed {
		fmt.Println("block slashed halting at %d\n",i)
	  }
	switch {

	case rand.Float64() < 0.2: 
			fmt.Println("⚠️ Block %d missed signing\n",i)
			penalities += 10
	case rand.Float64() < 0.1: 
            fmt.Println("❌ Block %d double signed - slashing triggered\n",i)
			slashed = true
			penalities += 100
	default:		
		fmt.Printf("Validator successfuly signed block %d\n",i)
		rewards += 20
	}
    time.Sleep(1*time.Second)
    }
	fmt.Println("\n 📊 final validator summary")
	fmt.printf("🔒 Slashed: %v\n",slashed)
	fmt.Printf(" Total Rewards %d\n",rewards)
	fmt.Printf("Total penalities %d",penalities)
}
