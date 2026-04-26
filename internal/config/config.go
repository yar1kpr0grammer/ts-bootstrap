package config

var BlockadeFileName = ".block_ts"

var NodeType = "module"
var IndexFileContent = `let message:string = "Hello world"
console.log(message)
`
var TsconfigContent = `{
 "compilerOptions": {
   // File Layout
   "rootDir": "./src",
   "outDir": "./dist",

   // Environment Settings
   "module": "nodenext",
   "target": "esnext",
   "types": ["node"],

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
