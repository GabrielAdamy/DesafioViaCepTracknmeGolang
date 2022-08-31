package dtos

type PageBody struct {
	Page          int64 `json:"page"`
	Limit         int64 `json:"limit"`
	Elements      int64 `json:"elements"`
	TotalElements int64 `json:"totalElements"`
	TotalPages    int64 `json:"totalPages"`
}

func NewPageBody(elements int, page int64, limit int64, count int64) PageBody {
	return PageBody{
		Elements:      int64(elements),
		Page:          page,
		Limit:         limit,
		TotalElements: count,
		TotalPages:    (count / limit) + 1,
	}
}

type EmployeePage struct {
	Content []*EmployeeDTO `json:"content"`
	PageBody
}

func NewEmployeePageWithResource(resources []*EmployeeDTO) *EmployeePage {
	return &EmployeePage{
		Content:  resources,
		PageBody: NewPageBody(len(resources), 0, int64(len(resources)), int64(len(resources))),
	}
}

func NewContractPage(resources []*EmployeeDTO, count int64, page int64, limit int64) *EmployeePage {
	return &EmployeePage{
		Content:  resources,
		PageBody: NewPageBody(len(resources), page, limit, count),
	}
}
