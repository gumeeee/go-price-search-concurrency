# Price Search

Este é um projeto de estudo em Go que simula a busca de preços em diferentes sites e calcula o preço médio em tempo real. O objetivo é demonstrar o uso de goroutines, canais e sincronização com `sync.WaitGroup`.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

```
price-search/
├── cmd/
│   └── main.go          # Ponto de entrada do programa
├── internal/
│   ├── fetcher/
│   │   └── price_fecher.go  # Lógica para buscar preços de diferentes sites
│   ├── models/
│   │   └── price.go         # Definição do modelo de dados PriceDetail
│   └── processor/
│       └── show_prices.go   # Processamento e exibição dos preços e média
├── go.mod               # Gerenciamento de dependências
└── README.md            # Documentação do projeto
```

## Funcionamento

O programa realiza as seguintes etapas:

1. **Busca de Preços**: 
   - A função `FetchPrices` no arquivo `price_fecher.go` utiliza goroutines para buscar preços de diferentes "sites" simultaneamente.
   - Cada "site" retorna um preço simulado com um atraso para simular o tempo de resposta.

2. **Processamento e Exibição**:
   - A função `ShowPricesAndAVG` no arquivo `show_prices.go` recebe os preços através de um canal e exibe os preços recebidos junto com a média atualizada em tempo real.

3. **Sincronização**:
   - O uso de `sync.WaitGroup` garante que todas as goroutines de busca de preços sejam concluídas antes de fechar o canal de preços.

## Como Executar

### Pré-requisitos

- [Go](https://golang.org/) instalado na máquina (versão 1.16 ou superior).

### Passos

1. Clone o repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd price-search
   ```

2. Execute o programa:
   ```bash
   go run cmd/main.go
   ```

3. O programa exibirá os preços recebidos de diferentes "sites" e a média calculada em tempo real, seguido do tempo total de execução.

### Exemplo de Saída

```
[11/05/2025 14:00:01] Price received from Site A: R$ 45.67 | Average Price at the moment: R$ 45.67
[11/05/2025 14:00:03] Price received from Site C: R$ 78.90 | Average Price at the moment: R$ 62.29
[11/05/2025 14:00:04] Price received from Site B: R$ 32.10 | Average Price at the moment: R$ 52.89
[11/05/2025 14:00:07] Price received from Site D: R$ 56.78 | Average Price at the moment: R$ 53.36
...

Tempo Total: 7.123456s
```

## Detalhes Técnicos

- **Goroutines**: Utilizadas para realizar buscas simultâneas de preços.
- **Canais**: Comunicação entre as goroutines de busca e a de processamento.
- **sync.WaitGroup**: Sincronização das goroutines para garantir que todas sejam concluídas antes de fechar o canal.

## Possíveis Melhorias

- Adicionar testes unitários para as funções principais.
- Implementar tratamento de erros para simular falhas em "sites".
- Permitir a configuração dinâmica dos "sites" e tempos de resposta.

## Licença

Este projeto é apenas para fins de estudo e não possui uma licença específica.