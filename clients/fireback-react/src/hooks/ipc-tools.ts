import { RemoteRequestOption } from "src/sdk/fireback";

export const ipcExecFn = (options: RemoteRequestOption) => {
  return function (method: string, url: string, body: any) {
    const finalUrl =
      url +
      "&token=" +
      options.headers?.authorization +
      "&workspaceId=" +
      ((options.headers && options.headers["workspace-id"]) || "");

    const action = getParameterByName("action", url);
    if (!action) {
      return;
    }

    const bridge = (window as any)?.go?.main?.AppIPCBridge;
    const ipcFunction = bridge && bridge[action];

    if (!ipcFunction) {
      return Promise.reject({
        error: {
          message: `IPC Function is not available. You cannot solve this problem, please report to us. ('${action}' missing)`,
        },
      });
    }

    return ipcFunction(JSON.stringify(body), finalUrl).then((re: any) => {
      // const el = document.getElementById("msg");
      // if (el) {
      //   el.innerHTML = "GOOD" + re;
      // }
      return JSON.parse(re);
    });
    // .catch((err: any) => {
    // const el = document.getElementById("msg");
    // if (el) {
    //   el.innerHTML = "err:" + `${err}`;
    // }
    // });
  };
};

export function getParameterByName(name: string, url = window.location.href) {
  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return "";
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}

// For future reference, if we wanted to switch to IPC
// const switchToIpc = () => {
//   fbContextFn(() => {
//     return ipcExecFn;
//   });
//   acContextFn(() => {
//     return ipcExecFn;
//   });
// };
// const { setExecFn: fbContextFn } = useContext(FirebackContext) as any;
// const { setExecFn: acContextFn } = useContext(AcademyContext) as any;

// const switchToHttp = () => {
//   fbContextFn(undefined);
//   acContextFn(undefined);
// };

// useEffect(() => {
//   fbContextFn(() => {
//     return (options: any) =>
//       mockExecFn(options, academyMockServer.current, t);
//   });
//   acContextFn(() => {
//     return (options: any) =>
//       mockExecFn(options, academyMockServer.current, t);
//   });
// }, []);

// const onRemoteChange = (mode: "ipc" | "remote") => {
// if (mode === "ipc") {
//   switchToIpc();
// }
// if (mode === "remote") {
//   switchToHttp();
// }
// };
