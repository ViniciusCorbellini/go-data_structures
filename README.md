# Estruturas de Dados em Go 🧠⚙️

Este repositório contém implementações simples de estruturas de dados usando a linguagem Go (Golang).  
Cada estrutura possui sua própria pasta com código modularizado, exemplos práticos e comandos para execução.


## 📚 Estruturas Implementadas

### 🔗 Lista Encadeada (Linked List)

Implementação de uma lista encadeada com suporte para:
- Inserção e Remoção de nós
- Acesso de Índices 
- Exibição da Lista completa
- Inversão da Lista
- Remoção de todas as ocorrências de um valor

### 📥 Fila (Queue)

Estrutura de fila (FIFO - First In First Out) com operações básicas:
- Enqueue (adicionar elemento)
- Dequeue (remover elemento)
- Peek (visualizar o primeiro da fila)
- Exibição da Fila completa
  
### 🗃️ Pilha (Stack)

Estrutura de pilha (LIFO - Last In First Out) com funcionalidades:
- Push (empilhar)
- Pop (desempilhar)
- Peek (visualizar o topo)
- Exibição da Pilha completa

## Estrutura do Projeto

```
GO_TUTORIAL/
│
├── cmd/
│   └── data_structures/
│       ├── linkedlist/
│       ├── queue/
│       └── stack/
│
├── data_structures/
│   ├── linkedlist/
│   │   ├── LinkedList.go
│   │   └── Node.go
│   ├── queue/
│   │   └── Queue.go
│   └── stack/
│       └── Stack.go
│
├── .gitignore
├── go.mod
└── README.txt
```

## Pré-requisitos
- Go instalado (versão 1.24.5 ou superior)  
- Terminal configurado com acesso ao comando `go`

## Como executar (No terminal)

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/ViniciusCorbellini/go-data_structures.git
   ```

2. **Mude o diretório para a pasta do projeto:**
   ```bash
   cd go-data_structures
   ```
   

3. **Execute o programa: Execute um dos comandos de acordo com a estrutura desejada**
   ```bash
   go run .\cmd\data_structures\linkedlist\
   ```
   ```bash
   go run .\cmd\data_structures\queue\
   ```
   ```bash
   go run .\cmd\data_structures\stack\
   ```

------------------------------------------------------------
✍️ Autor
------------------------------------------------------------

> Desenvolvido por Vinicius Souza Corbellini. Contribuições, sugestões e melhorias são sempre bem-vindas!