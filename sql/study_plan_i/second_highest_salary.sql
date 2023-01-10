Create table If Not Exists Employee (id int, salary int);
Truncate table Employee;
insert into Employee (id, salary) values ('1', '100');
insert into Employee (id, salary) values ('2', '200');
insert into Employee (id, salary) values ('3', '300');

select max(salary) as SecondHighestSalary from Employee
where salary<(select max(salary) from Employee);

select ifnull(
  (
    select distinct salary
    from Employee
    order by salary desc
    limit 1,1
  ), null
) as SecondHighestSalary