
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