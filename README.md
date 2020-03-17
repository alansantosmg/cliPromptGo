# cliPromptGo Package

CliPrompt é um componente reutilizável desenvolvido com objetivo de fornecer  prompts de entrada de dados de usuários.

O componente composto por funções que podem ser adicionada em aplicações GO que implementem algum tipo de CLI (interface de linha de comando).

Ao utilizar o CliPrompt o desenvolvedor fica liberado da criação de prompts e pode se
concentrar no desenvolvimento das funções de comando de sua CLI.

Para utilizar o componente é necessário baixar e importá-lo para sua aplicação:

## Donwload
Uma vez configurado seu ambiente Go utilize o comando:
```
go get https://github.com/alansantosmg/cliPromptGo
```

## Importação

Para importar use, conforme exemplo:
```
import (
"cliPrompt"
)
```

## Utilização
Na aplicação, invocar a função:

```
userStringInput(prompt)
```
A função userStringInput recebe um parâmetro prompt do tipo string e retorna um valor do tipo string referente ao que foi digitado pelo usuário no prompt de comando.

O parâmetro prompt permite personalizar a mensagem que será mostrada ao usuário no indicando a entrada no prompt de comando

### Exemplo:

```golang
package main

import (
"cliPrompt")

func main() {

  // Personaliza a mensagem da CLI para o usuário
  userPrompt = "Digite um comando: "

  // Invoca a função userStringInput
  // O que for digitado será guardado na variável userInput
  userInput = userStringInput(prompt)

  // Para fins pedagógicos somente, imprime o que foi retornado na tela
  fmt.Println("O comando digitado foi :", userInput )
  }
```

## Observações:

* Se o valor passado no parâmetro da função for "", o componente implementará a mensagem padrão `Digite: `.
* Atualmente o usuário pode abortar a aplicação teclando `Ctrl D`
* Se o usuário digitar somente espaços em branco e teclar `enter` o prompt continuará solicitando entrada.
* Se o usuário digitar `enter` sem nenhum outro caracter ou incluindo espaços em branco, o prompt continuará solicitando entrada.
* se o usuário digitar qualquer caracter exceto espaço em branco e teclar `enter` a função retornará o que foi digitado.

## TO-DO:
[x] vrs 0.9 Função userStringInput(prompt)
[ ] vrs 1.0 Reestruturar aplicação na forma de package
[ ] vrs 1.0 implementar função de conversão de entrada para tipo Int
[ ] vrs 1.1 Implementar função de conversão de entrada para tipo float64
[ ] vrs 1.x Implementar função de conversão de entrada para outros tipos
[ ] vrs 2.0 Implementar controle de bloqueio de sequências de escape
