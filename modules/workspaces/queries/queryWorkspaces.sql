SELECT
    *
FROM
    `user_role_workspace_entities`
    left join `workspace_entities` 
    on workspace_entities.unique_id = user_role_workspace_entities.workspace_id
WHERE @internalCondition
ORDER BY
    workspace_entities.Created desc
limit
    @limit offset @offset