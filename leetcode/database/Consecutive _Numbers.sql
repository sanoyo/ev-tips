-- https://qiita.com/Go-Noji/items/5baeb790ade4b57126ff
-- カウントして、集計してくれる
SELECT COUNT(Num)
FROM Logs
group by Num

-- https://www.dbonline.jp/mysql/select/index10.html
SELECT COUNT(Num) as ConsecutiveNums
FROM Logs
group by Num
having ConsecutiveNums > 3

-- answer
SELECT DISTINCT
    l1.Num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;