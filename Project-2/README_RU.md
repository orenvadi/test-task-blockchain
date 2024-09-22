
# Ethereum Mini Blockchain Indexer

## Обзор

Это приложение CLI извлекает информацию о блоках из блокчейна Ethereum и сохраняет ее в файл.

Собранная информация включает в себя:
- Номер блока
- Хэш блока
- Количество транзакций
- Временная метка блока

Приложение начинает выборку с номера блока, указанного с помощью флага (`--start`), и продолжает по нарастающей.


## Зависимости

- Go 1.16+
- Ethereum RPC (Infura, локальный узел Geth). Я пользовался этим `https://mainnet.infura.io/v3/8dcac4c10513450baedf76ff28e48bf0`.

## Установка

Для установки зависимостей:

``bash
go mod tidy
```

## Создание бинарного файла

``bash
go build main.go
```

## Запуск индексатора

``bash
go run main.go run --rpc=<ваш-rpc-url> --start=<номер блока> --out=<выходной файл>
```

### Пример

``bash
go run main.go run --rpc=https://mainnet.infura.io/v3/YOUR-PROJECT-ID --start=1 --out=blocks.log
```

Оно подключится к узлу Ethereum RPC и начнет получать данные о блоках начиная с 1, сохраняя информацию в `blocks.log`.

## Пример вывода

```
Number: 1
Hash: 0xabc...
TxCount: 10
Timestamp: 2023-07-10T10:30:00Z

Number: 2
Hash: 0xdef...
TxCount: 15
Timestamp: 2023-07-10T10:30:12Z
```
