select 
    count(*) total_items
from user_workspace_entities
left join workspace_role_entities on workspace_role_entities.workspace_id = user_workspace_entities.workspace_id
left join role_capabilities on role_capabilities .role_entity_unique_id = workspace_role_entities.role_id
left join capability_entities on capability_entities.unique_id = role_capabilities.capability_entity_unique_id
where user_workspace_entities.user_id = "(userId)"
