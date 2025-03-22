package payload

type (
	Pagination struct {
		Page  int  `query:"page" validate:"gte=0"`
		Limit int  `query:"limit" validate:"gte=0,lte=1000"`
		All   bool `query:"all"`
	}

	Params struct {
		OrderBy  string `query:"order_by"`
		SortedBy string `query:"sorted_by"`
	}
)
