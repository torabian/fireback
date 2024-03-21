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
  "prependScript": "func CastProductFieldsFromJson(schema *workspaces.JSON) ([]*ProductFields, *workspaces.IError) {\r\n    form := workspaces.JSONSchema{}\r\n\r\n    if err := form.FromJson(schema); err != nil {\r\n        return nil, workspaces.GormErrorToIError(err)\r\n    }\r\n\r\n    fields := []*ProductFields{}\r\n    for key, field := range workspaces.FlattenFields(\"\", form.Properties) {\r\n        key := key\r\n        field := field\r\n        fields = append(fields, \u0026ProductFields{\r\n            Type: \u0026field.Type,\r\n            Name: \u0026key,\r\n        })\r\n    }\r\n\r\n    return fields, nil\r\n}\r\n\r\n\r\nfunc ComputeValueFromInterface(row *ProductSubmissionValues, value interface{}) {\r\n\r\n\tswitch value := value.(type) {\r\n\tcase int64:\r\n\r\n\t\trow.ValueInt64 = \u0026value\r\n\tcase float64:\r\n\t\trow.ValueFloat64 = \u0026value\r\n\tcase string:\r\n\t\trow.ValueString = \u0026value\r\n\tcase bool:\r\n\t\trow.ValueBoolean = \u0026value\r\n\t}\r\n\r\n}\r\n\r\nfunc FindFieldId(fields []*ProductFields, fieldName string) string {\r\n\tfor _, field := range fields {\r\n\t\tif *field.Name == fieldName {\r\n\t\t\treturn field.UniqueId\r\n\t\t}\r\n\t}\r\n\treturn \"\"\r\n}\r\n\r\nfunc SubmergeDataObjectWithValuesArray(\r\n\tdata *workspaces.JSON,\r\n\tfields []*ProductFields,\r\n) []*ProductSubmissionValues {\r\n\r\n\titems := []*ProductSubmissionValues{}\r\n\r\n    if (data == nil ) {\r\n        return items\r\n    }\r\n\r\n\tvar data3 map[string]interface{}\r\n\t// var json = jsoniter.ConfigCompatibleWithStandardLibrary\r\n\t// json.UnmarshalFromString(data.String(), \u0026data3)\r\n\tjson.Unmarshal([]byte(data.String()), \u0026data3)\r\n\r\n\tfor k, v := range workspaces.FlattenData(data3, \"\") {\r\n\r\n\t\tfieldUniqueId := FindFieldId(fields, k)\r\n\t\tif fieldUniqueId == \"\" {\r\n\t\t\tcontinue\r\n\t\t}\r\n\r\n\t\trow := \u0026ProductSubmissionValues{\r\n\t\t\tProductFieldId: \u0026fieldUniqueId,\r\n\t\t}\r\n\t\tComputeValueFromInterface(row, v)\r\n\r\n\t\titems = append(items, row)\r\n\t}\r\n\r\n\treturn items\r\n}\r\n\r\nfunc ProductSubmissionCastFieldsToEavAndValidate(dto *ProductSubmissionEntity, query workspaces.QueryDSL) *workspaces.IError {\r\n    if dto == nil || dto.ProductId == nil {\r\n        return nil\r\n    }\r\n\tid := query.UniqueId\r\n\tquery.UniqueId = *dto.ProductId\r\n\tform, err := ProductActionGetOne(query)\r\n\tif err != nil {\r\n\t\treturn err\r\n\t}\r\n\r\n\tquery.UniqueId = id\r\n\r\n\tdto.Values = SubmergeDataObjectWithValuesArray(dto.Data, form.Fields)\r\n\r\n\tif err0 := workspaces.ValidateEavSchema(form.JsonSchema, dto.Data); err0 != nil {\r\n\t\treturn err0\r\n\t}\r\n\r\n\treturn nil\r\n}",
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