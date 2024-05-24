package dpfm_api_output_formatter

import (
	"data-platform-api-point-consumption-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToPointConsumptionType(rows *sql.Rows) (*[]PointConsumptionType, error) {
	defer rows.Close()
	pointConsumptionType := make([]PointConsumptionType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PointConsumptionType{}

		err := rows.Scan(
			&pm.PointConsumptionType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &pointConsumptionType, nil
		}

		data := pm
		pointConsumptionType = append(pointConsumptionType, PointConsumptionType{
			PointConsumptionType:	data.PointConsumptionType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &pointConsumptionType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.PointConsumptionType,
			&pm.Language,
			&pm.PointConsumptionTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			PointConsumptionType:		data.PointConsumptionType,
			Language:          			data.Language,
			PointConsumptionTypeName:	data.PointConsumptionTypeName,
			CreationDate:				data.CreationDate,
			LastChangeDate:				data.LastChangeDate,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
