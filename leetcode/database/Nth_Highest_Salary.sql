-- https://github.com/XiancaiTian/Leetcode-database/blob/master/Nth%20Highest%20Salary.sql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE skip INT;
  SET skip = N-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT DISTINCT Salary
      FROM Employee
      ORDER BY Salary DESC
      LIMIT 1 OFFSET skip
  );
END
