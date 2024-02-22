import { source } from "@/helpers/source";
import { useT } from "@/hooks/useT";
import { groupBy } from "lodash";
import { useContext } from "react";
import Link from "../link/Link";
import { ReactiveSearchContext } from "./ReactiveSearchContext";
import { IReactiveSearchResult } from "./ReactiveSearchDefinition";

export function ReactiveSearchResult({
  result,
  onComplete,
}: {
  result: IReactiveSearchResult[];
  onComplete: () => void;
}) {
  const t = useT();
  const renderGroup = groupBy(result, "group");
  const keys = Object.keys(renderGroup);

  return (
    <div className="reactive-search-result">
      {keys.length === 0 ? (
        <>{t.reactiveSearch.noResults}</>
      ) : (
        <ul>
          {keys.map((groupName, index) => {
            return (
              <li key={index}>
                <span className="result-group-name">{groupName}</span>
                <ul>
                  {renderGroup[groupName].map((inner, index2) => {
                    return (
                      <li key={inner.uniqueId}>
                        {inner.actionFn ? (
                          <Link onClick={onComplete} href={inner.uiLocation}>
                            {inner.icon && (
                              <img
                                className="result-icon"
                                src={source(inner.icon)}
                              />
                            )}
                            {inner.phrase}
                          </Link>
                        ) : null}
                      </li>
                    );
                  })}
                </ul>
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}
