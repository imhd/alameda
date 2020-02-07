package schemas

import (
	"errors"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"strconv"
	"strings"
)

type Column struct {
	Name       string
	Required   bool
	ColumnType ColumnType
	DataType   common.DataType
}

func NewColumn() *Column {
	column := Column{}
	return &column
}

func (p *Column) Parse(column string) error {
	trimmed := strings.TrimSpace(column)
	formatList := strings.Split(trimmed, "-")

	if len(formatList) == 4 {
		name := formatList[0]

		required, err := strconv.ParseBool(formatList[1])
		if err != nil {
			return errors.New("wrong format of column string")
		}

		columnType, err := strconv.ParseInt(formatList[2], 10, 64)
		if err != nil {
			return errors.New("wrong format of column string")
		}

		dataType, err := strconv.ParseInt(formatList[3], 10, 64)
		if err != nil {
			return errors.New("wrong format of column string")
		}

		p.Name = name
		p.Required = required
		p.ColumnType = ColumnType(columnType)
		p.DataType = common.DataType(dataType)
	}

	return nil
}

func (p *Column) String() string {
	values := make([]string, 0)
	values = append(values, p.Name)
	values = append(values, bool2String(p.Required))
	values = append(values, strconv.FormatInt(int64(p.ColumnType), 10))
	values = append(values, strconv.FormatInt(int64(p.DataType), 10))
	return strings.Join(values, "-")
}

func bool2String(b bool) string {
	if b == true {
		return "1"
	}
	return "0"
}
