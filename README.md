# ♟️ Jogo de Xadrez Online via terminal 

![Versão](https://img.shields.io/github/v/release/jhenriquem/gom-editor?label=vers%C3%A3o)
![Go Version](https://img.shields.io/github/go-mod/go-version/jhenriquem/gom-editor)
![Feito com Go](https://img.shields.io/badge/feito%20com-Go-00ADD8?logo=go)

 🧠 Decidi criar esse projeto para o [Hack Club Summer of Making 2025!](https://summer.hackclub.com/). Um servidor e um client de um jogo xadrez online desenvolvido com Go. O servidor permite a conexão de dois jogadores, gerencia o jogo e verifica a integridade da conexão usando um sistema de ping/pong. O client é responsável pela conexão com o sevidor e renderização do jogo no terminal do player. 

### Instalação 

- Acesse [release](https://github.com/jhenriquem/gom-editor/releases/tag/v0.1.1) e baixe o executavel referente ao seu sistema operacional
- Acessi-o via terminal e execute passando um nome para o seu jogador  ```chess-game.exe [name]```

#### 🚀 Visão Geral

- Servidor: Implementado em Go, utilizando WebSockets para comunicação em tempo real entre os jogadores.

- Cliente: Interface baseada em terminal que permite aos jogadores se conectarem ao servidor e jogarem partidas de xadrez.

- Tecnologias:
    - Go
    - WebSockets
    - tcell (para renderização no terminal)

## 📁 Estrutura do Projeto

```bash
chess-game/
├── client/
│   ├── ui/           # Renderização
│   └── main.go       # Ponto de entrada
├── model/
├── net/              # Conexões
├── server/
│   ├── handler/      # Gerenciamento de conexão e match
│   ├── game/         # Logica de jogo
│   └── main.go       # Ponto de entrada
├── go.mod
└── README.md (this file)

