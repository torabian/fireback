// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const WorkspaceTypeNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/workspace-type/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/workspace-type/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/workspace-type/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/workspace-types`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "workspace-type/edit/:uniqueId",
  Rcreate: "workspace-type/new",
  Rsingle: "workspace-type/:uniqueId",
  Rquery: "workspace-types",
};
