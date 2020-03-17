package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func userStringInput(prompt string) string {

	userPrompt := prompt

	if userPrompt == "" {
		prompt = "Digite: "
	}

	//Define variavel do tipo arquivos do sistma
	var f *os.File

	//Atribui stdIn para  userInput
	f = os.Stdin

	// cria Scanner de f
	scanner := bufio.NewScanner(f)

	// Mostra prompt de entrada para usuário
	io.WriteString(os.Stdout, prompt)

	//
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		// define variavel linha
		// pega valor da linha e retira espaços em branco
		line := bytes.TrimSpace(scanner.Bytes())
		// se linha sem espaços em branco tiver zerobytes (nenhum outro caracter
		if len(line) == 0 {

			//mostra prompt de entrada e continua loop
			io.WriteString(os.Stdout, prompt)
			continue
		} else {
			// se linha tiver mais de zero bytes sai do loop e retorna linha
			break

		}
	}
	return scanner.Text()
}

func main() {

	//recebe valor de prompt
	// será retirado quando prompt virar um pacote

	prompt := ""
	userInput := userStringInput(prompt)
	io.WriteString(os.Stdout, userInput)

}
