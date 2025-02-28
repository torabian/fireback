
select 
MAX(CASE WHEN depth_reverse == 0 THEN unique_id END) as 'unique_id0',
MAX(CASE WHEN depth_reverse == 0 THEN name END) as 'name0'
,
MAX(CASE WHEN depth_reverse == 1 THEN unique_id END) as 'unique_id1',
MAX(CASE WHEN depth_reverse == 1 THEN name END) as 'name1'
,
MAX(CASE WHEN depth_reverse == 2 THEN unique_id END) as 'unique_id2',
MAX(CASE WHEN depth_reverse == 2 THEN name END) as 'name2'
,
MAX(CASE WHEN depth_reverse == 3 THEN unique_id END) as 'unique_id3',
MAX(CASE WHEN depth_reverse == 3 THEN name END) as 'name3'
,
MAX(CASE WHEN depth_reverse == 4 THEN unique_id END) as 'unique_id4',
MAX(CASE WHEN depth_reverse == 4 THEN name END) as 'name4'
,
MAX(CASE WHEN depth_reverse == 5 THEN unique_id END) as 'unique_id5',
MAX(CASE WHEN depth_reverse == 5 THEN name END) as 'name5'
,
MAX(CASE WHEN depth_reverse == 6 THEN unique_id END) as 'unique_id6',
MAX(CASE WHEN depth_reverse == 6 THEN name END) as 'name6'
from (
    with result as (WITH RECURSIVE
        workspace_entities_cte(level, recordIndex, unique_id, parent_id,visibility,updated,created,description,name) AS (
        select * from (
            SELECT 0, ROW_NUMBER() OVER(ORDER BY workspace_entities.unique_id) AS recordIndex, workspace_entities.unique_id, workspace_entities.parent_id,workspace_entities.visibility,workspace_entities.updated,workspace_entities.created,workspace_entities.description,workspace_entities.name from workspace_entities
        )
        UNION ALL
        SELECT workspace_entities_cte.level+1,workspace_entities_cte.recordIndex,workspace_entities.unique_id, workspace_entities.parent_id,workspace_entities.visibility,workspace_entities.updated,workspace_entities.created,workspace_entities.description,workspace_entities.name
            FROM workspace_entities JOIN workspace_entities_cte ON workspace_entities.unique_id=workspace_entities_cte.parent_id
            ORDER BY 2 DESC
        )
    SELECT level, recordIndex, unique_id, parent_id,visibility,updated,created,description,name 
    ,ROW_NUMBER() OVER(ORDER BY workspace_entities_cte.recordIndex)
    FROM workspace_entities_cte
        order by recordIndex desc)
    select *, total - result.level - 1 as 'depth_reverse' from result
    left join (select count(*) total, recordIndex from result group by recordIndex) v
    on v.recordIndex = result.recordIndex
    order by result.recordIndex, depth_reverse asc
)
group by recordIndex
