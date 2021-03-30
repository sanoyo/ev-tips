

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