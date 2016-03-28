# Write your MySQL query statement below


Select 
    x.Name AS Department,
    E.Name AS Employee,
    E.Salary AS Salary
From 
    Employee E,
    (
        Select
            D.Name AS Name,
            QQ.DepartmentId AS DepartmentId,
            QQ.Salary AS Salary
        From
            Department D,
            (
            Select
                Top3.DepartmentId AS DepartmentId,
                Min(Top3.Salary) AS Salary
            From 
                (
                    Select
                        y.DepartmentId AS DepartmentId,
                        y.Salary AS Salary
                    From
                        (
                            Select Salary, DepartmentId,
                                @num    := if(@group = DepartmentId, if(@salary = Salary, @num, @num + 1), 1) as row_number,
                                @group  := DepartmentId,
                                @Salary := Salary
                            From
                                Employee,
                                (Select @group := NULL, @salary := NULL) AA
                            Order By
                                DepartmentId, Salary DESC
                        ) as y
                    Where
                        y.row_number <= 3
                ) AS Top3
            Group By
                Top3.DepartmentId
            ) QQ
        Where
            Salary IS NOT NULL
            and D.Id = QQ.DepartmentId
    ) AS x
Where
    E.DepartmentId = x.DepartmentId
    and E.Salary >= x.Salary

