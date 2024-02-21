select
    count(*) total_items
from
    fb_app_menu_entities
where
    parent_id is null