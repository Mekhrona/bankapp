package card

import (
	"bank/pkg/bank/types"
	"fmt")




	func ExamplePaymentSource() {

		cards:= [] types.Card{
			{
				PAN :     "6768 6845 3456",
				Balance:  58949,
				Active:   true,
			},
			{
				PAN :     "7848 6845 3456",
				Balance:  -89,
				Active:   true,
			},
			{
				PAN :     "9856 6845 3456",
				Balance:  899,
				Active:   false,
			},
			{
				PAN :     "6768 6845 3456",
				Balance:  89512,
				Active:   true,
			},
		}
		
			PaymentSources:=PaymentSources(cards)
			for _, paymentSource := range PaymentSources {
		
				fmt.Println(paymentSource.Balance)
			}
			
		
			//Output: 58949
			// 89512
			
		}

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