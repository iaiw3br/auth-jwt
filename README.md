# О проекте

Аутентификация пользователя, благодаря использованию двух токенов.

При логировании администратора, создаётся refresh_token (время жизни 1 час) и access_token (время жизни 1 минута).


В момент любого запроса, кроме логирования или разлогирования, если refresh_token жив, а access_token протух,
тогда оба токена обновляются.

##REST API <a name="restApi"></a>

### Login
**Вы отправляете** ваши учётные данные для входа.

**Вы получаете** refresh, access токены с которыми вы можете совершать действия.

```go
c.SetCookie("access", tokens["access_token"], 60, "/", "localhost", false, true)
c.SetCookie("refresh", tokens["refresh_token"], 3600, "/", "localhost", false, true)
```

Name | Value | Domain | Path | Expire | httpOnly | Secure
--- | --- | --- | --- |--- |--- |--- 
access | eyJhbGci.. | localhost | / | now() + 1 min | true | false
refresh | eyJhbGci.. | localhost | / | now() + 1 hour | true | false

**Request**

```json
POST http://localhost:8080/login

{
  "username": <username>,
  "password": <password>
}
```

**Success Response:**
```json
HTTP 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: access_token=eyJhbG...         
Set-Cookie: refresh_token=eyJhb...         
Date: Tue, 26 Oct 2021 19:13:44 GMT
Content-Length: 15
```

**Failed Response:**
```json
HTTP 401 Unauthorized
Content-Type: application/json; charset=utf-8
Date: Tue, 26 Oct 2021 19:13:44 GMT
Content-Length: 15

{
  "message": "the username or password is not correct"
}
```

### Logout
**Вы отправляете** намерения выйти из приложения.

**Вы получаете** очистку токенов в куки.

**Request**

```json
POST http://localhost:8080/logout

{
  "message": "success logout"
}
```

### Получить данные пользователя (username)
**Вы отправляете** имеющиеся токены в куки.

**Вы получаете** username пользователя.

**Request**

```json
GET http://localhost:8080/me

{
  "message": "<username>"
}
```

**Failed Response:**
```json
HTTP 401 Unauthorized
Content-Type: application/json; charset=utf-8
Date: Tue, 26 Oct 2021 19:13:44 GMT
Content-Length: 15

{
  "message": "invalid incoming data"
}
```

**Failed Response:**
```json
HTTP 401 Unauthorized
Content-Type: application/json; charset=utf-8
Date: Tue, 26 Oct 2021 19:13:44 GMT
Content-Length: 15

{
  "message": "invalid incoming data"
}
```

## Запуск приложения

Требуется:
1. Создать в `config/` файл config.yaml и добавить данные:
```yaml
SECRET_ACCESS: <any access secret>
SECRET_REFRESH: <any refresh secret>
USERNAME: <username>
PASSWORD: <password>
RUN_ADDRESS: "localhost"
PORT: ":8080"
```
2. Запустить файл `cmd/auth/main.go` с помощью команды `go run ./auth/main.go`

## Структура приложения
```
    cmd/auth/main.go    # Файл запуска 
    config/             # Конфиги
    internal
        error           # Выводимые ошибки
        middleware      # Миддлвар аутентификации
        model           # Модели
        route           # Роуты
        server          # Запуск сервера
        service         # Взаимодействие с конфигом
        token           # Токены
    pkg 
        auth            # Функции для работы с аутентификацией   
        user            # Функции для работы с пользователем
    .gitignore
    README.md
```