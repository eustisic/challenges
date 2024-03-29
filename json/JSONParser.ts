import { json } from "stream/consumers"
import { TokenType, Token, Tokens } from "./types"

const JSON_WHITESPACE = [' ', '\t', '\b', '\n', '\r']

// const JSON_SYNTAX = new RegExp(/[\w\d.,{}\[\]\-\"\':]/)

const JSON_SYNTAX = ["]", "]", "{", "}", ",", ":"]

const NUMBER_CHARACTERS = new RegExp(/[\d.\-e]/)

const TRUE_LENGTH = 4
const FALSE_LENGTH = 5
const UNDEFINED_LENGTH = 9

export class JSONParser {

    static InvalidJSON(char: any) {
      throw new Error(`Unexpected character: ${char}`)
    }

    static parse(jsonString: string) {
        let index = 0
        const tokens = JSONLexer.lex(jsonString)

        function parseArray() {
          const array: any[] = []
          let token = tokens[index]

          if (token === Tokens.BracketClose) {
            index++
            return array
          }

          while (true) {
            const json = parseTokens()
            array.push(json)

            token = token[index]

            if (token == Tokens.BracketClose) {
              index++
              break
            } else if (token !== Tokens.Comma) {
              JSONParser.InvalidJSON(token)
            } else {
              index++
            }
          }

          return array
        }

        function parseObject() {
          const obj: Record<string, any> = {}
          let token = tokens[index]

          if (token === Tokens.BraceClose) {
            index++
            return obj
          }

          while (true) {
            const key = tokens[index]

            if (typeof key === 'string') {
              index++
            } else {
              JSONParser.InvalidJSON(token)
            }

            if (tokens[index] !== Tokens.Colon) {
              JSONParser.InvalidJSON(token)
            }

            // move past colon and set key to value
            index++
            obj[key] = parseTokens()

            // should be at bracket closed
            token = tokens[index]

            if (token === Tokens.BraceClose) {
              return obj
            } else if (token !== Tokens.Comma) {
              JSONParser.InvalidJSON(token)
            }

            index++
          }
        }

        function parseTokens() {
          const token = tokens[index]
          let json

          if (token === Tokens.BraceOpen) {
            index++
            json = parseObject()
          } else if (token === Tokens.BracketOpen) {
            index++
            json = parseArray()
          } else {
            // if not brace or bracket it is a value
            json = token
            index++
          }
         
          return json
        }
            
        return parseTokens()
    }
}

export class JSONLexer {
  static lexString(str: string, tokens: any[]) {
    let json_string = ''

    if (str[0] === '"') {
      str = str.slice(1, str.length)
    } else {
      return str
    } 

    for (let char of str) {
      if (char === '"') {
        break
      } else {
        json_string += char
      }
    }

    if (json_string.length) {
      tokens.push(json_string)
      str = str.slice(json_string.length+1, str.length)
    }

    return str
  }

  static lexNumber(str: string, tokens: any[]) {
    let json_string = ''

    for (let char of str) {
      if (NUMBER_CHARACTERS.test(char)) {
       json_string += char
      } else {
        break
      }
    }

    if (json_string.length) {
      const json_number = Number(json_string)
      tokens.push(json_number)
      return str.slice(json_string.length, str.length)
    }
    
    return str
  }

  static lexBoolNull(str: string, tokens: any[]) {
    let true_str = str.slice(0, TRUE_LENGTH)
    let false_str = str.slice(0, FALSE_LENGTH)
    let undefined_str = str.slice(0, UNDEFINED_LENGTH)

    if (true_str === 'true') {
      tokens.push(true)
      return str.slice(TRUE_LENGTH, str.length)
    } else if (true_str === 'null') {
      tokens.push(null)
      return str.slice(TRUE_LENGTH, str.length)
    } else if (false_str === 'false') {
      tokens.push(false)
      return str.slice(FALSE_LENGTH, str.length)
    } else if (undefined_str === 'undefined') {
      tokens.push(undefined)
      return str.slice(UNDEFINED_LENGTH, str.length)
    }

    return str
  }

  static lex(str: string) {
    const tokens = []

    while (str.length) {
      str = this.lexString(str, tokens)

      str = this.lexNumber(str, tokens)

      str = this.lexBoolNull(str, tokens)

      if (JSON_WHITESPACE.includes(str[0])) {
        str = str.slice(1, str.length)
      } else if (JSON_SYNTAX.includes(str[0])) {
        tokens.push(str[0])
        str = str.slice(1, str.length)
      } else {
        JSONParser.InvalidJSON(str[0])
      }
    }

    return tokens
  }
}