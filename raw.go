package gorm

type RawSqls struct {
	RawSql []RawSql
}
type RawSql struct {
	Sql  string
	Vars []interface{}
}

func newRawSqls() *RawSqls {
	return &RawSqls{
		RawSql: []RawSql{},
	}
}
func (rss *RawSqls) add(rs ...RawSql) *RawSqls {
	for _, r := range rs {
		rss.RawSql = append(rss.RawSql, r)
	}
	return rss
}
