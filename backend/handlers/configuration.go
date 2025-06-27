package handlers

import (
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConfiguration(c *gin.Context) {
	_, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	type ChartVariable struct {
		Key  models.UserDataCode `json:"key"`
		Name string              `json:"name"`
		Unit string              `json:"unit"`
	}
	type ConfigurationChart struct {
		ID         int64           `json:"id"`
		Title      string          `json:"title"`
		Type       string          `json:"type"`
		YAxisTitle string          `json:"yAxisTitle"`
		Variables  []ChartVariable `json:"variables"`
	}
	type Configuration struct {
		Charts []ConfigurationChart `json:"charts"`
	}

	configuration := Configuration{
		Charts: []ConfigurationChart{
			{
				ID:         1,
				Title:      "Speeds",
				Type:       "line",
				YAxisTitle: "Value (km/h)",
				Variables: []ChartVariable{
					{
						Key:  models.UserDataCodeK0D,
						Name: models.UserDataItems[models.UserDataCodeK0D].FullName,
						Unit: models.UserDataItems[models.UserDataCodeK0D].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1001,
						Name: models.UserDataItems[models.UserDataCodeKFF1001].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1001].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1237,
						Name: models.UserDataItems[models.UserDataCodeKFF1237].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1237].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1263,
						Name: models.UserDataItems[models.UserDataCodeKFF1263].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1263].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1272,
						Name: models.UserDataItems[models.UserDataCodeKFF1272].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1272].Unit,
					},
				},
			},
			{
				ID:         2,
				Title:      "Engine performance and load",
				Type:       "line",
				YAxisTitle: "",
				Variables: []ChartVariable{
					{
						Key:  models.UserDataCodeK04,
						Name: models.UserDataItems[models.UserDataCodeK04].FullName,
						Unit: models.UserDataItems[models.UserDataCodeK04].Unit,
					},
					{
						Key:  models.UserDataCodeK43,
						Name: models.UserDataItems[models.UserDataCodeK43].FullName,
						Unit: models.UserDataItems[models.UserDataCodeK43].Unit,
					},
					//{
					//	Key:  "kc",
					//	Name: "Engine RPM",
					//	Unit: "rpm",
					//},
					{
						Key:  models.UserDataCodeKFF1225,
						Name: models.UserDataItems[models.UserDataCodeKFF1225].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1225].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1226,
						Name: models.UserDataItems[models.UserDataCodeKFF1226].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1226].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1273,
						Name: models.UserDataItems[models.UserDataCodeKFF1273].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1273].Unit,
					},
				},
			},
			{
				ID:         3,
				Title:      "Fuel consumption and efficiency",
				Type:       "line",
				YAxisTitle: "",
				Variables: []ChartVariable{
					{
						Key:  models.UserDataCodeKFF1201,
						Name: models.UserDataItems[models.UserDataCodeKFF1201].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1201].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1203,
						Name: models.UserDataItems[models.UserDataCodeKFF1203].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1203].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1207,
						Name: models.UserDataItems[models.UserDataCodeKFF1207].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1207].Unit,
					},
					{
						Key:  models.UserDataCodeKFF125A,
						Name: models.UserDataItems[models.UserDataCodeKFF125A].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF125A].Unit,
					},
					{
						Key:  models.UserDataCodeKFF125D,
						Name: models.UserDataItems[models.UserDataCodeKFF125D].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF125D].Unit,
					},
					{
						Key:  models.UserDataCodeKFF1271,
						Name: models.UserDataItems[models.UserDataCodeKFF1271].FullName,
						Unit: models.UserDataItems[models.UserDataCodeKFF1271].Unit,
					},
				},
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"configuration": configuration,
	})
}
