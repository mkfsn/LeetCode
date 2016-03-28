# Write your MySQL query statement below


Select
    D.Name,
    E.Name,
    E.Salary
From
    Department AS D,
    Employee AS E,
    (Select DepartmentId, MAX(Salary) AS Salary From Employee Group By DepartmentId) AS x
Where
    x.DepartmentId = E.DepartmentId
    and E.DepartmentId = D.Id
    and E.Salary = x.Salary
