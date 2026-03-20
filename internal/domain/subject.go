package domain

type Subject struct {
	CodSubject   int            `json:"cod_subject" db:"cod_subject"`
	Name         string         `json:"name" db:"name"`
	Semester     rune           `json:"semester" db:"semester"`
	CodDegree    int            `json:"cod_degree" db:"cod_degree"`
	Requirements []Requirements `json:"requirements" db:"-"`
}
