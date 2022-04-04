select
a.name AS customers
from customers AS a 
where a.id not in (
    select customerid from orders
)
