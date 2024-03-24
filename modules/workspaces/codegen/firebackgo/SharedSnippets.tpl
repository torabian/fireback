{{ define  "goimport" }}

{{ range $key, $value := .goimports }}
{{ if and ($value.Items) ($key) }}
import  "{{ $key}}"
{{ end }}
{{ end }}


{{ end }}


{{ define "golangtype" }}{{ if or (eq .Type "array") (eq .Type "many2many") }} []* {{ end }}{{ if or (eq .Type "object") (eq .Type "one") }} * {{ end }}{{ end }}

{{ define "validaterow" }}{{ if and (.Validate) (ne .Type "one") }} validate:"{{ .Validate }}" {{ end }}{{ end }}

{{ define "gormrow" }} {{ if .ComputedGormTag }} gorm:"{{ .ComputedGormTag }}" {{ end }} {{ end }}
{{ define "sqlrow" }} {{ if .ComputedSqlTag }} sql:"{{ .ComputedSqlTag }}" {{ end }} {{ end }}

// template for the type definition element for each field
{{ define "definitionrow" }}
  {{ $fields := index . 0 }}
  {{ $wsprefix := index . 1 }}
  {{ range $fields }}
    
    {{ if ne .Type "daterange" }}
    {{ .PublicName }} {{ if eq .Type "json" }} *{{ $wsprefix }} {{ end }} {{ template "golangtype" . }} {{ .ComputedType }} `json:"{{ if .Json }}{{.Json}}{{ else }}{{.PrivateName }}{{ end }}" yaml:"{{ if .Yaml }}{{.Yaml}}{{ else }}{{.PrivateName }}{{ end }}" {{ template "validaterow" . }} {{ template "gormrow" . }} {{ template "sqlrow" . }}{{ if .Translate }} translate:"true" {{ end }}`
    {{ end }}

    // Datenano also has a text representation
    {{ if eq .Type "datenano" }}
    {{ .PublicName }}Formatted string `json:"{{ .PrivateName }}Formatted" yaml:"{{ .PrivateName }}Formatted"`
    {{ end }}
    
    {{ if eq .Type "daterange" }}
    // Date range is a complex date storage
    {{ .PublicName }}Start {{ $wsprefix }}XDate `json:"{{ .PrivateName }}Start" yaml:"{{ .PrivateName }}Start"`
    {{ .PublicName }}End {{ $wsprefix }}XDate `json:"{{ .PrivateName }}End" yaml:"{{ .PrivateName }}End"`
    {{ .PublicName }} {{ $wsprefix }}XDateComputed `json:"{{ .PrivateName }}" yaml:"{{ .PrivateName }}" gorm:"-" sql:"-"`
    {{ end }}
    
    {{ if eq .Type "date" }}
    // Date range is a complex date storage
    {{ .PublicName }}DateInfo {{ $wsprefix }}XDateMetaData `json:"{{ .PrivateName }}DateInfo" yaml:"{{ .PrivateName }}DateInfo" sql:"-" gorm:"-"`
    {{ end }}
    
    
    {{ if eq .Type "text" }}
    {{ .PublicName }}Excerpt *string `json:"{{ .PrivateName }}Excerpt" yaml:"{{ .PrivateName }}Excerpt"`
    {{ end }}
    
    {{ if eq .Type "one" }}
        {{ if and (ne .Name "user") (ne .Name "workspace") }}
        {{ .PublicName }}Id *string `json:"{{ .PrivateName }}Id" yaml:"{{ .PrivateName }}Id"{{ if .IdFieldGorm }} gorm:"{{ .IdFieldGorm }}" {{ end }}{{ if .Validate }} validate:"{{ .Validate }}" {{ end }}`
        {{ end }}
    {{ end }}
    
    {{ if eq .Type "many2many" }}
    {{ .PublicName }}ListId []string `json:"{{ .PrivateName }}ListId" yaml:"{{ .PrivateName }}ListId" gorm:"-" sql:"-"`
    {{ end }}
    
    {{ if eq .Type "html" }}
    {{ .PublicName }}Excerpt * string `json:"{{ .PrivateName }}Excerpt" yaml:"{{ .PrivateName }}Excerpt"`
    {{ end }}
    
  {{ end }}
{{ end }}


{{ define "defaultgofields" }}
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"{{ if .GormMap.WorkspaceId }} gorm:"{{ .GormMap.WorkspaceId }}" {{ end }}{{ if eq .DistinctBy "workspace" }} gorm:"unique;not null;" {{ end }}`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"{{ if .GormMap.UserId }} gorm:"{{ .GormMap.UserId }}" {{ end }}`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
{{ end }}

{{ define "polyglottable" }}
  {{ if .e.HasTranslations }}

  type {{ .e.PolyglotName}} struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`

    {{ range .e.CompleteFields }}
      {{ if .Translate }}
        {{.PublicName}} string `yaml:"{{.Name}}" json:"{{.Name}}"`
      {{ end }}
    {{ end }}
  }

  {{ end }}
{{ end }}

{{/* Array or object operations */}}
{{ define "entitychildactions" }}
  {{ range .e.CompleteFields }}
    {{ if or (eq .Type "object") (eq .Type "array")}}
  
func {{ $.e.Upper }}{{ .PublicName }}ActionCreate(
  dto *{{ $.e.Upper }}{{ .PublicName }} ,
  query {{ $.wsprefix }}QueryDSL,
) (*{{ $.e.Upper }}{{ .PublicName }} , *{{ $.wsprefix }}IError) {

    dto.LinkerId = &query.LinkerId

    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = {{ $.wsprefix }}GetDbRef()
    } else {
        dbref = query.Tx
    }

    query.Tx = dbref
    if dto.UniqueId == "" {
        dto.UniqueId = {{ $.wsprefix }}UUID()
    }
    err := dbref.Create(&dto).Error
    if err != nil {
        err := {{ $.wsprefix }}GormErrorToIError(err)
        return dto, err
    }

    return dto, nil
}

func {{ $.e.Upper }}{{ .PublicName }}ActionUpdate(
    query {{ $.wsprefix }}QueryDSL,
    dto *{{ $.e.Upper }}{{ .PublicName }},
) (*{{ $.e.Upper }}{{ .PublicName }}, *{{ $.wsprefix }}IError) {

    dto.LinkerId = &query.LinkerId

    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = {{ $.wsprefix }}GetDbRef()
    } else {
        dbref = query.Tx
    }

    query.Tx = dbref
    err := dbref.UpdateColumns(&dto).Error
    if err != nil {
        err := {{ $.wsprefix }}GormErrorToIError(err)
        return dto, err
    }

    return dto, nil
}

func {{ $.e.Upper }}{{ .PublicName }}ActionGetOne(
    query {{ $.wsprefix }}QueryDSL,
) (*{{ $.e.Upper }}{{ .PublicName }} , *{{ $.wsprefix }}IError) {

    refl := reflect.ValueOf(&{{ $.e.Upper }}{{ .PublicName }} {})
    item, err := {{ $.wsprefix }}GetOneEntity[{{ $.e.Upper }}{{ .PublicName }} ](query, refl)
    return item, err
}

      {{ end }}
  {{ end }}
{{ end }}


{{ define "entityformatting" }}

func entity{{ .e.Upper }}Formatter(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
	if dto == nil {
		return
	}

	{{ range .e.CompleteFields }}
		{{ if or (eq .Type "datenano") }}
			dto.{{ .PublicName }}Formatted = {{ $.wsprefix }}FormatDateBasedOnQuery(dto.{{ .PublicName }}, query)
		{{ end }}

		{{ if or (eq .Type "daterange") }}
			dto.{{ .PublicName }} = {{ $.wsprefix }}ComputeDateRange(dto.{{ .PublicName }}Start, dto.{{ .PublicName }}End , query)
		{{ end }}
		{{ if or (eq .Type "date") }}
			dto.{{ .PublicName }}DateInfo = {{ $.wsprefix }}ComputeXDateMetaData(&dto.{{ .PublicName }}, query)
		{{ end }}
	{{ end }}

	if dto.Created > 0 {
		dto.CreatedFormatted = {{ .wsprefix }}FormatDateBasedOnQuery(dto.Created, query)
	}

	if dto.Updated > 0 {
		dto.CreatedFormatted = {{ .wsprefix }}FormatDateBasedOnQuery(dto.Updated, query)
	}
}

  {{ if .e.PostFormatter}}
  
