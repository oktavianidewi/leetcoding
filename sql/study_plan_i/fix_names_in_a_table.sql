Create table If Not Exists Users (user_id int, name varchar(40));
Truncate table Users;
insert into Users (user_id, name) values ('1', 'aLice');
insert into Users (user_id, name) values ('2', 'bOB');

select user_id, CONCAT(UPPER(LEFT(name, 1)), LOWER(SUBSTR(name, 2, LENGTH(name)))) as name
from Users order by 1;