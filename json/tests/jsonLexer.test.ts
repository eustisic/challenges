import fs from 'fs'
import path from 'path'
import { JSONLexer, JSONParser } from '../JSONParser'

const TEST_DIRECTORIES = ["step3", "step4"]

const EXPECTED: Record<string, any> = {
  step3: {
    valid: ['{', 'key1', ':', true, ',', "key2", ':', false, ',', "key3", ':', null, ',', "key4", ':', "value", ',', "key5", ':', 101, '}']
  },
  step4: {
    valid: [
      '{', "key", ':', "value", ',',
      "key-n", ':', 101, ',',
      "key-o", ':', '{','}', ',',
      "key-l", ':', '[',']', '}'
    ],
    valid2: [
      '{', "key", ':', "value", ',',
      "key-n", ':', 101, ',',
      "key-o", ':', '{',
        "inner key",':', "inner value",
      '}', ',',
      "key-l", ':', '[', "list value", ']', '}'
    ]
  }
}

const run = () => {
  for (let dir of TEST_DIRECTORIES) {
    const dirPath = path.join(`./tests/${dir}`)

    fs.readdir(dirPath, (err: any, files: any[]) => {
      if (err) {
        console.error('Error reading the directory')
        process.exit(1)
      }

      files.forEach(file => {
        const filePath = path.join(dirPath, file)

        fs.readFile(filePath, 'utf8', (err: any, data: any) => {
          if (err) {
            console.error(`Error reading the file ${file}:`, err.message)
            process.exit(1)
          }

          console.log(filePath)
          try {
            const tokens = JSONLexer.lex(data)
            const expected = EXPECTED[dir][file.split('.')[0]]
            
            console.log(tokens.every((ele, idx) => {
              ele !== expected[idx] ? console.log(ele, expected[idx]) : undefined
              return ele === expected[idx]
            }))
          } catch (parserErr) {
            if (file.includes('invalid')) {
              console.log(true)
            } else {
              
              console.log(parserErr)
            }
          }
        })
      })
    })
  }
}


/*

{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}
*/

run()