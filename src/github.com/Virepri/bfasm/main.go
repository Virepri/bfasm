package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/Virepri/bfasm/VarLexer"
	"github.com/Virepri/bfasm/Lexer"
	"github.com/Virepri/bfasm/SyntaxAnalysis"
	"github.com/Virepri/bfasm/Compiler"
)

func main(){
	//BFASM - a fun way to say a stupid thing
	if len(os.Args) >= 2 {
		if f, t := os.Open(os.Args[1]); t == nil {
			fi,_ := f.Stat()
			idat := make([]byte,fi.Size())
			f.Read(idat)
			file := string(idat)

			var Lexicons []Lexer.Token

			VarLexer.LexVars(file[:strings.Index(file,"!")])
			Lexicons = Lexer.Lex(file[strings.Index(file,"!")+1:])

			//fmt.Println(VarLexer.Variables)
			//fmt.Println(Lexicons)

			if !SyntaxAnalysis.AnalyzeSyntax(Lexicons,0,0) {
				return
			}

			if output, success := Compiler.Compile(Lexicons); success {
				fmt.Println(output)
				return
			}
			fmt.Println("Failed to compile file",os.Args[1])
		} else {
			fmt.Println("Could not open file",os.Args[1])
		}
	} else {
		fmt.Println("No file supplied")
	}
}
