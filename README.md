# Решение тестевого задания для стажёра Backend
# Сервис динамического сегментирования пользователей

### Проблема:

В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие.
На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.

### Задача:

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

### Решение:

Разработанный сервис умеет:
- Создать сегмент
- Удалить сегмент по имени
- Добавить пользователя в сегмент
- Удалить пользователя из сегмента
- Автоматически добавить в созданный сегмент N% пользователей от общего числа
- Создать файл с отчётом о добавлении/удалении пользователя из сегмента с точностью до месяца

### Запуск:

Для запуска сервера необходимо настроить окружение (создать .env файл или задать переменные окружения вручную, их список будет указан позже) и создать базу данных в PostgreSQL по схеме, указанной в файле schema.sql (также рекомендуется добавить несколько записей в таблицу users). При успешном запуске исполняемого файла в консоль не происходит никакого вывода.

Для запуска приложения в docker compose используйте команду `docker compose up`.
В некоторых ситуациях могут возникать ошибки при попытках подключения к postgres, однако, если доверять опыту, это связано с неверной последовательностью создания контейнеров докером. В таком случае остановите контейнеры и перезапустите их командой `docker compose up`.

### Переменные окружения:

Переменные окружения, которые должны быть обязательно заданы при запуске приложения

| ИМЯ | НАЗНАЧЕНИЕ |
| --- | ---------- |
| POSTGRES_USER | имя пользователя в базе данных |
| POSTGRES_DB | имя базы данных |
| POSTGRES_PORT | порт базы данных |
| POSTGRES_HOST | хост базы данных. инициализируйте эту переменную, если запскаете приложение без использования docker compose, иначе оставьте пустой |
| POSTGRES_PASSWORD | пароль пользователя базы данных |
| APP_PORT | порт, на котором должно быть запущено прииложение |

Оставшиеся две переменные окружения используются при запуске приложения с использованием docker compose.
| ИМЯ | НАЗНАЧЕНИЕ |
| --- | ---------- |
| PG_INNER_PORT | порт внутри докер контейнера, который слушает база данных |
| APP_OUTTER_PORT | внешний порт контейнера, который отображается на APP_PORT |

### Документация:

Swagger документациия доступна по URL `http://<host>:<port>/swagger/index.html`.
