package model

//ClassificationType determines the type of a transaction
type ClassificationType string

const (
	//Invalid classification
	Invalid ClassificationType = "Invalid"
	//Miscellaneous classification
	Miscellaneous = "Miscellaneous"
	//Food classification
	Food = "Food"
	//Salary classification
	Salary = "Salary"
	//Income classification
	Income = "Income"
	//RentAndDebt classification
	RentAndDebt = "RentAndDebt"
	//Entertainment classification
	Entertainment = "Entertainment"
	//Travel classification
	Travel = "Travel"
	//Clothes classification
	Clothes = "Clothes"
	//Debt classification
	Debt = "Debt"
	//Tax classification
	Tax = "Tax"
	//LastMonthBalance classification
	LastMonthBalance = "LastMonthBalance"
	//Unclassified classification
	Unclassified = "Unclassified"
)
