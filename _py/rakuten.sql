
-- create
CREATE TABLE phones (
  name VARCHAR NOT NULL UNIQUE,
  phone_number INT NOT NULL
);

CREATE TABLE calls (
  id INT NOT NULL,
  caller VARCHAR NOT NULL,
  callee VARCHAR NOT NULL,
  duration INT NOT NULL,
  unique(id)
);

-- insert
insert into phones values ('Jack', 1234);
insert into phones values ('Lena', 3333);
insert into phones values ('Mark', 9999);
insert into phones values ('Anna', 7582);
insert into calls values (25, 1234, 7582, 8);
insert into calls values (7, 9999, 7582, 1);
insert into calls values (18, 9999, 3333, 4);
insert into calls values (2, 7582, 3333, 3);
insert into calls values (3, 3333, 1234, 1);
insert into calls values (21, 3333, 1234, 1);

-- fetch

list_caller

 name | phone_number | duration 
------+--------------+----------
 Jack |         1234 |        8
 Lena |         3333 |        1
 Lena |         3333 |        1
 Anna |         7582 |        3
 Mark |         9999 |        1
 Mark |         9999 |        4

 name | phone_number | duration 
------+--------------+----------
 Jack |         1234 |        1
 Jack |         1234 |        1
 Lena |         3333 |        4
 Lena |         3333 |        3
 Anna |         7582 |        8
 Anna |         7582 |        1
 Mark |         9999 |         


 Example test:   (Second example test.)
Returned value: 
+---------+
|   Ginny |
|    Kate |
| Addison |
+---------+
WRONG ANSWER (row 1: got 'Ginny', expected 'Addison')

Your test case: 
insert into phones values ('Jack', 1234);
insert into phones values ('Lena', 3333);
insert into phones values ('Mark', 9999);
insert into phones values ('Anna', 7582);
insert into calls values (25, 1234, 7582, 8);
insert into calls values (7, 9999, 7582, 1);
insert into calls values (18, 9999, 3333, 4);
insert into calls values (2, 7582, 3333, 3);
insert into calls values (3, 3333, 1234, 1);
insert into calls values (21, 3333, 1234, 1);
