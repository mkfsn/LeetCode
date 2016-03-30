# Write your MySQL query statement below

Select
    x.day AS Day,
    ROUND(AVG(x.cnt), 2) AS 'Cancellation Rate'
From (
    Select 
        Trips.Request_at AS day,
        ( CASE
            WHEN Trips.Status = 'cancelled_by_client' THEN 1
            WHEN Trips.Status = 'cancelled_by_driver' THEN 1
            ELSE 0
        END ) AS cnt
    From
        Trips, Users
    Where
        Users.Role = 'client'
        and Users.Banned = 'No' 
        and Trips.Client_Id = Users_Id
        and (Trips.Request_at between '2013-10-01' and '2013-10-03')
) x
Group By x.day
