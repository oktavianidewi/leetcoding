Create table If Not Exists Activity (player_id int, device_id int, event_date date, games_played int);
Truncate table Activity;
insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-01', '5');
insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-05-02', '6');
insert into Activity (player_id, device_id, event_date, games_played) values ('2', '3', '2017-06-25', '1');
insert into Activity (player_id, device_id, event_date, games_played) values ('3', '1', '2016-03-02', '0');
insert into Activity (player_id, device_id, event_date, games_played) values ('3', '4', '2018-07-03', '5');


SELECT player_id, MIN(event_date) as 'first_login' FROM Activity
GROUP BY Player_id;

/*
// window function is not always better
// but better than self join


with ranked as (
  select player_id, event_date as first_login
  rank() over (partition by player_id order by event_date) as rank_event_date
  from Activity
)
select player_id, first_login
from ranked
where rank_event_date = 1

*/