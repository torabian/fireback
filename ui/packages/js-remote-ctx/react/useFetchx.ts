import { createContext, useContext } from "react";
import { FetchxContext } from "../common/fetchx";

const FetchxContextReact = createContext<FetchxContext | null>(null);

export const FetchxProvider = FetchxContextReact.Provider;

export function useFetchxContext(): FetchxContext | null {
  return useContext(FetchxContextReact);
}
