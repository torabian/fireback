SELECT
    count(*) total_items
FROM
    `user_role_workspace_entities`
    left join user_entities on user_entities.unique_id = user_role_workspace_entities.user_id
ORDER BY
    user_entities.Created desc
