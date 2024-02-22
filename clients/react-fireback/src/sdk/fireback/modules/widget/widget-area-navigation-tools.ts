// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const WidgetAreaNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/widget-area/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/widget-area/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/widget-area/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/widget-areas`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "widget-area/edit/:uniqueId",
  Rcreate: "widget-area/new",
  Rsingle: "widget-area/:uniqueId",
  Rquery: "widget-areas",
};
