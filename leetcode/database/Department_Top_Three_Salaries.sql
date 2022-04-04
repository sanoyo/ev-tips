-- 間違い
select
d.Name as Department,
e.Name as Employee,
e.Salary
from
Employee e
JOIN Department d ON e.DepartmentId = d.Id
order by Salary desc
limit 3

-- 回答
SELECT
    d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
FROM
    Employee e1
        JOIN
    Department d ON e1.DepartmentId = d.Id
WHERE
    3 > (SELECT
            COUNT(DISTINCT e2.Salary)
        FROM
            Employee e2
        WHERE
            e2.Salary > e1.Salary
                AND e1.DepartmentId = e2.DepartmentId
        )
;