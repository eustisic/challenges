import { json } from "stream/consumers"
import { TokenType, Token, Tokens } from "./types"

const JSON_WHITESPACE = " "

const JSON_SYNTAX = new RegExp(/[\w\d.,{}\[\]\-\"\':]/)

const NUMBER_CHARACTERS = new RegExp(/[\d.\-e]/)

export class JSONParser {

    static InvalidJSON() {
      throw new Error(`Invalid JSON`)
    }

    static parse(jsonString: string) {
        let index = 0

        // function parseString() {
            
        // }

        // function parseNumber() {
            
        // }

        // function parseBooleanOrNull() {
            
        // }

        // function parseArray() {
            
        // }

        function parseObject() {
           
        }

        function parseValue() {

          for (index < jsonString.length) {
            const char = jsonString[index]
            
            if (char === Tokens.BracketOpen) {
              
            }



            index++
          }
         
        }
            
        return parseValue()
    }
}

class JSONLexer {
  static lexString(str: string) {
    let json_string = ''

    if (str[0] === Tokens.Quote) {
      str = str.slice(1, str.length)
    } else {
      return {json_string, return_str: str}
    } 

    for (let char of str) {
      if (char === Tokens.Quote) {
        return {json_string, return_str: str}
      } else {
        json_string += char
      }
    }

    JSONParser.InvalidJSON()
  }

  static lexNumber(str: string) {
    let json_number = ''

    for (let char of str) {
      if (NUMBER_CHARACTERS.test(char)) {
       json_number += char
      } else {
        break
      }
    }

    

    
  }

  static lexBool(starting_str: string): {json_string: string | undefined, return_str: string} {
    return {json_string: undefined, return_str: starting_str}
  }

  static lexArray(starting_str: string): {json_string: string | undefined, return_str: string} {
    return {json_string: undefined, return_str: starting_str}
  }

  static lexObject(starting_str: string): {json_string: string | undefined, return_str: string} {
    return {json_string: undefined, return_str: starting_str}
  }



  static lex(str: string) {
    const tokens = []

    while (str.length) {
      let { json_string, return_str } = this.lexString(str) ?? {}

      if (json_string) {
        tokens.push(json_string)
      }

      if (str[0] === JSON_WHITESPACE) {
        str = str.slice(1, str.length)
      } else if (JSON_SYNTAX.test(str[0])) {
        tokens.push(str[0])
        str = str.slice(1, str.length)
      } else {
        JSONParser.InvalidJSON()
      }
    }

    return tokens
  }
}