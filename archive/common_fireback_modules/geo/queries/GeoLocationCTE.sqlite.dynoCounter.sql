select
    count(*) total_items
from
    fb_geo_location_entities
where
    parent_id is null