package main

type Contract struct {
	constractID int
	client      string
	amount      float64
}

type ContractManager struct {
	Contracts []Contract
}

func (cm *ContractManager) AddContract(cID int, client string, amount float64) {
	contract := Contract{
		constractID: cID,
		client:      client,
		amount:      amount,
	}

	cm.Contracts = append(cm.Contracts, contract)
}

func (cm ContractManager) TotalAmountForClient(client string) float64 {
	var totalAmount float64

	for _, v := range cm.Contracts {
		if v.client == client {
			totalAmount += v.amount
		}
	}

	return totalAmount
}
