SELECT
    count(*) total_items
FROM
    `fb_user_role_workspace_entities`
    left join fb_user_entities on fb_user_entities.unique_id = fb_user_role_workspace_entities.user_id
ORDER BY
    fb_user_entities.Created desc
