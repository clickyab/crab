package orm

import (
	"time"

	"fmt"

	"strings"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

var (
	defaultCTR = config.RegisterFloat64("crab.modules.campaign.ctr", .1, "default ctr in the system")
)

// CampaignKind is kind of campaign <web,app>
// @Enum{
// }
type CampaignKind string

const (
	// WebCampaign is web
	WebCampaign CampaignKind = "web"
	// AppCampaign is app
	AppCampaign CampaignKind = "app"
)

// CampaignType is type of campaign <vast,banner,native>
// @Enum{
// }
type CampaignType string

const (

	// BannerType of campaign
	BannerType CampaignType = "banner"
	// VastType   of campaign
	VastType CampaignType = "vast"
	// NativeType of campaign
	NativeType CampaignType = "native"
)

// Progress is progress of campaign
// @Enum{
// }
type Progress string

const (
	// ProgressInProgress is in progress
	ProgressInProgress Progress = "inprogress"
	// ProgressFinalized is finalized
	ProgressFinalized Progress = "finalized"
)

// CostType is type of campaign <cpm,cpc,cpa>
// @Enum{
// }
type CostType string

const (
	// CPM is cpm
	CPM CostType = "cpm"
	// CPC is cpc
	CPC CostType = "cpc"
	// CPA is cpa
	CPA CostType = "cpa"

	wh string = "where"
)

type base struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
}

// Campaign campaign model in database
// @Model {
//		table = campaigns
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Campaign struct {
	base
	CampaignBaseType
	CampaignStatus
	CampaignFinance
	UserID       int64           `json:"user_id" db:"user_id"`
	DomainID     int64           `json:"domain_id" db:"domain_id"`
	Exchange     bool            `json:"exchange" db:"exchange"`
	WhiteBlackID mysql.NullInt64 `json:"white_black_id" db:"white_black_id"`
	// WhiteBlackType true is whitelist
	WhiteBlackType  mysql.NullBool           `json:"white_black_type" db:"white_black_type"`
	WhiteBlackValue mysql.StringMapJSONArray `json:"-" db:"white_black_value"`
	Progress        Progress                 `json:"-" db:"progress"`
	Attributes      *CampaignAttributes      `json:"attributes,omitempty" db:"-"`
	ArchiveAt       mysql.NullTime           `json:"archive_at" db:"archive_at"`
}

