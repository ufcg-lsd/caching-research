# Footprint e Número de Requisições

Esse script pode ser usado para obter o número de diferentes itens (footprint), juntamente com a quantidade de requisições realizadas para dados de cada um dos tenants. É considerado que o dado esteja em formato `csv`.

## Formato de entrada do dado

O dado utilizado pelo script tem as seguintes colunas:
- ts: timestamp
- tenant: identificador único do tenant
- object: identificador único do objeto

## Formato de saída do dado

O dado resultante possui as seguintes colunas:
- tenant: identificador único do tenant
- footprint: quantidade de objetos diferentes do tenant
- requests: número total de requisições feitas ao tenant

## Como executar

`python3 footprint.py workload.csv`

