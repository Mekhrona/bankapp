package card

import (
	"bank/pkg/bank/types"
	"fmt")


func ExampleTotal() {

	cards:=[] types.Card{

		{
			Balance: 10_000_00,
			Active: false,
		},
		{
			Balance: 90_000_00,
			Active: false,
		},

		{
			Balance: 985,
			Active: true,
		},
	}

	fmt.Println(Total(cards))
	
	//Output: 985
	
}