# ♟️ Jogo de Xadrez Online via terminal 

 🧠 Decidi criar esse projeto para o [Hack Club Summer of Making 2025!](https://summer.hackclub.com/). Um servidor e um client de um jogo xadrez online desenvolvido com Go e WebSocket usando a biblioteca [gorilla/websocket](https://github.com/gorilla/websocket). O servidor permite a conexão de dois jogadores, gerencia o jogo e verifica a integridade da conexão usando um sistema de ping/pong. O client é responsável pela conexão com o sevidor e renderização do jogo no terminal do player. 


### 📌 Project progress
#### ✅ Etapas Concluídas

- [x] Criar servidor WebSocket básico 
- [x] Estruturar projeto com separação de responsabilidades (`game/`, `core/`, `server/`)
- [x] Criar lógica de pareamento entre dois jogadores
- [x] Criar tipo `Player` com controle de conexão
- [x] Implementar monitoramento de conexão com `ping/pong`
- [x] Encerrar a partida ao detectar desconexão
- [x] Estrutura basica do client

#### 🗒️ Notas 
- Provavelmente vou separar a struct `Game` do pacote game. Irei criar um pacote `models` que armazene esse tipo de dados. Modificarei o projeto para que se adpte a isso   
#### 🔜 Etapas Futuras

🧠 Lógica de Jogo
- [ ] Tratar o envio de jogadas pelo player
- [ ] Definir representação do tabuleiro de xadrez
- [ ] Implementar controle de turno e regras de movimento
- [ ] Validar jogadas no servidor
- [ ] Detectar xeque, xeque-mate e empate

⏱ Controle de Tempo

- [ ] Implementar cronômetro por jogador (ex: 10 minutos)
- [ ] Finalizar partida quando o tempo de um jogador acabar

💬 Comunicação

- [ ] Criar tipos de mensagens (ex: `move`, `chat`, `resign`, `timeout`)
- [ ] Validar e interpretar cada tipo de mensagem no servidor

## 📁 Estrutura do Projeto

```bash
chess-server/
├── cmd/
│   ├── client/main.go
│   └── server/main.go
├── pkg/
│   ├── protocol/
│   └── pieces/
├── server/
│   ├── match.go
│   ├── connection.go
│   └── core.go
├── client/
│   ├── match.go
│   ├── connection.go
│   └── core.go
├── game/
│   ├── game.go
│   ├── player.go
│   ├── ...
│   └── board.go 
├── go.mod
└── README.md (this file)

