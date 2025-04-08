select 
    user_entities.first_name,
    user_entities.last_name,
    user_entities.unique_id
FROM
    `user_role_workspace_entities`
    left join user_entities on user_entities.unique_id = user_role_workspace_entities.user_id
WHERE
    user_role_workspace_entities.workspace_id = "(workspaceId)"
ORDER BY
    user_entities.Created desc
limit
    @limit offset @offset
