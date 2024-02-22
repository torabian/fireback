import { RemoteRequestOption } from "@/definitions/JSONStyle";
import React, { useState } from "react";

export interface IRemoteQueryContext {
  setOptions: (options: RemoteRequestOption) => void;
  options: RemoteRequestOption;
}

export const RemoteQueryContext = React.createContext<IRemoteQueryContext>({
  setOptions() {},
  options: {},
});

export function RemoteQueryProvider({
  children,
  initialOptions,
}: {
  children: React.ReactNode;
  initialOptions: RemoteRequestOption;
}) {
  const [options, setOptions] = useState<any>(initialOptions);

  return (
    <RemoteQueryContext.Provider value={{ setOptions, options }}>
      {children}
    </RemoteQueryContext.Provider>
  );
}
