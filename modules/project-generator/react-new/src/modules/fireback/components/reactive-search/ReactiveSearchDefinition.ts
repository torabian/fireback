// Mirrors ReactiveSearchResultDtoType (modules/reactivesearch/ReactiveSearch.emi.yml's
// reactiveSearchResult dto) - was previously out of sync with it (had an unused
// "type" field, and was missing "group", even though ReactiveSearchResult.tsx's own
// groupBy(result, "group") already reads it at runtime).
export interface IReactiveSearchResult {
  description: string;
  group: string;
  icon: string;
  phrase: string;
  uniqueId: string;
  uiLocation: string;
  actionFn: string;
}
