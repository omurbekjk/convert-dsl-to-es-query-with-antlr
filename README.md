# convert-dsl-to-es-query-with-antlr
Convert custom sql like dsl query to es query

Basic example 1:

input: `"price=2000"`

Output: 
```
{
  "bool": {
    "must": {
      "term": {
        "price": "2000"
      }
    }
  }
}
```

Basic example 2:

input: `"price=2000 AND year=2021"`

Output: 

```
{
  "bool": {
    "must": [
      {
        "bool": {
          "must": {
            "term": {
              "year": "2021"
            }
          }
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

```

Complex example 1:

input: `"price=2000 AND year=2021 OR color=\"red\" OR (model=\"Tesla\")"`

Output: 

```
{
  "bool": {
    "should": [
      {
        "bool": {
          "must": {
            "term": {
              "model": "Tesla"
            }
          }
        }
      },
      {
        "bool": {
          "should": [
            {
              "bool": {
                "must": {
                  "term": {
                    "color": "red"
                  }
                }
              }
            },
            {
              "bool": {
                "must": [
                  {
                    "bool": {
                      "must": {
                        "term": {
                          "year": "2021"
                        }
                      }
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
          ]
        }
      }
    ]
  }
}

```

Any questions/bugs feel free to open an issue.

