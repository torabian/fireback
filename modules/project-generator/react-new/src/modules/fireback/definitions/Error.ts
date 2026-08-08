export interface ApiError {
  title?: string;
  message?: string;
  errors: { [key: string]: Array<string> };
}
