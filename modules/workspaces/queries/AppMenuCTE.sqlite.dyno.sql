WITH RECURSIVE
    fb_app_menu_entities_cte(level, unique_id, parent_id,visibility,updated,created,href,icon,label,active_matcher,apply_type) AS (
    select * from (
        SELECT 0, fb_app_menu_entities.unique_id, fb_app_menu_entities.parent_id,fb_app_menu_entities.visibility,fb_app_menu_entities.updated,fb_app_menu_entities.created,fb_app_menu_entities.href,fb_app_menu_entities.icon,fb_app_menu_entities.label,fb_app_menu_entities.active_matcher,fb_app_menu_entities.apply_type from fb_app_menu_entities
        where parent_id is null
        (internalCondition)
        limit @limit offset @offset
    )
    UNION ALL
    SELECT fb_app_menu_entities_cte.level+1,fb_app_menu_entities.unique_id, fb_app_menu_entities.parent_id,fb_app_menu_entities.visibility,fb_app_menu_entities.updated,fb_app_menu_entities.created,fb_app_menu_entities.href,fb_app_menu_entities.icon,fb_app_menu_entities.label,fb_app_menu_entities.active_matcher,fb_app_menu_entities.apply_type
        FROM fb_app_menu_entities JOIN fb_app_menu_entities_cte ON fb_app_menu_entities.parent_id=fb_app_menu_entities_cte.unique_id
        ORDER BY 2 DESC
    )
SELECT DISTINCT
    fb_app_menu_entities_cte.level,
    fb_app_menu_entities_cte.unique_id,
    fb_app_menu_entities_cte.parent_id,fb_app_menu_entities_cte.visibility,fb_app_menu_entities_cte.updated,fb_app_menu_entities_cte.created,fb_app_menu_entities_cte.href
,fb_app_menu_entities_cte.icon
,fb_app_menu_entity_polyglots.label
,fb_app_menu_entities_cte.active_matcher
,fb_app_menu_entities_cte.apply_type
    FROM fb_app_menu_entities_cte
LEFT JOIN fb_app_menu_entity_polyglots on fb_app_menu_entity_polyglots.linker_id = fb_app_menu_entities_cte.unique_id
and fb_app_menu_entity_polyglots.language_id = '(language)'