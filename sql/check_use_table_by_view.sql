
-- @table が view から参照されているかを確認する

set @database := 'my_db';
set @table := 'target_table';

select * from INFORMATION_SCHEMA.VIEWS 
where
view_definition like concat('%`', @database, '`.`', @table, '`%')
;
