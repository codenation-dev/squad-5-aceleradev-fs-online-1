package repository

import (
	"app/domain/builder"
	"app/domain/model"
	"app/domain/validator"
	"time"

	"github.com/go-xorm/xorm"
)

// YearMonthFormat formato de ano e mês
const YearMonthFormat = "2006-01"

// DasboardDB interface
type DasboardDB interface {
	GetData(q validator.DashboardRequest) (*model.Dashboard, error)
	ListCustomers(q validator.DashboardCustomerRequest) ([]model.DashboardCustomer, error)
}

// DasboardRepository struct
type DasboardRepository struct {
	DB *xorm.Engine
}

// GetData retornar os dados estatisticos do dashboard
func (r DasboardRepository) GetData(q validator.DashboardRequest) (*model.Dashboard, error) {
	now := time.Now()
	if q.MonthStart.IsZero() {
		q.MonthStart = time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	}
	if q.MonthEnd.IsZero() {
		q.MonthEnd = time.Date(now.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
	}

	sql := `select x.yearmonth,x.type,sum(x.count) count,sum(x.alerts) alerts,sum(x.news) news,sum(x.bigger) bigger
	from (
		select q.yearmonth,q.type,count(*) count,sum(q.alerts) alerts,0 news,0 bigger
		from (
			select to_char(a.created_at, 'YYYY-MM') yearmonth,a.type,(
				select count(*) from alert_user where alert_user.alert_id=a.id
			) alerts
			from alert a
		) q
		 group by q.yearmonth,q.type
		union all
		select to_char(c.created_at, 'YYYY-MM') yearmonth,4,0,0,count(*),0
		from customer c
		group by yearmonth
		union all
		select to_char(u.created_at, 'YYYY-MM') yearmonth,3,0,0,count(*),0
		from "user" u
		group by yearmonth
		union all
		select to_char(p.created_at, 'YYYY-MM') yearmonth,1,0,0,count(*),0
		from public_agent p
		group by yearmonth
		union all
		select z.yearmonth,1,0,0,0,count(*)
		from (
			select to_char(a.created_at, 'YYYY-MM') yearmonth,a.type,(
				select count(*) from alert_user where alert_user.alert_id=a.id
			) alerts
			from alert a
			where a.type=2 and not a.public_agent_id is null
		) z
		 group by z.yearmonth,z.type
	) x
	where x.yearmonth >= ? and x.yearmonth <= ?
	group by x.yearmonth,x.type
	order by x.yearmonth,x.type`
	results, err := r.DB.Query(sql, q.MonthStart.Format(YearMonthFormat), q.MonthEnd.Format(YearMonthFormat))
	if err != nil {
		return nil, err
	}

	sqlTotal := `select 4 "type",count(*) count
	from customer c
	union all
	select 3,count(*)
	from "user" u
	union all
	select 1,count(*)
	from public_agent p`

	totals, err := r.DB.Query(sqlTotal)
	if err != nil {
		return nil, err
	}

	d := builder.DashboardFromDB(results, totals)

	return d, nil
}

// ListCustomers lista os ultimos alertas de clientes
func (r DasboardRepository) ListCustomers(q validator.DashboardCustomerRequest) ([]model.DashboardCustomer, error) {
	var a []model.DashboardCustomer
	if q.Limit == 0 {
		q.Limit = 20
	}

	err := r.DB.Table([]string{"alert", "a"}).
		Select(`a.id,c.name,a.created_at datetime, a.type,
	CASE WHEN c.salary > p.salary THEN c.salary
		 ELSE p.salary
	end salary,
	(select count(*) from alert_user au where au.alert_id=a.id) users_quantity`).
		Join("LEFT", []string{"customer", "c"}, "a.customer_id = c.id").
		Join("LEFT", []string{"public_agent", "p"}, "a.public_agent_id = p.id").
		OrderBy("a.created_at DESC").
		Limit(q.Limit, q.Offset).
		Find(&a)

	if err != nil {
		return nil, err
	}

	return a, nil
}
