package parser

import (
	"github.com/olivere/elastic"
	"strings"
)

type MyDslQueryListener struct {
	*BaseDslQueryListener
	Stack []*elastic.BoolQuery
}

func (ql *MyDslQueryListener) ExitCompareExp(c *CompareExpContext) {
	boolQuery := elastic.NewBoolQuery()

	attrName := c.GetPropertyName().GetText()
	attrValue := strings.Trim(c.GetPropertyValue().GetText(), `\"`)
	// Based on operator type we build different queries, default is terms query(=)
	termsQuery := elastic.NewTermQuery(attrName, attrValue)
	boolQuery.Must(termsQuery)
	ql.Stack = append(ql.Stack, boolQuery)
}

func (ql *MyDslQueryListener) ExitAndLogicalExp(c *AndLogicalExpContext) {
	size := len(ql.Stack)
	right := ql.Stack[size-1]
	left := ql.Stack[size-2]
	ql.Stack = ql.Stack[:size-2] // Pop last two elements
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(right)
	boolQuery.Must(left)
	ql.Stack = append(ql.Stack, boolQuery)
}

func (ql *MyDslQueryListener) ExitOrLogicalExp(c *OrLogicalExpContext) {
	size := len(ql.Stack)
	right := ql.Stack[size-1]
	left := ql.Stack[size-2]
	ql.Stack = ql.Stack[:size-2] // Pop last two elements
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Should(right)
	boolQuery.Should(left)
	ql.Stack = append(ql.Stack, boolQuery)
}
