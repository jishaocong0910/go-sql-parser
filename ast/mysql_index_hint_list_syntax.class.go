package ast

type MySqlIndexHintListSyntax struct {
	*Syntax__
	*ListSyntax__[*MySqlIndexHintSyntax]
}

func (this *MySqlIndexHintListSyntax) accept(Visitor_) {}

func NewMySqlIndexHintListSyntax() *MySqlIndexHintListSyntax {
	s := &MySqlIndexHintListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*MySqlIndexHintSyntax](s)
	s.separator = ""
	return s
}
