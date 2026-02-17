# TCP Scanner Logic

Um scanner de portas TCP eficiente escrito em Go para identificar portas abertas em hosts remotos.

## 📋 Características

- Verificação rápida de portas TCP
- Suporte a scanning de múltiplas portas simultaneamente
- Detecção de banners de serviço
- Logging estruturado
- Tratamento de erros robusto

## 🚀 Instalação

```bash
git clone https://github.com/meluansantos/tcp-scanner-logic.git
cd tcp-scanner-logic
go mod download
```

## 💻 Uso

### Executar o scanner

```bash
go run cmd/scanner/main.go
```

### Build do executável

```bash
go build -o tcp-scanner cmd/scanner/main.go
./tcp-scanner
```

## 📁 Estrutura do Projeto

```
tcp-scanner-logic/
├── cmd/
│   └── scanner/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── scanner/
│   │   ├── scanner.go       # Lógica principal do scanner
│   │   └── result.go        # Estrutura de resultados
│   └── util/
│       └── logger.go        # Utilitários de logging
├── go.mod
└── README.md
```

## 🛠️ Desenvolvimento

### Pré-requisitos

- Go 1.23 ou superior

### Executar testes

```bash
go test ./...
```

### Formatar código

```bash
go fmt ./...
```

### Verificar código

```bash
go vet ./...
```

## 📝 Licença

Este projeto está sob a licença MIT.

## 👤 Autor

**Luan Rodrigues**

- Blog: [Luan Rodrigues](https://luansantos.net)
