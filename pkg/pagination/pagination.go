package pagination

type Pagination struct {
	Offset int
	Limit  int
}

func New(page int, limit int) Pagination {
	return Pagination{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}
}
