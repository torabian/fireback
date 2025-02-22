WITH RECURSIVE
    fb_geo_location_entities_cte(level, unique_id, parent_id,visibility,updated,created,name,code,type_id,status,flag,official_name) AS (
    select * from (
        SELECT 0, fb_geo_location_entities.unique_id, fb_geo_location_entities.parent_id,fb_geo_location_entities.visibility,fb_geo_location_entities.updated,fb_geo_location_entities.created,fb_geo_location_entities.name,fb_geo_location_entities.code,fb_geo_location_entities.type_id,fb_geo_location_entities.status,fb_geo_location_entities.flag,fb_geo_location_entities.official_name from fb_geo_location_entities
        where parent_id is null
        (internalCondition)
        limit @limit offset @offset
    )
    UNION ALL
    SELECT fb_geo_location_entities_cte.level+1,fb_geo_location_entities.unique_id, fb_geo_location_entities.parent_id,fb_geo_location_entities.visibility,fb_geo_location_entities.updated,fb_geo_location_entities.created,fb_geo_location_entities.name,fb_geo_location_entities.code,fb_geo_location_entities.type_id,fb_geo_location_entities.status,fb_geo_location_entities.flag,fb_geo_location_entities.official_name
        FROM fb_geo_location_entities JOIN fb_geo_location_entities_cte ON fb_geo_location_entities.parent_id=fb_geo_location_entities_cte.unique_id
        ORDER BY 2 DESC
    )
SELECT DISTINCT
    fb_geo_location_entities_cte.level,
    fb_geo_location_entities_cte.unique_id,
    fb_geo_location_entities_cte.parent_id,fb_geo_location_entities_cte.visibility,fb_geo_location_entities_cte.updated,fb_geo_location_entities_cte.created,fb_geo_location_entity_polyglots.name
,fb_geo_location_entities_cte.code
,fb_geo_location_entities_cte.type_id
,fb_geo_location_entities_cte.status
,fb_geo_location_entities_cte.flag
,fb_geo_location_entity_polyglots.official_name
    FROM fb_geo_location_entities_cte
LEFT JOIN fb_geo_location_entity_polyglots on fb_geo_location_entity_polyglots.linker_id = fb_geo_location_entities_cte.unique_id
and fb_geo_location_entity_polyglots.language_id = '(language)'