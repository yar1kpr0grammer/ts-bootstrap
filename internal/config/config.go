package config

var BlockadeFileName = ".block_ts"
var BlockadeMessage = "tsBootstrup is blocked in this directory"
var IndexFileContent = `let message:string = "Hello world"
console.log(message)
`
var ReadMeContent = "Чтобы запустить: `npm start`\n\n[Исходный код](https://github.com/yar1kpr0grammer)"
var TsconfigContent = `{
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
