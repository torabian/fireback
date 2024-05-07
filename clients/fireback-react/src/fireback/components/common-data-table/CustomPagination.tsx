import { localizeNumber } from "@/fireback/hooks/fonts";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useT } from "@/fireback/hooks/useT";
import classNames from "classnames";

export interface PageItem {
  label: string;
  page: number;
  active: boolean;
}

function calcVisiblePages({
  totalCount,
  pageSize,
  currentPage,
}: {
  totalCount: number;
  pageSize: number;
  currentPage: number;
}) {
  const totalPages = Math.ceil(totalCount / pageSize);
  const items: Array<{ label: string; page: number }> = [];

  let begin = currentPage - 2;
  let end = currentPage + 3;
  if (begin < 0) {
    begin = 0;
    end = 5;
  }
  if (end > totalPages) {
    end = totalPages;
    begin = totalPages - 5;
  }

  if (totalPages < 5) {
    begin = 0;
    end = totalPages;
  }

  for (let i = begin; i < end; i++) {
    items.push({ label: `${i + 1}`, page: i });
  }

  return items;
}

export function CustomPagination({
  totalCount,
  currentPage,
  onCurrentPageChange,
  onPageSizeChange,
  pageSize,
}: {
  totalCount: number;
  currentPage: number;
  onCurrentPageChange: (page: number) => void;
  pageSize: number;
  onPageSizeChange: () => void;
}) {
  const totalPages = Math.ceil(totalCount / pageSize);
  const t = useT();
  const pages = calcVisiblePages({ totalCount, pageSize, currentPage });
  const { locale } = useLocale();

  return (
    <nav>
      <ul className="pagination">
        <li
          className={classNames("page-item", { disabled: currentPage === 0 })}
        >
          <button
            onClick={() => onCurrentPageChange(currentPage - 1)}
            className="page-link"
            disabled={currentPage === 0}
          >
            {t.table.previous}
          </button>
        </li>
        {pages.map((page) => {
          return (
            <li
              key={page.page}
              className={classNames("page-item", {
                active: page.page === currentPage,
              })}
            >
              <button
                onClick={() => onCurrentPageChange(page.page)}
                className="page-link"
              >
                {localizeNumber(page.label, locale)}
              </button>
            </li>
          );
        })}

        <li
          className={classNames("page-item", {
            disabled: currentPage === totalPages - 1,
          })}
        >
          <button
            onClick={() => onCurrentPageChange(currentPage + 1)}
            className="page-link"
            disabled={currentPage === totalPages - 1}
          >
            {t.table.next}
          </button>
        </li>
      </ul>
    </nav>
  );
}
