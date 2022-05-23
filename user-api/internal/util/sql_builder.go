package util

type SqlBuildHelper struct {
}

func NewSqlBuildHelper() *SqlBuildHelper {
	return &SqlBuildHelper{}
}

func (s *SqlBuildHelper) AddCondition(value []interface{}, condition string) string {
	if len(value) > 0 {
		return condition
	}
	return ""
}
