import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CurrencyEntity,
} from "../currency/CurrencyEntity"
import {
    FileEntity,
} from "../drive/FileEntity"
import {
    BrandEntity,
} from "./BrandEntity"
import {
    CategoryEntity,
} from "./CategoryEntity"
import {
    ProductEntity,
    ProductFields,
} from "./ProductEntity"
import {
    TagEntity,
} from "./TagEntity"
// In this section we have sub entities related to this object
export class ProductSubmissionValues extends BaseEntity {
  public productField?: ProductFields | null;
      productFieldId?: string | null;
  public valueInt64?: number | null;
  public valueFloat64?: number | null;
  public valueString?: string | null;
  public valueBoolean?: boolean | null;
}
export class ProductSubmissionPrice extends BaseEntity {
  public stringRepresentationValue?: string | null;
  public variations?: ProductSubmissionPriceVariations[] | null;
}
export class ProductSubmissionPriceVariations extends BaseEntity {
  public currency?: CurrencyEntity | null;
      currencyId?: string | null;
  public amount?: number | null;
}
// Class body
export type ProductSubmissionEntityKeys =
  keyof typeof ProductSubmissionEntity.Fields;
export class ProductSubmissionEntity extends BaseEntity {
  public children?: ProductSubmissionEntity[] | null;
  public product?: ProductEntity | null;
      productId?: string | null;
  public data?: any | null;
  public values?: ProductSubmissionValues[] | null;
  public name?: string | null;
  public price?: ProductSubmissionPrice | null;
  public image?: FileEntity[] | null;
    imageListId?: string[] | null;
  public description?: string | null;
    public descriptionExcerpt?: string[] | null;
  public sku?: string | null;
  public brand?: BrandEntity | null;
      brandId?: string | null;
  public category?: CategoryEntity | null;
      categoryId?: string | null;
  public tags?: TagEntity[] | null;
    tagsListId?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submissions`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "product-submission/edit/:uniqueId",
      Rcreate: "product-submission/new",
      Rsingle: "product-submission/:uniqueId",
      Rquery: "product-submissions",
      rValuesCreate: "product-submission/:linkerId/values/new",
      rValuesEdit: "product-submission/:linkerId/values/edit/:uniqueId",
      editValues(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/${linkerId}/values/edit/${uniqueId}`;
      },
      createValues(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/${linkerId}/values/new`;
      },
      rPriceCreate: "product-submission/:linkerId/price/new",
      rPriceEdit: "product-submission/:linkerId/price/edit/:uniqueId",
      editPrice(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/${linkerId}/price/edit/${uniqueId}`;
      },
      createPrice(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-submission/${linkerId}/price/new`;
      },
  };
  public static definition = {
  "name": "productSubmission",
  "prependCreateScript": "ProductSubmissionCastFieldsToEavAndValidate(dto, query)",
  "prependUpdateScript": "ProductSubmissionCastFieldsToEavAndValidate(fields, query)",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "product",
      "type": "one",
      "target": "ProductEntity",
      "validate": "required",
      "computedType": "ProductEntity",
      "gormMap": {}
    },
    {
      "name": "data",
      "type": "json",
      "computedType": "any",
      "gormMap": {}
    },
    {
      "linkedTo": "ProductSubmissionEntity",
      "name": "values",
      "type": "array",
      "computedType": "ProductSubmissionValues[]",
      "gormMap": {},
      "fullName": "ProductSubmissionValues",
      "fields": [
        {
          "name": "productField",
          "type": "one",
          "target": "ProductFields",
          "rootClass": "ProductEntity",
          "computedType": "ProductFields",
          "gormMap": {}
        },
        {
          "name": "valueInt64",
          "type": "int64",
          "computedType": "number",
          "gormMap": {}
        },
        {
          "name": "valueFloat64",
          "type": "float64",
          "computedType": "number",
          "gormMap": {}
        },
        {
          "name": "valueString",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "valueBoolean",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        }
      ]
    },
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "linkedTo": "ProductSubmissionEntity",
      "name": "price",
      "type": "object",
      "computedType": "ProductSubmissionPrice",
      "gormMap": {},
      "fullName": "ProductSubmissionPrice",
      "fields": [
        {
          "name": "stringRepresentationValue",
          "type": "string",
          "computedType": "string",
          "gormMap": {},
          "sql": "-"
        },
        {
          "linkedTo": "ProductSubmissionPrice",
          "name": "variations",
          "type": "array",
          "computedType": "ProductSubmissionPriceVariations[]",
          "gormMap": {},
          "fullName": "ProductSubmissionPriceVariations",
          "fields": [
            {
              "name": "currency",
              "type": "one",
              "target": "CurrencyEntity",
              "module": "currency",
              "computedType": "CurrencyEntity",
              "gormMap": {}
            },
            {
              "name": "amount",
              "type": "float64",
              "computedType": "number",
              "gormMap": {}
            }
          ]
        }
      ]
    },
    {
      "name": "image",
      "type": "many2many",
      "target": "FileEntity",
      "module": "drive",
      "computedType": "FileEntity[]",
      "gormMap": {}
    },
    {
      "description": "Detailed description of the product",
      "name": "description",
      "type": "html",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "description": "Stock Keeping Unit code for the product",
      "name": "sku",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "description": "Brand of the product",
      "name": "brand",
      "type": "one",
      "target": "BrandEntity",
      "computedType": "BrandEntity",
      "gormMap": {}
    },
    {
      "description": "Main category the product belongs to",
      "name": "category",
      "type": "one",
      "target": "CategoryEntity",
      "computedType": "CategoryEntity",
      "gormMap": {}
    },
    {
      "description": "Tags",
      "name": "tags",
      "type": "many2many",
      "target": "TagEntity",
      "computedType": "TagEntity[]",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
          productId: 'productId',
      product$: 'product',
        product: ProductEntity.Fields,
      data: 'data',
      values$: 'values',
      values: {
  ...BaseEntity.Fields,
          productFieldId: 'productFieldId',
      productField$: 'productField',
        productField: ProductFields.Fields,
      valueInt64: 'valueInt64',
      valueFloat64: 'valueFloat64',
      valueString: 'valueString',
      valueBoolean: 'valueBoolean',
      },
      name: 'name',
      price$: 'price',
      price: {
  ...BaseEntity.Fields,
      stringRepresentationValue: 'stringRepresentationValue',
      variations$: 'variations',
      variations: {
  ...BaseEntity.Fields,
          currencyId: 'currencyId',
      currency$: 'currency',
        currency: CurrencyEntity.Fields,
      amount: 'amount',
      },
      },
        imageListId: 'imageListId',
      image$: 'image',
        image: FileEntity.Fields,
      description: 'description',
      sku: 'sku',
          brandId: 'brandId',
      brand$: 'brand',
        brand: BrandEntity.Fields,
          categoryId: 'categoryId',
      category$: 'category',
        category: CategoryEntity.Fields,
        tagsListId: 'tagsListId',
      tags$: 'tags',
        tags: TagEntity.Fields,
}
}