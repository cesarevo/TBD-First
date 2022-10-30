
## <a name="3!!"></a> PL/PGSQL запросы


1. [Выведите на экран любое сообщение.](#3_1)
2. [Выведите на экран текущую дату.](#3_2)
3. [Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран.](#3_3)
4. [Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.](#3_4)
5. [Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).](#3_5)
6. [Последовательность Коллатца. Берётся любое натуральное число. Если чётное - делим его на 2, если нечётное, то умножаем его на 3 и прибавляем 1. Такие действия выполняются до тех пор, пока не будет получена единица. Гипотеза заключается в том, что какое бы начальное число n не было выбрано, всегда получится 1 на каком-то шаге. Задания: написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.](#3_6)
7. [Числа Люка. Объявляем и присваиваем значение переменной - количество числе Люка. Вывести на экран последовательность чисел. Где L0 = 2, L1 = 1 ; Ln=Ln-1 + Ln-2 (сумма двух предыдущих чисел). Задания: написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел.](#3_7)
8. [Напишите функцию, которая возвращает количество человек родившихся в заданном году.](#3_8)
10. [Напишите функцию, которая возвращает количество человек с заданным цветом глаз.](#3_9)
11. [Напишите функцию, которая возвращает ID самого молодого человека в таблице.](#3_10)
12. [Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного. ИМТ.](#3_11)
13. [Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми.](#3_12)
14. [Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством.](#3_13)
15. [Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека.](#3_14)
16. [Напишите процедуру, которая позволяет актуализировать рост и вес человека.](#3_15)


<br></br>

### <a name="3_1"></a> 1. Выведите на экран любое сообщение.

#### `Запрос`

```SQL
CREATE OR REPLACE FUNCTION smth() returns varchar
AS $$
BEGIN 
	return 'рус en';
END
$$ language plpgsql;

SELECT smth();
```

<br></br>

### <a name="3_2"></a> 2. Выведите на экран текущую дату.

#### `Запрос`
```SQL
create or replace function date_today() returns varchar
as $$
begin
	return now();
end
$$ language plpgsql;

select date_today();
```

<br></br>

### <a name="3_3"></a> 3. Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран.

#### `Запрос`

```SQL
create or replace function math(x int, y int, OUT sum int, OUT multi int, OUT sub int, OUT div float) 
as $$
begin
    sum = x + y;
    multi = x * y;
    sub = x - y;
    div = x/y;
	
end;
$$ language plpgsql;

select * from math(10,10);
```

<br></br>

### <a name="3_4"></a> 4. Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.

#### `Запрос через if`

```SQL
create or replace function with_if(mark_facebook int) returns varchar
as $$
begin
	if mark_facebook = 5 then return 'Отлично';
	elsif mark_facebook = 4 then return 'Хорошо';
	elsif mark_facebook = 3 then return 'Удовлетворительно';
	elsif mark_facebook = 2 then return 'Неуд';
	else return 'введённая оценка не верна';
	end if;
	
end
$$ language plpgsql;

select with_if(1);
```

#### `Запрос через case'

```SQL
create or replace function with_case(mark_facebook int) returns varchar
as $$
begin
	case
		when mark_facebook = 5 then return 'Отлично';
		when mark_facebook = 4 then return 'Хорошо';
		when mark_facebook = 3 then return 'Удовлетворительно';
		when mark_facebook = 2 then return 'Неуд';
		else return 'введённая оценка не верна';
	end case;
end
$$ language plpgsql;

select with_case(2);
```

### <a name="3_5"></a> 5. Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).

#### 'loop'

```SQL

create or replace procedure pow_numbers()
language plpgsql
as $$
declare i int = 20;
begin
	loop
		exit when i > 30;
		raise notice 'number: %',i*i;
		i = i + 1;
	end loop;
end;
$$;
```
#### 'while'
```SQL
do $$
declare 
	i int := 20;
begin
	while i <= 30 loop
		raise notice 'number: %',i*i;
		i = i + 1;
	end loop;
end;
$$;
```
#### 'for'
```SQL

do $$
begin
	for i in 20..30 loop
		raise notice 'number: %',i*i;
	end loop;
end;
$$;
```
### <a name="3_6"></a> 6.Задания: написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.

#### `function`

```SQL
create or replace function collatz(x int) returns int
as $$
begin
	while x!=1 loop
		case 
			when x%2=0 then x = x/2;
			else x = x*3+1;
		end case;
	end loop;
	return x;
		
end
$$ language plpgsql;

select collatz(3);
```

#### 'procedure'
```SQL
CREATE OR REPLACE PROCEDURE collatz_procedure(INOUT x int)
LANGUAGE plpgsql
AS $$
BEGIN
	while x!=1 loop
		CASE 
			WHEN x%2=0 THEN x := x/2;
			else  x := x*3+1;
		END CASE;
	end loop;
END;
$$;

CALL collatz_procedure(5)
```

<br></br>



### <a name="3_7"></a> 7. Числа Люка

#### `Запрос`

```SQL
create or replace function luke_skywalker(n int) returns int
as $$
begin
	if n = 1 then return 2;
	elsif n = 2 then return 1;
	else return luke_skywalker(n - 1) + luke_skywalker(n - 2);
	end if;
	return 0;
end
$$ language plpgsql;


select luke_skywalker(5);
```
#### 'я не очень понял как здесь вывести всю последовательность чисел в числах Люка'
<br></br>

### <a name="3_8"></a> 8. Напишите функцию, которая возвращает количество человек родившихся в заданном году.

#### `Запрос`

```SQL
create or replace function peoples_birth(year int) returns int
as $$
declare counter int;
begin
	select 
		count(*) into counter
	from
		people
	where
		year(people.birth_date) = peoples_birth.year;
	return counter;
end
$$ language plpgsql;

select peoples_birth(1995); 
```

<br></br>


### <a name="3_9"></a> 9. Напишите функцию, которая возвращает количество человек с заданным цветом глаз.

#### `Запрос`

```SQL
create or replace function eyes(color varchar) returns int
as $$
declare counter int;
begin
	select 
		count(*) into counter
	from
		people
	where
		people.eyes = eyes.color;
	return counter;
end
$$ language plpgsql;

select eyes('brown'); 
```

<br></br>

### <a name="3_10"></a> 10. Напишите функцию, которая возвращает ID самого молодого человека в таблице.

#### `Запрос`

```SQL
create or replace function id_return() returns SETOF int
as $$
begin
	RETURN QUERY 
	select 
		people.id
	from
		people
	order by birth_date desc
	limit 1;
	
end
$$ language plpgsql;

select * from id_return(); 
```

<br></br>


<br></br>

### <a name="3_11"></a> 11. Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного.

#### `Запрос`

```SQL
create or replace procedure IMT(imt int)
as $$
declare
	pRT people%rowtype;
begin
	for pRT in select * from people
	loop
		if pRT.weight / (pRT.growth/100)^2 > imt
		then raise notice 'id: %, name: %, surname: %', pRT.id, pRT.name, pRT.surname;
		end if;
	end loop;
end
$$ language plpgsql;

call IMT(0);
```

<br></br>



<br></br>

### <a name="3_12"></a> 12. Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми.

#### `Запрос`

```SQL
create table parent_child(
people_id int references people(id),
child_id int references people(id)
);

insert into parent_child (people_id, child_id)
values (1, 3), (2, 3), (2, 6), (4, 5);
```
<br></br>

<br></br>

### <a name="3_13"></a> 13. Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством.

#### `Запрос`

```SQL
create or replace procedure AddPeople(name varchar, surname varchar, birth_date date,
				      growth real, weight real, eyes varchar, hair varchar,
				      child_id int, parent1_id int, parent2_id int)
as $$
declare
	person_id int;
begin
	insert into people (name, surname, birth_date, growth, weight, eyes, hair)
		values (name, surname, birth_date, growth, weight, eyes, hair) returning id into person_id;
	insert into parent_child (people_id, child_id)
		values (person_id, child_id);
	insert into parent_child (people_id, child_id)
		values (parent1_id, person_id);
	insert into parent_child (people_id, child_id)
		values (parent2_id, person_id);
end
$$ language plpgsql;

call AddPeople('Aleksey', 'Korshunov', '07.05.2002', 178, 81.2, 'green', 'black', 5, 3, 4)
```
<br></br>


<br></br>

### <a name="3_14"></a> 14. Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека

#### `Запрос`

```SQL
begin;
	alter table people
	add time_of_relevance timestamp not null default now();
commit;

```
<br></br>

<br></br>

### <a name="3_15"></a> 15. Напишите процедуру, которая позволяет актуализировать рост и вес человека.

#### `Запрос`

```SQL
create or replace procedure GrowthAndWeight(person_id int, newGrowth real, newWeight real)
as $$
begin
	update people
	set growth = newGrowth, weight = newWeight, actual_time = now()
	where people.id = person_id;
end
$$ language plpgsql;

call GrowthAndWeight(7, 180, 81);

```
<br></br>




