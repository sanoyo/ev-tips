select Request_at as Day, round(sum(Status <> 'completed')/count(Id),2) as `Cancellation Rate`
from Trips
where Client_Id not in (select Users_Id from Users where Banned = 'Yes')
    and Driver_Id not in (select Users_Id from Users where Banned = 'Yes')
    and Request_at between '2013-10-01' and '2013-10-03'
group by Day
order by Day
