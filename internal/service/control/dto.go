package control

type Create struct {
	StudentId int    `json:"student_id"`
	Status    string `json:"status"`
	UserId    int    `json:"user_id"`
}
