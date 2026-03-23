package main

var blockadeFileName = ".block_ts"
var blockadeMessage = "tsBootstrup is blocked in this directory"
var indexFileContent = `let message:string = "Hello world"
console.log(message)
`
var readMeContent = "Чтобы запустить: `npm start`\n\n[Исходный код](https://github.com/yar1kpr0grammer)"
var tsconfigContent = `{
 "compilerOptions": {
   // File Layout
   "rootDir": "./src",
   "outDir": "./dist",

   // Environment Settings
   "module": "nodenext",
   "target": "esnext",
   "types": [],

   // Other Outputs
   "sourceMap": true,
   "declaration": true,
   "declarationMap": true,

   // Stricter Typechecking Options
   "noUncheckedIndexedAccess": true,
   "exactOptionalPropertyTypes": true,

   // Recommended Options
   "strict": true,
   "jsx": "react-jsx",
   "verbatimModuleSyntax": true,
   "isolatedModules": true,
   "moduleDetection": "force",
   "skipLibCheck": true,
 }
}

`
