package model

type ClassifiedTransactionsMap struct {
	classifiedTransactions map[ClassificationType][]Transaction
}

func NewClassifiedTransactionsMap() ClassifiedTransactionsMap {
	return ClassifiedTransactionsMap{classifiedTransactions: make(map[ClassificationType][]Transaction)}
}

func (m ClassifiedTransactionsMap) Add(classificationType ClassificationType, transaction Transaction) {
	if _, ok := m.classifiedTransactions[classificationType]; !ok {
		m.classifiedTransactions[classificationType] = []Transaction{transaction}
		return
	}

	m.classifiedTransactions[classificationType] = append(m.classifiedTransactions[classificationType], transaction)
}

func (m ClassifiedTransactionsMap) Sum() (sum float64) {
	for _, transactions := range m.classifiedTransactions {
		for _, transaction := range transactions {
			sum += transaction.Amount
		}
	}

	return
}
