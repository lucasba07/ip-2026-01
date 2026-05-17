# Sistema de Gerenciamento de Pacientes (Health CRUD)

Um sistema web completo construído em Go (puro, sem frameworks) com PostgreSQL e interface em HTML/CSS para o gerenciamento de pacientes (operações CRUD).

## Estrutura do Projeto

```text
health-crud/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── Readme.md
├── app/
│   ├── main.go
│   ├── handlers/
│   │   ├── createPatientHandler.go
│   │   ├── deletePatientHandler.go
│   │   ├── searchPatientHandler.go
│   │   └── updatePatientHandler.go
│   └── utils/
│       ├── connectToDB.go
│       ├── createPatient.go
│       ├── deletePatient.go
│       ├── getPatientByCPF.go
│       └── updatePatient.go
├── db/
│   └── init/
│       └── init.sql
└── static/
    ├── index.html
    ├── forms/
    │   ├── createPatient.html
    │   ├── deletePatient.html
    │   ├── searchPatient.html
    │   └── updatePatient.html
    └── styles/
        ├── createPatient.style.css
        ├── deletePatient.style.css
        ├── index.style.css
        ├── searchPatient.style.css
        └── updatePatient.style.css
```

## Pré-requisitos

- **Go** (versão 1.21 ou superior)
- **PostgreSQL**
- **Docker e Docker Compose**

## Passo a passo de Instalação e Execução

1. **Clone o repositório** ou navegue até a pasta `health-crud`.

2. **Crie o arquivo `.env`** com base no exemplo fornecido:
   Crie um arquivo chamado `.env` na raiz da pasta `health-crud` e copie o conteúdo de `.env.example`:
   ```env
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=health_db
   DB_HOST=localhost
   DB_PORT=5432
   ```

3. **Inicie o banco de dados via Docker**:
   ```bash
   docker compose up -d
   ```
   Isso iniciará o PostgreSQL e criará a tabela `patients` automaticamente usando o script de inicialização `db/init/init.sql`.

4. **Instale as dependências do projeto Go**:
   ```bash
   go mod tidy
   ```

5. **Execute o servidor**:
   ```bash
   go run app/main.go
   ```
   O servidor estará disponível em `http://localhost:3000`.

## Tabela de Rotas

| Rota | Método | Descrição |
|------|--------|-----------|
| `/` | `GET` | Serve a página inicial (`index.html`) |
| `/create` | `POST` | Cria um novo paciente |
| `/search` | `POST` | Busca e exibe os dados de um paciente via CPF |
| `/update` | `POST` | Atualiza os dados de um paciente via CPF |
| `/delete` | `POST` | Deleta um paciente via CPF |

## Esquema do Banco de Dados

A tabela `patients` é criada automaticamente via script de inicialização com o seguinte esquema:

```sql
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    birth_date DATE NOT NULL,
    blood_type VARCHAR(3) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(150) UNIQUE,
    medical_notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Observações

- **Driver do Banco de Dados:** Este projeto utiliza o driver `github.com/lib/pq` para a comunicação com o PostgreSQL.
- **Contato/Suporte:** Em caso de dúvidas, sinta-se à vontade para abrir uma issue ou entrar em contato.
