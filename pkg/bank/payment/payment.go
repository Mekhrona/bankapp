package payment

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentsSource []types.PaymentSource
	for _, card := range cards {

		if card.Active && card.Balance > 0 {
			paymentSource := types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			}

			paymentsSource = append(paymentsSource, paymentSource)
		}

	}

	return paymentsSource
}

func Max(payments []types.Payment) types.Payment {
	maxAmount := payments[0].Amount
	maxPayment := payments[0]
	for _, payment := range payments {

		if payment.Amount > maxAmount {
			maxAmount = payment.Amount
			maxPayment = payment
		}
	}
	return maxPayment
}
