# Список использованных пакетов в проекте

- [github.com/ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)

    Простой в использовании пакет для чтения конфигураций с yaml файла и из
    переменных окружения. Ничего лишнего.

- [github.com/pkg/errors](https://github.com/pkg/errors)

    Всем известный пакет для лучшей обработки ошибок, обогащения ошибок
    дополнительным контекстом. С помощью этого пакета можно оборачивать
    ошибки другими ошибками, полезными сообщениями, а так же разворачивать
    цепочку обернутых ошибок для получения искомой ошибки.

- [go.uber.org/zap](https://go.uber.org/zap)

    Один из быстрейших логгеров с выводом структурированной информации, поддерживающий
    уровни логирования.

- [github.com/osamingo/jsonrpc/v2](https://github.com/osamingo/jsonrpc/v2)

    Лучшая библиотека, которую нашел для работы с JSONRPC.
    Прост в использовании, удобно реализовыватьпод него свои методы,
    избавляет от парсинга запросов.

- [github.com/stretchr/testify](https://github.com/stretchr/testify)

    Лидер среди пакетов тестирования. Содержит в себе ряд
    полезных пакетов, которые необходимы при тестировании,
    такие как assert, require, пакет mock для mock-тестирования
    suite для набора взаимосвязанных тестов.

- [github.com/jackc/pgx/v4](https://github.com/jackc/pgx/v4)
- [github.com/jackc/pgconn](https://github.com/jackc/pgconn)

    Один из самых популярных пакетов для работы с Postgres.
    До сих пор поддерживается, реализует множество дополнтельных фич
    Postgres, которые не стандартный database/sql не позволяет использовать.

- [github.com/testcontainers/testcontainers-go](https://github.com/testcontainers/testcontainers-go)
- [github.com/docker/go-connections](https://github.com/docker/go-connections)

    Пакет для запуска из программного кода докер-контейнеров и
    их последующей очистки. Очень удобно при интерграционном
    тестировании.

- [github.com/go-chi/chi/v5](https://github.com/go-chi/chi/v5)

    Легковесный, быстрый, гибконастраеваемый роутер. Совместим с net/http.
    Не тянет за собой много зависимостей, много удобных миддлваров.

- [github.com/google/uuid](https://github.com/google/uuid)

    Только лишь для генерации уникальных индентификаторов запросов.