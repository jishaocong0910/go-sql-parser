package ast

// 列表
type I_ListSyntax[T I_Syntax] interface {
	I_Syntax
	M_ListSyntax_() *M_ListSyntax[T]
	Add(T)
	Get(i int) T
}

type M_ListSyntax[T I_Syntax] struct {
	I         I_ListSyntax[T]
	separator string
	elements  []T
}

func (this *M_ListSyntax[T]) M_ListSyntax_() *M_ListSyntax[T] {
	return this
}

func (this *M_ListSyntax[T]) accept(iv I_Visitor) {
	for _, i := range this.elements {
		iv.m_Visitor_().visit(i)
	}
}

func (this *M_ListSyntax[T]) writeSql(builder *sqlBuilder) {
	if len(this.elements) > 0 {
		builder.writeSyntax(this.elements[0])
		for i := 1; i < len(this.elements); i++ {
			builder.writeStr(this.separator)
			builder.writeSpaceOrLf(this.I.(I_Syntax), false)
			builder.writeSyntax(this.elements[i])
		}
	}
}

func (this *M_ListSyntax[T]) Add(t T) {
	this.elements = append(this.elements, t)
}

func (this *M_ListSyntax[T]) Get(i int) T {
	var t T
	if i < this.Len() {
		t = this.elements[i]
	}
	return t
}

func (this *M_ListSyntax[T]) Len() int {
	return len(this.elements)
}

func ExtendListSyntax[T I_Syntax](i I_ListSyntax[T]) *M_ListSyntax[T] {
	return &M_ListSyntax[T]{I: i, separator: ","}
}
