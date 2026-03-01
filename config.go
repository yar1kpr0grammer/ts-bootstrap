package main

var indexFileContent = "console.log(\"Hello, world!\")"
var readMeContent = "Чтобы запустить: `npm start`"
var tsconfigContent = `
{
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
