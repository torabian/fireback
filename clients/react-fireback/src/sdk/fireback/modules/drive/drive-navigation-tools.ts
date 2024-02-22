// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const DriveNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/drive/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/drive/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/drive/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/drives`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "drive/edit/:uniqueId",
  Rcreate: "drive/new",
  Rsingle: "drive/:uniqueId",
  Rquery: "drives",
};
