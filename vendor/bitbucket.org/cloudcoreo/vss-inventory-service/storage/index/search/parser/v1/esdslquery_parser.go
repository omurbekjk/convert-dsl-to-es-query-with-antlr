// Code generated from EsDslQuery.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // EsDslQuery

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 53, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 22, 10, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 30, 10, 3, 12, 3, 14, 3, 33, 11, 3, 3, 4, 3,
	4, 3, 4, 5, 4, 38, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 45, 10, 5,
	3, 5, 3, 5, 5, 5, 49, 10, 5, 5, 5, 51, 10, 5, 3, 5, 2, 3, 4, 6, 2, 4, 6,
	8, 2, 2, 2, 58, 2, 10, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2,
	8, 50, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 3, 3, 2, 2, 2, 12, 13, 8, 3,
	1, 2, 13, 14, 7, 26, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 7, 27, 2, 2, 16,
	22, 3, 2, 2, 2, 17, 18, 5, 6, 4, 2, 18, 19, 7, 7, 2, 2, 19, 20, 5, 8, 5,
	2, 20, 22, 3, 2, 2, 2, 21, 12, 3, 2, 2, 2, 21, 17, 3, 2, 2, 2, 22, 31,
	3, 2, 2, 2, 23, 24, 12, 5, 2, 2, 24, 25, 7, 8, 2, 2, 25, 30, 5, 4, 3, 6,
	26, 27, 12, 4, 2, 2, 27, 28, 7, 9, 2, 2, 28, 30, 5, 4, 3, 5, 29, 23, 3,
	2, 2, 2, 29, 26, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31,
	32, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 37, 7, 28, 2,
	2, 35, 36, 7, 3, 2, 2, 36, 38, 7, 28, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38,
	3, 2, 2, 2, 38, 7, 3, 2, 2, 2, 39, 51, 7, 5, 2, 2, 40, 51, 7, 6, 2, 2,
	41, 51, 7, 29, 2, 2, 42, 51, 7, 30, 2, 2, 43, 45, 7, 4, 2, 2, 44, 43, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 48, 7, 31, 2, 2, 47,
	49, 7, 32, 2, 2, 48, 47, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2,
	2, 2, 50, 39, 3, 2, 2, 2, 50, 40, 3, 2, 2, 2, 50, 41, 3, 2, 2, 2, 50, 42,
	3, 2, 2, 2, 50, 44, 3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 9, 21, 29, 31, 37, 44,
	48, 50,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'-'", "", "'null'", "", "", "", "", "", "", "", "", "", "",
	"", "'ne'", "'gt'", "'lt'", "'ge'", "'le'", "'co'", "'sw'", "'ew'", "'('",
	"')'",
}
var symbolicNames = []string{
	"", "", "", "BOOLEAN", "NULL", "COMPARISON_OPERATOR", "AND", "OR", "O",
	"R", "A", "N", "D", "EQ", "E", "Q", "NE", "GT", "LT", "GE", "LE", "CO",
	"SW", "EW", "BR_OPEN", "BR_CLOSE", "ATTRNAME", "STRING", "DOUBLE", "INT",
	"EXP", "WS",
}

var ruleNames = []string{
	"start", "query", "attrPath", "attrValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type EsDslQueryParser struct {
	*antlr.BaseParser
}

func NewEsDslQueryParser(input antlr.TokenStream) *EsDslQueryParser {
	this := new(EsDslQueryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "EsDslQuery.g4"

	return this
}

// EsDslQueryParser tokens.
const (
	EsDslQueryParserEOF                 = antlr.TokenEOF
	EsDslQueryParserT__0                = 1
	EsDslQueryParserT__1                = 2
	EsDslQueryParserBOOLEAN             = 3
	EsDslQueryParserNULL                = 4
	EsDslQueryParserCOMPARISON_OPERATOR = 5
	EsDslQueryParserAND                 = 6
	EsDslQueryParserOR                  = 7
	EsDslQueryParserO                   = 8
	EsDslQueryParserR                   = 9
	EsDslQueryParserA                   = 10
	EsDslQueryParserN                   = 11
	EsDslQueryParserD                   = 12
	EsDslQueryParserEQ                  = 13
	EsDslQueryParserE                   = 14
	EsDslQueryParserQ                   = 15
	EsDslQueryParserNE                  = 16
	EsDslQueryParserGT                  = 17
	EsDslQueryParserLT                  = 18
	EsDslQueryParserGE                  = 19
	EsDslQueryParserLE                  = 20
	EsDslQueryParserCO                  = 21
	EsDslQueryParserSW                  = 22
	EsDslQueryParserEW                  = 23
	EsDslQueryParserBR_OPEN             = 24
	EsDslQueryParserBR_CLOSE            = 25
	EsDslQueryParserATTRNAME            = 26
	EsDslQueryParserSTRING              = 27
	EsDslQueryParserDOUBLE              = 28
	EsDslQueryParserINT                 = 29
	EsDslQueryParserEXP                 = 30
	EsDslQueryParserWS                  = 31
)

// EsDslQueryParser rules.
const (
	EsDslQueryParserRULE_start     = 0
	EsDslQueryParserRULE_query     = 1
	EsDslQueryParserRULE_attrPath  = 2
	EsDslQueryParserRULE_attrValue = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EsDslQueryParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EsDslQueryParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *EsDslQueryParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EsDslQueryParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.query(0)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EsDslQueryParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EsDslQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	*QueryContext
	propertyName  IAttrPathContext
	op            antlr.Token
	propertyValue IAttrValueContext
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetPropertyName() IAttrPathContext { return s.propertyName }

func (s *CompareExpContext) GetPropertyValue() IAttrValueContext { return s.propertyValue }

func (s *CompareExpContext) SetPropertyName(v IAttrPathContext) { s.propertyName = v }

func (s *CompareExpContext) SetPropertyValue(v IAttrValueContext) { s.propertyValue = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) COMPARISON_OPERATOR() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserCOMPARISON_OPERATOR, 0)
}

