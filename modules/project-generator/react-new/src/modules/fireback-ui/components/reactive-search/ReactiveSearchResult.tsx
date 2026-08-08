import { groupBy } from "lodash";
import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import Link from "../link/Link";
import { type IReactiveSearchResult } from "./ReactiveSearchDefinition";

export function ReactiveSearchResult({
  result,
  onComplete,
}: {
  result: IReactiveSearchResult[];
  onComplete: () => void;
}) {
  const s = useS(strings);
  const renderGroup = groupBy(result, "group");
  const keys = Object.keys(renderGroup);

  return (
    <div className="reactive-search-result">
      {keys.length === 0 ? (
        <>{s.reactiveSearch.noResults}</>
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
