import React, { useState } from "react";
import { IReactiveSearchResult } from "./ReactiveSearchDefinition";

export interface IReactiveSearchContext {
  result: Array<IReactiveSearchResult>;
  phrase: string;
  setResult: (data: IReactiveSearchResult[]) => void;
  appendResult: (data: IReactiveSearchResult) => void;
  setPhrase: (data: string) => void;
  reset: () => void;
}

export const ReactiveSearchContext =
  React.createContext<IReactiveSearchContext>({
    result: [],
    setResult() {},
    reset() {},
    appendResult() {},
    setPhrase() {},
    phrase: "",
  });

export function ReactiveSearchProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [phrase, setPhrase] = useState("");
  const [result, setResult] = useState<Array<IReactiveSearchResult>>([]);
  const appendResult = (result: IReactiveSearchResult) => {
    setResult((v) => [...v, result]);
  };

  const reset = () => {
    setPhrase("");
    setResult([]);
  };

  return (
    <ReactiveSearchContext.Provider
      value={{
        result,
        setResult,
        reset,
        appendResult,
        setPhrase,
        phrase,
      }}
    >
      {children}
    </ReactiveSearchContext.Provider>
  );
}
