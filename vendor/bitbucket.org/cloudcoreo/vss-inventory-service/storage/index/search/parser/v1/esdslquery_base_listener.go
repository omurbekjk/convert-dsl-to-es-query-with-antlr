// Code generated from EsDslQuery.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // EsDslQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEsDslQueryListener is a complete listener for a parse tree produced by EsDslQueryParser.
type BaseEsDslQueryListener struct{}

var _ EsDslQueryListener = &BaseEsDslQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEsDslQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEsDslQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEsDslQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEsDslQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseEsDslQueryListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseEsDslQueryListener) ExitStart(ctx *StartContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseEsDslQueryListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseEsDslQueryListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterAndLogicalExp is called when production andLogicalExp is entered.
func (s *BaseEsDslQueryListener) EnterAndLogicalExp(ctx *AndLogicalExpContext) {}

// ExitAndLogicalExp is called when production andLogicalExp is exited.
func (s *BaseEsDslQueryListener) ExitAndLogicalExp(ctx *AndLogicalExpContext) {}

// EnterBracketExp is called when production bracketExp is entered.
func (s *BaseEsDslQueryListener) EnterBracketExp(ctx *BracketExpContext) {}

// ExitBracketExp is called when production bracketExp is exited.
func (s *BaseEsDslQueryListener) ExitBracketExp(ctx *BracketExpContext) {}

// EnterOrLogicalExp is called when production orLogicalExp is entered.
func (s *BaseEsDslQueryListener) EnterOrLogicalExp(ctx *OrLogicalExpContext) {}

// ExitOrLogicalExp is called when production orLogicalExp is exited.
func (s *BaseEsDslQueryListener) ExitOrLogicalExp(ctx *OrLogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseEsDslQueryListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseEsDslQueryListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseEsDslQueryListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseEsDslQueryListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseEsDslQueryListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseEsDslQueryListener) ExitNull(ctx *NullContext) {}

// EnterString is called when production string is entered.
func (s *BaseEsDslQueryListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseEsDslQueryListener) ExitString(ctx *StringContext) {}

// EnterDouble is called when production double is entered.
func (s *BaseEsDslQueryListener) EnterDouble(ctx *DoubleContext) {}

// ExitDouble is called when production double is exited.
func (s *BaseEsDslQueryListener) ExitDouble(ctx *DoubleContext) {}

// EnterLong is called when production long is entered.
func (s *BaseEsDslQueryListener) EnterLong(ctx *LongContext) {}

// ExitLong is called when production long is exited.
func (s *BaseEsDslQueryListener) ExitLong(ctx *LongContext) {}
