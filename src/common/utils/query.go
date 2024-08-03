package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetQueryUUID(rows []uuid.UUID) string {
	if len(rows) == 0 {
		return ""
	}

	qryVal := ""
	for i := range rows {
		qryVal = fmt.Sprintf("%s%v,", qryVal, rows[i])
	}
	qryVal = strings.TrimRight(qryVal, ",")
	fmt.Println(qryVal)
	return qryVal
}
