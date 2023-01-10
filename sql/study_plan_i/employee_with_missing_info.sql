Create table If Not Exists Employees (employee_id int, name varchar(30));
Create table If Not Exists Salaries (employee_id int, salary int);
Truncate table Employees;
insert into Employees (employee_id, name) values ('2', 'Crew');
insert into Employees (employee_id, name) values ('4', 'Haven');
insert into Employees (employee_id, name) values ('5', 'Kristian');
Truncate table Salaries;
insert into Salaries (employee_id, salary) values ('5', '76071');
insert into Salaries (employee_id, salary) values ('1', '22517');
insert into Salaries (employee_id, salary) values ('4', '63539');


-- select * from Salaries

-- select employee_id from Employees
-- where employee_id not IN (select employee_id from Salaries)
-- union
-- select employee_id from Salaries
-- where employee_id not IN (select employee_id from Employees)
-- order by employee_id ASC

-- optimized 
select employee_id from (
select employee_id
from Employees
union all
select employee_id
from Salaries
) AS t
group by 1
having count(1) = 1
order by employee_id ASC

-- from (
--   select employee_id
-- )