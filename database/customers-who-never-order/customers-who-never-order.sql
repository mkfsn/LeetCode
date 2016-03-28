# Write your MySQL query statement below

Select C.Name From Customers C Where C.Id NOT IN (Select CustomerId From Orders);
