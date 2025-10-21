# Практическое задание № 5 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Установить и настроить PostgreSQL локально.
2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
3.	Выполнить параметризованные запросы INSERT и SELECT.
4.	Корректно работать с context, пулом соединений и обработкой ошибок.

Выполнение практического задания.

1.	Установить и настроить PostgreSQL локально.
   Для выполнения задания на сервер был установлен postgresSQL и Go
  	
<img width="558" height="122" alt="Снимок экрана 2025-10-21 021739" src="https://github.com/user-attachments/assets/68d77b0d-e9e2-480b-a9c2-0b5bc2b0a26a" />


2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
  Поссле был выполнен вход в пространство PostgresSQL

<img width="858" height="169" alt="Снимок экрана 2025-10-21 025839" src="https://github.com/user-attachments/assets/c2b401bb-659b-4d85-b44e-3a34eded63b1" />


3.	Выполнить параметризованные запросы INSERT и SELECT.
   Для выполнения задания была созда БД (todo) и таблиц (task), которая была заполнена первой записью

<img width="682" height="385" alt="Снимок экрана 2025-10-21 025936" src="https://github.com/user-attachments/assets/eab4b7e1-c95d-4bcc-8d38-1c76c63ed357" />


  После была выполнена подготовка проекта для выполнения практики
  
<img width="823" height="767" alt="Снимок экрана 2025-10-21 030415" src="https://github.com/user-attachments/assets/723c81d1-fe4c-4ad2-97d2-102d3edad4c3" />


  Затем была сделана структура проекта
  
<img width="577" height="146" alt="Снимок экрана 2025-10-21 030743" src="https://github.com/user-attachments/assets/7e86b5fd-4098-403d-89f1-77e6211b526f" />


  После был написан код в файле для db.gо, в котором будет происходит подключениие к БД и пулу
  
<img width="804" height="719" alt="Снимок экрана 2025-10-21 031040" src="https://github.com/user-attachments/assets/546b3fca-4890-414a-8d32-3a8eac274dbe" />


  После был написан код в файле для repository.go, в котором будут происходить запросы к БД
  
<img width="813" height="1003" alt="Снимок экрана 2025-10-21 031130" src="https://github.com/user-attachments/assets/738db387-b137-4d8b-bb98-eb12166b50e5" />


  После был написан код в файле для main.go, в котором будут происходить запуск проекта
  
<img width="639" height="978" alt="Снимок экрана 2025-10-21 031458" src="https://github.com/user-attachments/assets/22ddc53f-825c-481b-ac22-34432ff3c1e6" />

  Затем для проверки выполнения практики был осуществлен запуск проекта
  Через go run.
  
<img width="514" height="201" alt="Снимок экрана 2025-10-21 031657" src="https://github.com/user-attachments/assets/ed53e68a-b1a0-4c74-a5cf-34657aebb375" />


через psql

<img width="600" height="274" alt="Снимок экрана 2025-10-21 031859" src="https://github.com/user-attachments/assets/956dccef-0f69-4aec-b945-4930f51a1e92" />


# Проверочные задания 
1.	Реализуйте функцию ListDone
  Для реализации функции фильрации задач был написан расшиириин код в repository.go

<img width="708" height="389" alt="Снимок экрана 2025-10-21 105321" src="https://github.com/user-attachments/assets/a2cb4730-bad8-4c8d-aa68-b25d9c26602e" />


2.	Добавьте функцию FindByID
  Для реализации функции поиска в ID был написан расшиириин код в repository.go

<img width="701" height="220" alt="Снимок экрана 2025-10-21 110230" src="https://github.com/user-attachments/assets/1895b66d-9d7e-49df-8b8a-69b6904e7808" />


3.	Сделайте функцию массовой вставки CreateMany
 Для реализации функции массовой вставки был написан расширенный код в repository.go.

<img width="564" height="348" alt="Снимок экрана 2025-10-21 110255" src="https://github.com/user-attachments/assets/d8f0b3c8-1902-428e-b6eb-9b056161b0ca" />


4. Проверка выполнения всех заданий
 Для выполнения строк кода был обновлен файл main.go.

<img width="683" height="865" alt="Снимок экрана 2025-10-21 110629" src="https://github.com/user-attachments/assets/166a01b8-6939-44d4-83bf-48e116a0e558" />


  Для потверждения работоспособности было произведено выполнение проекта.

<img width="731" height="676" alt="Снимок экрана 2025-10-21 114130" src="https://github.com/user-attachments/assets/18843dcc-ff79-485f-97a6-e387142fde93" />