func {{ .e.Upper }}ItemsPostFormatter(entities []*{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
  for _, entity := range entities {
      {{ .e.PostFormatter }}(entity, query)
  }
} 
  {{ end }}

{{ end }}

{{ define "mockentityrow" }}
  {{ $fields := index . 0}}
  {{ $prefix := index . 1}}
  {{ range $fields}}

    {{ if or (eq .Type "string") (eq .Type "enum")}}
      {{ .PublicName }} : &stringHolder,
    {{ end }}
    
    {{ if or (eq .Type "int32") (eq .Type "int64") (eq .Type "int") }}
      {{ .PublicName }} : &int64Holder,
    {{ end }}
    
    {{ if or (eq .Type "float32") (eq .Type "float64")}}
      {{ .PublicName }} : &float64Holder,
    {{ end }}

  {{ end }}
{{ end }}

{{/* Used for generating mock data, useful for development or stress test */}}
{{ define "mockingentity" }}

func {{ .e.Upper }}MockEntity() *{{ .e.EntityName }} {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &{{ .e.EntityName }}{
		{{ template "mockentityrow" (arr .e.Fields "") }}
	}

	return entity
}

 
func {{ .e.Upper }}ActionSeeder(query {{ .wsprefix }}QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := {{ .e.Upper }}MockEntity()
		_, err := {{ .e.Upper }}ActionCreate(entity, query)
		if err == nil {
			successInsert++
		} else {
			fmt.Println(err)
			failureInsert++
		}

		bar.Add(1)
	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}

{{ end }}

{{ define "getEntityTranslateFields" }}

  {{ range .e.CompleteFields }}
    {{ if .Translate }}
    func (x*{{ $.e.EntityName}}) Get{{ .PublicName }}Translated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.{{ .PublicName }}
          }
        }
      }
      return ""
    }
    {{ end }}
  {{ end }}
{{ end }}


{{ define "entitySeederInit" }}

  func {{ .e.Upper }}ActionSeederInit(query {{ .wsprefix }}QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*{{ .e.EntityName }}{}
    tildaRef := "~"
    _ = tildaRef
    entity := &{{ .e.EntityName }}{

      {{ range .e.CompleteFields }}
        {{ if or (eq .Type  "string") (eq .Type  "enum") (eq .Type "") }}
          {{ .PublicName }}: &tildaRef,
        {{ end }}

        {{ if  eq .Type "object"  }}
          {{ .PublicName }}: &{{ $.e.Upper}}{{ .PublicName }}{},
        {{ end }}

        {{ if  eq .Type "array"  }}
          {{ .UpperPlural }}: []*{{ $.e.Upper}}{{ .PublicName }}{{"{{}}"}},
        {{ end }}

        {{ if  eq .Type "many2many"  }}
          {{ .PublicName }}ListId: []string{"~"},
          {{ .PublicName }}: []*{{ .TargetWithModule }}{{"{{}}"}},
        {{ end }}
      {{ end }}
    }

    data = append(data, entity)

    if format == "yml" || format == "yaml" {
      body, err = yaml.Marshal(data)
      if err != nil {
        log.Fatal(err)
      }
    }
    if format == "json" {
      body, err = json.MarshalIndent(data, "", "  ")
      if err != nil {
        log.Fatal(err)
      }

      file = strings.Replace(file, ".yml", ".json", -1)
    }

    os.WriteFile(file, body, 0644)
  }
{{ end }}


{{ define "entityAssociationCreate" }}
  func {{ .e.Upper }}AssociationCreate(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) error {

  {{ range .e.CompleteFields }}
    {{ if or (eq .Type "many2many") }}
      {
        if dto.{{ .PublicName }}ListId != nil && len(dto.{{ .PublicName }}ListId) > 0 {
          var items []{{ .TargetWithModule }}
          err := query.Tx.Where(dto.{{ .PublicName }}ListId).Find(&items).Error
          if err != nil {
              return err
          }

          err = query.Tx.Model(dto).Association("{{ .PublicName }}").Replace(items)
          if err != nil {
              return err
          }
        }
      }
    {{ end }}
  {{ end }}
    return nil
  }
{{ end }}


{{ define "entityRelationContentCreation" }}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func {{ .e.Upper }}RelationContentCreate(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) error {

  {{ range .e.CompleteFields }}
    {{ if and (eq .Type  "one") (eq .AllowCreate  true) (ne .Name  "") }}

  {
    if dto.{{ .PublicName }} != nil {
      dt, err := {{ .TargetWithModuleWithoutEntity}}ActionCreate(dto.{{ .PublicName }}, query);
      if err != nil {
        return err;
      }
      dto.{{ .PublicName }} = dt;
    }
  }
  {{ end }}
      
  {{ if and (eq .Type  "many2many") (eq .AllowCreate  true) (ne .Name  "") }}
  {
    if dto.{{ .PublicName }} != nil {
      

      dt, err := {{ .TargetWithModuleWithoutEntity}}ActionBatchCreateFn(dto.{{ .PublicName }}, query);
      if err != nil {
        return err;
      }
      dto.{{ .PublicName }} = dt;
    }
  }
  {{ end }}

{{ end }}
return nil
}

{{ end }}


{{ define "relationContentUpdate" }}
func {{ .e.Upper }}RelationContentUpdate(dto *{{ .e.EntityName}}, query {{ .wsprefix }}QueryDSL) error {

  {{ range .e.CompleteFields }}
    {{ if and (eq .Type  "one") (eq .AllowCreate  true) (ne .Name  "") }}
		{
			if dto.{{ .PublicName }} != nil {
			
				dt, err := {{ .TargetWithModuleWithoutEntity }}ActionUpdate(query, dto.{{ .PublicName }});
				if err != nil {
					return err;
				}
				dto.{{ .PublicName }} = dt;
			}
		}
		{{ end }}
    {{ if and (eq .Type "many2many") (eq .AllowCreate  true) (ne .Name  "") }}
		{
			if dto.{{ .PublicName }} != nil {

				cleanQuery := query
				for _, item := range dto.{{ .PublicName }} {
					cleanQuery.Query = "unique_id = " + item.UniqueId
					{{ .TargetWithModuleWithoutEntity }}ActionRemove(cleanQuery)
				}
	
 
				dt, err := {{ .TargetWithModule }}ActionBatchCreateFn(dto.{{ .PublicName }}, query);
				if err != nil {
					return err;
				}

				query.Tx.
					Model(&{{ $.e.EntityName}}{UniqueId: dto.UniqueId}).
					Association("{{ .PublicName }}").
					Replace(dt)
				
				dto.{{ .PublicName }} = dt;
			}
		}
		{{ end }}
	{{ end }}
	return nil
}
{{ end }}


{{ define "polyglot" }}
func {{ .e.Upper }}PolyglotCreateHandler(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
	if dto == nil {
		return
	}

  {{ if .e.HasTranslations}}
    {{ .wsprefix }}PolyglotCreateHandler(dto, &{{ .e.EntityName }}Polyglot{}, query)
  {{ end }}
}
{{ end }}

