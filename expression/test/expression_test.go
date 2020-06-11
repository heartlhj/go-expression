package test

import (
	"fmt"
	. "github.com/heartlhj/go-expression/expression"
	. "github.com/heartlhj/go-expression/expression/spel"
	"testing"
)

type Order struct {
	name string
	age  int
}

var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

//func init() {
//	context = StandardEvaluationContext{}
//}
//测试字符串下标
func TestStringIndex(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	expressionString := "#name[2]=='s'"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试数组
func TestIndex(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m1 := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	//切片
	//orders := make([]Order, 2)
	//数组
	orders := [2]Order{}
	orders[0] = Order{name: "lisi", age: 12}
	orders[1] = Order{name: "wang", age: 24}
	m1["code"] = orders
	m["order"] = m1
	context.SetVariables(m)
	expressionString := "#order.code[0].name=='lisi'"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试复合属性，对象里的字段
func TestCompound(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m["name"] = "lisi"
	m["age"] = 18
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m2["num"] = 12
	m1["code"] = m2
	m["order"] = m1
	context.SetVariables(m)
	expressionString := "#order.code.num==12"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试字符相等
func TestEQ(t *testing.T) {
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	expressionString := "#name=='lisi'"
	//expressionString := "#name" //返回lisi
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试与操作
func TestAnd(t *testing.T) {
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	expressionString := "#name=='lisi' && #age>=3"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试大于等于
func TestGT(t *testing.T) {
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	expressionString := "#age>=10"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试float类型 #num>=9f 会转为float64 #num>=9转为int
func TestFloat(t *testing.T) {
	var ageFloat float64
	ageFloat = 10
	m["num"] = ageFloat
	context.SetVariables(m)
	expressionString := "#num>=9f"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}
