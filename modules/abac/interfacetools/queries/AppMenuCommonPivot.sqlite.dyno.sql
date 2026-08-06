
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
        app_menu_entities_cte(level, recordIndex, unique_id, parent_id,visibility,updated,created,href,icon,label,active_matcher,apply_type) AS (
        select * from (
            SELECT 0, ROW_NUMBER() OVER(ORDER BY app_menu_entities.unique_id) AS recordIndex, app_menu_entities.unique_id, app_menu_entities.parent_id,app_menu_entities.visibility,app_menu_entities.updated,app_menu_entities.created,app_menu_entities.href,app_menu_entities.icon,app_menu_entities.label,app_menu_entities.active_matcher,app_menu_entities.apply_type from app_menu_entities
        )
        UNION ALL
        SELECT app_menu_entities_cte.level+1,app_menu_entities_cte.recordIndex,app_menu_entities.unique_id, app_menu_entities.parent_id,app_menu_entities.visibility,app_menu_entities.updated,app_menu_entities.created,app_menu_entities.href,app_menu_entities.icon,app_menu_entities.label,app_menu_entities.active_matcher,app_menu_entities.apply_type
            FROM app_menu_entities JOIN app_menu_entities_cte ON app_menu_entities.unique_id=app_menu_entities_cte.parent_id
            ORDER BY 2 DESC
        )
    SELECT level, recordIndex, unique_id, parent_id,visibility,updated,created,href,icon,label,active_matcher,apply_type 
    ,ROW_NUMBER() OVER(ORDER BY app_menu_entities_cte.recordIndex)
    FROM app_menu_entities_cte
        order by recordIndex desc)
    select *, total - result.level - 1 as 'depth_reverse' from result
    left join (select count(*) total, recordIndex from result group by recordIndex) v
    on v.recordIndex = result.recordIndex
    order by result.recordIndex, depth_reverse asc
)
group by recordIndex