// CampaignDataTable is the campaign full data in data table
// @DataTable {
//		url = /list
//		entity = campaign
//		view = campaign_list:self
//		checkable = false
//		multiselect = false
//		datefilter = cp.created_at
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignDataTableArray
//		_edit = campaign_edit:self
//		_copy = campaign_copy:self
// }
type CampaignDataTable struct {
	ID        int64     `json:"id" db:"id" type:"number"`
	CreatedAt time.Time `json:"created_at" db:"created_at" type:"date" sort:"true" map:"cp.created_at"`
	Active    bool      `json:"active" db:"active" type:"bool"`

	Kind CampaignKind `json:"kind" db:"kind" type:"enum" filter:"true" map:"cp.kind"`
	Type CampaignType `json:"type" db:"type" type:"enum" filter:"true" map:"cp.type"`

	Status  bool           `json:"status" db:"status" type:"bool"`
	StartAt time.Time      `json:"start_at" db:"start_at" type:"date" sort:"true"`
	EndAt   mysql.NullTime `json:"end_at" db:"end_at" type:"date"`
	Title   string         `json:"title" db:"title" type:"string" search:"true" map:"cp.title"`

	Budget     int64    `json:"budget" db:"budget" type:"number"`
	DailyLimit int64    `json:"daily_limit" db:"daily_limit" type:"number"`
	CostType   CostType `json:"cost_type" db:"cost_type" type:"enum" filter:"true" map:"cp.cost_type"`
	MaxBid     int64    `json:"max_bid" db:"max_bid" type:"number" sort:"true"`

	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,false"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,false"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,true"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,true"`
	TotalConv  int64   `json:"total_conv" db:"total_conv"`
	TotalCpc   int64   `json:"total_cpc" db:"total_cpc"`
	TotalCpm   int64   `json:"total_cpm" db:"total_cpm"`

	TotalSpent int64 `json:"total_spent" db:"-" graph:"total_spent,Total spent,line,false"`

	TodayImp   int64   `json:"today_imp" db:"today_imp"`
	TodayClick int64   `json:"today_click" db:"today_click"`
	TodayCtr   float64 `json:"today_ctr" db:"today_ctr"`

	ParentIDs   []int64          `db:"-" json:"-" visible:"false"`
	ParentEmail mysql.NullString `db:"parent_email" json:"parent_email"`
	OwnerEmail  string           `db:"owner_email" json:"owner_email" type:"string" search:"true" map:"owner.email"`
	OwnerID     int64            `db:"owner_id" json:"owner_id" visible:"false"`
	DomainID    int64            `db:"domain_id" json:"domain_id"`
	Actions     string           `db:"-" json:"_actions" visible:"false"`
}

// CampaignFinance is the financial
type CampaignFinance struct {
	Budget      int64                 `json:"budget" db:"budget"`
	DailyLimit  int64                 `json:"daily_limit" db:"daily_limit"`
	CostType    CostType              `json:"cost_type" db:"cost_type"`
	MaxBid      int64                 `json:"max_bid" db:"max_bid"`
	NotifyEmail mysql.StringJSONArray `json:"notify_email" db:"notify_email"`
}

// CampaignBaseType is fundamental data of campaign
type CampaignBaseType struct {
	Kind CampaignKind `json:"kind" db:"kind"`
	Type CampaignType `json:"type" db:"type"`
}

// CampaignStatus update campaign (stage one)
type CampaignStatus struct {
	Status   bool           `json:"status" db:"status"`
	StartAt  time.Time      `json:"start_at" db:"start_at"`
	EndAt    mysql.NullTime `json:"end_at" db:"end_at"`
	Title    string         `json:"title" db:"title" `
	Schedule ScheduleSheet  `json:"schedule" db:"-"`
}

// CampaignBase is minimum data for creating campaign (stage one)
type CampaignBase struct { // stage one create
	CampaignBaseType
	CampaignStatus
}

func (ca *Campaign) webMaxBid(c CampaignBase) {
	switch c.Type {
	case BannerType:
		ca.MaxBid = defaultWebBannerCPC.Int64()
	case VastType:
		ca.MaxBid = defaultWebVastCPC.Int64()
	case NativeType:
		ca.MaxBid = defaultWebNativeCPC.Int64()
	}
}
func (ca *Campaign) appMaxBid(c CampaignBase) {
	switch c.Type {
	case BannerType:
		ca.MaxBid = defaultAppBannerCPC.Int64()
	case VastType:
		ca.MaxBid = defaultAppVastCPC.Int64()
	case NativeType:
		ca.MaxBid = defaultAppNativeCPC.Int64()
	}
}

// FindCampaignByIDDomain return the Campaign base on its id and domain id
func (m *Manager) FindCampaignByIDDomain(id, d int64) (*Campaign, error) {
	var res Campaign
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=? AND domain_id=?", CampaignTableFull),
		id,
		d,
	)

	if err != nil {
		return nil, err
	}

	m.attachAttribute(&res)
	m.attachSchedule(&res)
	return &res, nil
}

// CampaignGraph is the campaign full data in data table
// @Graph {
//		url = /graph/all
//		entity = chart
//		view = campaign_graph:self
//		key = ID
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignGraph
// }
type CampaignGraph struct {
	OwnerEmail string       `db:"owner_email" json:"owner_email" type:"string" search:"true" map:"owner.email"`
	Kind       CampaignKind `json:"kind" db:"kind" type:"enum" filter:"true" map:"cp.kind"`
	Type       CampaignType `json:"type" db:"type" type:"enum" filter:"true" map:"cp.type"`
	Title      string       `json:"title" db:"title" type:"string" search:"true" map:"cp.title"`

	ID         int64   `json:"id" db:"id" type:"number"`
	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,false"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm" graph:"avg_cpm,Avg. CPM,line,false"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,false"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,true"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,true"`
	TotalSpent int64   `json:"total_spent" db:"total_spent" graph:"total_spent,Total spent,line,false"`
}

