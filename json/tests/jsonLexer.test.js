"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const JSONParser_1 = require("../JSONParser");
const TEST_DIRECTORIES = ["step3"];
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
                    try {
                        const tokens = JSONParser_1.JSONLexer.lex(data);
                        const expected = ['{', 'key1', ':', true, ',', "key2", ':', false, ',', "key3", ':', null, ',', "key4", ':', "value", ',', "key5", ':', 101, '}'];
                        console.log(tokens.every((ele, idx) => ele === expected[idx]));
                    }
                    catch (parserErr) {
                        if (file.includes('invalid')) {
                            console.log(true);
                        }
                        else {
                            console.log(err);
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
