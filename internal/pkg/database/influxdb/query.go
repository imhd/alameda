package influxdb

import (
	"fmt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"strings"
	"time"
)

type InfluxQuery struct {
	QueryCondition *common.QueryCondition
	Measurement    string
	Conditions     []string
}

func NewQuery(queryCondition *common.QueryCondition, measurement string) *InfluxQuery {
	influxQuery := InfluxQuery{
		QueryCondition: queryCondition,
		Measurement:    measurement,
		Conditions:     make([]string, 0),
	}
	if influxQuery.QueryCondition == nil {
		influxQuery.QueryCondition = &common.QueryCondition{}
	}
	return &influxQuery
}

func (p *InfluxQuery) SetTimeRange(startTime, endTime *time.Time) {
	p.QueryCondition.StartTime = startTime
	p.QueryCondition.EndTime = endTime
}

func (p *InfluxQuery) SetStep(step int) {
	if step != 0 {
		duration := time.Duration(step) * time.Second
		p.QueryCondition.StepTime = &duration
	} else {
		p.QueryCondition.StepTime = nil
	}
}

func (p *InfluxQuery) SetOrder(order common.Order) {
	p.QueryCondition.TimestampOrder = order
}

func (p *InfluxQuery) SetAggregateFunction() {

}

func (p *InfluxQuery) AppendSelects(selects []string) {
	if p.QueryCondition.Selects == nil {
		p.QueryCondition.Selects = make([]string, 0)
	}
	for _, s := range selects {
		p.QueryCondition.Selects = append(p.QueryCondition.Selects, s)
	}
}

func (p *InfluxQuery) AppendGroups(groups []string) {
	if p.QueryCondition.Groups == nil {
		p.QueryCondition.Groups = make([]string, 0)
	}
	for _, group := range groups {
		p.QueryCondition.Groups = append(p.QueryCondition.Groups, group)
	}
}

func (p *InfluxQuery) AppendCondition(keys, values, operators []string, dataTypes []common.DataType) {
	if p.QueryCondition.WhereCondition == nil {
		p.QueryCondition.WhereCondition = make([]*common.Condition, 0)
	}
	condition := common.Condition{
		Keys:      keys,
		Values:    values,
		Operators: operators,
		Types:     dataTypes,
	}
	p.QueryCondition.WhereCondition = append(p.QueryCondition.WhereCondition, &condition)
}

func (p *InfluxQuery) AppendConditionDirectly(condition string) {
	p.Conditions = append(p.Conditions, condition)
}

func (p *InfluxQuery) BuildQueryCmd() string {
	// SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] [GROUP_BY_clause] [ORDER_BY_clause] LIMIT_clause OFFSET <N> [SLIMIT_clause]
	cmd := fmt.Sprintf("SELECT %s FROM \"%s\" %s %s %s %s",
		p.selects(), p.Measurement, p.whereClause(),
		p.groupClause(), p.orderClause(), p.limitClause())
	return cmd
}

func (p *InfluxQuery) BuildDropCmd() string {
	cmd := fmt.Sprintf("DROP SERIES FROM \"%s\" %s", p.Measurement, p.whereClause())
	return cmd
}

func (p *InfluxQuery) selects() string {
	selects := make([]string, 0)
	if p.QueryCondition.Selects == nil {
		selects = append(selects, "*")
	} else {
		for _, s := range p.QueryCondition.Selects {
			if s == "*" {
				selects = append(selects, "*")
			} else {
				selects = append(selects, fmt.Sprintf(`"%s"`, s))
			}
		}
	}
	aggregateFunc := p.QueryCondition.AggregateOverTimeFunction
	if aggregateFunc != common.None {
		aggregateName := AggregateFuncMap[aggregateFunc]
		return fmt.Sprintf("%s(%s)", aggregateName, strings.Join(selects, ","))
	}
	return strings.Join(selects, ",")
}

func (p *InfluxQuery) whereClause() string {
	whereClause := make([]string, 0)
	if p.QueryCondition.WhereClause != "" {
		return p.QueryCondition.WhereClause
	}
	where := p.where(p.QueryCondition.WhereCondition)
	if where != "" {
		whereClause = append(whereClause, where)
	}
	timeRange := p.timeRange()
	if timeRange != "" {
		whereClause = append(whereClause, timeRange)
	}
	if len(whereClause) != 0 {
		return fmt.Sprintf("WHERE %s", strings.Join(whereClause, " AND "))
	}
	return ""
}

func (p *InfluxQuery) groupClause() string {
	groups := make([]string, 0)
	if p.QueryCondition.Groups != nil {
		for _, group := range p.QueryCondition.Groups {
			if strings.HasPrefix(group, "time(") {
				groups = append(groups, group)
			} else {
				groups = append(groups, fmt.Sprintf(`"%s"`, group))
			}
		}
		return fmt.Sprintf("GROUP BY %s", strings.Join(groups, ","))
	}
	return ""
}

func (p *InfluxQuery) orderClause() string {
	switch p.QueryCondition.TimestampOrder {
	case common.Asc:
		return "ORDER BY time ASC"
	case common.Desc:
		return "ORDER BY time DESC"
	default:
		return "ORDER BY time ASC"
	}
}

func (p *InfluxQuery) limitClause() string {
	if p.QueryCondition.Limit > 0 {
		return fmt.Sprintf("LIMIT %v", p.QueryCondition.Limit)
	}
	return ""
}

func (p *InfluxQuery) where(conditions []*common.Condition) string {
	where := make([]string, 0)
	for index := 0; index < len(conditions); index++ {
		condition := p.condition(conditions[index].Keys, conditions[index].Operators, conditions[index].Values, conditions[index].Types)
		where = append(where, condition)
	}
	if len(where) != 0 {
		return fmt.Sprintf("(%s)", strings.Join(where, " OR "))
	}
	return ""
}

func (p *InfluxQuery) condition(keys, operators, values []string, dataTypes []common.DataType) string {
	condition := make([]string, 0)
	for index := 0; index < len(keys); index++ {
		expression := p.expression(keys[index], operators[index], values[index], dataTypes[index])
		condition = append(condition, expression)
	}
	if len(condition) != 0 {
		return fmt.Sprintf("(%s)", strings.Join(condition, " AND "))
	}
	return ""
}

func (p *InfluxQuery) expression(key, operator, value string, dataType common.DataType) string {
	switch dataType {
	case common.Bool:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Int:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Int8:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Int16:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Int32:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Int64:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Uint:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Uint8:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Uint16:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Uint32:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Uint64:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Float32:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.Float64:
		return fmt.Sprintf("\"%s\"%s%s", key, operator, value)
	case common.String:
		return fmt.Sprintf("\"%s\"%s'%s'", key, operator, value)
	default:
		fmt.Println("not support")
		return fmt.Sprintf(`%s%s"%s"`, key, operator, value)
	}
}

func (p *InfluxQuery) timeRange() string {
	timeRange := make([]string, 0)
	if p.QueryCondition.StartTime != nil {
		timeRange = append(timeRange, fmt.Sprintf("time%s'%s'", ">=", p.QueryCondition.StartTime.Format(time.RFC3339)))
	}
	if p.QueryCondition.EndTime != nil {
		timeRange = append(timeRange, fmt.Sprintf("time%s'%s'", "<=", p.QueryCondition.EndTime.Format(time.RFC3339)))
	}
	return strings.Join(timeRange, " AND ")
}
