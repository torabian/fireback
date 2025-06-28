export const snippets = {
  "example1": `const example1 = () => {
    return openDrawer(() => <div>Hi, this is opened in a drawer</div>);
  }`,
  "example2": `const example2 = () => {
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
  }`,
  "example3": `const example3 = () => {
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
  }`,
  "example4": `const example4 = () => {
    example1();
    example1();
    example2();
    example2();
  }`,
  "example5": `const example5 = () => {
    openModal(({ close }) => {
      useEffect(() => {
        setTimeout(() => {
          close();
        }, 3000);
      }, []);

      return <span>I will disappear :)))))</span>;
    });
  }`,
  "example6": `const example6 = () => {
    const { close, id } = openModal(() => {
      return <span>I will disappear by outside :)))))</span>;
    });

    setTimeout(() => {
      alert(id);
      close();
    }, 2000);
  }`,
  "example7": `const example7 = () => {
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
  }`,
  "example8": `const example8 = () => {
    confirmDrawer({
      title: "Confirm",
      description: "Are you to confirm? You still can cancel",
      confirmLabel: "Confirm",
      cancelLabel: "Cancel",
    }).promise.then((result) => {
      console.log(10, result);
    });
  }`,
  "example9": `const example9 = () => {
    confirmModal({
      title: "Confirm",
      description: "Are you to confirm? You still can cancel",
      confirmLabel: "Confirm",
      cancelLabel: "Cancel",
    }).promise.then((result) => {
      console.log(10, result);
    });
  }`,
  "example10": `const example10 = () => {
    const { updateData, promise } = openDrawer(({ data }) => {
      return <span>Params: {JSON.stringify(data)}</span>;
    });

    const id = setInterval(() => {
      updateData({ c: ++counter.current } as any);
    }, 100);

    promise.finally(() => {
      clearInterval(id);
    });
  }`
};