// FillCampaignGraph is the function to handle
func (m *Manager) FillCampaignGraph(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	from, to time.Time) []CampaignGraph {
	res := make([]CampaignGraph, 0)

	query := fmt.Sprintf(`SELECT cd.daily_id as id,
	COALESCE(AVG(cd.cpc),0) AS avg_cpc,
	COALESCE(AVG(cd.cpm),0) AS avg_cpm,
	COALESCE(SUM(cd.click),0) AS total_click,
	COALESCE(SUM(cd.imp),0) AS total_imp,
	COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0) AS ctr,
	COALESCE(SUM(cd.cpc)+SUM(cd.cpm),0) AS total_spent
	FROM %s AS cp INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id `,
		CampaignTableFull, aaa.UserTableFull, aaa.ParentUserTableFull, aaa.UserTableFull, CampaignDetailTableFull)

	var where []string
	wh := " WHERE "

	where = append(where, fmt.Sprintf(`%s BETWEEN %d AND %d`, "cd.daily_id",
		libs.TimeToID(from),
		libs.TimeToID(to)))
	var params []interface{}
	params = append(params, pc.GetDomainID())

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}
	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	currentUserID := pc.GetID()
	highestScope := pc.GetCurrentScope()

	// find current user childes
	userManager := aaa.NewAaaManager()
	childes := userManager.GetUserChildesIDDomain(currentUserID, pc.GetDomainID())
	childes = append(childes, currentUserID)
	// self or parent
	if highestScope == permission.ScopeSelf {
		//check if parent or owner
		where = append(where, fmt.Sprintf("cp.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}
	//check for perm
	if len(where) > 0 {
		query += wh
	}
	query += strings.Join(where, " AND ")

	query += " GROUP BY cd.daily_id"
	_, err := m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	return res
}

// FillCampaignDataTableArray is the function to handle
func (m *Manager) FillCampaignDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	dateRange map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignDataTableArray, int64) {
	var params []interface{}
	var res CampaignDataTableArray
	var where []string
	todayInt := libs.TimeToID(time.Now())
	countQuery := fmt.Sprintf(`SELECT COUNT(cp.id) FROM %s AS cp
	INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id
	LEFT JOIN %s AS ycd ON (ycd.campaign_id=cp.id AND ycd.daily_id=%d)`,
		CampaignTableFull, aaa.UserTableFull, aaa.ParentUserTableFull, aaa.UserTableFull, CampaignDetailTableFull, CampaignDetailTableFull, todayInt)
	query := fmt.Sprintf(`SELECT cp.id AS id,
	cp.title,
	cp.kind,
	cp.daily_limit,
	cp.type,
	cp.status,
	cp.max_bid,
	cp.cost_type,
	cp.budget,
	cp.start_at,
	cp.end_at,
	cp.created_at,
	cp.domain_id AS domain_id,
	owner.email AS owner_email,
	owner.id AS owner_id,
	parent.email AS parent_email,
	COALESCE(AVG(cd.cpc),0) AS avg_cpc,
	COALESCE(AVG(cd.cpm),0)AS avg_cpm,
	COALESCE(SUM(cd.click),0) AS total_click,
	COALESCE(SUM(cd.imp),0) AS total_imp,
	COALESCE(SUM(cd.conv),0) AS total_conv,
	COALESCE(SUM(cd.cpc),0) AS total_cpc,
	COALESCE(SUM(cd.cpm),0) AS total_cpm,
	COALESCE(ycd.imp,0) AS today_imp,
	COALESCE(ycd.click,0) AS today_click
	FROM %s AS cp INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id
	LEFT JOIN %s AS ycd ON (ycd.campaign_id=cp.id AND ycd.daily_id=%d)`,
		CampaignTableFull, aaa.UserTableFull, aaa.ParentUserTableFull, aaa.UserTableFull, CampaignDetailTableFull, CampaignDetailTableFull, todayInt)

	//check for date range
	var dateRangeField string
	var from string
	var to string
	for key, val := range dateRange {
		dateRangeArr := strings.Split(key, "-")
		if len(dateRangeArr) == 2 {
			dateRangeField = dateRangeArr[1]
			if dateRangeArr[0] == "from" {
				from = val
			}
			if dateRangeArr[0] == "to" {
				to = val
			}
		}
	}
	if dateRangeField != "" && from != "" && to != "" {
		fromTime, err1 := time.Parse(time.RFC3339, from)
		toTime, err2 := time.Parse(time.RFC3339, to)

		if err1 == nil && err2 == nil {
			where = append(where,
				fmt.Sprintf(`%s BETWEEN "%s" AND "%s"`, dateRangeField,
					fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00"),
					toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")))
		}
	}

	params = append(params, pc.GetDomainID())
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}
	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}

	currentUserID := pc.GetID()
	highestScope := pc.GetCurrentScope()

	// find current user childes
	userManager := aaa.NewAaaManager()
	childes := userManager.GetUserChildesIDDomain(currentUserID, pc.GetDomainID())
	childes = append(childes, currentUserID)
	// self or parent
	if highestScope == permission.ScopeSelf {
		//check if parent or owner
		where = append(where, fmt.Sprintf("cp.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}
	//check for perm
	if len(where) > 0 {
		query += wh
		countQuery += wh
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")

	countQuery += " GROUP BY cp.id "
	query += " GROUP BY cp.id "

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
	for i := range res {
		res[i].Ctr = calculateCtr(res[i].TotalImp, res[i].TotalClick, defaultCTR.Float64())
		res[i].TodayCtr = calculateCtr(res[i].TodayImp, res[i].TodayClick, defaultCTR.Float64())
		res[i].TotalSpent = func() int64 {
			if res[i].CostType == CPC {
				return res[i].TotalCpc
			}
			return res[i].TotalCpm
		}()
	}

	return res, count
}

func calculateCtr(imp, click int64, def float64) float64 {
	if imp == 0 || click == 0 {
		return def
	}
	return float64(click) * 10 / float64(imp)
}
