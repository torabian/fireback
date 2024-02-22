import { localizeNumber } from "@/hooks/fonts";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import classNames from "classnames";

export function CustomPageSize({
  pageSizes,
  onPageSizeChange,
  currentPageSize,
}: {
  pageSizes: Array<number>;
  onPageSizeChange: (size: number) => void;
  currentPageSize: number;
}) {
  const t = useT();
  const { locale } = useLocale();

  return (
    <nav>
      <ul className="pagination">
        {pageSizes.map((page) => {
          return (
            <li
              onClick={() => onPageSizeChange(page)}
              key={page}
              className={classNames("page-item", {
                active: page === currentPageSize,
              })}
            >
              <button type="button" className="page-link">
                {localizeNumber(page + "", locale)}
              </button>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
