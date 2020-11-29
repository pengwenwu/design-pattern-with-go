package factory_method

// 被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct {

}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct {

}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}