-- 他のトランザクションが実行できないようにする
-- https://itsakura.com/sql-select-forupdate
select 項目 from テーブル
where 条件
for update;

select * from SYAIN
where ID = '1'
for update;

-- https://qiita.com/5zm/items/cc33fa9f591892222153
select * from customer where customer_code = '00000001' for update nowait;
