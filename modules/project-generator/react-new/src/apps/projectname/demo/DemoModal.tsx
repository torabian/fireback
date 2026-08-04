import { FormButton } from "@/modules/fireback/components/forms/form-button/FormButton";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { commonDialogs } from "@/modules/fireback/components/overlay/CommonOverlays";
import { useOverlay } from "@/modules/fireback/components/overlay/OverlayProvider";
import { useEffect, useRef, useState } from "react";
import { CodeViewer } from "./CodeViewer";
import { snippets } from "./DemoModal.snippets";
import { Showcase } from "./Showcase";

export function DemoModal() {
  const { openDrawer, openModal } = useOverlay();
  const { confirmDrawer, confirmModal } = commonDialogs();

  const example1 = () => {
    return openDrawer(() => <div>Hi, this is opened in a drawer</div>);
  };

  const example2 = () => {
    openDrawer(
      () => (
        <div>
          Hi, this is opened in a drawer, with a larger area and from left
        </div>
      ),
      {
        direction: "left",
        size: "40%",
      }
    );
  };

  const example3 = () => {
    openModal<string>(({ resolve }) => {
      const [value, setValue] = useState("");
      return (
        <form onSubmit={(e) => e.preventDefault()}>
          <span>
            If you enter <strong>ali</strong> in the box, you'll see the
            example1 opening
          </span>
          <FormText autoFocus value={value} onChange={(e) => setValue(e)} />
          <FormButton onClick={() => resolve(value)}>Okay</FormButton>
        </form>
      );
    }).promise.then(({ data }) => {
      if (data === "ali") {
        return example1();
      }
      alert(data);
    });
  };

  const example4 = () => {
    example1();
    example1();
    example2();
    example2();
  };

  const example5 = () => {
    openModal(({ close }) => {
      useEffect(() => {
        setTimeout(() => {
          close();
        }, 3000);
      }, []);

      return <span>I will disappear :)))))</span>;
    });
  };

  const example6 = () => {
    const { close, id } = openModal(() => {
      return <span>I will disappear by outside :)))))</span>;
    });

    setTimeout(() => {
      alert(id);
      close();
    }, 2000);
  };

  const example7 = () => {
    openModal(({ setOnBeforeClose }) => {
      const [dirty, setDirty] = useState(false);

      useEffect(() => {
        setOnBeforeClose?.(() => {
          if (!dirty) return true;
          return window.confirm("You have unsaved changes. Close anyway?");
        });
      }, [dirty]);

      return (
        <span>
          If you write anything here, it will be dirty and asks for quite.
          <input onChange={() => setDirty(true)} />
          {dirty ? "Will ask" : "Not dirty yet"}
        </span>
      );
    });
  };

  const example8 = () => {
    confirmDrawer({
      title: "Confirm",
      description: "Are you to confirm? You still can cancel",
      confirmLabel: "Confirm",
      cancelLabel: "Cancel",
    }).promise.then((result) => {
      console.log(10, result);
    });
  };

  const example9 = () => {
    confirmModal({
      title: "Confirm",
      description: "Are you to confirm? You still can cancel",
      confirmLabel: "Confirm",
      cancelLabel: "Cancel",
    }).promise.then((result) => {
      console.log(10, result);
    });
  };

  const counter = useRef(0);

  const example10 = () => {
    const { updateData, promise } = openDrawer(({ data }) => {
      return <span>Params: {JSON.stringify(data)}</span>;
    });

    const id = setInterval(() => {
      updateData({ c: ++counter.current } as any);
    }, 100);

    promise.finally(() => {
      clearInterval(id);
    });
  };

  return (
    <div>
      <h1>Demo Modals</h1>
      <p>
        Modals, Drawers are a major solved issue in the Fireback react.js. In
        here we make some examples. The core system is called `overlay`, can be
        used to show portals such as modal, drawer, alerts...
      </p>
      <hr />

      <Showcase>
        <h2>Opening a drawer</h2>
        <p>
          Every component can be shown as modal, or in a drawer in Fireback.
        </p>
        <button className="btn btn-sm btn-secondary" onClick={() => example1()}>
          Open a text in drawer
        </button>
        <CodeViewer codeString={snippets.example1} />
      </Showcase>
      <Showcase>
        <h2>Opening a drawer, from left</h2>
        <p>Shows a drawer from left, also larger</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example2()}>
          Open a text in drawer
        </button>
        <CodeViewer codeString={snippets.example2} />
      </Showcase>

      <Showcase>
        <h2>Opening a modal, and get result</h2>
        <p>
          You can open a modal or drawer, and make some operation in it, and
          send back the result as a promise.
        </p>
        <button className="btn btn-sm btn-secondary" onClick={() => example3()}>
          Open a text in drawer
        </button>
        <CodeViewer codeString={snippets.example3} />
      </Showcase>

      <Showcase>
        <h2>Opening multiple</h2>
        <p>You can open multiple modals, or drawers, doesn't matter.</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example4()}>
          Open 2 modal, and open 2 drawer
        </button>
        <CodeViewer codeString={snippets.example4} />
      </Showcase>

      <Showcase>
        <h2>Auto disappearing</h2>
        <p>A modal which disappears after 5 seconds</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example5()}>
          Run
        </button>
        <CodeViewer codeString={snippets.example5} />
      </Showcase>

      <Showcase>
        <h2>Control from outside</h2>
        <p>
          Sometimes you want to open a drawer, and then from outside component
          close it.
        </p>
        <button className="btn btn-sm btn-secondary" onClick={() => example6()}>
          Open but close from outside
        </button>
        <CodeViewer codeString={snippets.example6} />
      </Showcase>

      <Showcase>
        <h2>Prevent close</h2>
        <p>When a drawer or modal is open, you can prevent the close.</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example7()}>
          Open but ask before close
        </button>
        <CodeViewer codeString={snippets.example7} />
      </Showcase>
      <Showcase>
        <h2>Confirm Dialog (drawer)</h2>
        <p>There is a set of ready to use dialogs, such as confirm</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example8()}>
          Open the confirm
        </button>
        <CodeViewer codeString={snippets.example8} />
      </Showcase>
      <Showcase>
        <h2>Confirm Dialog (modal)</h2>
        <p>There is a set of ready to use dialogs, such as confirm</p>
        <button className="btn btn-sm btn-secondary" onClick={() => example9()}>
          Open the confirm
        </button>
        <CodeViewer codeString={snippets.example9} />
      </Showcase>
      <Showcase>
        <h2>Update params from outside</h2>
        <p>
          In rare cases, you might want to update the params from the outside.
        </p>
        <button
          className="btn btn-sm btn-secondary"
          onClick={() => example10()}
        >
          Open & Update name
        </button>
        <CodeViewer codeString={snippets.example10} />
      </Showcase>
      <br />
      <br />
      <br />
    </div>
  );
}
