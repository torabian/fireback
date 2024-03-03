import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type FileEntityKeys =
  keyof typeof FileEntity.Fields;
export class FileEntity extends BaseEntity {
  public name?: string | null;
  public diskPath?: string | null;
  public size?: number | null;
  public virtualPath?: string | null;
  public type?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/file/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/file/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/file/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/files`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "file/edit/:uniqueId",
      Rcreate: "file/new",
      Rsingle: "file/:uniqueId",
      Rquery: "files",
  };
  public static definition = {
  "name": "file",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "diskPath",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "size",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "virtualPath",
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
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      diskPath: 'diskPath',
      size: 'size',
      virtualPath: 'virtualPath',
      type: 'type',
}
}
