package card

import 

"bank/pkg/bank/types" 


var total types.Money=0
func Total(cards []types.Card)  types.Money{

	for _,card := range cards {

		if card.Balance>0 && card.Active {
		total+=card.Balance
		}	
	}
	return total
}