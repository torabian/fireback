SELECT
    *
FROM
    `fb_user_role_workspace_entities`
    left join `fb_workspace_entities` 
    on fb_workspace_entities.unique_id = fb_user_role_workspace_entities.workspace_id
WHERE @internalCondition
ORDER BY
    fb_workspace_entities.Created desc
limit
    @limit offset @offset