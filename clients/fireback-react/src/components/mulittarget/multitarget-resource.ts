// @ts-nocheck
/**
 * Store icons for different operating system here
 */

import { getOS } from "@/hooks/useHtmlClass";
const OS = getOS();

type IconOsMap = {
  [key: string]: {
    mac?: string;
    ios?: string;
    windows?: string;
    android?: string;
    linux?: string;
    web?: string;
    default?: string;
  };
};

const icons: IconOsMap = {
  edit: {
    default: "ios-theme/icons/edit.svg",
  },
  add: {
    default: "ios-theme/icons/add.svg",
  },
  cancel: {
    default: "ios-theme/icons/cancel.svg",
  },
  delete: {
    default: "ios-theme/icons/delete.svg",
  },
  entity: {
    default: "ios-theme/icons/entity.svg",
  },
  left: {
    default: "ios-theme/icons/left.svg",
  },
  menu: {
    default: "ios-theme/icons/menu.svg",
  },
  backup: {
    default: "ios-theme/icons/backup.svg",
  },
  right: {
    default: "ios-theme/icons/right.svg",
  },
  settings: {
    default: "ios-theme/icons/settings.svg",
  },
  user: {
    default: "ios-theme/icons/user.svg",
  },
  export: {
    default: "ios-theme/icons/export.svg",
  },
  up: {
    default: "ios-theme/icons/up.svg",
  },
  dataNode: {
    default: "ios-theme/icons/dnode.svg",
  },
  ctrlSheet: {
    default: "ios-theme/icons/ctrlsheet.svg",
  },
  gpio: {
    default: "ios-theme/icons/gpio.svg",
  },
  gpiomode: {
    default: "ios-theme/icons/gpiomode.svg",
  },
  gpiostate: {
    default: "ios-theme/icons/gpiostate.svg",
  },
  down: {
    default: "ios-theme/icons/down.svg",
  },
  turnoff: {
    default: "ios-theme/icons/turnoff.svg",
  },
  mqtt: {
    default: "ios-theme/icons/mqtt.svg",
  },
  cart: {
    default: "ios-theme/icons/cart.svg",
  },
  questionBank: {
    default: "ios-theme/icons/questions.svg",
  },
  dashboard: {
    default: "ios-theme/icons/dashboard.svg",
  },
  country: {
    default: "ios-theme/icons/country.svg",
  },
  order: {
    default: "ios-theme/icons/order.svg",
  },
  province: {
    default: "ios-theme/icons/province.svg",
  },
  city: {
    default: "ios-theme/icons/city.svg",
  },
  about: {
    default: "ios-theme/icons/about.svg",
  },
  sms: {
    default: "ios-theme/icons/sms.svg",
  },
  product: {
    default: "ios-theme/icons/product.svg",
  },
  discount: {
    default: "ios-theme/icons/discount.svg",
  },
  tag: {
    default: "ios-theme/icons/tag.svg",
  },
  category: {
    default: "ios-theme/icons/category.svg",
  },
  brand: {
    default: "ios-theme/icons/brand.svg",
  },
  form: {
    default: "ios-theme/icons/form.svg",
  },
};

export const osResources = {
  dashboard: icons.dashboard[OS]
    ? icons.dashboard[OS]
    : icons.dashboard.default,
  up: icons.up[OS] ? icons.up[OS] : icons.up.default,
  questionBank: icons.questionBank[OS]
    ? icons.questionBank[OS]
    : icons.questionBank.default,
  down: icons.down[OS] ? icons.down[OS] : icons.down.default,
  edit: icons.edit[OS] ? icons.edit[OS] : icons.edit.default,
  add: icons.add[OS] ? icons.add[OS] : icons.add.default,
  cancel: icons.cancel[OS] ? icons.cancel[OS] : icons.cancel.default,
  delete: icons.delete[OS] ? icons.delete[OS] : icons.delete.default,
  discount: icons.discount[OS] ? icons.discount[OS] : icons.discount.default,
  cart: icons.cart[OS] ? icons.cart[OS] : icons.cart.default,
  entity: icons.entity[OS] ? icons.entity[OS] : icons.entity.default,
  sms: icons.sms[OS] ? icons.sms[OS] : icons.sms.default,
  left: icons.left[OS] ? icons.left[OS] : icons.left.default,
  brand: icons.brand[OS] ? icons.brand[OS] : icons.brand.default,
  menu: icons.menu[OS] ? icons.menu[OS] : icons.menu.default,
  right: icons.right[OS] ? icons.right[OS] : icons.right.default,
  settings: icons.settings[OS] ? icons.settings[OS] : icons.settings.default,
  dataNode: icons.dataNode[OS] ? icons.dataNode[OS] : icons.dataNode.default,
  user: icons.user[OS] ? icons.user[OS] : icons.user.default,
  city: icons.city[OS] ? icons.city[OS] : icons.city.default,
  province: icons.province[OS] ? icons.province[OS] : icons.province.default,
  about: icons.about[OS] ? icons.about[OS] : icons.about.default,
  turnoff: icons.turnoff[OS] ? icons.turnoff[OS] : icons.turnoff.default,
  ctrlSheet: icons.ctrlSheet[OS]
    ? icons.ctrlSheet[OS]
    : icons.ctrlSheet.default,
  country: icons.country[OS] ? icons.country[OS] : icons.country.default,
  export: icons.export[OS] ? icons.export[OS] : icons.export.default,
  gpio: icons.ctrlSheet[OS] ? icons.ctrlSheet[OS] : icons.ctrlSheet.default,
  country: icons.country[OS] ? icons.country[OS] : icons.country.default,
  order: icons.order[OS] ? icons.order[OS] : icons.order.default,
  export: icons.export[OS] ? icons.export[OS] : icons.export.default,
  mqtt: icons.mqtt[OS] ? icons.mqtt[OS] : icons.mqtt.default,

  gpio: icons.gpio[OS] ? icons.gpio[OS] : icons.gpio.default,

  tag: icons.tag[OS] ? icons.tag[OS] : icons.tag.default,
  product: icons.product[OS] ? icons.product[OS] : icons.product.default,
  category: icons.category[OS] ? icons.category[OS] : icons.category.default,
  form: icons.form[OS] ? icons.form[OS] : icons.form.default,

  gpiomode: icons.gpiomode[OS] ? icons.gpiomode[OS] : icons.gpiomode.default,
  backup: icons.backup[OS] ? icons.backup[OS] : icons.backup.default,
  gpiostate: icons.gpiostate[OS]
    ? icons.gpiostate[OS]
    : icons.gpiostate.default,
};
