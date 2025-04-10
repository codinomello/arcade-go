# arcade-go
🍄 arcade-go - plataformer 2D usando raylib

## 🎮 Ideia geral
Um plataformer estrelando o Gopher 🐹, mascote da linguagem Go!

> O Gopher precisa coletar pacotes perdidos da linguagem Go que caíram em servidores bugados pelo mundo. Cada fase é um servidor, cheio de bugs, obstáculos e estruturas bizarras de código corrompido.

## 🔧 Mecânicas principais

### 🔳 Controles
- Andar para os lados
- Pular (com física)
- Agachar (pra passar por túneis)
- Pulo duplo (power-up)

### 📦 Itens colecionáveis
- **Pacotes (`package main`)**: Objetos principais do jogo. Colete todos pra liberar o próximo nível.
- **Goroutines**: Power-ups que deixam o Gopher mais rápido temporariamente.
- **Garbage Collectors**: Coletáveis que limpam bugs da fase (tipo moedas).

### 🐞 Inimigos (bugs)
- **NullPointer**: Anda em linha reta, toca = dano.
- **Deadlock**: Fica parado, mas bloqueia o caminho.
- **RaceCondition**: Corre aleatoriamente quando te vê.
- **Segfault**: Explode em partículas se você chegar perto, te empurrando.

### 💥 Power-ups
- **Go Mod**: Atualiza o Gopher (habilita dash ou pulo duplo)
- **Static Check**: Torna o Gopher invencível por 5 segundos

## 🧠 Desafios criativos
- Plataformas móveis (tipo *pipelines*)
- Portais que lembram canais Go (`chan`) — entra num portal, sai no outro
- Puzzles com semáforos (waitgroups) pra abrir portas

## 🌎 Estrutura de fases

### Exemplo de 3 fases:
1. **"localhost"** – Tutorial, fase fácil
2. **"prod-server-1"** – Inimigos e plataformas móveis
3. **"CI/CD-pipeline"** – Tudo em movimento, corrida contra o tempo

## 📦 Recursos gráficos e sonoros
- Gopher oficial: https://github.com/golang-samples/gopher-vector
- Sons de “compilação” para pegar itens
- Música estilo terminal retro / 8-bit
- Paleta de cores baseada em temas de editores (Monokai, Dracula, etc)

## 🌟 Estilo visual
- Visual flat 2D, estilo pixel ou vetorial
- Códigos como plano de fundo (scrolling com `DrawText`)
- Nuvens que são “blocos de código”
- Inimigos com carinhas tipo emoji de erro 🧟‍♂️

## ⚙️ Extensões futuras
- Sistema de conquistas (ex: "Você venceu um *race condition*!")
- Modo speedrun
- Editor de fases (níveis são arquivos `.go` com comentários especiais)