# Mean Inter-Reference Time

Esse script pode ser usado para obter a média de diferença de tempo nos acessos de um determinado item. É considerado que o dado esteja em formato `csv`.

## Formato de entrada do dado

O dado utilizado pelo script tem as seguintes colunas:
- tenant: identificador único do tenant
- object: identificador único do objeto
- ts: timestamp

## Formato de saída do dado

O dado resultante possui as seguintes colunas:
- object: identificador único do objeto
- MIRT: média da diferença de tempo dos acessos do item

## Como executar

`python3 mirt.py -load_path=workload.csv -output=MIRT.csv`

