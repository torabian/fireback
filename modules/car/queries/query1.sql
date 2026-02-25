select 
    name,
    field(first_name2, 'int64', 100),
    last_name
from table_1x;