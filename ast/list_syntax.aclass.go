package ast

// 列表
type ListSyntax_[T Syntax_] interface {
	ListSyntax_() *ListSyntax__[T]
	Syntax_

	Add(T)
	Get(i int) T
}

type ListSyntax__[T Syntax_] struct {
	I         ListSyntax_[T]
	separator string
	elements  []T
}

func (this *ListSyntax__[T]) ListSyntax_() *ListSyntax__[T] {
	return this
}

func (this *ListSyntax__[T]) accept(v_ Visitor_) {
	for _, i := range this.elements {
		v_.visitor_().visit(i)
	}
}

func (this *ListSyntax__[T]) writeSql(builder *sqlBuilder) {
	if len(this.elements) > 0 {
		builder.writeSyntax(this.elements[0])
		for i := 1; i < len(this.elements); i++ {
			builder.writeStr(this.separator)
			builder.writeSpaceOrLf(this.I.(Syntax_), false)
			builder.writeSyntax(this.elements[i])
		}
	}
}

func (this *ListSyntax__[T]) Add(t T) {
	this.elements = append(this.elements, t)
}

func (this *ListSyntax__[T]) Get(i int) T {
	var t T
	if i < this.Len() {
		t = this.elements[i]
	}
	return t
}

func (this *ListSyntax__[T]) Len() int {
	return len(this.elements)
}

func ExtendListSyntax[T Syntax_](i ListSyntax_[T]) *ListSyntax__[T] {
	return &ListSyntax__[T]{I: i, separator: ","}
}