{{ define "entityValidator" }}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func {{ .e.Upper }}Validator(dto *{{ .e.EntityName }}, isPatch bool) *{{ .wsprefix }}IError {
    err := {{ .wsprefix }}CommonStructValidatorPointer(dto, isPatch)

    {{ range .e.CompleteFields }}
      {{ if  eq .Type "array"  }}
        if dto != nil && dto.{{ .UpperPlural }} != nil {
          {{ $.wsprefix }}AppendSliceErrors(dto.{{ .UpperPlural }}, isPatch, "{{ .Name}}", err)
        }
      {{ end }}
    {{ end }}
    return err
  }
{{ end }}

{{ define "entitySanitize" }}
func {{ .e.Upper }}EntityPreSanitize(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

	{{ range .e.CompleteFields }}
		{{ if  eq .Type "html"  }}

			if (dto.{{ .PublicName }} != nil ) {
          {{ .PublicName }} := *dto.{{ .PublicName }}
          {{ .PublicName }}Excerpt := stripPolicy.Sanitize(*dto.{{ .PublicName }})
          {{ if ne .Unsafe true }}
            {{ .PublicName }} = ugcPolicy.Sanitize({{ .PublicName }})
            {{ .PublicName }}Excerpt = stripPolicy.Sanitize({{ .PublicName }}Excerpt)
          {{ end }}
          
        {{ .PublicName }}ExcerptSize, {{ .PublicName }}ExcerptSizeExists := {{ $.e.EntityName }}MetaConfig["{{ .PublicName }}ExcerptSize"]
        if {{ .PublicName }}ExcerptSizeExists {
          {{ .PublicName }}Excerpt = {{ $.wsprefix }}PickFirstNWords({{ .PublicName }}Excerpt, int({{ .PublicName }}ExcerptSize))
        } else {
          {{ .PublicName }}Excerpt = {{ $.wsprefix }}PickFirstNWords({{ .PublicName }}Excerpt, 30)
        }
        
        dto.{{ .PublicName }}Excerpt = &{{ .PublicName }}Excerpt
        dto.{{ .PublicName }} = &{{ .PublicName }}
      }
	    {{ end }}
	{{ end }}
}

{{ end }}



{{ define "beforecreatedtree" }}
  {{ $fields := index . 0 }}
  {{ $depth := index . 1 }}
  {{ $nextDepth := inc $depth }}
  {{ $fx := index . 2}}
  {{ $wsprefix := index . 3 }}
  {{ range $fields }}
    {{ $nextFx := (fx .PublicName $depth)}}
    {{ if  eq .Type "array"  }}
      if dto.{{ $fx }}{{ .PublicName }} != nil && len(dto.{{ $fx }}{{ .PublicName }}) > 0 {
        for index{{$depth}} := range dto.{{ $fx }}{{ .PublicName }} {
          if (dto.{{ $fx }}{{ .PublicName }}[index{{$depth}}].UniqueId == "") {
            dto.{{ $fx }}{{ .PublicName }}[index{{$depth}}].UniqueId = {{ $wsprefix }}UUID()
          }
          {{ template "beforecreatedtree" (arr .Fields $nextDepth $nextFx $wsprefix ) }}
        }
    }
    {{ end }} 
    {{ if  eq .Type "object"  }}
        if dto.{{ $fx }}{{ .PublicName }} != nil {
          dto.{{ $fx }}{{ .PublicName }}.UniqueId = {{ $wsprefix }}UUID()
        }
    {{ end }} 
  {{ end }}
{{ end }}

{{ define "entityBeforeCreateActions" }}
  func {{ .e.Upper }}EntityBeforeCreateAppend(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = {{ .wsprefix }}UUID()
    }
    
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId

    {{ .e.Upper }}RecursiveAddUniqueId(dto, query)
  }

  func {{ .e.Upper }}RecursiveAddUniqueId(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
    {{ template "beforecreatedtree" (arr .e.CompleteFields 0 "" .wsprefix) }}
  }
{{ end }}

{{ define "batchActionCreate" }}

func {{ .e.Upper}}ActionBatchCreateFn(dtos []*{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) ([]*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*{{ .e.EntityName }}{}
		for _, item := range dtos {
			s, err := {{ .e.Upper}}ActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
			
		}
		return items, nil
	}

	return dtos, nil;
}

{{ end }}

{{ define "entityActionCreate" }}

func {{ .e.Upper }}ActionCreateFn(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {

  {{ if .e.PrependCreateScript }}
    {{ .e.PrependCreateScript }}
  {{ end }}

	// 1. Validate always
	if iError := {{ .e.Upper }}Validator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	{{ .e.EntityName }}PreSanitize(dto, query)

	
	// 2. Append the necessary information about user, workspace
	{{ .e.EntityName }}BeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	{{ .e.Upper }}PolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	{{ .e.Upper }}RelationContentCreate(dto, query)

	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = {{ .wsprefix }}GetDbRef()
	} else {
		dbref = query.Tx
	}

	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := {{ .wsprefix }}GormErrorToIError(err)
		return dto, err
	}

	// 5. Create sub entities, objects or arrays, association to other entities
	{{ .e.Upper }}AssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire({{ .e.EventCreated }}, event.M{
		"entity":   dto,
		"entityKey": {{ .wsprefix }}GetTypeString(&{{ .e.EntityName }}{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})

	return dto, nil
}
{{ end }}


{{ define "entityActionGetAndQuery"}}
  func {{ .e.Upper }}ActionGetOne(query {{ .wsprefix }}QueryDSL) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    refl := reflect.ValueOf(&{{ .e.EntityName }}{})
    item, err := {{ .wsprefix }}GetOneEntity[{{ .e.EntityName }}](query, refl)

    {{ if .e.PostFormatter}}
		  {{ .e.PostFormatter}}(item, query)
		{{ end }}

    entity{{ .e.Upper }}Formatter(item, query)
    return item, err
  }
  {{ if ne .e.NoQuery true }}
  func {{ .e.Upper}}ActionQuery(query {{ .wsprefix }}QueryDSL) ([]*{{ .e.EntityName }}, *{{ .wsprefix }}QueryResultMeta, error) {
    refl := reflect.ValueOf(&{{ .e.EntityName }}{})
    items, meta, err := {{ .wsprefix }}QueryEntitiesPointer[{{ .e.EntityName }}](query, refl)

    {{ if .e.PostFormatter}}
      {{.e.Upper }}ItemsPostFormatter(items, query)
    {{ end }}
    
    for _, item := range items {
      entity{{ .e.Upper }}Formatter(item, query)
    }

    
    return items, meta, err
  }
  {{ end }}
{{ end }}

{{ define "queriesAndPivot" }}

  {{ if .e.HasExtendedQuer }}

  func {{ .e.Upper }}ActionExtendedQuery(query {{ .wsprefix }}QueryDSL) ([]*{{ .wsprefix }}PivotResult, *{{ .wsprefix }}QueryResultMeta, error) {

      items, meta, err := {{ .wsprefix }}UnsafeQuerySqlFromFs[{{ .wsprefix }}PivotResult](
        &queries.QueriesFs, "{{ .e.Upper }}Extended.sqlite.dyno", query,
      )

      return items, meta, err
  }
    
  {{ end }}

  {{ if .e.Cte }}
  func (dto *{{ .e.EntityName }}) Size() int {
    var size int = len(dto.Children)
    for _, c := range dto.Children {
      size += c.Size()
    }
    return size
  }

  func (dto *{{ .e.EntityName }}) Add(nodes ...*{{ .e.EntityName }}) bool {
    var size = dto.Size()
    for _, n := range nodes {
      if n.ParentId != nil && *n.ParentId == dto.UniqueId {
        dto.Children = append(dto.Children, n)
      } else {
        for _, c := range dto.Children {
          if c.Add(n) {
            break
          }
        }
      }
    }
    return dto.Size() == size+len(nodes)
  }

  func {{ .e.Upper }}ActionCommonPivotQuery(query {{ .wsprefix }}QueryDSL) ([]*{{ .wsprefix }}PivotResult, *{{ .wsprefix }}QueryResultMeta, error) {

    items, meta, err := {{ .wsprefix }}UnsafeQuerySqlFromFs[{{ .wsprefix }}PivotResult](
      &queries.QueriesFs, "{{ .e.Upper }}CommonPivot.sqlite.dyno", query,
    )

    return items, meta, err
  }

  func {{ .e.Upper }}ActionCteQuery(query {{ .wsprefix }}QueryDSL) ([]*{{ .e.EntityName }}, *{{ .wsprefix }}QueryResultMeta, error) {

    items, meta, err := {{ .wsprefix }}UnsafeQuerySqlFromFs[{{ .e.EntityName }}](
      &queries.QueriesFs, "{{ .e.Upper }}CTE.sqlite.dyno", query,
    )


    {{ if .e.PostFormatter}}
      {{ .e.PostFormatter }}(item, query)
      {{ end }}

    for _, item := range items {
      entity{{ .e.Upper }}Formatter(item, query)
    }


    var tree []*{{ .e.EntityName }}

    for _, item := range items {
      if item.ParentId == nil {
        root := item
        root.Add(items...)
        tree = append(tree, root)
      }
    }

    return tree, meta, err
  }
  {{ end }}

