# Documentação do Sistema: Jogo da Forca — Go 🎯

## 1. Visão Geral

O Jogo da Forca é um sistema simples desenvolvido em Go (Golang) que roda diretamente no terminal.  
O objetivo do usuário é adivinhar a palavra secreta antes que suas tentativas acabem, digitando uma letra por vez.  

---

## 2. Objetivo do Sistema

Permitir que o usuário jogue o clássico Jogo da Forca, garantindo:  
- Controle das tentativas restantes  
- Validação das letras digitadas  
- Feedback claro sobre acertos e erros  

---

## 3. Regras de Negócio

1. O usuário tem **6 tentativas** para adivinhar a palavra secreta.  
2. O usuário deve digitar **apenas uma letra por vez**.  
3. Cada erro resulta na perda de uma tentativa.  
4. O sistema mantém registro das letras já tentadas.  
5. Ao acertar todas as letras, o usuário vence.  
6. Ao acabar as tentativas, o jogo termina e revela a palavra secreta.  

---

## 4. Requisitos Funcionais (RF)

- **RF01:** O sistema deve exibir a palavra secreta com letras ocultas (`_`) e mostrar as letras corretas à medida que são descobertas.  
- **RF02:** O sistema deve permitir que o usuário digite apenas uma letra por vez.  
- **RF03:** O sistema deve verificar se a letra já foi tentada e informar o usuário.  
- **RF04:** O sistema deve atualizar as tentativas restantes a cada erro.  
- **RF05:** O sistema deve informar o resultado final do jogo (vitória ou derrota) e revelar a palavra secreta.

---

## 5. Requisitos Não Funcionais (RNF)

- **RNF01:** O sistema deve rodar no terminal, sem interface gráfica.  
- **RNF02:** O sistema deve ser compatível com Go (versão 1.18+ recomendada).  
- **RNF03:** As mensagens do sistema devem ser claras e amigáveis.  
- **RNF04:** O desempenho do jogo deve ser imediato, sem atrasos perceptíveis.  

---

## 6. Fluxo do Usuário / Casos de Uso

**Caso de Uso: Jogar Forca**  

1. O usuário executa o programa (`go run forca.go`).  
2. O sistema exibe a palavra secreta com letras ocultas (`_`) e o número de tentativas restantes.  
3. O usuário digita uma letra.  
4. O sistema verifica se a letra já foi tentada:  
   - Se sim, informa o usuário e solicita nova letra.  
   - Se não, avalia se a letra está na palavra secreta.  
5. O sistema atualiza as letras descobertas ou adiciona a letra à lista de erros.  
6. O sistema reduz o número de tentativas em caso de erro.  
7. Passos 3 a 6 se repetem até:  
   - O usuário adivinhar todas as letras (vitória).  
   - As tentativas acabarem (derrota, palavra revelada).  

---

## 7. Exemplo de Execução

```text
Palavra: _ _ _ _ _ _
Tentativas restantes: 6
Letras erradas: 
Digite uma letra: g
Palavra: g _ _ _ _ _
Tentativas restantes: 6
Letras erradas: 

```
## 8. Tecnologias Utilizadas

* Go (Golang)
* Terminal / Prompt de Comando

## 9. Possíveis Melhorias Futuras

* Adicionar palavras aleatórias de um banco de dados ou arquivo externo.
* Criar uma interface gráfica (GUI).
* Permitir múltiplos níveis de dificuldade (mais ou menos tentativas).
* Implementar dicas ou categorias de palavras.

## 10. Autor

* Daniela Soares — Estudante de Análise de Requisitos e Desenvolvedora em aprendizado
