select
    count(*) total_items
from user_workspace_entities
left join workspace_role_entities on workspace_role_entities.workspace_id = user_workspace_entities.workspace_id
left join role_entities on role_entities.unique_id = workspace_role_entities.role_id
left join lateral jsonb_array_elements_text(role_entities.capabilities_list_id) as rc(capability_id) on true
left join capability_entities on capability_entities.unique_id = rc.capability_id
where user_workspace_entities.user_id = '(userId)'