{{ end }}

{{ define "eventsAndMeta" }}

var {{ .e.AllUpper }}_EVENT_CREATED = "{{ .e.Name }}.created"
var {{ .e.AllUpper }}_EVENT_UPDATED = "{{ .e.Name }}.updated"
var {{ .e.AllUpper }}_EVENT_DELETED = "{{ .e.Name }}.deleted"

var {{ .e.AllUpper }}_EVENTS = []string{
	{{ .e.AllUpper }}_EVENT_CREATED,
	{{ .e.AllUpper }}_EVENT_UPDATED,
	{{ .e.AllUpper }}_EVENT_DELETED,
}

type {{ .e.Upper }}FieldMap struct {
 
	{{ range .e.CompleteFields }}
		{{ .PublicName }} {{ $.wsprefix }}TranslatedString `yaml:"{{ .Name }}"`
	{{ end }}
}

var {{ .e.EntityName }}MetaConfig map[string]int64 = map[string]int64{

    {{ range .e.CompleteFields }}
        {{ if or (eq .Type "html") (eq .Type "text")}}
            "{{ .PublicName }}ExcerptSize": {{ .ComputedExcerptSize }},
        {{ end }}
    {{ end }}
}

var {{ .e.EntityName }}JsonSchema = {{ .wsprefix }}ExtractEntityFields(reflect.ValueOf(&{{ .e.EntityName }}{}))

{{ end }}


{{ define "entityUpdateAction" }}
  func {{ .e.Upper}}ActionUpdateFn(query {{ .wsprefix }}QueryDSL, fields *{{ .e.EntityName }}) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    if fields == nil {
      return nil, {{ .wsprefix }}CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }


    {{ if .e.PrependUpdateScript }}
      {{ .e.PrependUpdateScript }}
    {{ end }}


    // 1. Validate always
    if iError := {{ .e.Upper }}Validator(fields, true); iError != nil {
      return nil, iError
    }

  
    // Let's not add this. I am not sure of the consequences
    // {{ .e.Upper }}RecursiveAddUniqueId(fields, query)


    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = {{ .wsprefix }}GetDbRef()

      var item *{{ .e.EntityName }}
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *{{ .wsprefix }}IError
        item, err = {{ .e.Upper }}UpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }

      })
      return item, {{ .wsprefix }}CastToIError(vf)
    } else {
      dbref = query.Tx
      return {{ .e.Upper }}UpdateExec(dbref, query, fields)
    }

  }

{{ end }}

{{ define "entityDeleteEntireChildrenRec" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}
  {{ $chained := index . 2 }}

  {{ range $fields }}

  {{ if or (eq .Type "object") (eq .Type "array") }}

  if dto{{ $chained }}{{ .PublicName }} != nil {
    q := query.Tx.
      Model(&dto{{ $chained }}{{ .PublicName }}).
      Where(&{{ $prefix }}{{ .PublicName }}{LinkerId: &dto{{ $chained }}UniqueId }).
      Delete(&{{ $prefix }}{{ .PublicName }}{})

    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
    {{ $newPrefix := print $prefix .PublicName  }}
    {{ $newChained := print $chained .PublicName "."   }}
    {{ template "entityDeleteEntireChildrenRec" (arr .CompleteFields $newPrefix $newChained)}}

  {{ end }}
 

  {{ end }}

{{ end }}

{{ define "entityDeleteEntireChildren" }}
func {{ .e.Upper}}DeleteEntireChildren(query {{ .wsprefix }}QueryDSL, dto *{{.e.EntityName }}) (*{{ .wsprefix }}IError) {
  {{ template "entityDeleteEntireChildrenRec" (arr .e.CompleteFields .e.Upper ".") }} 
  return nil
}

{{ end }}

{{ define "entityUpdateExec" }}

  func {{ .e.Upper }}UpdateExec(dbref *gorm.DB, query {{ .wsprefix }}QueryDSL, fields *{{.e.EntityName }}) (*{{.e.EntityName }}, *{{ .wsprefix }}IError) {
    uniqueId := fields.UniqueId

    query.TriggerEventName = {{ .e.EventUpdated }}

    {{.e.EntityName }}PreSanitize(fields, query)
    var item {{.e.EntityName }}
    q := dbref.
      Where(&{{.e.EntityName }}{UniqueId: uniqueId}).
      FirstOrCreate(&item)


    err := q.UpdateColumns(fields).Error
    if err != nil {

      return nil, {{ .wsprefix }}GormErrorToIError(err)
    }

    query.Tx = dbref
    {{ .e.Upper }}RelationContentUpdate(fields, query)

    {{ .e.Upper }}PolyglotCreateHandler(fields, query)

    if ero := {{ .e.Upper}}DeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }


    {{ range .e.CompleteFields }}
        {{ if or (eq .Type "object") }}
   
        if fields.{{ .PublicName }} != nil {

          linkerId := uniqueId

          q := dbref.
            Model(&item.{{ .PublicName }}).
            Where(&{{ $.e.Upper }}{{ .PublicName }}{LinkerId: &linkerId}).
            UpdateColumns(fields.{{ .PublicName }})

          err := q.Error
          if err != nil {
            return &item, {{ $.wsprefix }}GormErrorToIError(err)
          }

          if q.RowsAffected == 0 {
            fields.{{ .PublicName }}.UniqueId = {{ $.wsprefix }}UUID()
            fields.{{ .PublicName }}.LinkerId = &linkerId
            err := dbref.
              Model(&item.{{ .PublicName }}).Create(fields.{{ .PublicName }}).Error

            if err != nil {
              return &item, {{ $.wsprefix }}GormErrorToIError(err)
            }
          }
        }

      {{ end }}
    {{ end }}

    // @meta(update has many)

    {{ range .e.CompleteFields }}
      {{ if or (eq .Type "many2many") }}

        if fields.{{ .PublicName }}ListId  != nil {
          var items []{{.TargetWithModule}}

          if len(fields.{{ .PublicName }}ListId ) > 0 {
            dbref.
              Where(&fields.{{ .PublicName }}ListId ).
              Find(&items)
          }

          dbref.
            Model(&{{$.e.EntityName }}{UniqueId: uniqueId}).
            Association("{{ .PublicName }}").
            Replace(&items)
        }
      {{ end }}
    {{ end }}


    
    {{ $entityName := .e.Upper }}

    {{ range .e.CompleteFields }}
      {{ if or (eq .Type "array") }} 
    
      if fields.{{ .PublicName }} != nil {
       linkerId := uniqueId;
      
        {{ $m := .PublicName }}
        {{ range .Fields }}

          {{ if or (eq .Type "array") }}
          {

            items := []*{{ $entityName }}{{ $m }}{}
            
            dbref.
            Where(&{{ $entityName }}{{ $m }}{LinkerId: &linkerId}).
            Find(&items)
            
        
            for _, item := range items {
              
              {{ range .Fields }}
                {{ if or (eq .Type "object") }}

                  if item3.<%- toUpper(c) %> != nil {
                    item3.<%- toUpper(c) %>.UniqueId = {{ .wsprefix }}UUID()
                  }

                
                {{ end }}
              {{ end }}

              dbref.
              Where(&{{ $entityName }}{{ $m }}{{ .PublicName }} {LinkerId: &item.UniqueId}).
              Delete(&{{ $entityName }}{{ $m }}{{ .PublicName }} {})
            }
          }
          {{ end }}
          
        {{ end }}
           
        dbref.
          Where(&{{ $entityName }}{{ .PublicName }} {LinkerId: &linkerId}).
          Delete(&{{ $entityName }}{{ .PublicName }} {})
  
        for _, newItem := range fields.{{ .PublicName }} {

          {{ range .Fields }}
            {{ if or (eq .Type "object") }}
              if newItem.{{ .PublicName }} != nil {
                newItem.{{ .PublicName }}.UniqueId =  {{ $.wsprefix }}UUID()
              }
            {{ end }}
          {{ end }}

          newItem.UniqueId = {{ $.wsprefix }}UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    

      {{ end }}
    {{ end }}

  
    err = dbref.
      Preload(clause.Associations).
      Where(&{{.e.EntityName }}{UniqueId: uniqueId}).
      First(&item).Error

    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })

    if err != nil {
      return &item, {{ .wsprefix }}GormErrorToIError(err)
    }

    return &item, nil
  }

{{ end }}

{{ define "entityRemoveAndCleanActions" }}


var {{ .e.Upper }}WipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire {{ .e.TemplatesLower }} ",
	Action: func(c *cli.Context) error {
  
		query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
      ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_DELETE},
    })
		count, _ := {{ .e.Upper }}ActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