func (s *CompareExpContext) AttrValue() IAttrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrValueContext)
}

func (s *CompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterCompareExp(s)
	}
}

func (s *CompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitCompareExp(s)
	}
}

type AndLogicalExpContext struct {
	*QueryContext
	leftQuery  IQueryContext
	op         antlr.Token
	rightQuery IQueryContext
}

func NewAndLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndLogicalExpContext {
	var p = new(AndLogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *AndLogicalExpContext) GetOp() antlr.Token { return s.op }

func (s *AndLogicalExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndLogicalExpContext) GetLeftQuery() IQueryContext { return s.leftQuery }

func (s *AndLogicalExpContext) GetRightQuery() IQueryContext { return s.rightQuery }

func (s *AndLogicalExpContext) SetLeftQuery(v IQueryContext) { s.leftQuery = v }

func (s *AndLogicalExpContext) SetRightQuery(v IQueryContext) { s.rightQuery = v }

func (s *AndLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *AndLogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *AndLogicalExpContext) AND() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserAND, 0)
}

func (s *AndLogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterAndLogicalExp(s)
	}
}

func (s *AndLogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitAndLogicalExp(s)
	}
}

type BracketExpContext struct {
	*QueryContext
	leftBracket  antlr.Token
	rightBracket antlr.Token
}

func NewBracketExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExpContext {
	var p = new(BracketExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *BracketExpContext) GetLeftBracket() antlr.Token { return s.leftBracket }

func (s *BracketExpContext) GetRightBracket() antlr.Token { return s.rightBracket }

func (s *BracketExpContext) SetLeftBracket(v antlr.Token) { s.leftBracket = v }

func (s *BracketExpContext) SetRightBracket(v antlr.Token) { s.rightBracket = v }

func (s *BracketExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *BracketExpContext) BR_OPEN() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserBR_OPEN, 0)
}

func (s *BracketExpContext) BR_CLOSE() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserBR_CLOSE, 0)
}

func (s *BracketExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterBracketExp(s)
	}
}

func (s *BracketExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitBracketExp(s)
	}
}

type OrLogicalExpContext struct {
	*QueryContext
	leftQuery  IQueryContext
	op         antlr.Token
	rightQuery IQueryContext
}

func NewOrLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrLogicalExpContext {
	var p = new(OrLogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *OrLogicalExpContext) GetOp() antlr.Token { return s.op }

func (s *OrLogicalExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrLogicalExpContext) GetLeftQuery() IQueryContext { return s.leftQuery }

func (s *OrLogicalExpContext) GetRightQuery() IQueryContext { return s.rightQuery }

func (s *OrLogicalExpContext) SetLeftQuery(v IQueryContext) { s.leftQuery = v }

func (s *OrLogicalExpContext) SetRightQuery(v IQueryContext) { s.rightQuery = v }

func (s *OrLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *OrLogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *OrLogicalExpContext) OR() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserOR, 0)
}

func (s *OrLogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterOrLogicalExp(s)
	}
}

func (s *OrLogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitOrLogicalExp(s)
	}
}

