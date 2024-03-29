import fs from 'fs'
import path from 'path'
import { JSONLexer, JSONParser } from '../JSONParser'

const TEST_DIRECTORIES = ["step3"]

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

          try {
            const tokens = JSONLexer.lex(data)
            const expected = ['{', 'key1', ':', true, ',', "key2", ':', false, ',', "key3", ':', null, ',', "key4", ':', "value", ',', "key5", ':', 101, '}']
            console.log(tokens.every((ele, idx) => ele === expected[idx]))
          } catch (parserErr) {
            if (file.includes('invalid')) {
              console.log(true)
            } else {
              console.log(err)
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