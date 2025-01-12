type ErrorItem = { [key: string]: string };

export interface SecurityModel {}

export interface Module3Config {
  name?: string;
  type?: string;
  description?: string;
  default?: string;
  env?: string;
  fields?: Module3Config[];
}

export interface Module3 {
  path?: string;
  description?: string;
  version?: string;
  name?: string;
  entities?: Module3Entity[];
  tasks?: Module3Task[];
  dtos?: Module3DtoBase[];
  actions?: Module3Action[];
  macros?: Module3Macro[];
  remotes?: Module3Remote[];
  messages?: Module3Message;
}

export interface Module3Trigger {
  cron?: string;
}

export interface Module3Task {
  triggers?: Module3Trigger[];
  name: string;
  in?: Module3ActionBody;
}

export interface Module3Remote {
  method?: string;
  url?: string;
  out?: Module3ActionBody;
  responseFields?: Module3Field[];
  in?: Module3ActionBody;
  query?: Module3Field[];
  name?: string;
}

export interface Module3FieldOf {
  k?: string;
}

export interface Module3Macro {
  using?: string;
  name?: string;
  fields?: Module3Field[];
}

export interface Module3Field {
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
  ofType?: Module3FieldOf[];
  yaml?: string;
  idFieldGorm?: string;
  computedType?: string;
  computedTypeClass?: string;
  belongingEntityName?: string;
  matches?: Module3FieldMatch[];
  gorm?: string;
  gormMap?: GormOverrideMap;
  sql?: string;
  fullName?: string;
  fields?: Module3Field[];
}

export interface Module3FieldMatch {
  dto?: string;
}

export interface GormOverrideMap {
  workspaceId?: string;
  userId?: string;
}

export interface Security {
  model?: string;
}

export interface Module3Http {
  query?: boolean;
}

export interface Module3Permission {
  name?: string;
  key?: string;
  description?: string;
}

type Module3Message = { [key: string]: { [key: string]: string } };

export interface Module3Entity {
  permissions?: Module3Permission[];
  name?: string;
  distinctBy?: string;
  prependScript?: string;
  messages?: Module3Message;
  prependCreateScript?: string;
  prependUpdateScript?: string;
  noQuery?: boolean;
  access?: string;
  queryScope?: string;
  securityModel?: SecurityModel;
  http?: Module3Http;
  patch?: boolean;
  queries?: string[];
  get?: boolean;
  gormMap?: GormOverrideMap;
  query?: boolean;
  post?: boolean;
  importList?: string[];
  fields?: Module3Field[];
  c?: boolean;
  cliName?: string;
  cliShort?: string;
  description?: string;
  cte?: boolean;
  postFormatter?: string;
}

export interface Module3Dto extends Module3DtoBase {}
export interface Module3DtoBase {
  name?: string;
  importList?: string[];
  fields?: Module3Field[];
}

export interface Module3ActionBody {
  fields?: Module3Field[];
  dto?: string;
  entity?: string;
}

export interface Module3Action {
  cliName?: string;
  actionAliases?: string[];
  name?: string;
  url?: string;
  method?: string;
  query?: Module3Field[];
  fn?: string;
  description?: string;
  group?: string;
  format?: string;
  in?: Module3ActionBody;
  out?: Module3ActionBody;
  securityModel?: SecurityModel;
}
