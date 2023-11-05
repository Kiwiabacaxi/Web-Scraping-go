# Web-Scraping-go

Este é um projeto simples de web scraping escrito em Go. Ele raspa citações do site "http://quotes.toscrape.com/" e salva os resultados em um arquivo JSON. O projeto utiliza a biblioteca `colly` para o web scraping e segue os princípios da Clean Architecture.

## Estrutura do Projeto

O projeto é estruturado da seguinte maneira:

- `cmd/myapp/main.go`: Este é o ponto de entrada do programa. Ele configura o web scraper e inicia o processo de scraping.
- `cmd/dto/quote.go`: Este arquivo contém a definição do Data Transfer Object (DTO) para as citações.
- `pkg/controllers/quote.go`: Este arquivo contém a lógica de controle para o web scraping.
- `pkg/model/quote.go`: Este arquivo contém a definição do modelo de dados para as citações.
- `pkg/usecases/quote.go`: Este arquivo contém a lógica de negócios para o web scraping.
- `pkg/usecases/mappers.go`: Este arquivo contém funções para mapear entre modelos de dados e DTOs.
- `data/quotes.json`: Este é o arquivo onde as citações raspadas são salvas.

## Como Executar

Para executar o projeto, siga estas etapas:

1. Clone o repositório.
2. Navegue até a pasta do projeto.
3. Execute `go run cmd/myapp/main.go`.

Isso iniciará o processo de web scraping e salvará as citações raspadas no arquivo `data/quotes.json`.

## Testes

Este projeto inclui testes unitários. Os testes estão localizados nos arquivos `_test.go` correspondentes aos arquivos de código que eles testam.

Para executar os testes, siga estas etapas:

1. Navegue até a pasta do projeto.
2. Execute `go test ./...`.

Isso executará todos os testes em todos os pacotes do projeto.

Se você deseja adicionar mais testes, crie um arquivo `_test.go` no mesmo diretório do arquivo de código que você deseja testar (se ainda não existir) e adicione funções de teste a esse arquivo. Cada função de teste deve começar com a palavra `Test` e ter um parâmetro do tipo `*testing.T`.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.