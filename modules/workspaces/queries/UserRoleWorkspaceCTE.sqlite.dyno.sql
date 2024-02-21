WITH RECURSIVE
    fb_user_role_workspace_entities_cte(level, unique_id, parent_id,visibility,updated,created) AS (
    select * from (
        SELECT 0, fb_user_role_workspace_entities.unique_id, fb_user_role_workspace_entities.parent_id,fb_user_role_workspace_entities.visibility,fb_user_role_workspace_entities.updated,fb_user_role_workspace_entities.created from fb_user_role_workspace_entities
        where parent_id is null
        (internalCondition)
        limit @limit offset @offset
    )
    UNION ALL
    SELECT fb_user_role_workspace_entities_cte.level+1,fb_user_role_workspace_entities.unique_id, fb_user_role_workspace_entities.parent_id,fb_user_role_workspace_entities.visibility,fb_user_role_workspace_entities.updated,fb_user_role_workspace_entities.created
        FROM fb_user_role_workspace_entities JOIN fb_user_role_workspace_entities_cte ON fb_user_role_workspace_entities.parent_id=fb_user_role_workspace_entities_cte.unique_id
        ORDER BY 2 DESC
    )
SELECT DISTINCT
    fb_user_role_workspace_entities_cte.level,
    fb_user_role_workspace_entities_cte.unique_id,
    fb_user_role_workspace_entities_cte.parent_id,fb_user_role_workspace_entities_cte.visibility,fb_user_role_workspace_entities_cte.updated,fb_user_role_workspace_entities_cte.created
    FROM fb_user_role_workspace_entities_cte