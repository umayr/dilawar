package dilawar

func transaction(amount int, msg string, kind Type) error {
	return NewStore().Create(&Transaction{
		Amount:      amount,
		Type:        kind,
		Description: msg,
	})
}

func Credit(amount int, msg string) error {
	return transaction(amount, msg, TypeCredit)
}

func Debit(amount int, msg string) error {
	return transaction(amount, msg, TypeDebit)
}

func History() ([]Transaction, error) {
	return NewStore().List()
}

func Balance() (int, error) {
	items, err := History()
	if err != nil {
		return 0, err
	}

	amount := 0
	for _, v := range items {
		if v.Type == TypeCredit {
			amount -= v.Amount
		}
		if v.Type == TypeDebit {
			amount += v.Amount
		}
	}

	return amount, nil
}
