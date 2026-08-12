import { useExportActions } from "../components/action-menu/ActionMenu";

export enum KeyboardAction {
  NewEntity = "new_entity",
  SidebarToggle = "sidebarToggle",
  NewChildEntity = "new_child_entity",
  EditEntity = "edit_entity",
  ViewQuestions = "view_questions",
  ExportTable = "export_table",
  CommonBack = "common_back",
  StopStart = "StopStart",
  Delete = "delete",
  Select1Index = "select1_index",
  Select2Index = "select2_index",
  Select3Index = "select3_index",
  Select4Index = "select4_index",
  Select5Index = "select5_index",
  Select6Index = "select6_index",
  Select7Index = "select7_index",
  Select8Index = "select8_index",
  Select9Index = "select9_index",
  ToggleLock = "l",
}

export const NumericKeys = [
  KeyboardAction.Select1Index,
  KeyboardAction.Select2Index,
  KeyboardAction.Select3Index,
  KeyboardAction.Select4Index,
  KeyboardAction.Select5Index,
  KeyboardAction.Select6Index,
  KeyboardAction.Select7Index,
  KeyboardAction.Select8Index,
  KeyboardAction.Select9Index,
];

import { useApiOptions } from "./useApiOptions";

export function toBinaryString(data: any) {
  var ret: string[] = [];
  var len = data.length;
  var byte;
  for (var i = 0; i < len; i++) {
    byte = (data.charCodeAt(i) & 0xff) >>> 0;
    ret.push(String.fromCharCode(byte));
  }

  return ret.join("");
}

export function xhrStreamFile(
  path: string,
  method: string,
  token: string,
  workspaceId: string,
  roleId: string,
) {
  var xhr = new XMLHttpRequest();

  xhr.open(method, path);

  xhr.addEventListener(
    "load",
    function () {
      var data = toBinaryString(this.responseText);
      data = "data:application/text;base64," + btoa(data);
      document.location = data;
    },
    false,
  );

  xhr.setRequestHeader("Authorization", token);
  xhr.setRequestHeader("Workspace-Id", workspaceId);
  xhr.setRequestHeader("role-Id", roleId);
  xhr.overrideMimeType("application/octet-stream; charset=x-user-defined;");
  xhr.send(null);
}

export function useOctetDownload({ path }: { path: string }) {
  const options = useApiOptions();
  const h: any = options?.headers;

  const execute = () =>
    xhrStreamFile(
      options.prefix + "" + path,
      "POST",
      h.authorization || "",
      h["workspace-id"] || "",
      h["role-id"] || "",
    );

  return { execute };
}

export const useExportTools = ({ path }: { path?: string }) => {
  const options = useApiOptions();

  useExportActions(
    path
      ? () => {
          const h: any = options?.headers;

          xhrStreamFile(
            options.prefix + "" + path,
            "GET",
            h.authorization || "",
            h["workspace-id"] || "",
            h["role-id"] || "",
          );
        }
      : undefined,
    KeyboardAction.ExportTable,
  );
};
