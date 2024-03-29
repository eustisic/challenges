"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const JSONParser_1 = require("../JSONParser");
const TEST_DIRECTORIES = ["step3", "step4"];
const EXPECTED = {
    step3: {
        valid: ['{', 'key1', ':', true, ',', "key2", ':', false, ',', "key3", ':', null, ',', "key4", ':', "value", ',', "key5", ':', 101, '}']
    },
    step4: {
        valid: [
            '{', "key", ':', "value", ',',
            "key-n", ':', 101, ',',
            "key-o", ':', '{', '}', ',',
            "key-l", ':', '[', ']', '}'
        ],
        valid2: [
            '{', "key", ':', "value", ',',
            "key-n", ':', 101, ',',
            "key-o", ':', '{',
            "inner key", ':', "inner value",
            '}', ',',
            "key-l", ':', '[', "list value", ']', '}'
        ]
    }
};
const run = () => {
    for (let dir of TEST_DIRECTORIES) {
        const dirPath = path_1.default.join(`./tests/${dir}`);
        fs_1.default.readdir(dirPath, (err, files) => {
            if (err) {
                console.error('Error reading the directory');
                process.exit(1);
            }
            files.forEach(file => {
                const filePath = path_1.default.join(dirPath, file);
                fs_1.default.readFile(filePath, 'utf8', (err, data) => {
                    if (err) {
                        console.error(`Error reading the file ${file}:`, err.message);
                        process.exit(1);
                    }
                    console.log(filePath);
                    try {
                        const tokens = JSONParser_1.JSONLexer.lex(data);
                        const expected = EXPECTED[dir][file.split('.')[0]];
                        console.log(tokens.every((ele, idx) => {
                            ele !== expected[idx] ? console.log(ele, expected[idx]) : undefined;
                            return ele === expected[idx];
                        }));
                    }
                    catch (parserErr) {
                        if (file.includes('invalid')) {
                            console.log(true);
                        }
                        else {
                            console.log(parserErr);
                        }
                    }
                });
            });
        });
    }
};
/*

{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}
*/
run();
