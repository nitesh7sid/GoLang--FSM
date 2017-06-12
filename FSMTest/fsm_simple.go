package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

type SupplyChain struct {
	To  string
	FSM *fsm.FSM
}

func NewSupplyChain(to string) *SupplyChain {
	d := &SupplyChain{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "created", Src: []string{"idle"}, Dst: "manufactured"},
			{Name: "shipToDistributor", Src: []string{"manufactured"}, Dst: "distributor"},
			{Name: "shipToPharmacy", Src: []string{"distributor"}, Dst: "pharmacy"},
			{Name: "release", Src: []string{"pharmacy"}, Dst: "customer"},
		},
		fsm.Callbacks{
			"created": func(e *fsm.Event) {
				fmt.Printf("After Created in %s state\n",e.FSM.Current())
				//d.enterState(e) },
		},

		"shipToDistributor": func(e *fsm.Event){
			fmt.Printf("After Shipping to Distributor in %s state\n",e.FSM.Current())
		},

		"shipToPharmacy": func(e *fsm.Event){
			fmt.Printf("After Shipping to Pharmacy in %s state\n",e.FSM.Current())
		},

		"release": func(e *fsm.Event){
			fmt.Printf("Released. It is now out for market in %s state\n",e.FSM.Current())
		},

	},
)

	return d
}

func (d *SupplyChain) enterState(e *fsm.Event) {
	fmt.Printf("The %s is in %s state\n", d.To, e.Dst)
}

func main() {
	sc := NewSupplyChain("SupplyChain")

	err:= sc.FSM.Event("created")
	if err != nil {
		fmt.Println(err)
	}

	err = sc.FSM.Event("shipToDistributor")
	if err != nil {
		fmt.Println(err)
	}

	err = sc.FSM.Event("shipToPharmacy")
	if err != nil {
		fmt.Println(err)
	}

	err = sc.FSM.Event("release")
	if err != nil {
		fmt.Println(err)
	}

}
