// @ts-nocheck
import { createContext, useContext } from "react";
import { FetchxContext } from "../common/fetchx";
const FetchxContextReact = createContext<FetchxContext>(null);
export const FetchxProvider = FetchxContextReact.Provider;
export function useFetchxContext(): FetchxContext {
  return useContext(FetchxContextReact);
}
