package car

import "database/sql"

type Manifest struct {
	DB            *sql.DB
	FilterResolver func(string) (string, error)
}


func (m *Manifest) Query1(ctx Query1Context) ([]Query1Row, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return []Query1Row{}, err
		}
		ctx.Filter = filter
	}
	return Query1(m.DB, ctx)
}

