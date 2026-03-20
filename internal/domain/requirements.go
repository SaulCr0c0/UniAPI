package domain

type Requirements struct {
	CodSubject    int `json:"cod_subject" db:"cod_subject"`
	CodSubjectReq int `json:"cod_subject_req" db:"cod_subject_req"`
}
