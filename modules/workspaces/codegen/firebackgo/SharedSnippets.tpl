{{ define  "goimport" }}

{{ range $key, $value := .goimports }}
{{ if and ($value.Items) ($key) }}
import  "{{ $key}}"
{{ end }}
{{ end }}


{{ end }}


{{ define "golangtype" }}{{ if or (eq .Type "array") (eq .Type "many2many") }} []* {{ end }}{{ if or (eq .Type "embed")  (eq .Type "object") (eq .Type "one") }} * {{ end }}{{ end }}

{{ define "validaterow" }}{{ if and (.Validate) (ne .Type "one") }} validate:"{{ .Validate }}" {{ end }}{{ end }}

{{ define "gormrow" }} {{ if .ComputedGormTag }} gorm:"{{ .ComputedGormTag }}" {{ end }} {{ end }}
{{ define "sqlrow" }} {{ if .ComputedSqlTag }} sql:"{{ .ComputedSqlTag }}" {{ end }} {{ end }}
{{ define "useurl" }} url:"{{ .Name }}" {{ end }}

// template for the type definition element for each field
{{ define "definitionrow" }}
  {{ $fields := index . 0 }}
  {{ $wsprefix := index . 1 }}
  {{ $useUrl := false }}

  {{ $ok := safeIndex . 2 }}
  {{- if $ok }}
  {{ $useUrl = true }}
  {{- end }}

  {{ range $fields }}
    
    {{ if ne .Type "daterange" }}

    {{ if .Description }}
    {{ goComment .Description }}
    {{ end }}
    {{ .PublicName }} {{ if eq .Type "json" }} *{{ $wsprefix }} {{ end }} {{ template "golangtype" . }} {{ .ComputedType }} `json:"{{ if .Json }}{{.Json}}{{ else }}{{.PrivateName }}{{ end }}" yaml:"{{ if .Yaml }}{{.Yaml}}{{ else }}{{.PrivateName }}{{ end }}" {{ template "validaterow" . }} {{ template "gormrow" . }} {{ template "sqlrow" . }}{{ if .Translate }} translate:"true" {{ end }} {{ if $useUrl }} {{ template "useurl" . }} {{ end }}`
    {{ end }}

    {{ if eq .Type "datenano" }}
    // Datenano also has a text representation
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
        {{ .PublicName }}Id {{ $wsprefix }}String `json:"{{ .PrivateName }}Id" yaml:"{{ .PrivateName }}Id"{{ if .IdFieldGorm }} gorm:"{{ .IdFieldGorm }}" {{ end }}{{ if .Validate }} validate:"{{ .Validate }}" {{ end }}`
        {{ end }}
    {{ end }}
    
    {{ if eq .Type "many2many" }}
    {{ .PublicName }}ListId []string `json:"{{ .PrivateName }}ListId" yaml:"{{ .PrivateName }}ListId" gorm:"-" sql:"-"`
    {{ end }}
    
    {{ if eq .Type "html" }}
    {{ .PublicName }}Excerpt *string `json:"{{ .PrivateName }}Excerpt" yaml:"{{ .PrivateName }}Excerpt"`
    {{ end }}
    
  {{ end }}
{{ end }}


{{ define "defaultgofields" }}
  {{ $v := index . 0}}
  {{ $prefix := index . 1}}

    {{ if $v.DataFields.Essentials }}

    // Defines the visibility of the record in the table.
    // Visibility is a detailed topic, you can check all of the visibility values in workspaces/visibility.go
    // by default, visibility of record are 0, means they are protected by the workspace
    // which are being created, and visible to every member of the workspace
    Visibility       {{$prefix}}String                         `json:"visibility,omitempty" yaml:"visibility,omitempty"`

    // The unique-id of the workspace which content belongs to. Upon creation this will be designated
    // to the selected workspace by user, if they have write access. You can change this value
    // or prevent changes to it manually (on root features for example modifying other workspace)
    WorkspaceId      {{$prefix}}String                         `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"{{ if $v.GormMap.WorkspaceId }} gorm:"{{ $v.GormMap.WorkspaceId }}" {{ end }}{{ if eq $v.DistinctBy "workspace" }} gorm:"unique;not null;" {{ end }}`

    // The unique-id of the parent table, which this record is being linked to.
    // used internally for making relations in fireback, generally does not need manual changes
    // or modification by the developer or user. For example, if you have a object inside an object
    // the unique-id of the parent will be written in the child.
    LinkerId         {{$prefix}}String                         `json:"linkerId,omitempty" yaml:"linkerId,omitempty"`

    // Used for recursive or parent-child operations. Some tables, are having nested relations,
    // and this field makes the table self refrenceing. ParentId needs to exist in the table before
    // creating of modifying a record.
    ParentId         {{$prefix}}String                         `json:"parentId,omitempty" yaml:"parentId,omitempty"`

    // Makes a field deletable. Some records should not be deletable at all.
    // default it's true.
    IsDeletable         *bool                         `json:"isDeletable,omitempty" yaml:"isDeletable,omitempty" gorm:"default:true"`
    
    // Makes a field updatable. Some records should not be updatable at all.
    // default it's true.
    IsUpdatable         *bool                         `json:"isUpdatable,omitempty" yaml:"isUpdatable,omitempty" gorm:"default:true"`

    // The unique-id of the user which is creating the record, or the record belongs to.
    // Administration might want to change this to any user, by default Fireback fills
    // it to the current authenticated user.
    UserId           {{$prefix}}String                         `json:"userId,omitempty" yaml:"userId,omitempty"{{ if $v.GormMap.UserId }} gorm:"{{ $v.GormMap.UserId }}" {{ end }}`

    // General mechanism to rank the elements. From code perspective, it's just a number,
    // but you can sort it based on any logic for records to make a ranking, sorting.
    // they should not be unique across a table.
    Rank             {{$prefix}}Int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    {{ end }}

    {{ if $v.DataFields.PrimaryId }}

    // Primary numeric key in the database. This value is not meant to be exported to public
    // or be used to access data at all. Rather a mechanism of indexing columns internally
    // or cursor pagination in future releases of fireback, or better search performance.
    ID    uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty" yaml:"id,omitempty"`


    // Unique id of the record across the table. This value will be accessed from public APIs,
    // and many other places intead of numeric ID property.
    // Upon generation, a UUID automatically is being assigned, and if user has specified the
    // Unique id in the post body, it will be used. This mechanism allows to manage unsaved
    // content on front-end much easier than requiring parent to exists first.
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"unique;not null;size:100;" yaml:"uniqueId,omitempty"`
    {{ end }}
    
    {{ if $v.DataFields.NumericTimestamp }}

    // The time that the record has been created in nano-seconds.
    // the field will be automatically populated by gorm orm.
    Created          int64                           `json:"created,omitempty" yaml:"created,omitempty" gorm:"autoUpdateTime:nano"`

    // The time that the record has been updated in nano-seconds.
    // the field will be automatically populated by gorm orm.
    Updated          int64                           `json:"updated,omitempty" yaml:"updated,omitempty"`

    // The time that the record has been deleted softly (means the data still exists in database, but no longer visible to any feature) in nano seconds
    // you need to make sure check this field if writing custom sql queries.
    // the field will be automatically populated by gorm orm.
    Deleted          int64                           `json:"deleted,omitempty" yaml:"deleted,omitempty"`
    {{ end }}
    
    {{ if $v.DataFields.DateTimestamp }}
    // The time that the record has been updated in datetime.
    // the field will be automatically populated by gorm orm.
    Updated          *time.Time                           `json:"updated,omitempty" yaml:"updated,omitempty"`

    // The time that the record has been created in datetime.
    // the field will be automatically populated by gorm orm.
    Created          *time.Time                           `json:"created,omitempty" yaml:"created,omitempty"`

    // The time that the record has been deleted softly (means the data still exists in database, but no longer visible to any feature) in nano datatime
    // you need to make sure check this field if writing custom sql queries.
    // the field will be automatically populated by gorm orm.
    Deleted          *time.Time                           `json:"deleted,omitempty" yaml:"deleted,omitempty"`
    {{ end }}

    // Record creation date time formatting based on locale of the headers, or other
    // possible factors.
    CreatedFormatted string                          `json:"createdFormatted,omitempty" yaml:"createdFormatted,omitempty" sql:"-" gorm:"-"`

    // Record update date time formatting based on locale of the headers, or other
    // possible factors.
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" yaml:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
{{ end }}

{{ define "polyglottable" }}
  {{ if .e.HasTranslations }}

  type {{ .e.PolyglotName}} struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId,omitempty" yaml:"linkerId,omitempty"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId,omitempty" yaml:"languageId,omitempty"`

    {{ range .e.CompleteFields }}
      {{ if .Translate }}
        {{.PublicName}} string `yaml:"{{.Name}},omitempty" json:"{{.Name}},omitempty"`
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

    dto.LinkerId = {{ $.wsprefix }}NewString(query.LinkerId)

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

    dto.LinkerId = {{ $.wsprefix }}NewString(query.LinkerId)

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


  {{ if .e.DataFields.NumericTimestamp }}
	if dto.Created > 0 {
		dto.CreatedFormatted = {{ .wsprefix }}FormatDateBasedOnQuery(dto.Created, query)
	}

	if dto.Updated > 0 {
		dto.CreatedFormatted = {{ .wsprefix }}FormatDateBasedOnQuery(dto.Updated, query)
	}
  {{ end }}
}

  {{ if .e.PostFormatter}}
  
