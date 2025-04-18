SELECT
    count(*) total_items
FROM
    `user_role_workspace_entities`
    left join `workspace_entities` on workspace_entities.unique_id = user_role_workspace_entities.workspace_id
