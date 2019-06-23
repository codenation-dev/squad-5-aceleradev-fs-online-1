package builder

import (
	"app/domain/model"
	"strconv"
)

// DashboardFromDB converte retorno do banco em estrutura esperada
func DashboardFromDB(results []map[string][]byte, totals []map[string][]byte) *model.Dashboard {

	return &model.Dashboard{
		Data:   getResults(results),
		Totals: getTotals(totals),
	}
}

func getResults(results []map[string][]byte) []model.DashboardData {
	data := make(map[string]model.Alerts)

	for _, res := range results {
		v, _ := res["yearmonth"]
		ym := string(v)

		dr, dok := data[ym]
		if !dok {
			dr = model.Alerts{}
		}

		tp := getInt(res, "type")
		switch tp {
		case 1:
			dr.PublicAgent = model.PublicAgentData{
				CustomerQuantity:     getInt(res, "count"),
				NotifyQuantity:       getInt(res, "alerts"),
				BiggerSalaryQuantity: getInt(res, "bigger"),
				NewQuantity:          getInt(res, "news"),
			}
		case 2:
			dr.BiggerSalary = model.BiggerSalary{
				CustomerQuantity: getInt(res, "count"),
				NotifyQuantity:   getInt(res, "alerts"),
			}
		case 3:
			dr.BankEmployee = model.BankEmployee{
				CustomerQuantity: getInt(res, "count"),
				NotifyQuantity:   getInt(res, "alerts"),
				NewQuantity:      getInt(res, "news"),
			}
		case 4:
			dr.Clients = model.Clients{
				NewQuantity: getInt(res, "news"),
			}
		}

		data[ym] = dr
	}

	dds := make([]model.DashboardData, 0)
	for key, value := range data {
		dds = append(dds, model.DashboardData{
			Month:  key,
			Alerts: value,
		})
	}
	return dds
}

func getTotals(totals []map[string][]byte) model.Totals {
	t := model.Totals{}

	for _, res := range totals {

		tp := getInt(res, "type")
		switch tp {
		case 1:
			t.PublicAgentQuantity = getInt(res, "count")

		case 3:
			t.EmployeeQuantity = getInt(res, "count")

		case 4:
			t.CustomerQuantity = getInt(res, "count")
		}
	}

	return t
}

func getInt(res map[string][]byte, name string) int64 {
	t, has := res[name]
	if !has {
		return 0
	}

	tp, err := strconv.Atoi(string(t))
	if err != nil {
		return 0
	}
	return int64(tp)
}
