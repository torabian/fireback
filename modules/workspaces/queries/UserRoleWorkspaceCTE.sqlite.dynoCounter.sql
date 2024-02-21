select
    count(*) total_items
from
    fb_user_role_workspace_entities
where
    parent_id is null