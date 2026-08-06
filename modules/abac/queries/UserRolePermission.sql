SELECT DISTINCT
    uwe.workspace_id,
    uwe.user_id,
    wte.role_id AS role_id,
    rc.capability_id AS capability_id,
    'workspace_restrict' AS type,
    re.name AS role_name,
    we.name AS workspace_name
FROM user_workspace_entities uwe
LEFT JOIN workspace_entities we
    ON uwe.workspace_id = we.unique_id
LEFT JOIN workspace_type_entities wte
    ON wte.unique_id = we.type_id
LEFT JOIN role_entities re
    ON re.unique_id = wte.role_id
LEFT JOIN LATERAL jsonb_array_elements_text(re.capabilities_list_id) AS rc(capability_id) ON true

WHERE uwe.user_id = '(userId)'


UNION

SELECT DISTINCT
    uwe.workspace_id,
    uwe.user_id,
    wre.role_id,
    rc.capability_id AS capability_id,
    'account_restrict' AS type,
    re.name AS role_name,
    we.name AS workspace_name
FROM user_workspace_entities uwe
LEFT JOIN workspace_role_entities wre
    ON wre.workspace_id = uwe.workspace_id
LEFT JOIN workspace_entities we
    ON uwe.workspace_id = we.unique_id
LEFT JOIN role_entities re
    ON re.unique_id = wre.role_id
LEFT JOIN LATERAL jsonb_array_elements_text(re.capabilities_list_id) AS rc(capability_id) ON true
WHERE uwe.user_id = '(userId)'
