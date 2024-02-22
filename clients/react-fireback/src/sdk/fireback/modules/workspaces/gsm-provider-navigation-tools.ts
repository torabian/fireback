// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const GsmProviderNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/gsm-provider/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/gsm-provider/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/gsm-provider/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/gsm-providers`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "gsm-provider/edit/:uniqueId",
  Rcreate: "gsm-provider/new",
  Rsingle: "gsm-provider/:uniqueId",
  Rquery: "gsm-providers",
};
