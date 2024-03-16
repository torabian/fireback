{{ define  "tsimport" }}

{{ range $key, $value := .imports }}
{{ if $value.Items}}
import {
  {{ range $value.Items }}
    {{ .}},
  {{ end }}

} from "{{ $key}}"
{{ end }}
{{ end }}


{{ end }}



{{ define "matches" }}
  {{ range .Matches }}
    get {{$.Name}}As{{ .PublicName }}(): {{ .PublicName }} | null {
      return this.{{$.Name}} as any;
    }
  {{ end }}
{{ end }}


{{ define "routeUrl" }}

  {{ range .UrlParams}}
    computedUrl = computedUrl.replace("{{ .}}", (query as any)["{{ .}}".replace(":", "")])
  {{ end }}

{{ end }}

// template for the type definition element for each field
{{ define "definitionrow" }}
  {{ range . }}
  public {{ .PrivateName }}?: {{ .ComputedType }} | null;

  {{ if eq .Type "one" }}
    {{ if and (ne .PrivateName "user") (ne .PrivateName "workspace") }}
      {{ .PrivateName }}Id?: string | null;
    {{ end }}
  {{ end }}


  {{ if eq .Type "many2many" }}
    {{ .PrivateName }}ListId?: string[] | null;
  {{ end }}


  {{ if eq .Type "json" }}
      {{ template "matches" .}}
  {{ end }}

  {{ end }}
{{ end }}

// template for creating string schema of the element
{{ define "stringfield" }}
  {{ $row := index . 0 }}
  {{ $root := index . 1 }}
  ...BaseEntity.Fields,
  {{ range $row.CompleteFields }}

    {{ if or (eq .Type "array") (eq .Type "object") }}
      {{ .PrivateName }}$: '{{ .PrivateName }}',
      {{ .PrivateName }}: {
        {{ template "stringfield" (arr . $root) }}
      },
    {{ else if or (eq .Type "one") (eq .Type "many2many") }}

      {{ if eq .Type "one" }}
        {{ if and (ne .PrivateName "user") (ne .PrivateName "workspace") }}
          {{ .PrivateName }}Id: '{{ .PrivateName }}Id',
        {{ end }}
      {{ end }}
      {{ if eq .Type "many2many" }}
        {{ .PrivateName }}ListId: '{{ .PrivateName }}ListId',
      {{ end }}

      {{ .PrivateName }}$: '{{ .PrivateName }}',

      {{ if ne $root.ObjectName .Target}}
        {{ .PrivateName }}: {{ .Target }}.Fields,
      {{ end }}
    {{ else }}
      {{ .PrivateName }}: '{{ .PrivateName }}',
    {{ end }}
    
  {{ end }}
{{ end }}


{{ define "fbdefinition" }}

  public static definition = {{ .e.DefinitionJson }}

{{ end }}
{{ define "staticnavigation" }}


  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ .e.DashedName }}/edit/${uniqueId}`;
      },
      
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ .e.DashedName }}/new`;
      },

      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ .e.DashedName }}/${uniqueId}`;
      },
      
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ .e.DashedPluralName }}`;
      },

      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "{{ .e.DashedName }}/edit/:uniqueId",
      Rcreate: "{{ .e.DashedName }}/new",
      Rsingle: "{{ .e.DashedName }}/:uniqueId",
      Rquery: "{{ .e.DashedPluralName }}",

      {{ range .e.CompleteFields }}
      {{ if or (eq .Type "array") (eq .Type "object") }}

      r{{ .PublicName}}Create: "{{ $.e.DashedName }}/:linkerId/{{ .DashedName}}/new",
      r{{ .PublicName}}Edit: "{{ $.e.DashedName }}/:linkerId/{{ .DashedName}}/edit/:uniqueId",
      edit{{ .PublicName}}(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ $.e.DashedName }}/${linkerId}/{{ .DashedName}}/edit/${uniqueId}`;
      },
      create{{ .PublicName}}(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/{{ $.e.DashedName }}/${linkerId}/{{ .DashedName}}/new`;
      },
      
      {{ end }}
      {{ end }}
  };

{{ end }}

{{ define "staticfield" }}

public static Fields = {
    {{ template "stringfield" (arr .e .e) }}
}
  

{{ end }}

{{ define "actionStringFields" }}
  {{ range . }}

    {{ if or (eq .Type "array") (eq .Type "object") }}
      {{ .PrivateName }}$: '{{ .PrivateName }}',
      {{ .PrivateName }}: {
        {{ template "actionStringFields" .Fields }}
      },
    {{ else if or (eq .Type "one") (eq .Type "many2many") }}

      {{ if eq .Type "one" }}
        {{ if and (ne .PrivateName "user") (ne .PrivateName "workspace") }}
          {{ .PrivateName }}Id: '{{ .PrivateName }}Id',
        {{ end }}
      {{ end }}
      {{ if eq .Type "many2many" }}
        {{ .PrivateName }}ListId: '{{ .PrivateName }}ListId',
      {{ end }}

      {{ .PrivateName }}$: '{{ .PrivateName }}',
      {{ .PrivateName }}: {{ .Target }}.Fields,
    {{ else }}
      {{ .PrivateName }}: '{{ .PrivateName }}',
    {{ end }}
    

  {{ end }}
{{ end }}


{{ define "actionDtoFields" }}
public static Fields = {
    {{ template "actionStringFields" .}}
}
{{ end }}