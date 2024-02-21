WITH RECURSIVE
    fb_workspace_entities_cte(level, unique_id, parent_id,visibility,updated,created,description,name) AS (
    select * from (
        SELECT 0, fb_workspace_entities.unique_id, fb_workspace_entities.parent_id,fb_workspace_entities.visibility,fb_workspace_entities.updated,fb_workspace_entities.created,fb_workspace_entities.description,fb_workspace_entities.name from fb_workspace_entities
        where parent_id is null
        (internalCondition)
        limit @limit offset @offset
    )
    UNION ALL
    SELECT fb_workspace_entities_cte.level+1,fb_workspace_entities.unique_id, fb_workspace_entities.parent_id,fb_workspace_entities.visibility,fb_workspace_entities.updated,fb_workspace_entities.created,fb_workspace_entities.description,fb_workspace_entities.name
        FROM fb_workspace_entities JOIN fb_workspace_entities_cte ON fb_workspace_entities.parent_id=fb_workspace_entities_cte.unique_id
        ORDER BY 2 DESC
    )
SELECT DISTINCT
    fb_workspace_entities_cte.level,
    fb_workspace_entities_cte.unique_id,
    fb_workspace_entities_cte.parent_id,fb_workspace_entities_cte.visibility,fb_workspace_entities_cte.updated,fb_workspace_entities_cte.created,fb_workspace_entities_cte.description
,fb_workspace_entities_cte.name
    FROM fb_workspace_entities_cte