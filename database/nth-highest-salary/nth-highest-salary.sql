CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      
      Select E.Salary
      From ( 
        Select
            r.Salary AS Salary,
            @rownum := @rownum + 1 AS position
        From (Select Salary From Employee Group By Salary) AS r
        Join (Select @rownum := 0) t
        Order By r.Salary Desc
      ) E
      Where E.position = N
  );
END
