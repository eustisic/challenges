class JSONParser {
    static parse(jsonString) {
        let index = 0;

        function nextChar() {
            return jsonString[index++];
        }

        function parseString() {
            let str = '';
            let char;
            while ((char = nextChar()) !== '"') {
                str += char;
            }
            return str;
        }

        function parseNumber() {
            let numStr = '';
            let char;
            while (/[\d.eE+-]/.test(char = nextChar())) {
                numStr += char;
            }
            index--; // Move back one character to leave non-number character for next parsing
            return parseFloat(numStr);
        }

        function parseBooleanOrNull() {
            const word = jsonString.substr(index, 4);
            if (word === 'true') {
                index += 4;
                return true;
            } else if (word === 'fals') {
                index += 5; // Move forward 5 characters to skip 'false'
                return false;
            } else if (word === 'null') {
                index += 4;
                return null;
            }
            throw new Error('Invalid input at position ' + index);
        }

        function parseArray() {
            const arr = [];
            let char;
            while ((char = nextChar()) !== ']') {
                if (char === ',') {
                    continue; // Ignore commas between elements
                }
                index--; // Move back one character to start parsing the next element
                arr.push(parseValue());
            }
            return arr;
        }

        function parseObject() {
            const obj = {};
            let key, value;
            let char;
            while ((char = nextChar()) !== '}') {
                if (char === ',') {
                    continue; // Ignore commas between key-value pairs
                }
                if (char !== '"') {
                    throw new Error('Expected a key at position ' + index);
                }
                key = parseString();
                char = nextChar(); // Skip ':' character
                if (char !== ':') {
                    throw new Error('Expected ":" at position ' + index);
                }
                value = parseValue();
                obj[key] = value;
            }
            return obj;
        }

        function parseValue() {
            const char = nextChar();
            if (char === '{') {
                return parseObject();
            } else if (char === '[') {
                return parseArray();
            } else if (char === '"') {
                return parseString();
            } else if (/\d/.test(char)) {
                index--; // Move back one character to parse number correctly
                return parseNumber();
            } else if (/[tfn]/.test(char)) {
                index--; // Move back one character to parse boolean or null correctly
                return parseBooleanOrNull();
            }
            throw new Error('Unexpected character at position ' + index);
        }

        return parseValue();
    }
}

// Example usage
const jsonString = '{"name": "John", "age": 30, "city": "New York"}';
const jsonObject = JSONParser.parse(jsonString);

