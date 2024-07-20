import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PostCategoryEntity,
} from "./PostCategoryEntity"
import {
    PostTagEntity,
} from "./PostTagEntity"
// In this section we have sub entities related to this object
// Class body
export type PostEntityKeys =
  keyof typeof PostEntity.Fields;
export class PostEntity extends BaseEntity {
  public children?: PostEntity[] | null;
  public title?: string | null;
  public content?: string | null;
    public contentExcerpt?: string[] | null;
  public category?: PostCategoryEntity | null;
      categoryId?: string | null;
  public tags?: PostTagEntity[] | null;
    tagsListId?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/post/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/post/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/post/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/posts`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "post/edit/:uniqueId",
      Rcreate: "post/new",
      Rsingle: "post/:uniqueId",
      Rquery: "posts",
  };
  public static definition = {
  "name": "post",
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
      "description": "Main category the product belongs to",
      "name": "category",
      "type": "one",
      "target": "PostCategoryEntity",
      "computedType": "PostCategoryEntity",
      "gormMap": {}
    },
    {
      "description": "Tags",
      "name": "tags",
      "type": "many2many",
      "target": "PostTagEntity",
      "computedType": "PostTagEntity[]",
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
        category: PostCategoryEntity.Fields,
        tagsListId: 'tagsListId',
      tags$: 'tags',
        tags: PostTagEntity.Fields,
}
}