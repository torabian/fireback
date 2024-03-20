import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
export class ProductFields extends BaseEntity {
  public product?: ProductEntity | null;
      productId?: string | null;
  public name?: string | null;
  public type?: string | null;
}
// Class body
export type ProductEntityKeys =
  keyof typeof ProductEntity.Fields;
export class ProductEntity extends BaseEntity {
  public children?: ProductEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public uiSchema?: any | null;
  public jsonSchema?: any | null;
  public fields?: ProductFields[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/product/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/products`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "product/edit/:uniqueId",
      Rcreate: "product/new",
      Rsingle: "product/:uniqueId",
      Rquery: "products",
      rFieldsCreate: "product/:linkerId/fields/new",
      rFieldsEdit: "product/:linkerId/fields/edit/:uniqueId",
      editFields(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product/${linkerId}/fields/edit/${uniqueId}`;
      },
      createFields(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product/${linkerId}/fields/new`;
      },
  };
  public static definition = {
  "name": "product",
  "prependScript": "func CastProductFieldsFromJson(schema *workspaces.JSON) ([]*ProductFields, *workspaces.IError) {\n    form := workspaces.JSONSchema{}\n\n    if err := form.FromJson(schema); err != nil {\n        return nil, workspaces.GormErrorToIError(err)\n    }\n\n    fields := []*ProductFields{}\n    for key, field := range workspaces.FlattenFields(\"\", form.Properties) {\n        key := key\n        field := field\n        fields = append(fields, \u0026ProductFields{\n            Type: \u0026field.Type,\n            Name: \u0026key,\n        })\n    }\n\n    return fields, nil\n}\n\n\nfunc ComputeValueFromInterface(row *ProductSubmissionValues, value interface{}) {\n\n\tswitch value := value.(type) {\n\tcase int64:\n\n\t\trow.ValueInt64 = \u0026value\n\tcase float64:\n\t\trow.ValueFloat64 = \u0026value\n\tcase string:\n\t\trow.ValueString = \u0026value\n\tcase bool:\n\t\trow.ValueBoolean = \u0026value\n\t}\n\n}\n\nfunc FindFieldId(fields []*ProductFields, fieldName string) string {\n\tfor _, field := range fields {\n\t\tif *field.Name == fieldName {\n\t\t\treturn field.UniqueId\n\t\t}\n\t}\n\treturn \"\"\n}\n\nfunc SubmergeDataObjectWithValuesArray(\n\tdata *workspaces.JSON,\n\tfields []*ProductFields,\n) []*ProductSubmissionValues {\n\n\titems := []*ProductSubmissionValues{}\n\n    if (data == nil ) {\n        return items\n    }\n\n\tvar data3 map[string]interface{}\n\t// var json = jsoniter.ConfigCompatibleWithStandardLibrary\n\t// json.UnmarshalFromString(data.String(), \u0026data3)\n\tjson.Unmarshal([]byte(data.String()), \u0026data3)\n\n\tfor k, v := range workspaces.FlattenData(data3, \"\") {\n\n\t\tfieldUniqueId := FindFieldId(fields, k)\n\t\tif fieldUniqueId == \"\" {\n\t\t\tcontinue\n\t\t}\n\n\t\trow := \u0026ProductSubmissionValues{\n\t\t\tProductFieldId: \u0026fieldUniqueId,\n\t\t}\n\t\tComputeValueFromInterface(row, v)\n\n\t\titems = append(items, row)\n\t}\n\n\treturn items\n}\n\nfunc ProductSubmissionCastFieldsToEavAndValidate(dto *ProductSubmissionEntity, query workspaces.QueryDSL) *workspaces.IError {\n    if dto == nil || dto.ProductId == nil {\n        return nil\n    }\n\tid := query.UniqueId\n\tquery.UniqueId = *dto.ProductId\n\tform, err := ProductActionGetOne(query)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tquery.UniqueId = id\n\n\tdto.Values = SubmergeDataObjectWithValuesArray(dto.Data, form.Fields)\n\n\tif err0 := workspaces.ValidateEavSchema(form.JsonSchema, dto.Data); err0 != nil {\n\t\treturn err0\n\t}\n\n\treturn nil\n}",
  "prependCreateScript": "\n\tif fields, err := CastProductFieldsFromJson(dto.JsonSchema); err != nil {\n\t\treturn nil, err\n\t} else {\n\t\tdto.Fields = fields\n\t}\n\t",
  "prependUpdateScript": "\n\tif fields2, err := CastProductFieldsFromJson(fields.JsonSchema); err != nil {\n\t\treturn nil, err\n\t} else {\n\t\tfields.Fields = fields2\n\t}\n\t",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "description",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "uiSchema",
      "type": "json",
      "computedType": "any",
      "gormMap": {}
    },
    {
      "name": "jsonSchema",
      "type": "json",
      "computedType": "any",
      "gormMap": {}
    },
    {
      "linkedTo": "ProductEntity",
      "name": "fields",
      "type": "array",
      "computedType": "ProductFields[]",
      "gormMap": {},
      "fullName": "ProductFields",
      "fields": [
        {
          "name": "product",
          "type": "one",
          "target": "ProductEntity",
          "computedType": "ProductEntity",
          "gormMap": {}
        },
        {
          "name": "name",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "type",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        }
      ]
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
      uiSchema: 'uiSchema',
      jsonSchema: 'jsonSchema',
      fields$: 'fields',
      fields: {
  ...BaseEntity.Fields,
          productId: 'productId',
      product$: 'product',
      name: 'name',
      type: 'type',
      },
}
}