func {{ .e.Upper }}ActionRemove(query {{ .wsprefix }}QueryDSL) (int64, *{{ .wsprefix }}IError) {
	refl := reflect.ValueOf(&{{ .e.EntityName }}{})
	query.ActionRequires = []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper}}_DELETE}
	return {{ .wsprefix }}RemoveEntity[{{ .e.EntityName }}](query, refl)
}

func {{ .e.Upper }}ActionWipeClean(query {{ .wsprefix }}QueryDSL) (int64, error) {
 
	var err error;
	var count int64 = 0;
	
	  {{ range .e.CompleteFields }}
        {{ if or (eq .Type "object") (eq .Type "array")}}
			{
				subCount, subErr := {{ $.wsprefix }}WipeCleanEntity[{{ $.e.Upper }}{{ .PublicName }}]()
				if (subErr != nil) {
					fmt.Println("Error while wiping '{{ $.e.Upper }}{{ .PublicName }}'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
		{{ end }}
    {{ end }}
	
	{
		subCount, subErr := {{ .wsprefix }}WipeCleanEntity[{{ .e.EntityName }}]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping '{{ .e.EntityName }}'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

{{ end }}

{{ define "entityBulkUpdate" }}

  func {{ .e.Upper}}ActionBulkUpdate(
    query {{ .wsprefix }}QueryDSL, dto *{{ .wsprefix }}BulkRecordRequest[{{ .e.EntityName }}]) (
    *{{ .wsprefix }}BulkRecordRequest[{{ .e.EntityName }}], *{{ .wsprefix }}IError,
  ) {
    result := []*{{ .e.EntityName }}{}
    err := {{ .wsprefix }}GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := {{ .e.Upper}}ActionUpdate(query, record)

        if err != nil {
          return err
        } else {
          result = append(result, item)
        }

      }

      return nil
    })

    if err == nil {
      return dto, nil
    }

    return nil, err.(*{{ .wsprefix }}IError)
  }
{{ end }}

{{ define "entityDistinctOperations" }}

  {{ if or (eq .e.DistinctBy "user") (eq .e.DistinctBy "workspace")}}
  func {{ .e.Upper }}DistinctActionUpdate(
    query {{ .wsprefix }}QueryDSL,
    fields *{{ .e.EntityName }},
  ) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    query.UniqueId = query.UserId
    entity, err := {{ .e.Upper }}ActionGetOne(query)

    if err != nil || entity.UniqueId == "" {
      fields.UniqueId = query.UserId
      return {{ .e.Upper }}ActionCreateFn(fields, query)
    } else {

      fields.UniqueId = query.UniqueId
      return {{ .e.Upper }}ActionUpdateFn(query, fields)
    }
  }

  func {{ .e.Upper }}DistinctActionGetOne(
    query {{ .wsprefix }}QueryDSL,
  ) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    query.UniqueId = query.UserId
    entity, err := {{ .e.Upper }}ActionGetOne(query)

    if err != nil && err.HttpCode == 404 {
      return &{{ .e.EntityName }}{}, nil
    }

    return entity, err
  }
  {{ end }}

{{ end }}

{{ define "entityExtensions" }}


func (x *{{ .e.EntityName }}) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

{{ end }}

{{ define "entityMeta" }}

var {{ .e.EntityName}}Meta = {{ .wsprefix }}TableMetaData{
	EntityName:    "{{.e.Upper}}",
	ExportKey:    "{{.e.DashedPluralName}}",
	TableNameInDb: "fb_{{.e.AllLower}}_entities",
	EntityObject:  &{{ .e.EntityName}}{},
	ExportStream: {{ .e.Upper }}ActionExportT,
	ImportQuery: {{ .e.Upper }}ActionImport,
}

{{ end }}


{{ define "entityImportExport" }}

func {{ .e.Upper }}ActionExport(
	query {{ .wsprefix }}QueryDSL,
) (chan []byte, *{{ .wsprefix }}IError) {
	return {{ .wsprefix }}YamlExporterChannel[{{ .e.EntityName}}](query, {{ .e.Upper }}ActionQuery, {{ .e.Upper }}PreloadRelations)
}


func {{ .e.Upper }}ActionExportT(
	query {{ .wsprefix }}QueryDSL,
) (chan []interface{}, *{{ .wsprefix }}IError) {
	return {{ .wsprefix }}YamlExporterChannelT[{{ .e.EntityName}}](query, {{ .e.Upper }}ActionQuery, {{ .e.Upper }}PreloadRelations)
}


func {{ .e.Upper }}ActionImport(
	dto interface{}, query {{ .wsprefix }}QueryDSL,
) *{{ .wsprefix }}IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content {{ .e.EntityName}}
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return {{ .wsprefix }}CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
 
	json.Unmarshal(cx, &content)
 
	_, err := {{ .e.Upper }}ActionCreate(&content, query)

	return err
}

{{ end }}



