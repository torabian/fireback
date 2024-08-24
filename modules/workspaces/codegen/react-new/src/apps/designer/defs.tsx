type ErrorItem = { [key: string]: string };

export interface SecurityModel {}

export interface Module2 {
  path?: string;
  description?: string;
  version?: string;
  name?: string;
  entities?: Module2Entity[];
  tasks?: Module2Task[];
  dtos?: Module2DtoBase[];
  actions?: Module2Action[];
  macros?: Module2Macro[];
  remotes?: Module2Remote[];
  messages?: Module2Message;
}

export interface Module2Task {
  cron?: string;
  name: string;
  in?: Module2ActionBody;
}

export interface Module2Remote {
  method?: string;
  url?: string;
  out?: Module2ActionBody;
  responseFields?: Module2Field[];
  in?: Module2ActionBody;
  query?: Module2Field[];
  name?: string;
}

export interface Module2FieldOf {
  k?: string;
}

export interface Module2Macro {
  using?: string;
  name?: string;
  fields?: Module2Field[];
}

export interface Module2Field {
  recommended?: boolean;
  linkedTo?: string;
  description?: string;
  name?: string;
  type?: string;
  primitive?: string;
  target?: string;
  rootClass?: string;
  validate?: string;
  excerptSize?: number;
  default?: any;
  translate?: boolean;
  unsafe?: boolean;
  allowCreate?: boolean;
  module?: string;
  provider?: string;
  json?: string;
  ofType?: Module2FieldOf[];
  yaml?: string;
  idFieldGorm?: string;
  computedType?: string;
  computedTypeClass?: string;
  belongingEntityName?: string;
  matches?: Module2FieldMatch[];
  gorm?: string;
  gormMap?: GormOverrideMap;
  sql?: string;
  fullName?: string;
  fields?: Module2Field[];
}

export interface Module2FieldMatch {
  dto?: string;
}

export interface GormOverrideMap {
  workspaceId?: string;
  userId?: string;
}

export interface Security {
  model?: string;
}

export interface Module2Http {
  query?: boolean;
}

export interface Module2Permission {
  name?: string;
  key?: string;
  description?: string;
}

type Module2Message = { [key: string]: { [key: string]: string } };

export interface Module2Entity {
  permissions?: Module2Permission[];
  name?: string;
  distinctBy?: string;
  prependScript?: string;
  messages?: Module2Message;
  prependCreateScript?: string;
  prependUpdateScript?: string;
  noQuery?: boolean;
  access?: string;
  queryScope?: string;
  securityModel?: SecurityModel;
  http?: Module2Http;
  patch?: boolean;
  queries?: string[];
  get?: boolean;
  gormMap?: GormOverrideMap;
  query?: boolean;
  post?: boolean;
  importList?: string[];
  fields?: Module2Field[];
  c?: boolean;
  cliName?: string;
  cliShort?: string;
  cliDescription?: string;
  cte?: boolean;
  postFormatter?: string;
}

export interface Module2DtoBase {
  name?: string;
  importList?: string[];
  fields?: Module2Field[];
}

export interface Module2ActionBody {
  fields?: Module2Field[];
  dto?: string;
  entity?: string;
}

export interface Module2Action {
  actionName?: string;
  cliName?: string;
  actionAliases?: string[];
  name?: string;
  url?: string;
  method?: string;
  query?: Module2Field[];
  fn?: string;
  description?: string;
  group?: string;
  format?: string;
  in?: Module2ActionBody;
  out?: Module2ActionBody;
  securityModel?: SecurityModel;
}
