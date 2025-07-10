package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	palavraSecreta := "golang"
	letrasDescobertas := make([]string, len(palavraSecreta))
	for i := range letrasDescobertas {
		letrasDescobertas[i] = "_"
	}

	tentativas := 6
	letrasErradas := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for tentativas > 0 {
		fmt.Println("\nPalavra:", strings.Join(letrasDescobertas, " "))
		fmt.Println("Tentativas restantes:", tentativas)
		fmt.Println("Letras erradas:", letrasErradas)
		fmt.Print("Digite uma letra: ")

		scanner.Scan()
		letra := strings.ToLower(scanner.Text())

		if len(letra) != 1 || letra < "a" || letra > "z" {
			fmt.Println("Digite apenas uma letra válida.")
			continue
		}

		if strings.Contains(strings.Join(letrasDescobertas, ""), letra) || contains(letrasErradas, letra) {
			fmt.Println("Você já tentou essa letra.")
			continue
		}

		acertou := false
		for i, l := range palavraSecreta {
			if string(l) == letra {
				letrasDescobertas[i] = letra
				acertou = true
			}
		}

		if !acertou {
			letrasErradas = append(letrasErradas, letra)
			tentativas--
			fmt.Println("Letra incorreta!")
		}

		if strings.Join(letrasDescobertas, "") == palavraSecreta {
			fmt.Println("\n🎉 Parabéns! Você acertou a palavra:", palavraSecreta)
			return
		}
	}

	fmt.Println("\n💀 Fim de jogo! A palavra era:", palavraSecreta)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
