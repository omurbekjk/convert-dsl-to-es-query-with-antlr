grammar DslQuery;

options {
language = Go;
}

start: query;

query
   : leftBracket = '(' query rightBracket = ')'                             #bracketExp
   | leftQuery=query op=AND rightQuery=query                                #andLogicalExp
   | leftQuery=query op=OR rightQuery=query                                 #orLogicalExp
   | propertyName=attrPath op=COMPARISON_OPERATOR propertyValue=attrValue   #compareExp
   ;

attrPath
   : ATTRNAME ('.' ATTRNAME)?
   ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

attrValue
   : BOOLEAN           #boolean
   | NULL              #null
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

BOOLEAN
   : 'true' | 'false'
   ;

NULL
   : 'null'
   ;

COMPARISON_OPERATOR: EQ;

AND: A N D;
OR: O R;

O  : 'O'|'o';
R  : 'R'|'r';
A  : 'A'|'a';
N  : 'N'|'n';
D  : 'D'|'d';
EQ : '='| E Q;
E  : 'E'|'e';
Q  : 'Q'|'q';
NE : 'ne' ;
GT : 'gt' ;
LT : 'lt' ;
GE : 'ge' ;
LE : 'le' ;
CO : 'co' ;
SW : 'sw' ;
EW : 'ew' ;
BR_OPEN: '(';
BR_CLOSE: ')';

ATTRNAME
   : ALPHA ATTR_NAME_CHAR*
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;
// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;
// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

WS: [\t ]+ -> skip;
