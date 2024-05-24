package requests

type PointConsumptionType struct {
	PointConsumptionType	string	`json:"PointConsumptionType"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
