export type JsonQuery = any;

export interface Context {
  url: string;
  token: string;
  workspaceId: string;
  body: any;
  acceptLanguage: string;
  method: string;
  itemsPerPage: number;
  startIndex?: number;
  paramValues: Array<string>;
}
