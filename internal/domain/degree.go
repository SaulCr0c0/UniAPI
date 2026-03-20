package domain

type Degree struct {
	CodDegree  int    `json:"cod_degree" db:"cod_degree"`
	NameDegree string `json:"name_degree" db:"name_degree"`
}