func (p *EsDslQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *EsDslQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, EsDslQueryParserRULE_query, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EsDslQueryParserBR_OPEN:
		localctx = NewBracketExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)

			var _m = p.Match(EsDslQueryParserBR_OPEN)

			localctx.(*BracketExpContext).leftBracket = _m
		}
		{
			p.SetState(12)
			p.query(0)
		}
		{
			p.SetState(13)

			var _m = p.Match(EsDslQueryParserBR_CLOSE)

			localctx.(*BracketExpContext).rightBracket = _m
		}

	case EsDslQueryParserATTRNAME:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)

			var _x = p.AttrPath()

			localctx.(*CompareExpContext).propertyName = _x
		}
		{
			p.SetState(16)

			var _m = p.Match(EsDslQueryParserCOMPARISON_OPERATOR)

			localctx.(*CompareExpContext).op = _m
		}
		{
			p.SetState(17)

			var _x = p.AttrValue()

			localctx.(*CompareExpContext).propertyValue = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				localctx.(*AndLogicalExpContext).leftQuery = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EsDslQueryParserRULE_query)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(22)

					var _m = p.Match(EsDslQueryParserAND)

					localctx.(*AndLogicalExpContext).op = _m
				}
				{
					p.SetState(23)

					var _x = p.query(4)

					localctx.(*AndLogicalExpContext).rightQuery = _x
				}

			case 2:
				localctx = NewOrLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				localctx.(*OrLogicalExpContext).leftQuery = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EsDslQueryParserRULE_query)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(25)

					var _m = p.Match(EsDslQueryParserOR)

					localctx.(*OrLogicalExpContext).op = _m
				}
				{
					p.SetState(26)

					var _x = p.query(3)

					localctx.(*OrLogicalExpContext).rightQuery = _x
				}

			}

		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EsDslQueryParserRULE_attrPath
	return p
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EsDslQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) AllATTRNAME() []antlr.TerminalNode {
	return s.GetTokens(EsDslQueryParserATTRNAME)
}

func (s *AttrPathContext) ATTRNAME(i int) antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserATTRNAME, i)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterAttrPath(s)
	}
}

func (s *AttrPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitAttrPath(s)
	}
}

func (p *EsDslQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EsDslQueryParserRULE_attrPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(EsDslQueryParserATTRNAME)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EsDslQueryParserT__0 {
		{
			p.SetState(33)
			p.Match(EsDslQueryParserT__0)
		}
		{
			p.SetState(34)
			p.Match(EsDslQueryParserATTRNAME)
		}

	}

	return localctx
}

// IAttrValueContext is an interface to support dynamic dispatch.
type IAttrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrValueContext differentiates from other interfaces.
	IsAttrValueContext()
}

type AttrValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrValueContext() *AttrValueContext {
	var p = new(AttrValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EsDslQueryParserRULE_attrValue
	return p
}

func (*AttrValueContext) IsAttrValueContext() {}

func NewAttrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrValueContext {
	var p = new(AttrValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EsDslQueryParserRULE_attrValue

	return p
}

func (s *AttrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrValueContext) CopyFrom(ctx *AttrValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AttrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanContext struct {
	*AttrValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.AttrValueContext = NewEmptyAttrValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AttrValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitBoolean(s)
	}
}

type NullContext struct {
	*AttrValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.AttrValueContext = NewEmptyAttrValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AttrValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserNULL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitNull(s)
	}
}

type StringContext struct {
	*AttrValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.AttrValueContext = NewEmptyAttrValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AttrValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitString(s)
	}
}

type DoubleContext struct {
	*AttrValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	p.AttrValueContext = NewEmptyAttrValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AttrValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserDOUBLE, 0)
}

func (s *DoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterDouble(s)
	}
}

func (s *DoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitDouble(s)
	}
}

type LongContext struct {
	*AttrValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	p.AttrValueContext = NewEmptyAttrValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AttrValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserINT, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(EsDslQueryParserEXP, 0)
}

func (s *LongContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.EnterLong(s)
	}
}

func (s *LongContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EsDslQueryListener); ok {
		listenerT.ExitLong(s)
	}
}

func (p *EsDslQueryParser) AttrValue() (localctx IAttrValueContext) {
	localctx = NewAttrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EsDslQueryParserRULE_attrValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EsDslQueryParserBOOLEAN:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(EsDslQueryParserBOOLEAN)
		}

	case EsDslQueryParserNULL:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(EsDslQueryParserNULL)
		}

	case EsDslQueryParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(EsDslQueryParserSTRING)
		}

	case EsDslQueryParserDOUBLE:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Match(EsDslQueryParserDOUBLE)
		}

	case EsDslQueryParserT__1, EsDslQueryParserINT:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EsDslQueryParserT__1 {
			{
				p.SetState(41)
				p.Match(EsDslQueryParserT__1)
			}

		}
		{
			p.SetState(44)
			p.Match(EsDslQueryParserINT)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(45)
				p.Match(EsDslQueryParserEXP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *EsDslQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EsDslQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
