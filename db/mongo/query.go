package mongo

type Query interface {
	All(interface{}) error
	One(interface{}) error
	Sort(...string) Query
	Select(interface{}) Query
	Limit(int) Query
	Skip(int) Query
	Count() (int, error)
}

func (m *mgodb) All(value interface{}) error {
	defer m.close()
	return m.q.All(value)
}

func (m *mgodb) One(value interface{}) error {
	defer m.close()
	return m.q.One(value)
}

func (m *mgodb) Sort(sort ...string) Query {
	m.q = m.q.Sort(sort...)
	return m
}
func (m *mgodb) Select(selector interface{}) Query {
	m.q = m.q.Select(selector)
	return m
}

func (m *mgodb) Limit(limit int) Query {
	m.q = m.q.Limit(limit)
	return m
}

func (m *mgodb) Skip(skip int) Query {
	m.q = m.q.Skip(skip)
	return m
}

func (m *mgodb) Count() (int, error) {
	defer m.close()
	return m.q.Count()
}
