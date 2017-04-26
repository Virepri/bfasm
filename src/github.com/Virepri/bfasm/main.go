package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/Virepri/bfasm/VarLexer"
	"github.com/Virepri/bfasm/Lexer"
)

func main(){
	//BFASM - a fun way to say a stupid thing
	if len(os.Args) >= 2 {
		if f, t := os.Open(os.Args[1]); t == nil {
			fi,_ := f.Stat()
			idat := make([]byte,fi.Size())
			f.Read(idat)
			file := string(idat)

			var Lexicons []Lexer.Lexicon

			//TODO: Actually make sure this checks if it's a variables section or not.
			VarLexer.LexVars(file[:strings.Index(file,"!")])
			Lexicons = Lexer.Lex(file[strings.Index(file,"!")+1:])

			fmt.Println(VarLexer.Variables)
			fmt.Println(Lexicons)
		} else {
			fmt.Println("Could not open file",os.Args[1])
		}
	} else {
		fmt.Println("No file supplied")
	}
}
