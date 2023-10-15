package todo

type Todos []*Todo

func (s Todos) LastID() uint64 {
	var id uint64
	for _, t := range s {
		if t.ID > id {
			id = t.ID
		}
	}
	return id
}

func (s Todos) Find(id uint64) *Todo {
	for _, t := range s {
		if t.ID == id {
			return t
		}
	}
	return nil
}

func (s Todos) Select(f func(*Todo) bool) Todos {
	var res Todos
	for _, t := range s {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}

func (s Todos) RemoveByID(id uint64) Todos {
	return s.Select(func(t *Todo) bool { return t.ID != id })
}
