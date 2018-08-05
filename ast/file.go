package ast

type File struct {
	Declarations []Declaration
	Comments     []Comment
}

type Comment struct {
	Text string
}
