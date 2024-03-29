import fs from 'fs'
import path from 'path'
import { JSONParser } from '../JSONParser'

const TEST_DIRECTORIES = ["step1"]

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
            console.log(`${file}:\n`, JSONParser.parse(data))
          } catch (parserErr) {
            console.error(parserErr)
          }
        })
      })
    })
  }
}

run()