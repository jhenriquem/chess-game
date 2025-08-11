# ♟️ Jogo de Xadrez Online via terminal 

 🧠 Decidi criar esse projeto para o [Hack Club Summer of Making 2025!](https://summer.hackclub.com/). Um servidor e um client de um jogo xadrez online desenvolvido com Go e WebSocket usando a biblioteca [gorilla/websocket](https://github.com/gorilla/websocket). O servidor permite a conexão de dois jogadores, gerencia o jogo e verifica a integridade da conexão usando um sistema de ping/pong. O client é responsável pela conexão com o sevidor e renderização do jogo no terminal do player. 


## 📁 Estrutura do Projeto

```bash
chess-game/
├── client/
│   └── main.go       # Ponto de entrada
├── model/
├── server/
│   ├── handler/      #
│   ├── game/         # 
│   └── main.go       # Ponto de entrada
├── go.mod
└── README.md (this file)

