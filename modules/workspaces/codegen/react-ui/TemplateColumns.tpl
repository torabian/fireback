import { enTranslations } from "@/translations/en";
import { TemplateEntityFields } from "src/sdk/xsdk/modules/xmodule/xtypefields";

export const columns = (t: typeof enTranslations) => [
  {
    name: TemplateEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  <% for (let field of fields) { let name = field.name ; %>
    <% if (field.type.includes("string")) {  %>
      {
        name: TemplateEntity.Fields.<%- name %>,
        title: t.templates.<%- name %>,
        width: 100,
      },    
    <% } %>
  <% } %>
 
];
