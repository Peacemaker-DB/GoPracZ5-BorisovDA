# Практическое задание № 5 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Установить и настроить PostgreSQL локально.
2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
3.	Выполнить параметризованные запросы INSERT и SELECT.
4.	Корректно работать с context, пулом соединений и обработкой ошибок.

Выполнение практического задания.

1.	Установить и настроить PostgreSQL локально.
   Для выполнения задания на сервер был установлен postgresSQL и Go
  	
<img width="558" height="122" alt="Снимок экрана 2025-10-21 021739" src="https://github.com/user-attachments/assets/6dc4ad09-e753-49c8-bbd3-e91b471172a6" />

2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
  Поссле был выполнен вход в пространство PostgresSQL

<img width="858" height="169" alt="Снимок экрана 2025-10-21 025839" src="https://github.com/user-attachments/assets/df158d14-9e1b-4930-ba8d-b7204f73935d" />

3.	Выполнить параметризованные запросы INSERT и SELECT.
   Для выполнения задания была созда БД (todo) и таблиц (task), которая была заполнена первой записью

<img width="682" height="385" alt="Снимок экрана 2025-10-21 025936" src="https://github.com/user-attachments/assets/a47c32d0-9d14-441b-84de-1d5732d88fea" />

  После была выполнена подготовка проекта для выполнения практики
  
<img width="823" height="767" alt="Снимок экрана 2025-10-21 030415" src="https://github.com/user-attachments/assets/6848b284-ba2e-484b-87d5-1f2d4a41f66e" />

  Затем была сделана структура проекта
  
<img width="577" height="146" alt="Снимок экрана 2025-10-21 030743" src="https://github.com/user-attachments/assets/011a7db1-db3f-4504-bae5-ea19b0b051d0" />

  После был написан код в файле для db.gо, в котором будет происходит подключениие к БД и пулу
  
<img width="804" height="719" alt="Снимок экрана 2025-10-21 031040" src="https://github.com/user-attachments/assets/02891e4d-2919-4865-891d-fd2e0dab0e8e" />

  После был написан код в файле для repository.go, в котором будут происходить запросы к БД
  
<img width="813" height="1003" alt="Снимок экрана 2025-10-21 031130" src="https://github.com/user-attachments/assets/f41e4a89-917f-4d52-89d0-138096ffe71a" />

  После был написан код в файле для main.go, в котором будут происходить запуск проекта
  
<img width="639" height="978" alt="Снимок экрана 2025-10-21 031458" src="https://github.com/user-attachments/assets/0f881097-bfe7-4c30-90b4-386af06782e9" />

  Затем для проверки выполнения практики был осуществлен запуск проекта
  Через go run.
  
<img width="514" height="201" alt="Снимок экрана 2025-10-21 031657" src="https://github.com/user-attachments/assets/6c376aed-d1d7-4384-97c4-35c1d1a2d150" />

через psql

<img width="600" height="274" alt="Снимок экрана 2025-10-21 031859" src="https://github.com/user-attachments/assets/dc73dd1c-5de5-4ab5-8cb0-3142314066ba" />

# Проверочные задания 
1.	Реализуйте функцию ListDone
  Для реализации функции фильрации задач был написан расшиириин код в repository.go

<img width="708" height="389" alt="Снимок экрана 2025-10-21 105321" src="https://github.com/user-attachments/assets/1a67db68-87f7-45e2-8cd8-e33b91fd0ff4" />

3.	Добавьте функцию FindByID
  Для реализации функции поиска в ID был написан расшиириин код в repository.go

<img width="701" height="220" alt="Снимок экрана 2025-10-21 110230" src="https://github.com/user-attachments/assets/707e008d-30b2-493d-a5c7-060aca077a3b" />

4.	Сделайте функцию массовой вставки CreateMany
 Для реализации функции массовой вставки был написан расширенный код в repository.go.

<img width="564" height="348" alt="Снимок экрана 2025-10-21 110255" src="https://github.com/user-attachments/assets/d28ca3a9-8b86-4964-919a-5edd5fd400b4" />

5. Проверка выполнения всех заданий
 Для выполнения строк кода был обновлен файл main.go.

<img width="683" height="865" alt="Снимок экрана 2025-10-21 110629" src="https://github.com/user-attachments/assets/e6253816-6e26-4181-ab5f-8206793d93d1" />

  Для потверждения работоспособности было произведено выполнение проекта.

<img width="731" height="676" alt="Снимок экрана 2025-10-21 114130" src="https://github.com/user-attachments/assets/1878012a-ee57-4cae-900a-c0b3df43f800" />

