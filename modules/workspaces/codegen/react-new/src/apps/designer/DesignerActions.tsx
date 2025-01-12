import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module3Action } from "./defs";
import { Arrow } from "./Arrow";
import { ActionEditor } from "./ActionEditor";

export function DesignerActions({ content, setContent }: any) {
  const addAction = () => {
    content.actions = [{ name: "" }, ...(content.actions || [])];

    setContent({ ...content });
    setItem("item-0", { initialEntered: true });
  };

  const deleteActionAt = (index: number) => {
    content.actions = content.actions?.filter((_, index2) => index !== index2);
    setContent({ ...content });
  };

  const updateAction = (action: Module3Action, index: number) => {
    setContent((c) => {
      c.actions = (c.actions || []).map((origin, ind) =>
        ind === index ? action : origin
      );
      return { ...c };
    });
  };

  const providerValue = useAccordionProvider({
    allowMultiple: true,
    transition: true,
    transitionTimeout: 250,
  });

  // Destructuring `toggle` and `toggleAll` from `providerValue`
  const { setItem } = providerValue;

  return (
    <div className="actions-designer">
      <h1 className="section-bar">
        <span>Actions ({content?.actions?.length || 0})</span>
        <button className="btn btn-sm btn-primary" onClick={addAction}>
          Add action
        </button>
      </h1>

      {!content.actions || content.actions?.length === 0 ? (
        <span>No actions in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(content.actions || [])?.map((action, index) => {
          return (
            <AccordionItem
              key={index}
              itemKey={`item-${index}`}
              itemID={`item-${index}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{action.name}</div>
                  <div className="description">{action.description}</div>
                  <Arrow isUpward={state.isEnter} />
                  <button
                    className="mr-2"
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteActionAt(index);
                    }}
                  >
                    Del
                  </button>
                </div>
              )}
            >
              <ActionEditor
                action={action}
                onChange={(action) => updateAction(action, index)}
              />
            </AccordionItem>
          );
        })}
      </ControlledAccordion>
    </div>
  );
}