{{ define "entityInteractiveCliFlag" }}
  {{ $fields := index . 0}}
  {{ $prefix := index . 1}}

  {{ range $fields }}

  {{ if or (eq .Type "string") (eq .Type "enum") (eq .Type "") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
		Usage:    "{{ .ComputedCliDescription}}",
		Type: "string",
	},
	{{ end }}
  {{ if or (eq .Type "int64") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
		Usage:    "{{ .Name}}",
		Type: "int64",
	},
	{{ end }}
  {{ if or (eq .Type "float64") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
		Usage:    "{{ .Name}}",
		Type: "float64",
	},
	{{ end }}
  {{ if or (eq .Type "bool") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
		Usage:    "{{ .Name}}",
		Type: "bool",
	},
	{{ end }}
{{ end }}

{{ end }}

{{ define "dtoCliFlag" }}
  {{ $fields := index . 0}}
  {{ $prefix := index . 1}}

  {{ range $fields }}
    {{ if or (eq .Type "object")}}
      {{ template "entityCommonCliFlag" (arr .Fields $prefix)}}
    {{ end }}

    {{ if or (eq .Type "string") (eq .Type "enum") (eq .Type "text") (eq .Type "html") (eq .Type "")}}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .ComputedCliDescription}}",
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "daterange") }}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-start",
      Required: {{ .IsRequired }},
      Usage:    "{{ .ComputedCliDescription}}",
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-end",
      Required: {{ .IsRequired }},
      Usage:    "{{ .ComputedCliDescription}}",
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "date") }}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .ComputedCliDescription}}",
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}

    {{ if or (eq .Type "int64")}}
    &cli.Int64Flag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .Name }}",
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
 
    {{ if or (eq .Type "float64")}}
    &cli.Float64Flag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .Name }}",
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
    
    {{ if or (eq .Type "bool")}}
    &cli.BoolFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .Name }}",
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "one")}}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-id",
      Required: {{ .IsRequired }},
      Usage:    "{{ .Name }}",
    },
    {{ end }}
    
    {{ if or (eq .Type "array") (eq .Type "many2many")}}
    &cli.StringSliceFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    "{{ .Name }}",
    },
    {{ end }}

  {{ end }}
{{ end }}

{{ define "entityCommonCliFlag" }}
  {{ $fields := index . 0}}
  {{ $prefix := index . 1}}

  &cli.StringFlag{
    Name:     "{{ $prefix }}wid",
    Required: false,
    Usage:    "Provide workspace id, if you want to change the data workspace",
  },
  &cli.StringFlag{
    Name:     "{{ $prefix }}uid",
    Required: false,
    Usage:    "uniqueId (primary key)",
  },
  &cli.StringFlag{
    Name:     "{{ $prefix }}pid",
    Required: false,
    Usage:    " Parent record id of the same type",
  },

  {{ template "dtoCliFlag" (arr $fields $prefix)}}

{{ end }}

{{ define "cliFlags" }}
var {{ .e.Upper }}CommonCliFlags = []cli.Flag{
  {{ template "entityCommonCliFlag" (arr .e.CompleteFields "") }}
}

var {{ .e.Upper }}CommonInteractiveCliFlags = []{{ .wsprefix }}CliInteractiveFlag{
  {{ template "entityInteractiveCliFlag" (arr .e.CompleteFields "")}}
}

var {{ .e.Upper }}CommonCliFlagsOptional = []cli.Flag{
  {{ template "entityCommonCliFlag" (arr .e.CompleteFields "") }}
}
{{ end }}

{{ define "entityCliCommands" }}

  var {{ .e.Upper }}CreateCmd cli.Command = {{.e.AllUpper}}_ACTION_POST_ONE.ToCli()

  var {{ .e.Upper }}CreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
      })

      entity := &{{ .e.EntityName }}{}

      for _, item := range {{ .e.Upper }}CommonInteractiveCliFlags {

        if !item.Required && c.Bool("all") == false {
          continue
        }

        result := {{ .wsprefix }}AskForInput(item.Name, "")

        {{ .wsprefix }}SetFieldString(entity, item.StructField, result)

      }

      if entity, err := {{ .e.Upper }}ActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {

        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }

  var {{ .e.Upper }}UpdateCmd cli.Command = cli.Command{

    Name:    "update",
    Aliases: []string{"u"},
    Flags: {{ .e.Upper }}CommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {

      query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},
      })

      entity := Cast{{ .e.Upper }}FromCli(c)

      if entity, err := {{ .e.Upper }}ActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {

        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }

      return nil
    },
  }

{{ end }}



{{ define "recursiveGetEnums" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}

  {{ range $fields }}

    {{ if or (eq .Type "enum")}}
      
var {{$prefix}}{{ .PublicName}} = new{{$prefix}}{{ .PublicName}}()

func new{{$prefix}}{{ .PublicName}}() *x{{$prefix}}{{ .PublicName}} {
	return &x{{$prefix}}{{ .PublicName}}{
    {{ range .OfType }}
      {{ .KeyUpper }}: "{{ .Key }}",
    {{ end }}
	}
}

type x{{$prefix}}{{ .PublicName}} struct {
	{{ range .OfType }}
    {{ .KeyUpper }} string
  {{ end }}
}

	  {{ end }}
  {{ end }}
{{ end }}


{{ define "entityCliCastRecursive" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}

  {{ range $fields }}

    {{ if or (eq .Type "string") (eq .Type "enum") (eq .Type "html") (eq .Type "text") (eq .Type "") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }} = &value
      }
	  {{ end }}
    {{ if or (eq .Type "one") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}-id") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}-id")
        template.{{ .PublicName }}Id = &value
      }
	  {{ end }}
    {{ if or (eq .Type "daterange") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}-start") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}-start")
        template.{{ .PublicName }}Start.Scan(value)
      }
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}-end") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}-end")
        template.{{ .PublicName }}End.Scan(value)
      }
	  {{ end }}
    {{ if or (eq .Type "date") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }}.Scan(value)
      }
	  {{ end }}
    {{ if or (eq .Type "many2many") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.String("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }}ListId = strings.Split(value, ",")
      }
	  {{ end }}
  {{ end }}
{{ end }}

{{ define "entityCastFromCli" }}

func (x* {{ .e.ObjectName }}) FromCli(c *cli.Context) *{{ .e.ObjectName }} {
	return Cast{{ .e.Upper }}FromCli(c)
}

func Cast{{ .e.Upper }}FromCli (c *cli.Context) *{{ .e.ObjectName }} {
	template := &{{ .e.ObjectName }}{}

	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}

	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	
	{{ template "entityCliCastRecursive" (arr .e.CompleteFields "")}}

	return template
}

{{ end }}

{{ define "dtoCastFromCli" }}

func Cast{{ .e.Upper }}FromCli (c *cli.Context) *{{ .e.ObjectName }} {
	template := &{{ .e.ObjectName }}{}

	{{ template "entityCliCastRecursive" (arr .e.CompleteFields "")}}

	return template
}

{{ end }}

{{ define "entityMockAndSeeders" }}

  func {{ .e.Upper }}SyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{},
      {{ .e.Upper }}ActionCreate,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }

  {{ if .hasSeeders }}
  func {{ .e.Upper }}SyncSeeders() {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{WorkspaceId: {{ .wsprefix }}USER_SYSTEM},
      {{ .e.Upper }}ActionCreate,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  {{ end }}

  {{ if .hasMocks }}
  func {{ .e.Upper }}ImportMocks() {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{},
      {{ .e.Upper }}ActionCreate,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  {{ end }}

  func {{ .e.Upper }}WriteQueryMock(ctx {{ .wsprefix }}MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := {{ .wsprefix }}QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := {{ .e.Upper }}ActionQuery(f)
      result := {{ .wsprefix }}QueryEntitySuccessResult(f, items, count)
      {{ .wsprefix }}WriteMockDataToFile(lang, "", "{{ .e.Upper }}", result)
    }
  }
{{ end }}

{{ define "entityCliImportExportCmd" }}

var {{ .e.Upper }}ImportExportCommands = []cli.Command{
	{

		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
		},
		Action: func(c *cli.Context) error {
			query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
      })
			{{ .e.Upper }}ActionSeeder(query, c.Int("count"))

			return nil
		},
	},
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The address of file you want the csv be exported to",
				Value: "{{ .e.Template }}-seeder.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
      query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
      })

			{{ .e.Upper }}ActionSeederInit(query, c.String("file"), c.String("format"))
			return nil
		},
	},
	{
		Name:    "validate",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Validates an import file, such as yaml, json, csv, and gives some insights how the after import it would look like",
				Value: "{{ .e.Template }}-seeder-{{ .e.Template }}.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of {{ .e.DashedPluralName }}, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]{{ .e.EntityName }}{}
			{{ .wsprefix }}ReadYamlFile(c.String("file"), data)

			fmt.Println(data)
			return nil
		},
	},
	{{ if .hasSeeders }}
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := {{ .wsprefix }}GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {

				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}

			return nil
		},
	},
	cli.Command{
		Name:  "sync",
		Usage: "Tries to sync the embedded content into the database, the list could be seen by 'list' command",
		Action: func(c *cli.Context) error {

			{{ .wsprefix }}CommonCliImportEmbedCmd(c,
				{{ .e.Upper }}ActionCreate,
				reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
				&seeders.ViewsFs,
			)

			return nil
		},
	},
	{{ end }}
	{{ if .hasMocks }}
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := {{ .wsprefix }}GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
					fmt.Println(err.Error())
				} else {

					f, _ := json.MarshalIndent(entity, "", "  ")
					fmt.Println(string(f))
				}

				return nil
			},
		},
		cli.Command{
			Name:  "msync",
			Usage: "Tries to sync mocks into the system",
			Action: func(c *cli.Context) error {

				{{ .wsprefix }}CommonCliImportEmbedCmd(c,
					{{ .e.Upper }}ActionCreate,
					reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
					&mocks.ViewsFs,
				)

				return nil
			},
		},
	{{ end }}
  {{ if .hasMetas }}
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append({{ .wsprefix }}CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
	
			{{ .wsprefix }}CommonCliExportCmd(c,
				{{ .e.Upper }}ActionQuery,
				reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"{{ .e.Upper }}FieldMap.yml",
				{{ .e.Upper }}PreloadRelations,
			)
	
			return nil
		},
	},
  {{ end }}
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				{{ .wsprefix }}CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			{{ .e.Upper }}CommonCliFlagsOptional...,
		),

		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
	
			{{ .wsprefix }}CommonCliImportCmdAuthorized(c,
				{{ .e.Upper }}ActionCreate,
				reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
				c.String("file"),
        &{{ .wsprefix }}SecurityModel{
					ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
				},
        func() {{ .e.EntityName }} {
					v := Cast{{ .e.Upper }}FromCli(c)
					return *v
				},
			)
	
			return nil
		},
	},
}

