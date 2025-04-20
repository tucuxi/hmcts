package persistence

import "github.com/tucuxi/hmcts/pkg/domain"

type Repository struct {
	cases map[int]domain.Case
}

func NewRepository() *Repository {
	return &Repository{
		cases: make(map[int]domain.Case),
	}
}

func (r *Repository) Add(c domain.Case) bool {
	if _, exists := r.cases[c.Id]; exists {
		return false
	}
	r.cases[c.Id] = c
	return true
}

func (r *Repository) Get(id int) (domain.Case, bool) {
	c, exists := r.cases[id]
	return c, exists
}

func (r *Repository) List() []domain.Case {
	res := []domain.Case{}
	for _, c := range r.cases {
		res = append(res, c)
	}
	return res
}

func (r *Repository) Remove(id int) bool {
	if _, exists := r.cases[id]; !exists {
		return false
	}
	delete(r.cases, id)
	return true
}
