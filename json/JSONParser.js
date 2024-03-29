"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.JSONLexer = exports.JSONParser = void 0;
const JSON_WHITESPACE = [' ', '\t', '\b', '\n', '\r'];
// const JSON_SYNTAX = new RegExp(/[\w\d.,{}\[\]\-\"\':]/)
const JSON_SYNTAX = ["]", "]", "{", "}", ",", ":"];
const NUMBER_CHARACTERS = new RegExp(/[\d.\-e]/);
const TRUE_LENGTH = 4;
const FALSE_LENGTH = 5;
const UNDEFINED_LENGTH = 9;
class JSONParser {
    static InvalidJSON(char) {
        throw new Error(`Unexpected character: ${char}`);
    }
    static parse(jsonString) {
        let index = 0;
        function parseArray() {
        }
        function parseObject() {
        }
        // function parseValue() {
        //   for (index < jsonString.length) {
        //     const char = jsonString[index]
        //     if (char === Tokens.BracketOpen) {
        //     }
        //     index++
        //   }
        // }
        // return parseValue()
    }
}
exports.JSONParser = JSONParser;
class JSONLexer {
    static lexString(str, tokens) {
        let json_string = '';
        if (str[0] === '"') {
            str = str.slice(1, str.length);
        }
        else {
            return str;
        }
        for (let char of str) {
            if (char === '"') {
                break;
            }
            else {
                json_string += char;
            }
        }
        if (json_string.length) {
            tokens.push(json_string);
            str = str.slice(json_string.length + 1, str.length);
        }
        // if (json_string[json_string.length - 1] !== Tokens.Quote) {
        //   JSONParser.InvalidJSON(json_string[json_string.length - 1])
        // }
        return str;
    }
    static lexNumber(str, tokens) {
        let json_string = '';
        for (let char of str) {
            if (NUMBER_CHARACTERS.test(char)) {
                json_string += char;
            }
            else {
                break;
            }
        }
        if (json_string.length) {
            const json_number = Number(json_string);
            tokens.push(json_number);
            return str.slice(json_string.length, str.length);
        }
        return str;
    }
    static lexBoolNull(str, tokens) {
        let true_str = str.slice(0, TRUE_LENGTH);
        let false_str = str.slice(0, FALSE_LENGTH);
        let undefined_str = str.slice(0, UNDEFINED_LENGTH);
        if (true_str === 'true') {
            tokens.push(true);
            return str.slice(TRUE_LENGTH, str.length);
        }
        else if (true_str === 'null') {
            tokens.push(null);
            return str.slice(TRUE_LENGTH, str.length);
        }
        else if (false_str === 'false') {
            tokens.push(false);
            return str.slice(FALSE_LENGTH, str.length);
        }
        else if (undefined_str === 'undefined') {
            tokens.push(undefined);
            return str.slice(UNDEFINED_LENGTH, str.length);
        }
        return str;
    }
    static lex(str) {
        const tokens = [];
        while (str.length) {
            str = this.lexString(str, tokens);
            str = this.lexNumber(str, tokens);
            str = this.lexBoolNull(str, tokens);
            // str = this.lexArray(str, tokens)
            // str = this.lexObject(str, tokens)
            if (JSON_WHITESPACE.includes(str[0])) {
                str = str.slice(1, str.length);
            }
            else if (JSON_SYNTAX.includes(str[0])) {
                tokens.push(str[0]);
                str = str.slice(1, str.length);
            }
            else {
                JSONParser.InvalidJSON(str[0]);
            }
        }
        return tokens;
    }
}
exports.JSONLexer = JSONLexer;
