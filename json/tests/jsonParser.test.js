"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const JSONParser_1 = require("../JSONParser");
const TEST_DIRECTORIES = ["step1"];
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
                        console.log(`${file}:\n`, JSONParser_1.JSONParser.parse(data));
                    }
                    catch (parserErr) {
                        console.error(parserErr);
                    }
                });
            });
        });
    }
};
run();
