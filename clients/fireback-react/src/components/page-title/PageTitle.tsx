export function PageTitle({
  title,
  description,
  children,
}: {
  title: string;
  description?: string;
  children?: any;
}) {
  return (
    <div className="page-title">
      <div className="row g-0">
        <div className="col-lg-7">
          <h1>
            <PageTitleManager />
          </h1>
          <p className="lead">{description}</p>
        </div>
        {children}
      </div>
    </div>
  );
}

/**
 * Action menu stands for those menus, which can accept some buttons, and change
 * based on the page user is.
 */
import React, { useContext, useEffect, useState } from "react";

export function PageTitleManager() {
  const t = useContext(PageTitleContext);

  return <span>{t.ref.title}</span>;
}

export type onTriggerFn = (action: string) => void;

export interface PageTitleRef {
  title: string;
}

export type SetPageTitleFn = (menuName: string) => void;

export interface IPageTitleContext {
  ref: { title: string };
  setPageTitle: SetPageTitleFn;
  removePageTitle: (menuName: string) => void;
}

export const PageTitleContext = React.createContext<IPageTitleContext>({
  setPageTitle() {},
  removePageTitle() {},
  ref: {
    title: "",
  },
});

export interface PageTitleOptions {
  onTrigger: (actionKey: string) => void;
}

export function usePageTitle(title?: string) {
  const t = useContext(PageTitleContext);

  useEffect(() => {
    t.setPageTitle(title || "");

    return () => {
      t.removePageTitle("");
    };
  }, [title]);
}

export function PageTitleProvider({ children }: { children: React.ReactNode }) {
  const [title, setTitle] = useState("");

  const setPageTitle: SetPageTitleFn = (title) => {
    document.title = title;
    setTitle(title);
  };

  const removePageTitle = () => {
    document.title = "";
    setTitle("");
  };

  return (
    <PageTitleContext.Provider
      value={{
        ref: {
          title,
        },
        setPageTitle,
        removePageTitle,
      }}
    >
      {children}
    </PageTitleContext.Provider>
  );
}
