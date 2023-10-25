## API Multithreading GoExpert

Esta é uma aplicação Go que oferece uma API para consultar informações de endereços usando a API ViaCep e BrasilAPI. A aplicação utiliza concorrência para melhorar o desempenho ao buscar informações de endereços de diferentes fontes simultaneamente.

### Configuração

Antes de executar a aplicação, você deve configurar as variáveis de ambiente no arquivo `.env` ou no ambiente em que a aplicação será executada. Aqui está um exemplo de configuração:

```env
WEB_SERVER_PORT=:8000
VIACEP_HOST_API=https://viacep.com.br/ws/
BRASILCEP_HOST_API=https://brasilapi.com.br/api/cep/v1/
DB_FILE_PATH=./internal/infra/database/database.db
```
Os valores das variáveis de ambiente VIACEP_HOST_API e BRASILCEP_HOST_API são obrigatórios e devem ser os fornecidos acima (a menos que essas URLs sofram alterações externas). Por outro lado, as variáveis WEB_SERVER_PORT e DB_FILE_PATH podem ser definidas conforme a preferência do usuário.

### Iniciando a Aplicação

Para iniciar o servidor da aplicação, execute o seguinte comando no diretório raiz do projeto:

```bash
go run main.go
```

### Endpoints da API

#### `GET /myaddress/{cep}`

Este endpoint permite a consulta de informações de endereço para um CEP específico. O servidor buscará as informações de endereço usando as APIs ViaCep e BrasilAPI simultaneamente e retornará os resultados da fonte mais rápida.

**Exemplo de Requisição:**
```
GET /myaddress/12345678
```

**Exemplo de Resposta ViaCep:**
```json
{
    "id": "c8c3d10d-2a9d-4c2c-9ed8-e108c33c780c",
    "cep": "06448-190",
    "logradouro": "Rua Belém",
    "complemento": "",
    "bairro": "Jardim do Líbano",
    "localidade": "Barueri",
    "uf": "SP",
    "ibge": "3505708",
    "gia": "2069",
    "ddd": "11",
    "siafi": "6213"
}
```
**Exemplo de Resposta BrasilCep:**
```json
{
    "id": "dec336e6-7e75-4f3f-9418-60c0b2ed2f88",
    "cep": "06448190",
    "state": "SP",
    "city": "Barueri",
    "neighborhood": "Jardim do Líbano",
    "street": "Rua Belém",
    "service": "viacep"
}
```

### Estrutura do Código

- **`main.go`**: O ponto de entrada da aplicação. Configura a aplicação e inicia o servidor web.
- **`config`**: Contém a lógica para carregar a configuração da aplicação.
- **`internal/infra/database`**: Contém a lógica para inicialização e interação com o banco de dados com Gorm e SQLite, bem como as interfaces e implementações para as operações do banco de dados.
- **`internal/infra/webservers/handlers`**: Contém os manipuladores de solicitação HTTP para os endpoints da API.
- **`internal/infra/webservers/api`**: Contém a lógica para chamar as APIs externas (ViaCep e BrasilAPI) para buscar informações de endereço.
- **`internal/utils`**: Contém funções utilitárias, como a função para adicionar corpo JSON à resposta HTTP.

### Considerações Importantes

- A aplicação utiliza concorrência para melhorar o desempenho, buscando informações de diferentes fontes de forma simultânea.
- Certifique-se de configurar corretamente as variáveis de ambiente antes de iniciar a aplicação para garantir que as chamadas de API externas funcionem conforme esperado.
- As respostas da API estão no formato JSON e incluem um ID gerado para cada entrada de endereço.

Espero que este README seja útil! Se você tiver alguma dúvida ou precisar de mais ajuda, não hesite em perguntar.