func {{ .e.Upper }}ItemsPostFormatter(entities []*{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
  for _, entity := range entities {
      {{ .e.PostFormatter }}(entity, query)
  }
} 
  {{ end }}

{{ end }}

{{/* Used for generating mock data, useful for development or stress test */}}
{{ define "mockingentity" }}
 



func {{ .e.Upper }}ActionSeederMultiple(query {{ .wsprefix }}QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	batchSize := 100
	bar := progressbar.Default(int64(count))

	// Collect entities in batches
	var entitiesBatch []*{{ .e.Upper }}Entity

	for i := 1; i <= count; i++ {
		entity := {{ .e.Upper }}Actions.SeederInit()
		entitiesBatch = append(entitiesBatch, entity)

		// When batch size is reached, perform the batch insert
		if len(entitiesBatch) == batchSize || i == count {
			// Insert batch
			_, err := {{ .e.Upper }}Actions.MultiInsert(entitiesBatch, query)
			if err == nil {
				successInsert += len(entitiesBatch)
			} else {
				fmt.Println(err)
				failureInsert += len(entitiesBatch)
			}

			// Clear the batch after insert
			entitiesBatch = nil
		}
		bar.Add(1)
	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}

 
func {{ .e.Upper }}ActionSeeder(query {{ .wsprefix }}QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := {{ .e.Upper }}Actions.SeederInit()
		_, err := {{ .e.Upper }}Actions.Create(entity, query)
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

func (x *{{ .e.EntityName }}) Seeder() string {
	obj := {{ .e.Upper }}Actions.SeederInit()
  v, _ := json.MarshalIndent(obj, "", "  ")
  
  return string(v)
}

func {{ .e.Upper }}ActionSeederInitFn() *{{ .e.EntityName }} {
  
  entity := &{{ .e.EntityName }}{

    {{ range .e.CompleteFields }}

      {{ if  eq .Type "embed"  }}
        {{ .PublicName }}: &{{ $.e.Upper}}{{ .PublicName }}{},
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

  return entity
}
{{ end }}


{{ define "entityAssociationCreate" }}
  func {{ .e.Upper }}AssociationCreate(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) error {

  {{ range .e.CompleteFields }}
    {{ if or (eq .Type "many2many") }}
      {
        if dto.{{ .PublicName }}ListId != nil && len(dto.{{ .PublicName }}ListId) > 0 {
          var items []{{ .TargetWithModule }}
          // this operation is based on unique_id not primary key
          op := query.Tx.Where(dto.{{ .PublicName }}ListId)
          for _, item := range dto.{{ .PublicName }}ListId {
            op = op.Or("unique_id = ?", item)
          }
          err := op.Find(&items).Error

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
      dt, err := {{ .TargetWithModuleWithoutEntity}}Actions.Create(dto.{{ .PublicName }}, query);
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
			
				dt, err := {{ .TargetWithModuleWithoutEntity }}Actions.Update(query, dto.{{ .PublicName }});
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
					{{ .TargetWithModuleWithoutEntity }}Actions.Remove(cleanQuery)
				}
	
 
				dt, err := {{ .TargetWithModuleWithoutEntity }}ActionBatchCreateFn(dto.{{ .PublicName }}, query);
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


{{ define "generateQuery" }}
  {{ $fields := index . 0 }}
  {{ $wsprefix := index . 1 }}
  /*
  {{ range $fields }}
  // {{ .Name }}
  {{ end }}
  */
{{ end }}

{{ define "polyglot" }}
func {{ .e.Upper }}PolyglotUpdateHandler(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
	if dto == nil {
		return
	}

  {{ if .e.HasTranslations}}
    {{ .wsprefix }}PolyglotUpdateHandler(dto, &{{ .e.EntityName }}Polyglot{}, query)
  {{ end }}
}
{{ end }}

{{ define "asks" }}

// Creates a set of natural language queries, which can be used with
// AI tools to create content or help with some tasks

var {{ .e.Upper }}AskCmd cli.Command = cli.Command{
	Name:  "nlp",
	Usage: "Set of natural language queries which helps creating content or data",
  Subcommands: []cli.Command{
		{
			Name:  "sample",
			Usage: "Asks for generating sample by giving an example data",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "format",
					Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
					Value: "yaml",
				},
				&cli.IntFlag{
					Name:  "count",
					Usage: "How many samples to ask",
					Value: 30,
				},
			},
			Action: func(c *cli.Context) error {
				v := &{{ .e.Upper }}Entity{}

				format := c.String("format")
				request := "\033[1m" + `
I need you to create me an array of exact signature as the example given below,
with at least ` + fmt.Sprint(c.String("count")) + ` items, mock the content with few words, and guess the possible values
based on the common sense. I need the output to be a valid ` + format + ` file.

Make sure you wrap the entire array in 'items' field. Also before that, I provide some explanation of each field:



{{ template "describeFieldRecursively" (arr .e.CompleteFields "")}}

And here is the actual object signature:

` + v.Seeder() + `

`
				fmt.Println(request)
				return nil
			},
		},
	},
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

	{{ range .e.CompleteFields }}
		{{ if  eq .Type "html"  }}

			if (dto.{{ .PublicName }} != "" ) {
          {{ .PublicName }} := dto.{{ .PublicName }}
          {{ .PublicName }}Excerpt := {{ $.wsprefix}}StripPolicy.Sanitize(dto.{{ .PublicName }})
          {{ if ne .Unsafe true }}
            {{ .PublicName }} = {{ $.wsprefix}}UgcPolicy.Sanitize({{ .PublicName }})
            {{ .PublicName }}Excerpt = {{ $.wsprefix}}StripPolicy.Sanitize({{ .PublicName }}Excerpt)
          {{ end }}
          
        {{ .PublicName }}ExcerptSize, {{ .PublicName }}ExcerptSizeExists := {{ $.e.EntityName }}MetaConfig["{{ .PublicName }}ExcerptSize"]
        if {{ .PublicName }}ExcerptSizeExists {
          {{ .PublicName }}Excerpt = {{ $.wsprefix }}PickFirstNWords({{ .PublicName }}Excerpt, int({{ .PublicName }}ExcerptSize))
        } else {
          {{ .PublicName }}Excerpt = {{ $.wsprefix }}PickFirstNWords({{ .PublicName }}Excerpt, 30)
        }
        
        dto.{{ .PublicName }}Excerpt = &{{ .PublicName }}Excerpt
        dto.{{ .PublicName }} = {{ .PublicName }}
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
    
    {{ if .e.DataFields.Essentials }}
    dto.WorkspaceId = {{ .wsprefix }}NewString(query.WorkspaceId)
    dto.UserId = {{ .wsprefix }}NewString(query.UserId)
    {{ end }}

    {{ .e.Upper }}RecursiveAddUniqueId(dto, query)
  }

  func {{ .e.Upper }}RecursiveAddUniqueId(dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) {
    {{ template "beforecreatedtree" (arr .e.CompleteFields 0 "" .wsprefix) }}
  }
{{ end }}

{{ define "batchActionCreate" }}


/*
*

	Batch inserts, do not have all features that create
	operation does. Use it with unnormalized content,
	or read the source code carefully.

  This is not marked as an action, because it should not be available publicly
  at this moment.

*
*/
func {{ .e.Upper}}MultiInsertFn(dtos []*{{ .e.Upper}}Entity, query {{ .wsprefix }}QueryDSL) ([]*{{ .e.Upper}}Entity, *{{ .wsprefix }}IError) {
	if len(dtos) > 0 {

		for index := range dtos {
			{{ .e.Upper}}EntityPreSanitize(dtos[index], query)
			{{ .e.Upper}}EntityBeforeCreateAppend(dtos[index], query)
		}
		var dbref *gorm.DB = nil
		if query.Tx == nil {
			dbref = {{ .wsprefix }}GetDbRef()
		} else {
			dbref = query.Tx
		}
		query.Tx = dbref
		err := dbref.Create(&dtos).Error

		if err != nil {
			return nil, {{ .wsprefix }}GormErrorToIError(err)
		}
	}
	return dtos, nil
}

func {{ .e.Upper}}ActionBatchCreateFn(dtos []*{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL) ([]*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*{{ .e.EntityName }}{}
		for _, item := range dtos {
			s, err := {{ .e.Upper}}Actions.Create(item, query)
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

	// 3. Create other entities if we want select from them
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


{{ define "entityMemory"}}
var {{ .e.Name }}MemoryItems []*{{ .e.Upper }}Entity = []*{{ .e.Upper }}Entity{}

func {{ .e.Upper }}EntityIntoMemory() {
	q := {{ .wsprefix }}QueryDSL{
		ItemsPerPage: 500,
		StartIndex:   0,
	}
	_, qrm, _ := {{ .e.Upper }}Actions.Query(q)
	for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
		items, _, _ := {{ .e.Upper }}Actions.Query(q)
		{{ .e.Name }}MemoryItems = append({{ .e.Name }}MemoryItems, items...)
		i += q.ItemsPerPage
		q.StartIndex = i
	}
}


func {{ .e.Upper }}MemGet(id uint) *{{ .e.Upper }}Entity {
	for _, item := range {{ .e.Name }}MemoryItems {
		if item.ID == id {
			return item
		}
	}

	return nil
}


func {{ .e.Upper }}MemJoin(items []uint) []*{{ .e.Upper }}Entity {

	res := []*{{ .e.Upper }}Entity{}
	for _, item := range items {
		v := {{ .e.Upper }}MemGet(item)

		if v != nil {
			res = append(res, v)
		}
	}

	return res
}


{{ end }}


{{ define "entityActionGetAndQuery"}}
  func {{ .e.Upper }}ActionGetOneFn(query {{ .wsprefix }}QueryDSL) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    refl := reflect.ValueOf(&{{ .e.EntityName }}{})
    item, err := {{ .wsprefix }}GetOneEntity[{{ .e.EntityName }}](query, refl)

    {{ if .e.PostFormatter}}
		  {{ .e.PostFormatter}}(item, query)
		{{ end }}

    entity{{ .e.Upper }}Formatter(item, query)
    return item, err
  }
  func {{ .e.Upper }}ActionGetByWorkspaceFn(query {{ .wsprefix }}QueryDSL) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
    refl := reflect.ValueOf(&{{ .e.EntityName }}{})
    item, err := {{ .wsprefix }}GetOneByWorkspaceEntity[{{ .e.EntityName }}](query, refl)

    {{ if .e.PostFormatter}}
		  {{ .e.PostFormatter}}(item, query)
		{{ end }}

    entity{{ .e.Upper }}Formatter(item, query)
    return item, err
  }

  func {{ .e.Upper}}ActionQueryFn(query {{ .wsprefix }}QueryDSL) ([]*{{ .e.EntityName }}, *{{ .wsprefix }}QueryResultMeta, error) {
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
      if n.ParentId.Valid && n.ParentId.String == dto.UniqueId {
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
    refl := reflect.ValueOf(&{{.e.EntityName}}{})
    items, meta, err := {{ .wsprefix }}ContextAwareVSqlOperation[{{ .e.EntityName }}](
      refl, &queries.QueriesFs, "{{ .e.Upper }}Cte.vsql", query,
    )


    {{ if .e.PostFormatter}}
      {{ .e.PostFormatter }}(item, query)
      {{ end }}

    for _, item := range items {
      entity{{ .e.Upper }}Formatter(item, query)
    }


    var tree []*{{ .e.EntityName }}

    for _, item := range items {
      if !item.ParentId.Valid {
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
      return nil, {{ $.wsprefix }}Create401Error(&{{ .wsprefix }}WorkspacesMessages.BodyIsMissing, []string{})
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
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
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

    // If the entity is distinct by workspace, then the Query.WorkspaceId
    // which is selected is being used as the condition for create or update
    // if not, the unique Id is being used

    {{ if (eq .e.DistinctBy "workspace") }}
      cond2 := &{{.e.EntityName }}{WorkspaceId: {{ .wsprefix }}NewString(query.WorkspaceId)}
    {{ else if (eq .e.DistinctBy "user") }}
      cond2 := &{{.e.EntityName }}{UserId: {{ .wsprefix }}NewString(query.UserId)}
    {{ else }}
      cond2 := &{{.e.EntityName }}{UniqueId: uniqueId}
    {{ end }}

    q := dbref.
      Where(cond2).
      FirstOrCreate(&item)

    err := q.UpdateColumns(fields).Error
    if err != nil {

      return nil, {{ .wsprefix }}GormErrorToIError(err)
    }

    query.Tx = dbref
    {{ .e.Upper }}RelationContentUpdate(fields, query)

    {{ .e.Upper }}PolyglotUpdateHandler(fields, query)

    if ero := {{ .e.Upper}}DeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }


    {{ range .e.CompleteFields }}
        {{ if or (eq .Type "object") }}
   
        if fields.{{ .PublicName }} != nil {

          linkerId := uniqueId

          q := dbref.
            Model(&item.{{ .PublicName }}).
            Where(&{{ $.e.Upper }}{{ .PublicName }}{LinkerId: {{ .wsprefix }}NewString(linkerId) }).
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
              Where("unique_id IN ?", fields.{{ .PublicName }}ListId ).
              Find(&items)
          }

          dbref.
            Model(&{{$.e.EntityName }}{UniqueId: uniqueId }).
            Association("{{ .PublicName }}").
            Clear()

          dbref.
            Model(&{{$.e.EntityName }}{UniqueId: uniqueId }).
            Where(&{{$.e.EntityName }}{UniqueId: uniqueId }).
            Association("{{ .PublicName }}").
            Replace(items)
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
            Where(&{{ $entityName }}{{ $m }}{LinkerId: {{ $.wsprefix }}NewString(linkerId)}).
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
              Where(&{{ $entityName }}{{ $m }}{{ .PublicName }} {LinkerId: {{ $.wsprefix }}NewString(item.UniqueId)}).
              Delete(&{{ $entityName }}{{ $m }}{{ .PublicName }} {})
            }
          }
          {{ end }}
          
        {{ end }}
           
        dbref.
          Where(&{{ $entityName }}{{ .PublicName }} {LinkerId: {{ $.wsprefix }}NewString(linkerId)}).
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
          newItem.LinkerId = {{ $.wsprefix }}NewString(linkerId)
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
      
      {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
      ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
      {{ end }}

      {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
      AllowOnRoot: true,

      {{ end }}
    })
		count, _ := {{ .e.Upper }}ActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

func {{ .e.Upper }}ActionRemoveFn(query {{ .wsprefix }}QueryDSL) (int64, *{{ .wsprefix }}IError) {
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
        item, err := {{ .e.Upper}}Actions.Update(query, record)

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

    {{ if or (eq .e.DistinctBy "workspace")}}
      entity, err := {{ .e.Upper }}Actions.GetByWorkspace(query)
      // Because we are updating by workspace, the unique id and workspace id
      // are important to be the same.
      fields.UniqueId = query.WorkspaceId
      fields.WorkspaceId = {{ .wsprefix }}NewString(query.WorkspaceId)
    {{ else if (eq .e.DistinctBy "user" )}}
      entity, err := {{ .e.Upper }}Actions.GetOne(query)

      // It's distinct by user, then unique id and user needs to be equal
      fields.UniqueId = query.UserId
      fields.UserId = &query.UserId
    {{ end }}


    if err != nil || entity.UniqueId == "" {
      return {{ .e.Upper }}Actions.Create(fields, query)
    } else {
      return {{ .e.Upper }}Actions.Update(query, fields)
    }
  }

  func {{ .e.Upper }}DistinctActionGetOne(
    query {{ .wsprefix }}QueryDSL,
  ) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {

    {{ if or (eq .e.DistinctBy "workspace")}}
      // Get's by workspace
      entity, err := {{ .e.Upper }}Actions.GetByWorkspace(query)
    {{ else }}
      // This needs to be fixed for distinct by user or workspace/user
      query.UniqueId = query.UserId
      entity, err := {{ .e.Upper }}Actions.GetOne(query)
    {{ end }}

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
  {{ if .e.Table }}
	TableNameInDb: "{{ .e.Table }}",
  {{ else }}
	TableNameInDb: "{{.e.AllLower}}_entities",
  {{ end }}
	EntityObject:  &{{ .e.EntityName}}{},
	ExportStream: {{ .e.Upper }}ActionExportT,
	ImportQuery: {{ .e.Upper }}ActionImport,
}

{{ end }}


{{ define "entityImportExport" }}

func {{ .e.Upper }}ActionExport(
	query {{ .wsprefix }}QueryDSL,
) (chan []byte, *{{ .wsprefix }}IError) {
	return {{ .wsprefix }}YamlExporterChannel[{{ .e.EntityName}}](query, {{ .e.Upper }}Actions.Query, {{ .e.Upper }}PreloadRelations)
}


func {{ .e.Upper }}ActionExportT(
	query {{ .wsprefix }}QueryDSL,
) (chan []interface{}, *{{ .wsprefix }}IError) {
	return {{ .wsprefix }}YamlExporterChannelT[{{ .e.EntityName}}](query, {{ .e.Upper }}Actions.Query, {{ .e.Upper }}PreloadRelations)
}


func {{ .e.Upper }}ActionImport(
	dto interface{}, query {{ .wsprefix }}QueryDSL,
) *{{ .wsprefix }}IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content {{ .e.EntityName}}
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return {{ $.wsprefix }}Create401Error(&{{ $.wsprefix }}WorkspacesMessages.InvalidContent, []string{})
	}
 
	json.Unmarshal(cx, &content)
 
	_, err := {{ .e.Upper }}Actions.Create(&content, query)

	return err
}

{{ end }}



{{ define "entityInteractiveCliFlag" }}
  {{ $fields := index . 0}}
  {{ $prefix := index . 1}}

  {{ range $fields }}

  {{ if or (eq .Type "string") (eq .Type "enum") (eq .Type "string?") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
		Recommended: {{ .IsRecommended }},
		Usage:    `{{ .ComputedCliDescription}}`,
		Type: "string",
	},
	{{ end }}
  {{ if or (eq .Type "int64") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
    Recommended: {{ .IsRecommended }},
    Usage:    `{{ .ComputedCliDescription}}`,
		Type: "int64",
	},
	{{ end }}
  {{ if or (eq .Type "float64") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
    Recommended: {{ .IsRecommended }},
    Usage:    `{{ .ComputedCliDescription}}`,
		Type: "float64",
	},
	{{ end }}
  {{ if or (eq .Type "bool") }}
	{
		Name:     "{{ $prefix }}{{ .Name}}",
		StructField:     "{{ $prefix }}{{ .PublicName }}",
		Required: {{ .IsRequired }},
    Recommended: {{ .IsRecommended }},
		Usage:    `{{ .ComputedCliDescription}}`,
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
  		Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "daterange") }}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-start",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-end",
      Required: {{ .IsRequired }},
  		Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "date") }}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: `{{ .Default }}`,
      {{ end }}
    },
    {{ end }}

    {{ if or (eq .Type "int64")}}
    &cli.Int64Flag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
      Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
 
    {{ if or (eq .Type "float64")}}
    &cli.Float64Flag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
    
    {{ if or (eq .Type "bool")}}
    &cli.BoolFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
      {{ if .Default }}
      Value: {{ .Default }},
      {{ end }}
    },
    {{ end }}
   
    {{ if or (eq .Type "one")}}
    &cli.StringFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}-id",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
    },
    {{ end }}
    
    {{ if or (eq .Type "array") (eq .Type "many2many")}}
    &cli.StringSliceFlag{
      Name:     "{{ $prefix }}{{ .ComputedCliName }}",
      Required: {{ .IsRequired }},
		  Usage:    `{{ .ComputedCliDescription}}`,
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
    Usage:    "Unique Id - external unique hash to query entity",
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
    Usage: "Creates a new entity, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},

        {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
        {{ end }}

        {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
        AllowOnRoot: true,

        {{ end }}
      })

      entity := &{{ .e.EntityName }}{}
      {{ .wsprefix }}PopulateInteractively(entity, c, {{ .e.Upper }}CommonInteractiveCliFlags)

      if entity, err := {{ .e.Upper }}Actions.Create(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {

        f, _ := yaml.Marshal(entity)
			  fmt.Println({{ .wsprefix }}FormatYamlKeys(string(f)))
      }
    },
  }

  var {{ .e.Upper }}UpdateCmd cli.Command = cli.Command{

    Name:    "update",
    Aliases: []string{"u"},
    Flags: {{ .e.Upper }}CommonCliFlagsOptional,
    Usage:   "Updates entity by passing the parameters",
    Action: func(c *cli.Context) error {

      query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},

        {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
        {{ end }}

        {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
        AllowOnRoot: true,

        {{ end }}
      })

      entity := Cast{{ .e.Upper }}FromCli(c)

      if entity, err := {{ .e.Upper }}Actions.Update(query, entity); err != nil {
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


{{ define "describeFieldRecursively" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}

  {{ range $fields }}

{{ .PublicName }}: (type: {{ .Type }}) Description: {{ .Description }}

  {{ end }}
{{ end }}

{{ define "entityCliCastRecursive" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}
  {{ $wsprefix := index . 2 }}

  {{ range $fields }}

    {{ if or (eq .Type "string") (eq .Type "enum") (eq .Type "html") (eq .Type "text")}}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        template.{{ .PublicName }} = c.String("{{ $prefix }}{{ .ComputedCliName }}")
      }
	  {{ end }}
    {{ if or (eq .Type "string?") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        template.{{ .PublicName }} = {{ $wsprefix }}NewStringAutoNull(c.String("{{ $prefix }}{{ .ComputedCliName }}"))
      }
	  {{ end }}
    {{ if or (eq .Type "int64") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.Int64("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }} = value
      }
	  {{ end }}
    {{ if or (eq .Type "float64") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.Float64("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }} = value
      }
	  {{ end }}
    {{ if or (eq .Type "bool") (eq .Type "boolean") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}") {
        value := c.Bool("{{ $prefix }}{{ .ComputedCliName }}")
        template.{{ .PublicName }} = value
      }
    {{ end }}
    {{ if or (eq .Type "one") }}
      if c.IsSet("{{ $prefix }}{{ .ComputedCliName }}-id") {
        template.{{ .PublicName }}Id = {{ $prefix }}NewStringAutoNull(c.String("{{ $prefix }}{{ .ComputedCliName }}-id"))
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
      } {{ if endsWithDto .Target }} {{ else }}  else {
        template.{{ .PublicName }}ListId = {{ $wsprefix }}CliInteractiveSearchAndSelect(
          "Select {{ .PublicName }}",
          {{ .PublicName }}ActionQueryString,
        )
      } {{ end }}
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

  {{ if .e.DataFields.Essentials }}
	if c.IsSet("pid") {
		template.ParentId = {{ .wsprefix}}NewStringAutoNull(c.String("pid"))
	}
  {{ end }}
	
	{{ template "entityCliCastRecursive" (arr .e.CompleteFields "" .wsprefix)}}

	return template
}

{{ end }}

{{ define "dtoCastFromCli" }}

func Cast{{ .e.Upper }}FromCli (c *cli.Context) *{{ .e.ObjectName }} {
	template := &{{ .e.ObjectName }}{}

	{{ template "entityCliCastRecursive" (arr .e.CompleteFields "" .wsprefix)}}

	return template
}

{{ end }}

{{ define "entityMockAndSeeders" }}

  func {{ .e.Upper }}SyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{},
      {{ .e.Upper }}Actions.Create,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }

  func {{ .e.Upper }}SyncSeeders() {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{WorkspaceId: {{ .wsprefix }}USER_SYSTEM},
      {{ .e.Upper }}Actions.Create,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      {{ .e.Name }}SeedersFs,
      []string{},
      true,
    )
  }

  func {{ .e.Upper }}ImportMocks() {
    {{ .wsprefix }}SeederFromFSImport(
      {{ .wsprefix }}QueryDSL{},
      {{ .e.Upper }}Actions.Create,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }

  func {{ .e.Upper }}WriteQueryMock(ctx {{ .wsprefix }}MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := {{ .wsprefix }}QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := {{ .e.Upper }}Actions.Query(f)
      result := {{ .wsprefix }}QueryEntitySuccessResult(f, items, count)
      {{ .wsprefix }}WriteMockDataToFile(lang, "", "{{ .e.Upper }}", result)
    }
  }
{{ end }}

{{ define "entityCliImportExportCmd" }}

var {{ .e.Upper }}ImportExportCommands = []cli.Command{
  {{ if eq .e.Features.HasMockAction true }}
	{

		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
			&cli.BoolFlag{
				Name:  "batch",
				Usage: "Multiple insert into database mode. Might miss children and relations at the moment",
			},
		},
		Action: func(c *cli.Context) error {
			query := {{ .wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, &{{ .wsprefix }}SecurityModel{
        ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
        {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
        {{ end }}

        {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
        AllowOnRoot: true,

        {{ end }}
      })

      if c.Bool("batch") {
			  {{ .e.Upper }}ActionSeederMultiple(query, c.Int("count"))
			} else {
			  {{ .e.Upper }}ActionSeeder(query, c.Int("count"))
      }

			return nil
		},
	},
  {{ end }}

  
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
			seed := {{ .e.Upper }}Actions.SeederInit()

      {{ .wsprefix }}CommonInitSeeder(strings.TrimSpace(c.String("format")), seed)
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
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
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
	cli.Command{
		Name:  "slist",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := {{ .wsprefix }}GetSeederFilenames({{ .e.Name }}SeedersFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {

				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}

			return nil
		},
	},
	cli.Command{
		Name:  "ssync",
		Usage: "Tries to sync the embedded content into the database, the list could be seen by 'slist' command",
		Action: func(c *cli.Context) error {

			{{ .wsprefix }}CommonCliImportEmbedCmd(c,
				{{ .e.Upper }}Actions.Create,
				reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
				{{ .e.Name }}SeedersFs,
			)

			return nil
		},
	},
  {{ if eq .e.Features.HasMsyncActions true }}
  cli.Command{
    Name:  "mlist",
    Usage: "Prints the list of embedded mocks into the app",
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
        {{ .e.Upper }}Actions.Create,
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
      return {{ .wsprefix }}CommonCliExportCmd2(c,
        {{ .e.Upper }}EntityStream,
        reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
        c.String("file"),
        &metas.MetaFs,
        "{{ .e.Upper }}FieldMap.yml",
        {{ .e.Upper }}PreloadRelations,
      )
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
				{{ .e.Upper }}Actions.Create,
				reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
				c.String("file"),
        &{{ .wsprefix }}SecurityModel{
					ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
          {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
          ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
          {{ end }}

          {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
          AllowOnRoot: true,

          {{ end }}
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
      {{.e.AllUpper}}_ACTION_QUERY.ToCli(),
      {{.e.AllUpper}}_ACTION_TABLE.ToCli(),
      {{ if ne .e.Access "read" }}

      {{ .e.Upper }}CreateCmd,
      {{ .e.Upper }}UpdateCmd,
      {{ .e.Upper }}AskCmd,
      {{ .e.Upper }}CreateInteractiveCmd,
      {{ .e.Upper }}WipeCmd,
      {{ .wsprefix }}GetCommonRemoveQuery(
        reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
        {{ .e.Upper }}Actions.Remove,
      ),
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
    commands := append({{ .e.Upper }}ImportExportCommands, {{ .e.Upper }}CliCommands...)

    return cli.Command{
      Name:        "{{ .e.ComputedCliName }}",
      {{ if .e.CliShort }}
      ShortName:   "{{ .e.CliShort }}",
      {{ end }}
      Description: "{{ .e.Upper }}s module actions",
      Usage:       `{{ .e.ComputedCliDescription }}`,
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: commands,
    }
  }

{{ end }}

{{ define "entityHttp" }}

var {{.e.AllUpper}}_ACTION_TABLE = {{ .wsprefix }}Module3Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  {{ .wsprefix }}CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: {{ .e.Upper }}Actions.Query,
  CliAction: func(c *cli.Context, security *{{ .wsprefix }}SecurityModel) error {
    {{ .wsprefix }}CommonCliTableCmd2(c,
      {{ .e.Upper }}Actions.Query,
      security,
      reflect.ValueOf(&{{ .e.EntityName }}{}).Elem(),
    )

    return nil
  },
}

var {{.e.AllUpper}}_ACTION_QUERY = {{ .wsprefix }}Module3Action{
  Method: "GET",
  Url:    "/{{ .e.DashedPluralName }}",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      {{ .wsprefix }}HttpQueryEntity(c, {{ .e.Upper }}Actions.Query)
    },
  },
  Format: "QUERY",
  Action: {{ .e.Upper }}Actions.Query,
  ResponseEntity: &[]{{ .e.EntityName }}{},
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
  CliAction: func(c *cli.Context, security *{{ .wsprefix }}SecurityModel) error {
		{{ .wsprefix }}CommonCliQueryCmd2(
			c,
			{{ .e.Upper }}Actions.Query,
			security,
		)
		return nil
	},
	CliName:       "query",
	Name:    "query",
	ActionAliases: []string{"q"},
	Flags:         {{ .wsprefix }}CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}

{{ if .e.Cte }}
var {{.e.AllUpper}}_ACTION_QUERY_CTE = {{ .wsprefix }}Module3Action{
  Method: "GET",
  Url:    "/cte-{{ .e.DashedPluralName }}",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
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
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}
{{ end }}

var {{.e.AllUpper}}_ACTION_EXPORT = {{ .wsprefix }}Module3Action{
  Method: "GET",
  Url:    "/{{ .e.DashedPluralName }}/export",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
    {{ end }}
    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
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
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}

var {{.e.AllUpper}}_ACTION_GET_ONE = {{ .wsprefix }}Module3Action{
  Method: "GET",
  Url:    "/{{ .e.Template }}/:uniqueId",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_QUERY},
    {{ end }}
    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      {{ .wsprefix }}HttpGetEntity(c, {{ .e.Upper }}Actions.GetOne)
    },
  },
  Format: "GET_ONE",
  Action: {{ .e.Upper }}Actions.GetOne,
  ResponseEntity: &{{ .e.EntityName }}{},
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}

{{ if ne .e.Access "read" }}
var {{.e.AllUpper}}_ACTION_POST_ONE = {{ .wsprefix }}Module3Action{
  Name:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new {{ .e.Name }}",
  Flags: {{ .e.Upper }}CommonCliFlags,
  Method: "POST",
  Url:    "/{{ .e.Template }}",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_CREATE},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
    AllowOnRoot: true,

    {{ end }}
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      {{ .wsprefix }}HttpPostEntity(c, {{ .e.Upper }}Actions.Create)
    },
  },
  CliAction: func(c *cli.Context, security *{{ .wsprefix }}SecurityModel) error {
    result, err := {{ .wsprefix }}CliPostEntity(c, {{ .e.Upper }}Actions.Create, security)
    {{ .wsprefix }}HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: {{ .e.Upper }}Actions.Create,
  Format: "POST_ONE",
  RequestEntity: &{{ .e.EntityName }}{},
  ResponseEntity: &{{ .e.EntityName }}{},
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
  In: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}

var {{.e.AllUpper}}_ACTION_PATCH = {{ .wsprefix }}Module3Action{
  Name:    "update",
  ActionAliases: []string{"u"},
  Flags: {{ .e.Upper }}CommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/{{ .e.Template }}",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
    AllowOnRoot: true,

    {{ end }}
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      {{ .wsprefix }}HttpUpdateEntity(c, {{ .e.Upper }}Actions.Update)
    },
  },
  Action: {{ .e.Upper }}Actions.Update,
  RequestEntity: &{{ .e.EntityName }}{},
  ResponseEntity: &{{ .e.EntityName }}{},
  Format: "PATCH_ONE",
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
  In: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}


var {{.e.AllUpper}}_ACTION_PATCH_BULK = {{ .wsprefix }}Module3Action{
  Method: "PATCH",
  Url:    "/{{ .e.DashedPluralName }}",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
    AllowOnRoot: true,

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
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
  In: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}
var {{.e.AllUpper}}_ACTION_DELETE = {{ .wsprefix }}Module3Action{
  Method: "DELETE",
  Url:    "/{{ .e.Template }}",
  Format: "DELETE_DSL",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_DELETE},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
    AllowOnRoot: true,

    {{ end }}
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      {{ .wsprefix }}HttpRemoveEntity(c, {{ .e.Upper }}Actions.Remove)
    },
  },
  Action: {{ .e.Upper }}Actions.Remove,
  RequestEntity: &{{ .wsprefix }}DeleteRequest{},
  ResponseEntity: &{{ .wsprefix }}DeleteResponse{},
  TargetEntity: &{{ .e.EntityName }}{},
}

