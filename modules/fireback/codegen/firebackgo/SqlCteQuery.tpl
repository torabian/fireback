{{"{{"}} if .IsSqlite {{"}}"}}
    {{"{{"}} if .IsCounter {{"}}"}}
    select
        count(*) total_items
    from
        template_entities
    where
        parent_id is null

    {{"{{"}} else {{"}}"}}

    WITH RECURSIVE
        template_entities_cte(level, unique_id, {{ join .e.GetSqlFieldNames "," }}) AS (
        select * from (
            SELECT 0, template_entities.unique_id, {{ join .e.GetSqlFields "," }} from template_entities
            where parent_id is null
            (internalCondition)
            (extraCondition)
            limit @limit offset @offset
        )
        UNION ALL
        SELECT template_entities_cte.level+1,template_entities.unique_id, {{ join .e.GetSqlFields "," }}
            FROM template_entities JOIN template_entities_cte ON template_entities.parent_id=template_entities_cte.unique_id
            ORDER BY 2 DESC

        )
    SELECT DISTINCT
        template_entities_cte.level,
        template_entities_cte.unique_id,
        {{ join .e.GetSqlFieldNamesAfter "," }}
        
        FROM template_entities_cte

    {{ if .e.HasTranslations }}

    LEFT JOIN template_entity_polyglots on template_entity_polyglots.linker_id = template_entities_cte.unique_id
    and template_entity_polyglots.language_id = '(language)'

    {{ end}}
    {{"{{"}} end {{"}}"}}
{{"{{"}} end {{"}}"}}



{{"{{"}} if .IsMysql {{"}}"}}
    {{"{{"}} if .IsCounter {{"}}"}}
    select
        count(*) total_items
    from
        template_entities
    where
        parent_id is null

    {{"{{"}} else }}
        with
            template_entities_cte as (
                select * from template_entities
            )
        select 
            template_entities_cte.unique_id,
            {{ join .e.GetSqlFieldNamesAfter "," }} 
        from template_entities_cte
        {{ if .e.HasTranslations }}
            LEFT JOIN template_entity_polyglots on template_entity_polyglots.linker_id = template_entities_cte.unique_id
            and template_entity_polyglots.language_id = '(language)'
        {{ end}}
    {{"{{"}} end {{"}}"}}
{{"{{"}} end {{"}}"}}

