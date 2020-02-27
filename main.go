package main

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/omurbekjk/convert-dsl-to-es-query-with-antlr/parser"
)

func main() {
	fmt.Println("Starting here")
	query := "price=2000 OR model=\"hyundai\" AND (color=\"red\" OR color=\"blue\")"
	stream := antlr.NewInputStream(query)
	lexer := parser.NewDslQueryLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	dslParser := parser.NewDslQueryParser(tokenStream)
	tree := dslParser.Start()

	listener := &parser.MyDslQueryListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	esQuery := listener.Stack[0]

	src, err := esQuery.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
}

/**		Generated es query
{
  "bool": {
    "should": [
      {
        "bool": {
          "must": [
            {
              "bool": {
                "should": [
                  {
                    "bool": {
                      "must": {
                        "term": {
                          "color": "blue"
                        }
                      }
                    }
                  },
                  {
                    "bool": {
                      "must": {
                        "term": {
                          "color": "red"
                        }
                      }
                    }
                  }
                ]
              }
            },
            {
              "bool": {
                "must": {
                  "term": {
                    "model": "hyundai"
                  }
                }
              }
            }
          ]
        }
      },
      {
        "bool": {
          "must": {
            "term": {
              "price": "2000"
            }
          }
        }
      }
    ]
  }
}

*/
