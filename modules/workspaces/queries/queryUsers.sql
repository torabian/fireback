select 
    fb_user_entities.first_name,
    fb_user_entities.last_name,
    fb_user_entities.unique_id
FROM
    `fb_user_role_workspace_entities`
    left join fb_user_entities on fb_user_entities.unique_id = fb_user_role_workspace_entities.user_id
WHERE
    fb_user_role_workspace_entities.workspace_id = "(workspaceId)"
ORDER BY
    fb_user_entities.Created desc
limit
    @limit offset @offset
