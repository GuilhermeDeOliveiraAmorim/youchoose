# IMDb Web Scraping Service

Este projeto é um serviço de web scraping desenvolvido em Go, destinado a extrair informações de filmes do site IMDb. O serviço utiliza as bibliotecas `Colly` e `GoQuery` para realizar o scraping e manipulação dos dados.

## Funcionalidades

- **Extração de Dados do Filme**: O serviço coleta informações como título, ano de lançamento, duração, gêneros, resumo, poster e avaliação no IMDb.
- **Interface Simples**: Com apenas o ID do filme no IMDb, o serviço retorna um JSON com todos os dados extraídos.

## Tecnologias Utilizadas

- **Linguagem**: Go
- **Bibliotecas**:
  - `Colly`: Para navegação e scraping web.
  - `GoQuery`: Para seleção e manipulação de elementos HTML, semelhante ao jQuery.
  - Go's `net/http`: Para requisições HTTP.

## Instalação

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/usuario/imdb-web-scraping.git
   cd imdb-web-scraping
