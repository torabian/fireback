// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const CapabilityNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/capability/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/capability/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/capability/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/capabilities`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "capability/edit/:uniqueId",
  Rcreate: "capability/new",
  Rsingle: "capability/:uniqueId",
  Rquery: "capabilities",
};
