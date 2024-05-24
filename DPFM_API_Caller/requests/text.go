package requests

type Text struct {
	PointConsumptionType		string  `json:"PointConsumptionType"`
	Language          			string  `json:"Language"`
	PointConsumptionTypeName	string  `json:"PointConsumptionTypeName"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
