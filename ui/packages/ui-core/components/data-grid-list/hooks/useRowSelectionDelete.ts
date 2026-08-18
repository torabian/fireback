import { useEffect, useState } from "react";
import {
  type IMenuActionItem,
  useMenuTools,
} from "../../action-menu/ActionMenu";
import { commonDialogs } from "../../overlay/CommonOverlays";
import { strings } from "../../strings/translations";
import { httpErrorHanlder } from "../../../hooks/api";
import { osResources } from "../../../hooks/resources";
import { KeyboardAction } from "../../../hooks/useExportTools";
import { useKeyCombination } from "../../../hooks/useKeyPress";
import { useS } from "../../../hooks/useS";

// Selection + delete, ported as-is from useDatatableFiltering's deleteItems/
// deleteAction/selection wiring (ui-core/hooks/useDatatableFiltering.tsx) - same
// confirm-dialog copy, same "resolved-but-failed response" unwrapping bug fix, same
// action-menu integration and Delete-key shortcut, same submitDelete(uniqueIds)
// contract. Split out as its own hook (rather than folded into DataGridList
// itself) so it can be dropped or swapped independently of paging/filtering/
// scrolling, per DataGridList's one-hook-per-feature layout.
export function useRowSelectionDelete({
  submitDelete,
  onRecordsDeleted,
}: {
  submitDelete?: (body: string, opts: any) => Promise<any>;
  onRecordsDeleted?: (items: string[]) => void;
}) {
  const s = useS(strings);
  const { confirmModal } = commonDialogs();
  const { addActions, removeActionMenu } = useMenuTools();

  const [selection, setSelection] = useState<string[]>([]);

  const deleteItems = async () => {
    confirmModal({
      title: s.confirm,
      confirmLabel: s.common.yes,
      cancelLabel: s.common.no,
      description: s.deleteConfirmMessage,
    })
      .promise.then(({ type }) => {
        if (type !== "resolved" || !submitDelete) {
          return;
        }

        return submitDelete(
          JSON.stringify({ uniqueIds: selection }),
          null as any,
        ).then((response: any) => {
          // A rejected delete (e.g. WorkspaceAwareDeleteAction refusing to delete
          // the root workspace) resolves rather than rejects - our fetch layer
          // hands back a {error: {...}} envelope instead of throwing - so this
          // has to be checked explicitly rather than assuming .then() only ever
          // runs on real success.
          const errorInfo = response?.error?.toJSON?.() ?? response?.error;
          if (
            errorInfo?.message ||
            errorInfo?.messageTranslated ||
            errorInfo?.errors?.length
          ) {
            httpErrorHanlder({ error: errorInfo }, s);
            return;
          }

          setSelection([]);
          onRecordsDeleted?.(selection);
        });
      })
      .catch((err) => httpErrorHanlder(err, s));
  };

  const deleteAction = (): IMenuActionItem => ({
    label: s.deleteAction,
    onSelect() {
      deleteItems();
    },
    icon: osResources.delete,
    uniqueActionKey: "GENERAL_DELETE_ACTION",
  });

  useEffect(() => {
    if (selection.length > 0 && typeof submitDelete !== "undefined") {
      return addActions("table-selection", [deleteAction()]);
    }
    removeActionMenu("table-selection");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [selection]);

  useKeyCombination(KeyboardAction.Delete, () => {
    if (selection.length > 0 && typeof submitDelete !== "undefined") {
      deleteItems();
    }
  });

  return { selection, setSelection, deleteItems };
}

export type RowSelectionDelete = ReturnType<typeof useRowSelectionDelete>;
