# Write your MySQL query statement below

Select 
    L.Num
From (
    Select 
        grp, MAX(samecount) AS cnt, Num
    From (
        Select
            (@group := @group + IF(@curnum = t.Num, 0, 1)) grp,
            (@same := IF(@curnum = t.Num, @same + 1, 1)) samecount,
            (@curnum := t.Num) tmp,
            t.Num AS Num
        From
            Logs AS t,
            (Select @group := 0, @same := 0) t2
        Order By id
    ) x
    Group By grp
    Having cnt >= 3
) L
Group by L.Num
