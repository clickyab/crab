package orm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// CampaignLog is the campaign log data in data-table
// @DataTable {
//		url = /log/:id
//		entity = CampaignLog
//		checkable = false
//		datefilter = created_at
//		multiselect = false
//		view = log_campaign:self
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignLog
// }
type CampaignLog struct {
	CreatedAt         time.Time              `json:"created_at" db:"created_at" visible:"true"`
	Action            aaa.AuditActionType    `json:"action" db:"action" visible:"true"`
	ImpersonatorEmail mysql.NullString       `json:"impersonator_email" db:"impersonator_email" visible:"true"`
	ManipulatorEmail  mysql.NullString       `json:"manipulator_email" db:"manipulator_email" visible:"true"`
	OwnerEmail        mysql.NullString       `json:"owner_email" db:"owner_email" visible:"false"`
	Data              mysql.GenericJSONField `json:"data" db:"data" visible:"false"`

	CampaignName string       `json:"campaign_name" visible:"true"`
	Kind         CampaignKind `json:"kind" visible:"true"`
	Strategy     Strategy     `json:"strategy" visible:"true"`
	StartAt      string       `json:"start_at" visible:"true"`
	EndAt        string       `json:"end_at" visible:"true"`
	TotalBudget  float64      `json:"total_budget" visible:"false"`
	DailyBudget  float64      `json:"daily_budget" visible:"true"`
	MaxBid       float64      `json:"max_bid" visible:"true"`

	OwnerID   int64   `db:"-" json:"-" visible:"false"`
	DomainID  int64   `db:"-" json:"-" visible:"false"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"-" visible:"false"`
}

// FillCampaignLog is the function to handle
func (m *Manager) FillCampaignLog(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignLogArray, int64, error) {

	var params []interface{}
	var res CampaignLogArray
	var where []string
	var whereLike []string

	val, ok := contextparams["id"]
	if !ok {
		return nil, 0, errors.New("")
	}
	intVal, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return nil, 0, errors.New("")
	}

	params = append(params, intVal)
	params = append(params, "campaign")

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for date filter
	if from != "" && to != "" {
		fromArr := strings.Split(from, "*")
		toArr := strings.Split(to, "*")
		fromTime, err := time.Parse("2006-01-02 15:04:05", fromArr[1])
		if err != nil {
			return nil, 0, errors.New("")
		}
		toTime, err := time.Parse("2006-01-02 15:04:05", toArr[1])
		if err != nil {
			return nil, 0, errors.New("")
		}
		where = append(where, fmt.Sprintf(`a.%s BETWEEN ? AND ?`, fromArr[0]))
		params = append(params, fromTime, toTime)
	}

	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "c.domain_id"))
	params = append(params, pc.GetDomainID())

	highestScope := pc.GetCurrentScope()
	if highestScope == permission.ScopeSelf {
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "log_campaign", pc.GetDomainID())
		childes = append(childes, pc.GetID())
		where = append(where, fmt.Sprintf("c.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}

	for column, val := range search {
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}

	if len(whereLike) > 0 {
		wl := "(" + strings.Join(whereLike, " OR ") + ")"
		where = append(where, wl)
	}

	var conds string
	if len(where) > 0 {
		conds += wh
	}
	conds += strings.Join(where, " AND ")

	countQuery := fmt.Sprintf(`SELECT 
	COUNT(a.id) FROM %s AS a 
	INNER JOIN %s AS c ON (c.id=a.target_id AND a.target_id=? AND a.target_model=?) %s`,
		aaa.AuditLogTableFull,
		CampaignTableFull,
		conds,
	)
	query := fmt.Sprintf(`SELECT
	a.created_at AS created_at,
	al.data AS data,
	a.action AS action,
	owner.email AS owner_email,
	manipulator.email AS manipulator_email,
	impersonator.email AS impersonator_email FROM %s AS a
	INNER JOIN %s AS al ON al.audit_log_id=a.id
	INNER JOIN %s AS c ON (c.id=a.target_id AND a.target_id=? AND a.target_model=?)
	INNER JOIN %s AS owner ON owner.id=c.user_id
	LEFT JOIN %s AS impersonator ON impersonator.id=a.impersonator_id
	JOIN %s AS manipulator ON manipulator.id=a.user_id %s
	
`, aaa.AuditLogTableFull,
		aaa.AuditLogDetailTableFull,
		CampaignTableFull,
		aaa.UserTableFull,
		aaa.UserTableFull,
		aaa.UserTableFull,
		conds)

	limit := c
	offset := (p - 1) * c
	if sort != "" {
		query += fmt.Sprintf(" ORDER BY %s %s ", sort, order)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)
	count, err := m.GetRDbMap().SelectInt(countQuery, params...)
	assert.Nil(err)

	_, err = m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	fillCampaignLogData(res)

	return res, count, nil
}

func fillCampaignLogData(res CampaignLogArray) CampaignLogArray {
	for i := range res {
		strategy, ok := res[i].Data["strategy"]
		if ok {
			res[i].Strategy = Strategy(strategy.(string))
		}
		startAt, ok := res[i].Data["start_at"]
		if ok {
			res[i].StartAt = startAt.(string)
		}

		endAt, ok := res[i].Data["end_at"]
		if ok {
			res[i].EndAt = endAt.(string)
		}

		kind, ok := res[i].Data["kind"]
		if ok {
			res[i].Kind = CampaignKind(kind.(string))
		}

		campaignName, ok := res[i].Data["title"]
		if ok {
			res[i].CampaignName = campaignName.(string)
		}

		totalBudget, ok := res[i].Data["total_budget"]
		if ok {
			res[i].TotalBudget = totalBudget.(float64)
		}

		dailyBudget, ok := res[i].Data["daily_budget"]
		if ok {
			res[i].DailyBudget = dailyBudget.(float64)
		}

		maxBid, ok := res[i].Data["max_bid"]
		if ok {
			res[i].MaxBid = maxBid.(float64)
		}

	}
	return res
}
