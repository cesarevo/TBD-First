-------------1------------

SELECT name, surname
FROM student
WHERE score BETWEEN 4 AND 4.5


SELECT name, surname
FROM student
WHERE score>=4 and score<=4.5


--------------2--------------
SELECT * 
FROM student
WHERE CAST (student.n_group as varchar)  like '2%';


-------------3---------------
SELECT * 
FROM student
Order by student.n_group DESC, student.name;


-------------4----------------
SELECT * 
FROM student
WHERE score>4
Order by student.score DESC;


-------------5----------------
SELECT * 
FROM student
WHERE score>4
Order by student.score DESC;


-------------6----------------

-------------7----------------
SELECT *
FROM student
WHERE score>4.5
ORDER BY score DESC;

-----------------8--------------
SELECT *
FROM student
WHERE score>4.5
ORDER BY score DESC
LIMIT 5;


-------------------9---------------
SELECT hobby_name,
CASE
WHEN risk>=8 THEN 'очень высокий'
WHEN risk<8 and risk>=6 THEN 'высокий'
WHEN risk<6 and risk>=4 THEN 'средний'
WHEN risk<4 and risk>=2 THEN 'низкий'
WHEN risk<2  THEN 'очень низкий'
END
FROM hobby;


-------------------10---------------
SELECT *
FROM hobby
ORDER BY risk DESC LIMIT 3;
