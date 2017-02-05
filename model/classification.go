package model

//Classification determines the classification of a transaction
type Classification struct {
	Client  string
	Purpose string
	Type    ClassificationType
}
