База данных использовалась postgresql, поднималось через докер контейнер.
Скрипт создания базы данных находится в data_base/000001_init.up.sql.
Данные для подключения к бд и порт приложения содержаться в config/config.yml.
http запросы тестировались через программу postman.
Запросы отправлялись в формате json.

1) Создание платежа: отправляется POST запрос. 
URL запроса выглядит следующим образом: locahost:порт_докера/api/item/ 
Тело запроса выглядит слеующим обраом:
{
    "user_id":"2",
	"email":"2@ya.ru",
	"sum": "1000",
	"currensy": "rub"
}
user_id - ID пользователя.
email - email пользователя.
sum - сумма пополнения.
currensy - валюта пополнения.

Данные пользователя(id и email), запись о котром содержиться в базе данных.
В ответе содержиться ID платежа и сообщение "траназакция успешно создана". 
Некоторые платежи получают статус Ошибка.

2) Изменение статуса платежа: отправляется PUT запрос. 
URL запроса выглядит следующим образом: locahost:порт_докера/api/item/
Тело запроса выглядит слеующим обраом:
{
    "id":"2",
	"discription":"Успех"
}
id - ID платежа, запись о котром содержиться в базе данных.
discription - новый статус платежа (Успех или Неуспех).

В случае если платеж нахоится в статусе который можно изменить в ответе будет сообщение "ОК",
в случае если платеж нахоится в статусе который нельзя изменить в ответе будет сообщение "Ошибка"

3) Проверка статуса платжа по его ID: отправляется GET запрос. 
URL запроса выглядит следующим образом locahost:порт_докера/api/item/check
Тело запроса выглядит слеующим обраом:
{
    "id":"2",
}
id - ID платежа, запись о котром содержиться в базе данных.

В ответе бует вся информация по патежу, если нужно сделать только отображение статуса, поменять не трудно.

4) Получение списка всех платежей пользователя по его ID: отправляется GET запрос. 
URL запроса выглядит следующим образом locahost:порт_докера/api/list/id
Тело запроса выглядит слеующим обраом:
{
    "user_id":"2"
}
user_id - ID пользователя, запись о котром содержиться в базе данных.

В ответе будет информаци по всем платажам пользователя с данным ID.

5) Получение списка всех платежей пользователя по его emai: отправляется GET запрос. 
URL запроса выглядит следующим образом locahost:порт_докера/api/item/email
Тело запроса выглядит слеующим обраом:
{
	"email":"2@ya.ru"
}
email - email пользователя, запись о котром содержиться в базе данных.

В ответе будет информаци по всем платажам пользователя с данным email.

6) Отмена платажа по его ID: отправляется PUT запрос. 
URL запроса выглядит следующим образом locahost:порт_докера/api/item/
Тело запроса выглядит слеующим обраом:
{
    "id":"1",
	"discription":"Отменён"
}
id - ID платежа, запись о котром содержиться в базе данных.
discription - слово "Отменён".

В случае если платеж нахоится в статусе который можно именить в ответе бует сообщение "ОК",
в случае если платеж нахоится в статусе который нелья именить в ответе бует сообщение "Ошибка"

Если нужно отдельно реализовать отмену, можно сделать PUT запрос, URL запроса  будет выглядит следующим образом locahost:порт_докера/api/item/cancel
Тело запроса будет выглядеть слеующим обраом:
{
    "id":"1"
}
Всё по механике всё анналогично изменению статуса платежа, только не будет указыватся новый статус платажа, а в sql запросе сразу будет указан статус "Отменён" 