import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module2Config } from "./defs";
import { ConfigEditor } from "./ConfigEditor";

import { Arrow } from "./Arrow";
import { deburr } from "lodash";

function tokenize(keyword: string): string {
  return deburr(keyword).toLowerCase();
}

function searchInMemory(config: Module2Config[], search): Module2Config[] {
  if (!search) {
    return config;
  }

  return config.filter((config) =>
    [config?.name, config.description]
      .map(tokenize)
      .join(" - ")
      .includes(tokenize(search))
  );
}

export function Designerconfig({ content, setContent, search }: any) {
  const addConfig = () => {
    content.config = [{ name: "" }, ...(content.config || [])];

    setContent({ ...content });
    setItem("item-0", { initialEntered: true });
  };

  const deleteConfigAt = (index: number) => {
    content.config = content.config?.filter((_, index2) => index !== index2);
    setContent({ ...content });
  };

  const updateConfig = (config: Module2Config, index: number) => {
    setContent((c) => {
      c.config = (c.config || []).map((origin, ind) =>
        ind === index ? config : origin
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
  const config = searchInMemory(content.config, search);

  return (
    <div className="config-designer">
      <h1 className="section-bar">
        <span>config ({content?.config?.length || 0})</span>{" "}
        <button className="btn btn-sm btn-primary" onClick={addConfig}>
          Add config
        </button>
      </h1>

      {!config || config?.length === 0 ? (
        <span>No config in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(config || [])?.map((config, index) => {
          return (
            <AccordionItem
              key={config.name + "_" + config.description}
              itemKey={`item-${config.name + "_" + config.description}`}
              itemID={`item-${config.name + "_" + config.description}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{config.name}</div>
                  <div className="description">{config.description}</div>
                  <Arrow isUpward={state.isEnter} />
                  <button
                    className="mr-2"
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteConfigAt(index);
                    }}
                  >
                    Del
                  </button>
                </div>
              )}
            >
              <ConfigEditor
                config={config}
                onChange={(config) => updateConfig(config, index)}
              />
            </AccordionItem>
          );
        })}
      </ControlledAccordion>
    </div>
  );
}
