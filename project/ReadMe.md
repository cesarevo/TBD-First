# Немножко теории

### Объекту можно назначить нового владельца с помощью команды ALTER для соответствующего типа объекта, например:
```SQl
ALTER TABLE имя_таблицы OWNER TO новый_владелец;
```
### Суперпользователь может делать это без ограничений, а обычный пользователь — только если он является одновременно текущим владельцем объекта (или членом роли владельца) и членом новой роли.

### Для назначения прав применяется команда GRANT. Например, если в базе данных есть роль role и таблица table, право на изменение таблицы можно дать этой роли так:
```SQl
GRANT UPDATE ON имя_таблицы TO название_роли;
```

### Если вместо конкретного права написать ALL, роль получит все права, применимые для объекта этого типа.

### Для назначения права всем ролям в системе можно использовать специальное имя «роли»: PUBLIC. Также для упрощения управления ролями, когда в базе данных есть множество пользователей, можно настроить «групповые» роли

<br></br>

```SQL
CREATE TABLE public.cars_type (
cars_type_id serial NOT NULL,
type_name varchar NOT NULL,
CONSTRAINT cars_type_pk PRIMARY KEY (cars_type_id)
);
```

### Назначаю нового владельца
### и даю права доступа
```SQl
ALTER TABLE public.cars_type OWNER TO postgres;
GRANT ALL ON TABLE public.cars_type TO postgres;
```

```SQL
CREATE TABLE public.cars_type (
cars_type_id serial4 NOT NULL,
type_name varchar NOT NULL,
CONSTRAINT cars_type_pk PRIMARY KEY (cars_type_id)
);
```