{{ end }}

{{ define "entityCliActionsCmd" }}

    var {{ .e.Upper }}CliCommands []cli.Command = []cli.Command{

      {{ .wsprefix }}GetCommonQuery2({{ .e.Upper }}ActionQuery, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
      }),
      {{ .wsprefix }}GetCommonTableQuery(reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(), {{ .e.Upper }}ActionQuery),
    {{ if ne .e.Access "read" }}

          {{ .e.Upper }}CreateCmd,
          {{ .e.Upper }}UpdateCmd,
          {{ .e.Upper }}CreateInteractiveCmd,
          {{ .e.Upper }}WipeCmd,
          {{ .wsprefix }}GetCommonRemoveQuery(reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(), {{ .e.Upper }}ActionRemove),

          {{ if .e.HasExtendedQuer }}
              {{ .wsprefix }}GetCommonExtendedQuery({{ .e.Upper }}ActionExtendedQuery),
          {{ end }}

          {{ if .e.Cte}}
              {{ .wsprefix }}GetCommonCteQuery({{ .e.Upper }}ActionCteQuery),
              {{ .wsprefix }}GetCommonPivotQuery({{ .e.Upper }}ActionCommonPivotQuery),
          {{ end }}
    
    {{ end }}
  }

  func {{ .e.Upper }}CliFn() cli.Command {
    {{ .e.Upper }}CliCommands = append({{ .e.Upper }}CliCommands, {{ .e.Upper }}ImportExportCommands...)

    return cli.Command{
      Name:        "{{ .e.ComputedCliName }}",
      {{ if .e.CliShort }}
      ShortName:   "{{ .e.CliShort }}",
      {{ end }}
      Description: "{{ .e.Upper }}s module actions (sample module to handle complex entities)",
      Usage:       "{{ .e.ComputedCliDescription }}",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: {{ .e.Upper }}CliCommands,
    }
  }

{{ end }}

{{ define "entityHttp" }}

