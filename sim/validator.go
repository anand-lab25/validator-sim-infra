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
	rewardGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:"validator_reward_total",
		Help: "total reward earned by Validators",
	})
	penaltyGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:"validator_penalty_total",
		Help:"total penalty levied on a block",
	})
	slashedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:"validator_slashed",
		Help:"Slashing status 1 if slashed otherwise 0",
	})
)


func init() {
    prometheus.MustRegister(rewardGauge,penaltyGauge,slashedGauge)
}



func main() {
	rand.Seed(time.Now().UnixNano())
	const totalBlocks =10
	var slashed bool 
	var rewards, penalities int 
	go func() {
		http.Handle("/metrics",promhttp.Handler())
		// log.Println(" Exposing metrics at http://www.localhost:2112/metrics")
		// log.Fatal(http.ListenAndServe("0.0.0.0:2112",nil))
		http.ListenAndServe(":2112", nil)
	}()
	for i:=1; i <= 10;i++ {
      if slashed {
		fmt.Printf("block slashed halting at %d\n",i)
	  }
	switch {

	case rand.Float64() < 0.2: 
			fmt.Printf("⚠️ Block %d missed signing\n",i)
			penalities += 10
	case rand.Float64() < 0.1: 
            fmt.Printf("❌ Block %d double signed - slashing triggered\n",i)
			slashed = true
			penalities += 100
	default:		
		fmt.Printf("Validator successfuly signed block %d\n",i)
		rewards += 20
	}
    time.Sleep(1*time.Second)
    }
	rewardGauge.Set(float64(rewards))
	penaltyGauge.Set(float64(penalities))
	if slashed {
		slashedGauge.Set(1)
	} else {
		slashedGauge.Set(0)
	}
	fmt.Println("\n 📊 final validator summary")
	fmt.Printf("🔒 Slashed: %v\n",slashed)
	fmt.Printf(" Total Rewards %d\n",rewards)
	fmt.Printf("Total penalities %d",penalities)
	select {} // keeps the program running
}


