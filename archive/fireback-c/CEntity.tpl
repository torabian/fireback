#ifndef template_tools_<%- implementation ? 'c' : 'h' %>
#define template_tools_<%- implementation ? 'c' : 'h' %>
#include "cJSON.h"
#include "esp-tools.c"
#include "sqlite3.h"
// inside fireback use this:
// #include <cjson/cJSON.h>
 

char * {{ .e.UnderscoreName }}_list(struct {{ .e.UnderscoreName }}_t *items)
{{ if ne .implementation "skip" }}
{
    int length = *(&items + 1) - items;
    cJSON *body = cJSON_CreateObject();
    cJSON_AddNumberToObject(body, "totalItems", length);
    cJSON_AddNumberToObject(body, "totalAvailableItems", length);
    cJSON *json = cJSON_AddArrayToObject(body, "items");

    int i;
    for (i = 0; i < length; i++)
    {
            cJSON_AddItemToArray(json, {{ .e.UnderscoreName }}_json_object(&items[i]));
    }

    char *res = cJSON_Print(body);
    cJSON_free(body);

    return res;
}
{{ end }}
 
char* {{ .e.UnderscoreName }}_to_sql_insert(struct {{ .e.UnderscoreName }}_t *data)
{{ if ne .implementation "skip" }}
{
    char* query = sqlite3_mprintf(
        {{ .e.UnderscoreName }}_insert_sql,
        <% let count = 0; for (const field of sql.fields) { count ++; %>
            <% if (!cSupportedFields.includes(field.type)) continue %>

            data-><%- field.cName %> == NULL ? <%- field.type === "TEXT" ? '""' : '0' %> : <%- field.type === "TEXT" ? '' : '*' %> data-><%- field.cName %><%- count === sql.fields.length ? '' : ',' %>
        <% } %>
    );

    return query;
};
{{ end }}


static char * {{ .e.UnderscoreName }}_sql_query(sqlite3 *db1, const char * sql, int *count)
{{ if ne .implementation "skip" }}
{

    cJSON *body = cJSON_CreateObject();
    cJSON *json = cJSON_AddArrayToObject(body, "items");

    sqlite3_stmt *stmt;
    sqlite3_prepare_v2(db1, sql, -1, &stmt, NULL);

	printf("Got results:\n");
	
    int record = 0;
    while (sqlite3_step(stmt) != SQLITE_DONE) {
        int i;
        {{ .e.UnderscoreName }}_t dto;
		int num_cols = sqlite3_column_count(stmt);
		
		for (i = 0; i < num_cols; i++)
		{
            char * name = sqlite3_column_name(stmt, i);
            {{ range .e.CompleteFields }}

                if (strcmp(name, "<%- field.name %>") == 0) {
                    {{ if or (eq .Type "string") (eq .Type "enum")}}
                        dto.<%- field.cName %> = (const char *)sqlite3_column_text(stmt, i);
                    {{ end }}
                    {{ if or (eq .Type "int32") (eq .Type "int64") (eq .Type "int") }}

                        // We need to free this some how and I do not know how.
                        double *x = malloc(sizeof(double));
                        *x = (double) sqlite3_column_int(stmt, i);
                        dto.<%- field.cName %> = x;
                    {{ end }}
                }
            {{ end }}
        }
        record++;
        cJSON_AddItemToArray(json, {{ .e.UnderscoreName }}_json_object(&dto));
    }
    *count = (int*)record;
    cJSON_AddNumberToObject(body, "totalItems", *count);
    cJSON_AddNumberToObject(body, "totalAvailableItems", *count);

    char *res = cJSON_Print(body);
    cJSON_free(body);
    return res;
};
{{ end }}



esp_err_t template_query(httpd_req_t *req)
{{ if ne .implementation "skip" }}
{
    int response;

    int count = 0;
    char * result = template_sql_query(dbref, "select * from fb_template_entities limit 1", &count);

    httpd_resp_set_type(req, "application/json");
    response = httpd_resp_send(req, result, HTTPD_RESP_USE_STRLEN);
    return response;
}
{{ end }}
esp_err_t template_wipe(httpd_req_t *req)
{{ if ne .implementation "skip" }}
{
    int response;

    test_sqlite("delete from fb_template_entities where 1 <> 2", dbref);

    httpd_resp_set_type(req, "application/json");
    response = httpd_resp_send(req, "", HTTPD_RESP_USE_STRLEN);
    return response;
}
{{ end }}
esp_err_t template_post(httpd_req_t *req)
{{ if ne .implementation "skip" }}
{
    int response;

    int count = 0;
    char * body = read_body_as_string(req);
    template_t dto = json_template(body);
    char * sql = template_to_sql_insert(&dto);
    test_sqlite(sql , dbref);
    httpd_resp_set_type(req, "application/json");
    response = httpd_resp_send(req, "", HTTPD_RESP_USE_STRLEN);
    return response;
}
{{ end }}
#endif