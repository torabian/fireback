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
  {{ if .Description }}
  /**
  {{ typescriptComment .Description }}
  */
  {{ end }}
  public {{ .PrivateName }}?: {{ .ComputedType }} | null;

  {{ if eq .Type "one" }}
    {{ if and (ne .PrivateName "user") (ne .PrivateName "workspace") }}
      {{ .PrivateName }}Id?: string | null;
    {{ end }}
  {{ end }}


  {{ if eq .Type "many2many" }}
    {{ .PrivateName }}ListId?: string[] | null;
  {{ end }}

  {{ if or (eq .Type "html") (eq .Type "text") }}
    public {{ .PrivateName }}Excerpt?: string[] | null;
  {{ end }}
  
  {{ if or (eq .Type "daterange") }}
    public {{ .PrivateName }}Start?: string[] | null;
    public {{ .PrivateName }}End?: string[] | null;
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
  {{ $prefix := index . 2 }}
  ...BaseEntity.Fields,
  {{ range $row.CompleteFields }}
    {{ $newPrefix := print $prefix .PrivateName "." }}

    {{ if or (eq .Type "daterange") }}
      {{ .PrivateName }}Start: `{{ $prefix}}{{ .PrivateName }}Start`,
      {{ .PrivateName }}End: `{{ $prefix}}{{ .PrivateName }}End`,
    {{ end }}

    {{ if or (eq .Type "array") }}
      {{ .PrivateName }}$: `{{ $prefix}}{{ .PrivateName }}`,
      {{ .PrivateName }}At: (index: number) => {
        return {
          $: `{{ $prefix}}{{ .PrivateName }}[${index}]`,

          {{ $newPrefix := print $prefix .PrivateName "[${index}]." }}
          {{ template "stringfield" (arr . $root $newPrefix) }}
        };
      },
    {{ else if or (eq .Type "object") (eq .Type "embed") }}
      {{ .PrivateName }}$: '{{ $prefix}}{{ .PrivateName }}',
      {{ .PrivateName }}: {
        {{ template "stringfield" (arr . $root $newPrefix) }}
      },
    {{ else if or (eq .Type "one") (eq .Type "many2many") }}

      {{ if eq .Type "one" }}
        {{ if and (ne .PrivateName "user") (ne .PrivateName "workspace") }}
          {{ .PrivateName }}Id: `{{ $prefix}}{{ .PrivateName }}Id`,
        {{ end }}
      {{ end }}
      {{ if eq .Type "many2many" }}
        {{ .PrivateName }}ListId: `{{ $prefix}}{{ .PrivateName }}ListId`,
      {{ end }}
      
      {{ if or (eq .Type "html") (eq .Type "text") }}
        {{ .PrivateName }}Excerpt: `{{ $prefix}}{{ .PrivateName }}Excerpt`,
      {{ end }}

      {{ .PrivateName }}$: `{{ $prefix}}{{ .PrivateName }}`,

      {{ if ne $root.ObjectName .Target }}
        {{ .PrivateName }}: {{ .Target }}.Fields,
      {{ end }}
    {{ else }}
      {{ .PrivateName }}: `{{ $prefix}}{{ .PrivateName }}`,
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
      {{ if or (eq .Type "array") (eq .Type "object") (eq .Type "embed")  }}

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
    {{ template "stringfield" (arr .e .e "") }}
}
  

{{ end }}

{{ define "actionStringFields" }}
  {{ range . }}

    {{ if or (eq .Type "array") (eq .Type "object") (eq .Type "embed") }}
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

      {{ if or (eq .Type "html") (eq .Type "text") }}
        {{ .PrivateName }}Excerpt: '{{ .PrivateName }}Excerpt',
      {{ end }}

      {{ if or (eq .Type "daterange") }}
        {{ .PrivateName }}Start: '{{ .PrivateName }}Start',
        {{ .PrivateName }}End: '{{ .PrivateName }}End',
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