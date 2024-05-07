import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PageCategoryEntity,
} from "./PageCategoryEntity"
import {
    PageTagEntity,
} from "./PageTagEntity"
// In this section we have sub entities related to this object
// Class body
export type PageEntityKeys =
  keyof typeof PageEntity.Fields;
export class PageEntity extends BaseEntity {
  public children?: PageEntity[] | null;
  public title?: string | null;
  public content?: string | null;
    public contentExcerpt?: string[] | null;
  public category?: PageCategoryEntity | null;
      categoryId?: string | null;
  public tags?: PageTagEntity[] | null;
    tagsListId?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/page/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/page/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/page/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/pages`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "page/edit/:uniqueId",
      Rcreate: "page/new",
      Rsingle: "page/:uniqueId",
      Rquery: "pages",
  };
  public static definition = {
  "name": "page",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "title",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "content",
      "type": "html",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "description": "Main category",
      "name": "category",
      "type": "one",
      "target": "PageCategoryEntity",
      "computedType": "PageCategoryEntity",
      "gormMap": {}
    },
    {
      "description": "Tags",
      "name": "tags",
      "type": "many2many",
      "target": "PageTagEntity",
      "computedType": "PageTagEntity[]",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      title: 'title',
      content: 'content',
          categoryId: 'categoryId',
      category$: 'category',
        category: PageCategoryEntity.Fields,
        tagsListId: 'tagsListId',
      tags$: 'tags',
        tags: PageTagEntity.Fields,
}
}