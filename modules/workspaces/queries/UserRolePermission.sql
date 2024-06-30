select DISTINCT fb_user_workspace_entities.workspace_id, fb_user_workspace_entities.user_id,
 role_id, fb_role_capabilities.capability_entity_unique_id capability_id, 
 
    "workspace_restrict" "type"
from fb_user_workspace_entities
left join fb_workspace_entities on 
  fb_user_workspace_entities.workspace_id = fb_workspace_entities.unique_id
  left join fb_workspace_type_entities on
  fb_workspace_type_entities.unique_id = fb_workspace_entities.type_id
 left join fb_role_capabilities on 
      fb_role_capabilities.role_entity_unique_id = fb_workspace_type_entities.role_id
    where fb_user_workspace_entities.user_id = "(userId)"
union

select 
    DISTINCT
    fb_user_workspace_entities.workspace_id,
    fb_user_workspace_entities.user_id,
    fb_workspace_role_entities.role_id,
    fb_role_capabilities.capability_entity_unique_id capability_id,
    "account_restrict" "type"

from fb_user_workspace_entities
left join fb_workspace_role_entities on fb_workspace_role_entities.workspace_id = fb_user_workspace_entities.workspace_id
left join fb_role_capabilities on fb_role_capabilities .role_entity_unique_id = fb_workspace_role_entities.role_id
left join fb_capability_entities on fb_capability_entities.unique_id = fb_role_capabilities.capability_entity_unique_id
    where fb_user_workspace_entities.user_id = "(userId)"


order by "type" desc

-- select 
--     fb_user_workspace_entities.workspace_id,
--     fb_user_workspace_entities.user_id,
--     fb_workspace_role_entities.role_id,
--     fb_role_capabilities.capability_entity_unique_id capability_id

-- from fb_user_workspace_entities
-- left join fb_workspace_role_entities on fb_workspace_role_entities.workspace_id = fb_user_workspace_entities.workspace_id
-- left join fb_role_capabilities on fb_role_capabilities .role_entity_unique_id = fb_workspace_role_entities.role_id
-- left join fb_capability_entities on fb_capability_entities.unique_id = fb_role_capabilities.capability_entity_unique_id
-- where fb_user_workspace_entities.user_id = "(userId)"
