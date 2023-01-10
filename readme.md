## Booking backend

### Переменные окружения


#### Backend 
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_LOGLEVEL                                 | Уровень логирования (debug, info, warn, error, fatal) |           | warn                                |         
| TN_LOG_OUTPUT                               | Формат вывода логов (console/json)                    |           | json                                |           
| TN_IP                                       | Ip, на котором запускается сервер                     |           | 0.0.0.0                             |           
| TN_GRPC_PORT                                | Порт для grpc сервера                                 |           | 8080                                |           
| TN_HTTP_PORT                                | Порт для http сервера                                 |           | 3015                                |  

#### Postgres                                                                                                                 
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_POSTGRES_DSN                             | DSN для доступа к postgres                            |     *     |                                     |           
| TN_MIGRATION_MIGRATIONS                     | Путь к папке с миграциями                             |     *     |                                     |           
| TN_MIGRATION_VERSION_TABLE                  | Требуемая версия миграций                             |           | public.schema_version               |           

#### Profile   
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_PROFILE_EMPLOYEE_API_URL                 | URL API ЦП                                            |     *     |                                     |           
| TN_PROFILE_EMPLOYEE_API_TOKEN               | API token (employee)                                  |     *     |                                     |           
| TN_PROFILE_CLIENT_ID                        | Client ID в ЦП                                        |     *     |                                     |           
| TN_PROFILE_CLIENT_SECRET                    | Client secret в ЦП                                    |     *     |                                     |           
| TN_PROFILE_AUTH_URL                         | URL для авторизации пользователя                      |     *     |                                     |           
| TN_PROFILE_TOKEN_URL                        | URL для получения токена                              |     *     |                                     |           
| TN_PROFILE_REDIRECT_URL                     | OAuth redirect url                                    |     *     |                                     |           
| TN_PROFILE_SCOPES                           | Необходимые claims для работы приложения              |           | phone.read,employee.read,email.read | 

#### Sessions                                                                                                                                           
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_ACCESS_EXPIRY                            | Время жизни access токена                             |           | 1h                                  |           
| TN_REFRESH_EXPIRY                           | Время жизни refresh токена                            |           | 48h                                 |           
| TN_SECRET                                   | Secret для подписи токена                             |     *     |                                     |

#### S3
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_AWS_S3_REGION                            | Регион s3                                             |           | ru-central1                         |           
| TN_AWS_S3_ENDPOINT                          | Хост s3                                               |           | storage.yandexcloud.net             |           
| TN_AWS_S3_FILES_BUCKET                      | Название бакета для файлов                            |     *     |                                     |           
| TN_AWS_S3_ACCESS_KEY_ID                     | AWS key                                               |     *     |                                     |           
| TN_AWS_S3_SECRET_ACCESS_KEY                 | AWS secret                                            |     *     |                                     |           
| TN_AWS_S3_TEMPORARY_FILES_PREFIX®           | Префикс для временных файлов                          |           | tmp/                                |           

®На данный префикс в s3 настраивается жизненный цикл - удаление через некоторый промежуток времени.

#### Email sender                                                                                                                                        
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_SKIP_TLS                                 | Флаг на проверку TLS                                  |           | false                               |
| TN_EMAIL_HOST                               | Хост SMTP-сервера                                     |           |                                     |
| TN_EMAIL_PORT                               | Порт SMTP-сервера                                     |           |                                     |
| TN_EMAIL_USER                               | Пользователь SMTP-сервера                             |           |                                     |
| TN_EMAIL_PASSWORD                           | Пароль пользователя SMTP-сервера                      |           |                                     |
| TN_EMAIL_FROM                               | Системный email                                       |           | nicktretyakov@gmail.com                 |

#### Notifications                               
| ENV                                         | Описание                                              | Required  | Default                             |
|---------------------------------------------|-------------------------------------------------------|-----------|-------------------------------------|
| TN_REPORTS_CHECKER_SCHEDULER_DAY            | Настройки крона для создания уведомлений              |           | 8                                   |
| TN_REPORTS_CHECKER_SCHEDULER_HOUR           | Настройки крона для создания уведомлений              |           | 07                                  |
| TN_REPORTS_CHECKER_SCHEDULER_MINUTES        | Настройки крона для создания уведомлений              |           | 30                                  |
| TN_MISSED_REPORTS_CHECKER_SCHEDULER_DAY     | Настройки крона для создания уведомлений              |           | 10                                  |
| TN_MISSED_REPORTS_CHECKER_SCHEDULER_HOUR    | Настройки крона для создания уведомлений              |           | 07                                  |
| TN_MISSED_REPORTS_CHECKER_SCHEDULER_MINUTES | Настройки крона для создания уведомлений              |           | 30                                  |
| TN_ENABLE_CREATE_NOTIFICATIONS              | Включено ли создание уведомлений                      |           | true                                |
| TN_SEND_EMAIL_NOTIFICATIONS_PERIOD          | Период отправки email уведомлений                     |           | 10m                                 |
| TN_SEND_LIFE_NOTIFICATIONS_PERIOD           | Период отправки life уведомлений                      |           | 10m                                 |
| TN_URL_BASE_NOTIFICATION                    | URL адрес для формирования ссылки на проект/продукт   |           | https://tn-office-dev.tages.dev     |
| TN_BOOKING_VIEW_FRONT_ROUTE_NOTIFICATION    | Фронтовый роут для бронирований                           |           | /booking/[[uuid]]                   |
| TN_ROOM_VIEW_FRONT_ROUTE_NOTIFICATION    | Фронтовый роут для комнат                          |           | /room/[[uuid]]                   |
| TN_BOT_API_URL                              | URL адрес api бота                                    |           |                                     |
| TN_BOT_ID                                   | Идентификатор бота                                    |           |                                     |
| TN_BOT_TOKEN                                | Токен бота                                            |           |                                     |


Примеры заполнения переменных приведены в файле local/vars.sh.