select
    count(*) total_items
from
    fb_workspace_entities
where
    parent_id is null