package typedef

/*
	FieldTag映射对应的json值到struct字段
*/
type Teacher struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Grade    Grade     `json:"grade"`
	Students []Student `json:"students"`
}

type Grade struct {
	School	string	`json:"school"`
	Grade	int		`json:"grade"`
}

type Student struct {
	Name	string	`json:"name"`
	Age		int		`json:"age"`
	Sex		uint8	`json:"sex"`
}

var JsonStr = `{
	"name": "pyl",
	"age": 33,
	"grade": {
		"school": "bit",
		"grade": 1
	},
	"students": [
		{"name": "zy", "age": 34, "sex": 0},
		{"name": "py", "age": 39, "sex": 1}
	]
}`