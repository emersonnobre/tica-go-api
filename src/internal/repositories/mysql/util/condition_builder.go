package util

import (
	"fmt"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
)

func BuildConditionsString(filters []repositories.Filter) string {
	result := "WHERE "
	for index, filter := range filters {
		if filter.Partial {
			result += fmt.Sprintf("%s like %s", filter.Key, transformToSqlValue(filter))
		} else {
			result += fmt.Sprintf("%s = %s", filter.Key, transformToSqlValue(filter))
		}
		if index != len(filters)-1 {
			result += " AND "
		}
	}
	return result
}

func transformToSqlValue(filter repositories.Filter) string {
	if filter.IsString {
		if filter.Partial {
			return fmt.Sprintf("'%%%s%%'", filter.Value)
		}
		return fmt.Sprintf("'%s'", filter.Value)
	}
	return filter.Value
}
