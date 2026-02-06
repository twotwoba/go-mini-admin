package utils

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// IsDuplicateError 判断是否为唯一索引冲突错误
func IsDuplicateError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	return false
}
