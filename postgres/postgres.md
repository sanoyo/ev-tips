
## distinct on
```sql
select distinct on (c1) * from t1;
```

https://qiita.com/katsunory/items/4b2faa753cce65d8cba0


distinct したカラムで order byすること

```sql
select distinct city from weather order by city;
```


https://www.postgresql.jp/document/8.0/html/tutorial-select.html

## IN
```sql
select * from country where country in ('Japan','Zambia');
```

```sql
select * from city where country_id in (select country_id from country where country in ('Japan','Zambia'));
```

https://qiita.com/ponsuke0531/items/fdfa8a4b1df8715917cb



## any
下記のサンプルで1, ,2 3のいづれかが真であれば真を返す

```sql
select * from m_school where school_code = any(array['1','2','3']);
```

## all
下記のサンプルで1, ,2 3の全てが真であれば真を返す

```sql
select * from m_school where school_code = all(array['1','2','3']);
```

## Ref
https://www.postgresql.jp/document/9.0/html/functions-comparisons.html
https://qiita.com/choplin/items/9d5e2ff8721fb9509bf8
https://qiita.com/nuko_yokohama/items/db2b221dd32463fe4fa8


## EXISTS
### 結合する場合
exists内のSQLで値が存在しないときは、外側のSQLは実行されない。


```
select * from users
where exists (
	select * from sales
	where sales.id = 2)
	and romaji = 'sato';
```


### 結合しない場合
相関副問合せ。
上との違いは、外側のSQL(主問い合わせ）の項目(users.id)があること。

```
select * from users
where 
    romaji like '%a%'and
    exists (
	select 1 from salesgit st
	where sales.id = users.id);
```

https://itsakura.com/sql-exists
http://www.sql-reference.com/select/subquery_exists.html


## case when
### 単純CASE式
`gender`が男であれば、1として扱うということを意味する。

```
CASE gender
    WHEN '男' THEN 1
    WHEN '女' THEN 2
    ELSE 99
END
```

### 検索CASE式
```
CASE
    WHEN gender = '男' THEN 1
    WHEN gender = '女' THEN 2
    ELSE 99
END


SELECT
    id,
    name,
    CASE WHEN age >= 60 THEN 1 ELSE 0 END AS senior
FROM member
LIMIT 20
```
https://qiita.com/sfp_waterwalker/items/acc7f95f6ab5aa5412f3


## ON CONFLICT
### DO NOTHING
ON CONFLICT DO NOTHINGは代替の動作として、単に行の挿入をしなくなるだけ。
すでにデータが存在する場合は、作

```
ISNERT INTO テーブル名 VALUES (1,1,1) ON CONFLICT DO NOTHING;
```

### ON CONSTRAINT
```
INSERT INTO distributors (did, dname) VALUES (9, 'Antwerp Design')
ON CONFLICT ON CONSTRAINT distributors_pkey DO NOTHING;
```

https://www.postgresql.jp/document/9.5/html/sql-insert.html
https://dev.classmethod.jp/articles/postgresql-9-5-new-function-upsert-use/