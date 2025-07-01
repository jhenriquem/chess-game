# ♟️ Servidor de Xadrez Online 

 🧠 Um projeto de estudos. Decidi criar esse projeto para o [Hack Club Summer of Making 2025!](https://summer.hackclub.com/). Um servidor de xadrez online desenvolvido com Go e WebSocket usando a biblioteca [gorilla/websocket](https://github.com/gorilla/websocket). O servidor permite a conexão de dois jogadores, gerencia o jogo e verifica a integridade da conexão usando um sistema de ping/pong.


### 📌 Project progress
#### ✅ Etapas Concluídas

- [x] Criar servidor WebSocket básico 
- [x] Estruturar projeto com separação de responsabilidades (`game/`, `core/`, `server/`)
- [x] Criar lógica de pareamento entre dois jogadores
- [x] Criar tipo `Player` com controle de conexão
- [x] Implementar monitoramento de conexão com `ping/pong`
- [x] Encerrar a partida ao detectar desconexão
- [x] Enviar mensagens entre os jogadores (relay)
 
 #### 🗒️ Notas
 - [ ] Melhora a logica de desconexão entre jogadores  


#### 🔜 Etapas Futuras

🧠 Lógica de Jogo

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
├── main.go
├── server/
│   └── server.go
├── game/
│   ├── game.go
│   ├── player.go
│   ├── ...
│   └── board.go 
├── go.mod
└── README.md (this file)

