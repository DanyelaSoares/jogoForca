package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Função auxiliar que verifica se uma letra já está na lista
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {

	// Palavras para sorteio
	palavras := []string{"golang", "codigo", "backend", "dev", "docker", "terminal", "api"}

	// Gera uma palavra aleatória
	rand.Seed(time.Now().UnixNano())
	palavraSecreta := palavras[rand.Intn(len(palavras))]

	// Inicializa as letras descobertas
	letrasDescobertas := make([]string, len(palavraSecreta))
	for i := range letrasDescobertas {
		letrasDescobertas[i] = "_"
	}

	tentativas := 6
	letrasErradas := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🎯 Bem-vindo ao Jogo da Forca!")
	fmt.Println("Adivinhe a palavra secreta!")

	for tentativas > 0 {

		fmt.Println("\n--------------------------------")
		fmt.Println("Palavra:", strings.Join(letrasDescobertas, " "))
		fmt.Println("Tentativas restantes:", tentativas)
		fmt.Println("Letras erradas:", letrasErradas)
		fmt.Print("Digite uma letra: ")

		scanner.Scan()
		letra := strings.ToLower(scanner.Text())

		// Input inválido
		if len(letra) != 1 || letra < "a" || letra > "z" {
			fmt.Println("⚠ Digite apenas uma letra válida (a-z).")
			continue
		}

		// Verifica repetição
		if strings.Contains(strings.Join(letrasDescobertas, ""), letra) || contains(letrasErradas, letra) {
			fmt.Println("🔁 Você já tentou essa letra.")
			continue
		}

		// Verifica letra encontrada
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
			fmt.Println("❌ Letra incorreta!")
		} else {
			fmt.Println("✔ Boa! Você acertou uma letra!")
		}

		// Verifica vitória
		if strings.Join(letrasDescobertas, "") == palavraSecreta {
			fmt.Println("\n🎉 Parabéns!!! Você acertou a palavra:", palavraSecreta)
			fmt.Println("--------------------------------")
			return
		}
	}

	// Caso acabe as tentativas
	fmt.Println("\n--------------------------------")
	fmt.Println("💀 Fim de jogo!")
	fmt.Println("A palavra era:", palavraSecreta)
	fmt.Println("--------------------------------")
}
