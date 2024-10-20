export class BaseEntity {
  public visibility?: string | null = null;
  public parentId?: string | null = null;
  public linkerId?: string | null = null;
  public workspaceId?: string | null = null;
  public linkedId?: string | null = null;
  public uniqueId?: string | null = null;
  public userId?: string | null = null;
  public updated?: number | null = null;
  public created?: number | null = null;
  public createdFormatted?: string | null = null;
  public updatedFormatted?: string | null = null;
  static Fields = {
    visibility: 'visibility',
    parentId: 'parentId',
    linkerId: 'linkerId',
    workspaceId: 'workspaceId',
    linkedId: 'linkedId',
    uniqueId: 'uniqueId',
    userId: 'userId',
    updated: 'updated',
    created: 'created',
    updatedFormatted: 'updatedFormatted',
    createdFormatted: 'createdFormatted',
  };
}

export class BaseDto {}
