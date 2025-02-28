select DISTINCT user_workspace_entities.workspace_id, user_workspace_entities.user_id,
 role_id, role_capabilities.capability_entity_unique_id capability_id, 
 
    "workspace_restrict" "type"
from user_workspace_entities
left join workspace_entities on 
  user_workspace_entities.workspace_id = workspace_entities.unique_id
  left join workspace_type_entities on
  workspace_type_entities.unique_id = workspace_entities.type_id
 left join role_capabilities on 
      role_capabilities.role_entity_unique_id = workspace_type_entities.role_id
    where user_workspace_entities.user_id = "(userId)"
union

select 
    DISTINCT
    user_workspace_entities.workspace_id,
    user_workspace_entities.user_id,
    workspace_role_entities.role_id,
    role_capabilities.capability_entity_unique_id capability_id,
    "account_restrict" "type"

from user_workspace_entities
left join workspace_role_entities on workspace_role_entities.workspace_id = user_workspace_entities.workspace_id
left join role_capabilities on role_capabilities .role_entity_unique_id = workspace_role_entities.role_id
left join capability_entities on capability_entities.unique_id = role_capabilities.capability_entity_unique_id
    where user_workspace_entities.user_id = "(userId)"


order by "type" desc

-- select 
--     user_workspace_entities.workspace_id,
--     user_workspace_entities.user_id,
--     workspace_role_entities.role_id,
--     role_capabilities.capability_entity_unique_id capability_id

-- from user_workspace_entities
-- left join workspace_role_entities on workspace_role_entities.workspace_id = user_workspace_entities.workspace_id
-- left join role_capabilities on role_capabilities .role_entity_unique_id = workspace_role_entities.role_id
-- left join capability_entities on capability_entities.unique_id = role_capabilities.capability_entity_unique_id
-- where user_workspace_entities.user_id = "(userId)"