{{ if or (eq .e.DistinctBy "user") (eq .e.DistinctBy "workspace")}}
var {{.e.AllUpper}}_ACTION_DISTINCT_PATCH_ONE = {{ .wsprefix }}Module3Action{
  Method: "PATCH",
  Url:    "/{{ .e.Template }}/distinct",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_UPDATE_DISTINCT_{{ .e.DistinctByAllUpper}}},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.WriteOnRoot) }}
    AllowOnRoot: true,

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
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
  In: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}

var {{.e.AllUpper}}_ACTION_DISTINCT_GET_ONE = {{ .wsprefix }}Module3Action{
  Method: "GET",
  Url:    "/{{ .e.Template }}/distinct",
  SecurityModel: &{{ .wsprefix }}SecurityModel{
    {{ if ne $.e.QueryScope "public" }}
    ActionRequires: []{{ .wsprefix }}PermissionInfo{PERM_ROOT_{{ .e.AllUpper }}_GET_DISTINCT_{{ .e.DistinctByAllUpper}}},
    {{ end }}

    {{ if and (.e.SecurityModel) (.e.SecurityModel.ResolveStrategy) }}
    ResolveStrategy: "{{ .e.SecurityModel.ResolveStrategy }}",
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
  Out: &{{ .wsprefix }}Module3ActionBody{
		Entity: "{{ .e.EntityName }}",
	},
}
{{ end }}

{{ range .e.CompleteFields }}
  {{ if or (eq .Type "object") (eq .Type "array")}}
    var {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_PATCH = {{ $.wsprefix }}Module3Action{
      Method: "PATCH",
      Url:    "/{{ $.e.Template }}/:linkerId/{{ .DashedName }}/:uniqueId",
      SecurityModel: &{{ $.wsprefix }}SecurityModel{
        {{ if ne $.e.QueryScope "public" }}
        ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_UPDATE},
        {{ end }}

        {{ if and ($.e.SecurityModel) ($.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ $.e.SecurityModel.ResolveStrategy }}",
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
      Out: &{{ $.wsprefix }}Module3ActionBody{
        Entity: "{{ $.e.Upper }}{{ .PublicName }}",
      },
      In: &{{ $.wsprefix }}Module3ActionBody{
        Entity: "{{ $.e.Upper }}{{ .PublicName }}",
      },
    }
    var {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_GET = {{ $.wsprefix }}Module3Action {
      Method: "GET",
      Url:    "/{{ $.e.Template }}/{{ .DashedName }}/:linkerId/:uniqueId",
      SecurityModel: &{{ $.wsprefix }}SecurityModel{
        {{ if ne $.e.QueryScope "public" }}
        ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_QUERY},
        {{ end }}

        {{ if and ($.e.SecurityModel) ($.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ $.e.SecurityModel.ResolveStrategy }}",
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
      Out: &{{ $.wsprefix }}Module3ActionBody{
        Entity: "{{ $.e.Upper }}{{ .PublicName }}",
      },
    }
    var {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_POST = {{ $.wsprefix }}Module3Action{
      Method: "POST",
      Url:    "/{{ $.e.Template }}/:linkerId/{{ .DashedName }}",
      SecurityModel: &{{ $.wsprefix }}SecurityModel{
        {{ if ne $.e.QueryScope "public" }}
        ActionRequires: []{{ $.wsprefix }}PermissionInfo{PERM_ROOT_{{ $.e.AllUpper }}_CREATE},
        {{ end }}

        {{ if and ($.e.SecurityModel) ($.e.SecurityModel.ResolveStrategy) }}
        ResolveStrategy: "{{ $.e.SecurityModel.ResolveStrategy }}",
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
      Out: &{{ $.wsprefix }}Module3ActionBody{
        Entity: "{{ $.e.Upper }}{{ .PublicName }}",
      },
      In: &{{ $.wsprefix }}Module3ActionBody{
        Entity: "{{ $.e.Upper }}{{ .PublicName }}",
      },
    }
  {{ end }}
{{ end }}

{{ end }}
  /**
  *	Override this function on {{ .e.EntityName }}Http.go,
  *	In order to add your own http
  **/
  var Append{{ .e.Upper }}Router = func(r *[]{{ .wsprefix }}Module3Action) {}
 
  func Get{{ .e.Upper }}Module3Actions() []{{ .wsprefix }}Module3Action {

    routes := []{{ .wsprefix }}Module3Action{
      {{ if .e.Cte }}
        {{.e.AllUpper}}_ACTION_QUERY_CTE,
      {{ end }}
      {{.e.AllUpper}}_ACTION_QUERY,
      {{.e.AllUpper}}_ACTION_EXPORT,
      {{.e.AllUpper}}_ACTION_GET_ONE,

      {{ if ne .e.Access "read" }}
      {{.e.AllUpper}}_ACTION_POST_ONE,
      {{.e.AllUpper}}_ACTION_PATCH,
      {{.e.AllUpper}}_ACTION_PATCH_BULK,
      {{.e.AllUpper}}_ACTION_DELETE,
     
      {{ if or (eq .e.DistinctBy "user") (eq .e.DistinctBy "workspace")}}
      {{.e.AllUpper}}_ACTION_DISTINCT_PATCH_ONE,
      {{.e.AllUpper}}_ACTION_DISTINCT_GET_ONE,
      {{ end }}

      {{ end }}

      {{ range .e.CompleteFields }}
        {{ if or (eq .Type "object") (eq .Type "array")}}
          {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_PATCH,
          {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_GET,
          {{ $.e.AllUpper }}_{{ .AllUpper }}_ACTION_POST,
        {{ end }}
      {{ end }}
    }
   
    // Append user defined functions
    Append{{ .e.Upper }}Router(&routes)

    return routes

  }

{{ end }}


{{ define "entityPermissions" }}

var PERM_ROOT_{{ .e.AllUpper }}_DELETE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.delete",
  Name: "Delete {{ .e.HumanReadable }}",
}

var PERM_ROOT_{{ .e.AllUpper }}_CREATE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.create",
  Name: "Create {{ .e.HumanReadable }}",
}

var PERM_ROOT_{{ .e.AllUpper }}_UPDATE = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.update",
  Name: "Update {{ .e.HumanReadable }}",
}

var PERM_ROOT_{{ .e.AllUpper }}_QUERY = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.query",
  Name: "Query {{ .e.HumanReadable }}",
}

{{ if .e.DistinctBy}}
  var PERM_ROOT_{{ .e.AllUpper }}_GET_DISTINCT_{{ .e.DistinctByAllUpper}} = {{ .wsprefix }}PermissionInfo{
    CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.get-distinct-{{ .e.DistinctByAllLower}}",
    Name: "Get {{ .e.HumanReadable }} Distinct",
  }

  var PERM_ROOT_{{ .e.AllUpper }}_UPDATE_DISTINCT_{{ .e.DistinctByAllUpper}} = {{ .wsprefix }}PermissionInfo{
    CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.update-distinct-{{ .e.DistinctByAllLower}}",
    Name: "Update {{ .e.HumanReadable }} Distinct",
  }

{{ end }}
var PERM_ROOT_{{ .e.AllUpper }} = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ .ctx.RelativePathDot}}.{{ .e.AllLower }}.*",
  Name: "Entire {{ .e.HumanReadable }} actions (*)",
}


{{ range .e.Permissions }}
var PERM_ROOT_{{ $.e.AllUpper }}_{{ .AllUpper }} = {{ $.wsprefix }}PermissionInfo{
  CompleteKey: "root.{{ $.ctx.RelativePathDot}}.{{ $.e.AllLower }}.{{ .AllLower }}",
  Name: "{{ .AllUpper }}",
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

{{ define "messageCode" }}

{{ $name := index . 0 }}
{{ $messages := index . 1 }}
{{ $wsprefix := index . 2 }}

type {{ $name }}Code string

const (
{{- range $key, $items := $messages }}
	{{ upper $key }} {{ $name }}Code = "{{ upper $key }}"
{{- end }}
) 

var {{ upper $name }}Messages = new{{ upper $name }}MessageCode()

func new{{ upper $name }}MessageCode() *{{ $name }}Msgs {
	return &{{ $name }}Msgs{

    {{- range $key, $items := $messages }}
      {{ upper $key }}: {{ $wsprefix }}ErrorItem{
        "$": "{{ upper $key }}",
        {{- range $lang, $value := $items }}
          "{{ $lang }}": "{{ $value }}",
        {{- end }}
      },
    {{- end }}
	}
}

type {{ $name }}Msgs struct {
  {{- range $key, $items := $messages }}
      {{ upper $key }} {{ $wsprefix }}ErrorItem
  {{- end }}
}


{{ end }}


{{ define "commonFieldsAndDto "}}
// Hi, I am here
{{ end }}


{{ define "actions-section" }}
// using shared actions here
  {{ $actions := index . 0 }}
  {{ $wsprefix := index . 1 }}
  {{ $name := index . 2 }}
  {{ $remoteQueryChildren := index . 3 }}
  {{ $childrenIn := index . 4 }}
  {{ $childrenOut := index . 5 }}



{{ range $remoteQueryChildren }}

  {{ range .}}

  type {{ .FullName }} struct {
    {{ template "definitionrow" (arr .Fields $wsprefix) }}
  }

  {{ end }}
{{ end }}

{{ range $childrenIn }}

  {{ range .}}

  type {{ .FullName }} struct {
    {{ template "definitionrow" (arr .Fields $wsprefix) }}
  }

  {{ end }}
{{ end }}

{{ range $childrenOut }}

  {{ range .}}

  type {{ .FullName }} struct {
    {{ template "definitionrow" (arr .Fields $wsprefix) }}
  }

  {{ end }}
{{ end }}


{{ range $actions }}

  {{ if .SecurityModel }}
  var {{ .Upper }}SecurityModel = &{{ $wsprefix }}SecurityModel{
    ActionRequires: []{{ $wsprefix }}PermissionInfo{ 
        {{ range .SecurityModel.ActionRequires }}
            {
              CompleteKey: "{{ .CompleteKey }}",
            },
        {{ end }}
    },
  }
  {{ else }}
  var {{ .Upper }}SecurityModel *{{ $wsprefix }}SecurityModel = nil
  {{ end }}


  {{ if .Query }}
    type {{ upper .Name }}Query struct {
      {{ template "definitionrow" (arr .Query $wsprefix true) }}
    }
  {{ end }}



  {{ if .In }}
    {{ if .In.Fields }}
      type {{ .Upper }}ActionReqDto struct {
          {{ template "definitionrow" (arr .In.Fields $wsprefix) }}
      }

      func ( x * {{ .Upper }}ActionReqDto) RootObjectName() string {
        return "{{ $name }}"
      }

      var {{ .Upper }}CommonCliFlagsOptional = []cli.Flag{
        {{ template "dtoCliFlag" (arr .In.Fields "") }}
      }

      func {{ .Upper }}ActionReqValidator(dto *{{ .Upper }}ActionReqDto) *{{ $wsprefix }}IError {
        err := {{ $wsprefix }}CommonStructValidatorPointer(dto, false)

        {{ range .In.Fields }}
          {{ if  eq .Type "array"  }}
            if dto != nil && dto.{{ .UpperPlural }} != nil {
              {{ $wsprefix }}AppendSliceErrors(dto.{{ .UpperPlural }}, false, "{{ .Name}}", err)
            }
          {{ end }}
        {{ end }}
        return err
      }

      func Cast{{ .Upper }}FromCli (c *cli.Context) *{{ .Upper }}ActionReqDto {
        template := &{{ .Upper }}ActionReqDto{}

        {{ template "entityCliCastRecursive" (arr .In.Fields "" $wsprefix)}}

        return template
      }
    {{ end }}
  {{ end }}

  {{ if .Out }}
    {{ if .Out.Fields }}
      type {{ .Upper }}ActionResDto struct {
        {{ template "definitionrow" (arr .Out.Fields $wsprefix) }}
      }

      func ( x * {{ .Upper }}ActionResDto) RootObjectName() string {
        return "{{ $name }}"
      }
    {{ end }}
  {{ end }}

  {{ if or (eq .Method "reactive")}}
    var {{ .Upper }}ActionImp = {{ $wsprefix }}DefaultEmptyReactiveAction
  {{ else }}
    type {{ .Name }}ActionImpSig func(
      {{ if .ComputeRequestEntity }}{{ if ne .ActionReqDto "nil" }}req {{ .ActionReqDto }}, {{ end}}{{end}}
      q {{ $wsprefix }}QueryDSL) ({{ .ActionResDto }},
      {{ if (eq .FormatComputed "QUERY") }} *{{ $wsprefix }}QueryResultMeta, {{ end }}
      *{{ $wsprefix }}IError,
    )
    var {{ .Upper }}ActionImp {{ .Name }}ActionImpSig
  {{ end }}


  {{ if or (eq .Method "reactive")}}
    // Reactive action does not have that
  {{ else }}
    func {{ .Upper }}ActionFn(
      {{ if .ComputeRequestEntity }}
          {{ if ne .ActionReqDto "nil" }}req {{ .ActionReqDto }}, {{ end}}
      {{ end }}
      q {{ $wsprefix }}QueryDSL,
    ) (
        {{ .ActionResDto }},
        {{ if (eq .FormatComputed "QUERY") }} *{{ $wsprefix }}QueryResultMeta, {{ end }}
        *{{ $wsprefix }}IError,
    ) {
      if {{ .Upper }}ActionImp == nil {
          return {{ if (eq .ActionResDto "string")}} "" {{ else }} nil {{ end }}, {{ if (eq .FormatComputed "QUERY") }} nil, {{ end }} nil
      }
      return {{ .Upper }}ActionImp({{ if .ComputeRequestEntity }}{{ if ne .ActionReqDto "nil" }}req, {{ end}}{{ end}} q)
    }
  {{ end }}

  
  var {{ .Upper }}ActionCmd cli.Command = cli.Command{
    Name:  "{{ .ComputedCliName }}",
    Usage: `{{ .Description }}`,
      {{ if (eq .FormatComputed "QUERY") }}
      Flags: {{ $wsprefix }}CommonQueryFlags,
      {{ end }}

      {{ if .In }}
          {{ if .In.Fields }}
          Flags: {{ .Upper }}CommonCliFlagsOptional,
          {{ else if .In.Entity }}
          Flags: {{ .In.EntityPure }}CommonCliFlagsOptional,
          {{ end }}
      {{ end }}
    Action: func(c *cli.Context) {

      query := {{ $wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, {{ .Upper }}SecurityModel)

          {{ if or (ne .Method "reactive")}}

          {{ if .In }}
              {{ if .In.Fields }}
              dto := Cast{{ .Upper }}FromCli(c)
              {{ else if .In.Entity }}
              dto := Cast{{ .In.EntityPure }}FromCli(c)
              {{ end }}
          {{ end }}


          {{ if or (eq .FormatComputed "QUERY")}}
      result, _, err := {{ .Upper }}ActionFn(query)
          {{ else if or (eq .ComputeRequestEntity "") }}
      result, err := {{ .Upper }}ActionFn(query)
          {{ else }}
      result, err := {{ .Upper }}ActionFn(dto, query)
          {{ end }}

      {{ $wsprefix }}HandleActionInCli(c, result, err, map[string]map[string]string{})
          {{ else }}
          {{ $wsprefix }}CliReactivePipeHandler(query, {{ .Upper }}ActionImp)
          {{end}}
    },
  }


{{ end }}


func {{ $name }}CustomActions() []{{ $wsprefix }}Module3Action {
	routes := []{{ $wsprefix }}Module3Action{
        {{ range $actions }}
		{
			Method: "{{ .MethodAllUpper }}",
			Url:    "{{ .ComputedUrl }}",
            SecurityModel: {{ .Upper }}SecurityModel,
            Name: "{{ .Name }}",
            Description: "{{ .Description }}",
			Handlers: []gin.HandlerFunc{
                {{ if or (eq .Method "reactive")}}
                {{ $wsprefix }}ReactiveSocketHandler({{ .Upper }}ActionImp),
                {{ else }}
				func(c *gin.Context) {
                    // {{ .FormatComputed }} - {{ .Method }}
                    
                    {{ if or (eq .FormatComputed "POST") (eq .Method "POST") (eq .Method "post") }}
                        {{ $wsprefix }}HttpPostEntity(c, {{ .Upper }}ActionFn)
                    {{ end }}
                    {{ if or (eq .FormatComputed "QUERY")}}
                        {{ $wsprefix }}HttpQueryEntity2(c, {{ .Upper }}ActionFn)
                    {{ end }}
                    
                    {{ if or (eq .FormatComputed "GET_ONE")}}
                        {{ $wsprefix }}HttpGetEntity(c, {{ .Upper }}ActionFn)
                    {{ end }}
                },
                {{ end }}
			},
			Format:         "{{ .FormatComputed }}",
            {{ if or (ne .Method "reactive")}}
			Action:         {{ .Upper }}ActionFn,
            {{end}}
            {{ if .ComputeResponseEntity }}
			ResponseEntity: {{.ComputeResponseEntity}},
            Out: &{{ $wsprefix }}Module3ActionBody{
                Entity: "{{ .ComputeResponseEntityS }}",
            },
            {{ end }}
            {{ if .ComputeRequestEntity}}
			RequestEntity: {{.ComputeRequestEntity}},
            In: &{{ $wsprefix }}Module3ActionBody{
                Entity: "{{ .ComputeRequestEntityS }}",
            },
            {{ end }}
		},
        {{ end }}
	}
    
	return routes
}


{{ end }}