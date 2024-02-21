select 
    count(*) total_items
from fb_user_workspace_entities
left join fb_workspace_role_entities on fb_workspace_role_entities.workspace_id = fb_user_workspace_entities.workspace_id
left join fb_role_capabilities on fb_role_capabilities .role_entity_unique_id = fb_workspace_role_entities.role_id
left join fb_capability_entities on fb_capability_entities.unique_id = fb_role_capabilities.capability_entity_unique_id
where fb_user_workspace_entities.user_id = "(userId)"
