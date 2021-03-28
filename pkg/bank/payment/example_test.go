package payment

import (
	"bank/pkg/bank/types"	
	"fmt"
)





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

func ExampleMax() {
	payments := [] types.Payment{
		{
			ID: 78,
			Amount: 108,
		},
		{
			ID: 99,
			Amount: 108,
		},

	}

	max:=Max(payments)
		fmt.Println(max.ID)

	//Output:78

}