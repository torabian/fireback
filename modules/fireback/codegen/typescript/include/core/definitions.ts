export class BaseEntity {
  /**
   * Visibility is how each record of the database are accessible
   * For example, visiblility system means it's a system property and a part of the
   * entire app and should be visible to everyone.
   * when visibility is user, then it means only the user which created must see it
   * also visiblity can be public, which allows to make the record publicly available
   * and workspace, which is a default visiblity.
   */
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
    /**
     * Contains 'visibility' string
     */
    visibility: "visibility",
    parentId: "parentId",
    linkerId: "linkerId",
    workspaceId: "workspaceId",
    linkedId: "linkedId",
    uniqueId: "uniqueId",
    userId: "userId",
    updated: "updated",
    created: "created",
    updatedFormatted: "updatedFormatted",
    createdFormatted: "createdFormatted",
  };
}

export class BaseDto {}
