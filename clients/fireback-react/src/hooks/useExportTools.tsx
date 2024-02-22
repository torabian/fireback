import { useT } from "@/hooks/useT";

import { useExportActions } from "@/components/action-menu/ActionMenu";
import { KeyboardAction } from "@/definitions/definitions";
import { RemoteQueryContext } from "@/sdk/fireback/core/react-tools";
import { useContext } from "react";

function toBinaryString(data: any) {
  var ret = [];
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
  roleId: string
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
    false
  );

  xhr.setRequestHeader("Authorization", token);
  xhr.setRequestHeader("Workspace-Id", workspaceId);
  xhr.setRequestHeader("role-Id", roleId);
  xhr.overrideMimeType("application/octet-stream; charset=x-user-defined;");
  xhr.send(null);
}

export function useOctetDownload({ path }: { path: string }) {
  const { options } = useContext(RemoteQueryContext);
  const h: any = options?.headers;

  const execute = () =>
    xhrStreamFile(
      options.prefix + "" + path,
      "POST",
      h.authorization || "",
      h["workspace-id"] || "",
      h["role-id"] || ""
    );

  return { execute };
}

export const useExportTools = ({ path }: { path?: string }) => {
  const t = useT();
  const { options } = useContext(RemoteQueryContext);

  useExportActions(
    path
      ? () => {
          const h: any = options?.headers;

          xhrStreamFile(
            options.prefix + "" + path,
            "GET",
            h.authorization || "",
            h["workspace-id"] || "",
            h["role-id"] || ""
          );
        }
      : undefined,
    KeyboardAction.ExportTable
  );
};
