// Code generated from DslQuery.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // DslQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDslQueryListener is a complete listener for a parse tree produced by DslQueryParser.
type BaseDslQueryListener struct{}

var _ DslQueryListener = &BaseDslQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDslQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDslQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDslQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDslQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDslQueryListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDslQueryListener) ExitStart(ctx *StartContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseDslQueryListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseDslQueryListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterAndLogicalExp is called when production andLogicalExp is entered.
func (s *BaseDslQueryListener) EnterAndLogicalExp(ctx *AndLogicalExpContext) {}

// ExitAndLogicalExp is called when production andLogicalExp is exited.
func (s *BaseDslQueryListener) ExitAndLogicalExp(ctx *AndLogicalExpContext) {}

// EnterBracketExp is called when production bracketExp is entered.
func (s *BaseDslQueryListener) EnterBracketExp(ctx *BracketExpContext) {}

// ExitBracketExp is called when production bracketExp is exited.
func (s *BaseDslQueryListener) ExitBracketExp(ctx *BracketExpContext) {}

// EnterOrLogicalExp is called when production orLogicalExp is entered.
func (s *BaseDslQueryListener) EnterOrLogicalExp(ctx *OrLogicalExpContext) {}

// ExitOrLogicalExp is called when production orLogicalExp is exited.
func (s *BaseDslQueryListener) ExitOrLogicalExp(ctx *OrLogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseDslQueryListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseDslQueryListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseDslQueryListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseDslQueryListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseDslQueryListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseDslQueryListener) ExitNull(ctx *NullContext) {}

// EnterString is called when production string is entered.
func (s *BaseDslQueryListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseDslQueryListener) ExitString(ctx *StringContext) {}

// EnterDouble is called when production double is entered.
func (s *BaseDslQueryListener) EnterDouble(ctx *DoubleContext) {}

// ExitDouble is called when production double is exited.
func (s *BaseDslQueryListener) ExitDouble(ctx *DoubleContext) {}

// EnterLong is called when production long is entered.
func (s *BaseDslQueryListener) EnterLong(ctx *LongContext) {}

// ExitLong is called when production long is exited.
func (s *BaseDslQueryListener) ExitLong(ctx *LongContext) {}
