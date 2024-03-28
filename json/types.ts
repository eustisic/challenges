export type TokenType =
  | "BraceOpen"
  | "BraceClose"
  | "BracketOpen"
  | "BracketClose"
  | "String"
  | "Number"
  | "Comma"
  | "Colon"
  | "True"
  | "False"
  | "Null"

export interface Token {
  type: TokenType
  value: string
}

export enum Tokens {
  BraceOpen = '{',
  BraceClose = '}',
  BracketOpen = '[',
  BracketClose = ']',
  Quote = '"', 
  Dash = '-',
  Comma = ',',
  Colon = ':',
}