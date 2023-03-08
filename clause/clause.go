package clause

import "strings"

type Type int

const (
	INSERT = iota
	VALUES
	SELECT
	WHERE
	LIMIT
	ORDER
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]any
}

func (c *Clause) Set(name Type, vars ...any) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]any)
	}
	sql, vars := generatorMap[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

func (c *Clause) Build(orders ...Type) (string, []any) {
	var sqls []string
	var vars []any
	for _, v := range orders {
		if s, ok := c.sql[v]; ok {
			sqls = append(sqls, s)
			vars = append(vars, c.sqlVars[v]...)
		}
	}

	return strings.Join(sqls, " "), vars
}
