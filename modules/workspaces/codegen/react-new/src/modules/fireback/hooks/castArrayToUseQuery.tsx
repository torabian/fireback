/**
 * Often it's needed to search through a constant array and make selection on front-end
 * side. Although this is not prefered, this function would take an array, and make it
 * look like a useQuery result from backend.
 *
 * At this version, search functionality is not added but you can add it later on
 * @param items
 * @returns
 */
export const castArrayToUseQuery = <T extends object>(items: T[]) => {
  return function useHook() {
    return {
      query: {
        data: {
          data: {
            items,
          },
          totalAvailableItems: (items || []).length,
        },
      },
    };
  };
};
