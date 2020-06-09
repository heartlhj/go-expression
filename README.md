# expression  EL表达式



## Install

```go
go get github.com/heartlhj/go-expression

```

## 功能

 - 字符串的提取和比较

```go
    context := spel.StandardEvaluationContext{}
    m := make(map[string]interface{})
    m["name"] = "lisi"
    m["age"] = 18
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#name=='lisi'"
    //expressionString := "#name" //返回lisi
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

 - 数字比较，支持int和 float64

```go
    context := spel.StandardEvaluationContext{}
    m := make(map[string]interface{})
    m["name"] = "lisi"
    m["age"] = 18
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#age>=10"
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

float64比较

```go
    context := spel.StandardEvaluationContext{}
    m := make(map[string]interface{})
    var ageFloat float64
    ageFloat = 10
    m["num"] = ageFloat
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#num>=9f"
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

 - 与（&&）或 （||）操作
 

```go
    context := spel.StandardEvaluationContext{}
    m := make(map[string]interface{})
    m["name"] = "lisi"
    m["age"] = 18
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#name=='lisi' && #age>=3"
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

 - 复合结构，迭代提取

```go
    context := spel.StandardEvaluationContext{}
    context.AddPropertyAccessor(spel.MapAccessor{})
    m := make(map[string]interface{})
    m["name"] = "lisi"
    m["age"] = 18
    m1 := make(map[string]interface{})
    m2 := make(map[string]interface{})
    m2["num"] = 12
    m1["code"] = m2
    m["order"] = m1
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#order.code.num==12"
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

 - 数组和切片提取

```go
    context := spel.StandardEvaluationContext{}
    context.AddPropertyAccessor(spel.MapAccessor{})
    m := make(map[string]interface{})
    m["name"] = "lisi"
    m["age"] = 18
    m1 := make(map[string]interface{})
    //切片
    //orders := make([]Order, 2)
    //数组
    orders := [2]Order{}
    orders[0] = Order{name: "lisi", age: 12}
    orders[1] = Order{name: "wang", age: 24}
    m1["code"] = orders
    m["order"] = m1
    context.SetVariables(m)
    parser := SpelExpressionParser{}
    expressionString := "#order.code[0].name=='lisi'"
    valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
```

