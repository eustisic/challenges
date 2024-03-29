import fs from 'fs'
import path from 'path'
import { JSONParser } from '../JSONParser'

const TEST_DIRECTORIES = ["step1", "step2", "step3", "step4"]

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
            const parserJson = JSONParser.parse(data)
            const refJson = JSON.parse(data)
            const result = JSON.stringify(refJson) === JSON.stringify(parserJson)
            if (!result) {
              console.log(refJson)
              console.log(parserJson)
            }
            console.log(`${filePath}: `, result)
          } catch (parserErr) {
            if (file.includes('invalid')) {
              console.log(`${filePath}: `, true)
            } else {
              console.error(filePath)
              console.error(parserErr)
            }
          }
        })
      })
    })
  }
}

run()