{{ if ne .e.Access "read" }}
var {{.e.AllUpper}}_ACTION_POST_ONE = {{ .wsprefix }}Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new {{ .e.Name }}",
    Flags: {{ .e.Upper }}CommonCliFlags,
    Method: "POST",
    Url:    "/{{ .e.Template }}",
    SecurityModel: &{{ .wsprefix }}SecurityModel{
      {{ if ne $.e.QueryScope "public" }}
      ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
      {{ end }}
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        {{ .wsprefix }}HttpPostEntity(c, {{ .e.Upper }}ActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *{{ .wsprefix }}SecurityModel) error {
      result, err := {{ .wsprefix }}CliPostEntity(c, {{ .e.Upper }}ActionCreate, security)
      {{ .wsprefix }}HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: {{ .e.Upper }}ActionCreate,
    Format: "POST_ONE",
    RequestEntity: &{{ .e.EntityName }}{},
    ResponseEntity: &{{ .e.EntityName }}{},
  }
{{ end }}
  /**
  *	Override this function on {{ .e.EntityName }}Http.go,
  *	In order to add your own http
  **/
  var Append{{ .e.Upper }}Router = func(r *[]{{ .wsprefix }}Module2Action) {}
 
  func Get{{ .e.Upper }}Module2Actions() []{{ .wsprefix }}Module2Action {

    routes := []{{ .wsprefix }}Module2Action{
      {{ if .e.Cte }}
      {
        Method: "GET",
        Url:    "/cte-{{ .e.DashedPluralName }}",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpQueryEntity(c, {{ .e.Upper }}ActionCteQuery)
          },
        },
        Format: "QUERY",
        Action: {{ .e.Upper }}ActionCteQuery,
        ResponseEntity: &[]{{ .e.EntityName }}{},
      },
      {{ end }}

       {
        Method: "GET",
        Url:    "/{{ .e.DashedPluralName }}",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpQueryEntity(c, {{ .e.Upper }}ActionQuery)
          },

        },
        Format: "QUERY",
        Action: {{ .e.Upper }}ActionQuery,
        ResponseEntity: &[]{{ .e.EntityName }}{},
      },
      {
        Method: "GET",
        Url:    "/{{ .e.DashedPluralName }}/export",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpStreamFileChannel(c, {{ .e.Upper }}ActionExport)
          },
        },
        Format: "QUERY",
        Action: {{ .e.Upper }}ActionExport,
        ResponseEntity: &[]{{ .e.EntityName }}{},
      },
      {
        Method: "GET",
        Url:    "/{{ .e.Template }}/:uniqueId",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpGetEntity(c, {{ .e.Upper }}ActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: {{ .e.Upper }}ActionGetOne,
        ResponseEntity: &{{ .e.EntityName }}{},
      },

      {{ if ne .e.Access "read" }}
      {{.e.AllUpper}}_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: {{ .e.Upper }}CommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/{{ .e.Template }}",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpUpdateEntity(c, {{ .e.Upper }}ActionUpdate)
          },
        },
        Action: {{ .e.Upper }}ActionUpdate,
        RequestEntity: &{{ .e.EntityName }}{},
        Format: "PATCH_ONE",
        ResponseEntity: &{{ .e.EntityName }}{},
      },
      {
        Method: "PATCH",
        Url:    "/{{ .e.DashedPluralName }}",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpUpdateEntities(c, {{ .e.Upper }}ActionBulkUpdate)
          },
        },
        Action: {{ .e.Upper }}ActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &{{ .wsprefix }}BulkRecordRequest[{{ .e.EntityName }}]{},
        ResponseEntity: &{{ .wsprefix }}BulkRecordRequest[{{ .e.EntityName }}]{},
      },
      {
        Method: "DELETE",
        Url:    "/{{ .e.Template }}",
        Format: "DELETE_DSL",
        SecurityModel: &{{ .wsprefix }}SecurityModel{
          {{ if ne $.e.QueryScope "public" }}
          ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_DELETE},
          {{ end }}
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            {{ .wsprefix }}HttpRemoveEntity(c, {{ .e.Upper }}ActionRemove)
          },
        },
        Action: {{ .e.Upper }}ActionRemove,
        RequestEntity: &{{ .wsprefix }}DeleteRequest{},
        ResponseEntity: &{{ .wsprefix }}DeleteResponse{},
        TargetEntity: &{{ .e.EntityName }}{},
      },

        {{ if or (eq .e.DistinctBy "user") (eq .e.DistinctBy "workspace")}}
          {
            Method: "PATCH",
            Url:    "/{{ .e.Template }}/distinct",
            SecurityModel: &{{ .wsprefix }}SecurityModel{
              {{ if ne $.e.QueryScope "public" }}
              ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE_DISTINCT_{{ .e.DistinctByAllUpper}}},
              {{ end }}
            },
            Handlers: []gin.HandlerFunc{
              func (c *gin.Context) {
                {{ .wsprefix }}HttpUpdateEntity(c, {{ .e.Upper }}DistinctActionUpdate)
              },
            },
            Action: {{ .e.Upper }}DistinctActionUpdate,
            Format: "PATCH_ONE",
            RequestEntity: &{{ .e.EntityName }}{},
            ResponseEntity: &{{ .e.EntityName }}{},
          },
          {
            Method: "GET",
            Url:    "/{{ .e.Template }}/distinct",
            SecurityModel: &{{ .wsprefix }}SecurityModel{
              {{ if ne $.e.QueryScope "public" }}
              ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_GET_DISTINCT_{{ .e.DistinctByAllUpper}}},
              {{ end }}
            },
            Handlers: []gin.HandlerFunc{
              func (c *gin.Context) {
                {{ .wsprefix }}HttpGetEntity(c, {{ .e.Upper }}DistinctActionGetOne)
              },
            },
            Action: {{ .e.Upper }}DistinctActionGetOne,
            Format: "GET_ONE",
            ResponseEntity: &{{ .e.EntityName }}{},
          },
        {{ end }}

      {{ end }}

      {{ range .e.CompleteFields }}
        {{ if or (eq .Type "object") (eq .Type "array")}}
          {
            Method: "PATCH",
            Url:    "/{{ $.e.Template }}/:linkerId/{{ .DashedName }}/:uniqueId",
            SecurityModel: &{{ $.wsprefix }}SecurityModel{
              {{ if ne $.e.QueryScope "public" }}
              ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_UPDATE},
              {{ end }}
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                {{ $.wsprefix }}HttpUpdateEntity(c, {{ $.e.Upper }}{{ .PublicName }}ActionUpdate)
              },
            },
            Action: {{ $.e.Upper }}{{ .PublicName }}ActionUpdate,
            Format: "PATCH_ONE",
            RequestEntity: &{{ $.e.Upper }}{{ .PublicName }}{},
            ResponseEntity: &{{ $.e.Upper }}{{ .PublicName }}{},
          },
          {
            Method: "GET",
            Url:    "/{{ $.e.Template }}/{{ .DashedName }}/:linkerId/:uniqueId",
            SecurityModel: &{{ $.wsprefix }}SecurityModel{
              {{ if ne $.e.QueryScope "public" }}
              ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_QUERY},
              {{ end }}
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                {{ $.wsprefix }}HttpGetEntity(c, {{ $.e.Upper }}{{ .PublicName }}ActionGetOne)
              },
            },
            Action: {{ $.e.Upper }}{{ .PublicName }}ActionGetOne,
            Format: "GET_ONE",
            ResponseEntity: &{{ $.e.Upper }}{{ .PublicName }}{},
          },
          {
            Method: "POST",
            Url:    "/{{ $.e.Template }}/:linkerId/{{ .DashedName }}",
            SecurityModel: &{{ $.wsprefix }}SecurityModel{
              {{ if ne $.e.QueryScope "public" }}
              ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_CREATE},
              {{ end }}
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                {{ $.wsprefix }}HttpPostEntity(c, {{ $.e.Upper }}{{ .PublicName }}ActionCreate)
              },
            },
            Action: {{ $.e.Upper }}{{ .PublicName }}ActionCreate,
            Format: "POST_ONE",
            RequestEntity: &{{ $.e.Upper }}{{ .PublicName }}{},
            ResponseEntity: &{{ $.e.Upper }}{{ .PublicName }}{},
          },

        {{ end }}
      {{ end }}
    }
   
    // Append user defined functions
    Append{{ .e.Upper }}Router(&routes)

    return routes

  }

  func Create{{ .e.Upper }}Router(r *gin.Engine) []{{ .wsprefix }}Module2Action {

    httpRoutes := Get{{ .e.Upper }}Module2Actions()

    {{ .wsprefix }}CastRoutes(httpRoutes, r)
    {{ .wsprefix }}WriteHttpInformationToFile(&httpRoutes, {{ .e.EntityName }}JsonSchema, "{{ .e.Template }}-http", "{{ .m.Path }}")
    {{ .wsprefix }}WriteEntitySchema("{{ .e.EntityName }}", {{ .e.EntityName }}JsonSchema, "{{ .m.Path }}")

    return httpRoutes
  }
{{ end }}


{{ define "entityPermissions" }}

var PERM_ROOT_{{ .e.AllUpper }}_DELETE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/delete",
}

var PERM_ROOT_{{ .e.AllUpper }}_CREATE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/create",
}

var PERM_ROOT_{{ .e.AllUpper }}_UPDATE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/update",
}

var PERM_ROOT_{{ .e.AllUpper }}_QUERY = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/query",
}

{{ if .e.DistinctBy}}
  var PERM_ROOT_{{ .e.AllUpper }}_GET_DISTINCT_{{ .e.DistinctByAllUpper}} = {{ .wsprefix }}PermissionInfo{
    CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/get-distinct-{{ .e.DistinctByAllLower}}",
  }

  var PERM_ROOT_{{ .e.AllUpper }}_UPDATE_DISTINCT_{{ .e.DistinctByAllUpper}} = {{ .wsprefix }}PermissionInfo{
    CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/update-distinct-{{ .e.DistinctByAllLower}}",
  }

{{ end }}
var PERM_ROOT_{{ .e.AllUpper }} = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ .m.Path}}/{{ .e.AllLower }}/*",
}


{{ range .e.Permissions }}
var PERM_ROOT_{{ $.e.AllUpper }}_{{ .AllUpper }} = {{ $.wsprefix }}PermissionInfo{
  CompleteKey: "root/{{ $.m.Path}}/{{ $.e.AllLower }}/{{ .AllLower }}",
}

{{ end }}

var ALL_{{ .e.AllUpper }}_PERMISSIONS = []{{ .wsprefix }}PermissionInfo{
	PERM_ROOT_{{ .e.AllUpper }}_DELETE,
	PERM_ROOT_{{ .e.AllUpper }}_CREATE,
	PERM_ROOT_{{ .e.AllUpper }}_UPDATE,
  {{ if .e.DistinctBy}}
    PERM_ROOT_{{ .e.AllUpper }}_GET_DISTINCT_{{ .e.DistinctByAllUpper}},
    PERM_ROOT_{{ .e.AllUpper }}_UPDATE_DISTINCT_{{ .e.DistinctByAllUpper}},
  {{ end }}
	PERM_ROOT_{{ .e.AllUpper }}_QUERY,
	PERM_ROOT_{{ .e.AllUpper }},
  {{ range .e.Permissions }}
  PERM_ROOT_{{ $.e.AllUpper }}_{{ .AllUpper }},
  {{ end }}
}
{{ end }}