import { ReactNode } from "react";

export const Showcase = ({ children }: { children: ReactNode }) => {
  return <div style={{ marginBottom: "70px" }}>{children}</div>;
};
