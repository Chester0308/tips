
-- クロスジョインで大量 insert

insert into oauth_access_tokens(
   user_id
  ,name
)
select
   @id := @id + 1
  ,concat('name_', @id)
from
    (SELECT @id := 0) AS t,
    users as t1,
    users as t2,
    users as t3
limit 1000
;
