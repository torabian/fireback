// @ts-nocheck
/**
 * functions and constants which helps to build
 * nextjs or react router dom navigation operations
 * for an entity.
 */

export const KeyboardShortcutNavigationTools = {
  edit(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/keyboard-shortcut/edit/${uniqueId}`;
  },

  create(locale?: string) {
    return `${locale ? "/" + locale : ""}/keyboard-shortcut/new`;
  },

  single(uniqueId: string, locale?: string) {
    return `${locale ? "/" + locale : ""}/keyboard-shortcut/${uniqueId}`;
  },

  query(params: any = {}, locale?: string) {
    return `${locale ? "/" + locale : ""}/keyboard-shortcuts`;
  },

  /**
   * Use R series while building router in CRA or nextjs, or react navigation for react Native
   * Might be useful in Angular as well.
   **/
  Redit: "keyboard-shortcut/edit/:uniqueId",
  Rcreate: "keyboard-shortcut/new",
  Rsingle: "keyboard-shortcut/:uniqueId",
  Rquery: "keyboard-shortcuts",
};
