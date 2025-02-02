# HTTP server for working with calendar 
Последняя работа с моего обучения в техношколе Wildberries.Проект представляет собой сервер для работы с календарем, 
изначально были предумостренны исключительно API методы, но я решил самостоятельно доделать пользовательский интерфейс, 
а также сервис для работы с ApacheKafka: `Consumer` и `Producer`.  
Приложение соответствует чистой архитектуре.
## API endpoints: 
Базовый url: `http://localhost:8080/api/`
### - POST/AddUser
**Пример ответа** (`200 OK`):
```json
{
  "User":
    {
      "id":6,
      "key":"aoaoa12",
      "events":null
    }
}
```
#### Примечание
Id генерируется базой данных во избежание коллизий.  
### - POST/AddEvent  
**Пример ответа** (`200 OK`):
```json
    {
      "Event":
        {
          "id":12,
          "date":"2026-11-09",
          "description":"Сделать что-то важное"
        }
    }
  ```
#### Примечание
Id генерируется базой данных во избежание коллизий.
### - POST/UpdateEventDate  
**Пример ответа** (`200 OK`):
```json
    {
      "UpdatedEvent":
       {
         "id":6,
         "date":"2025-01-31",
         "description":"go sleep"
       }
   }
```
### - POST/UpdateEventDescription  
  В качестве ответа аналогично возвращается обновленное событие.
### - POST/delete_event
  **Пример ответа** (`200 OK`):
```json
    {
      "DeletedEvent":
       {
         "id":3,
         "date":"2025-01-16",
         "description":"Сделать что-то невероятно важное!"
       }
   }
```
### - GET/events_for_day 
### - GET/events_for_week 
### - GET/events_for_month  
Каждый из этих GET запросов в случае удачи возвращает список событий  
**Пример ответа** (`200 OK`):
```json
  {
  "Events":[
    {
      "id":8,
      "date":"2025-01-28",
      "description":"проснутся да полить цветы"
    },
    {
      "id":10,
      "date":"2025-01-28",
      "description":"пойти в магаз купить еды"
    },
    {
      "id":11,
      "date":"2025-01-28",
      "description":"придти домой сварить еды"
    },
    {
      "id":12,
      "date":"2025-01-28",
      "description":"покушать "
    }
  ]}
```
## ApacheKafka:
Используется топик TestTopic, всё необходимое можно найти в Dockerfile.  
- `KafkaConsumer`  
Принимает и логирует сообщения из топика.
- `KafkaProducer`  
Логирует и отправляет сообщения в топик.  
### Пример сообщения:
```json
  {
  "id": 123,
  "key": "ahaha000",
      "events": [
      {
          "id": 1,
          "date": "2025-07-20",
          "description": "Встреча с командой"
      },
      {
          "id": 2,
          "date": "2025-07-25",
          "description": "Презентация проекта"
      }]
  }
  ```
#### Примечание:  
Модель данных для сообщений представлена в файле `Infrastucture/kafka/model.json`
## User Interface:  
Базовый url: `http://localhost:8080/ui/`  
Веб приложение для взаимодействия с календарем, реализованы все те же методы представленные в API.  
В качестве frontend части воспользовался bootstrap.
### Примерный результат работы: 
![image](https://github.com/user-attachments/assets/65cabedb-6280-49d4-a7c8-9549599097eb)
## Дополнительно:
Все методы презентации и взаимодействия проходят через middleware для логирования запросов.
Логирование реализовано через библиотеку UberZap, логи записываются в файл app.log.
Также все приложение, включая очередь сообщений и базу данных, контейнеризированно.
 



