# Write your MySQL query statement below

Select W.Id From Weather W, Weather T Where T.Date = DATE_SUB(W.Date, Interval 1 Day) and W.Temperature > T.Temperature;
