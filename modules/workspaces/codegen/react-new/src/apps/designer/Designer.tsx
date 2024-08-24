import { useEffect, useRef, useState } from "react";
import { Module2, Module2Entity } from "./defs";
import { EntityEditor } from "./EntityEditor";
import {
  Accordion,
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";

import { mockYaml } from "./mock";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { Arrow } from "./Arrow";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";
const yaml = require("js-yaml");

function sendUpdate(content: any) {
  if (typeof vscode == "undefined" || !vscode?.postMessage) {
    return;
  }
  vscode.postMessage({ command: "updateFile", content });
}

const vscode = (window as any).acquireVsCodeApi
  ? (window as any).acquireVsCodeApi()
  : null;

export function Designer() {
  const [content, setContent] = useState<Module2>({});

  const parseYaml = (yamlString: string) => {
    try {
      const parsed = yaml.load(yamlString);
      setContent(parsed);
    } catch (e) {
      console.error("Error parsing YAML:", e);
    }
  };

  const lockReflex = useRef(true);

  useEffect(() => {
    // Use workspaces module as a predefined data
    if (!vscode) {
      parseYaml(mockYaml);
    }

    window.addEventListener("message", (event) => {
      if (event.data.command === "updateContent") {
        lockReflex.current = true;
        parseYaml(event.data.content);

        setTimeout(() => {
          lockReflex.current = false;
        }, 50);
      }
    });

    // For demo purpose

    // To send messages to the extension
  }, []);

  useEffect(() => {
    if (lockReflex.current) {
      return;
    }
    sendUpdate(yaml.dump(content));
  }, [content]);

  const addEntity = () => {
    content.entities = [{ name: "" }, ...(content.entities || [])];

    setContent({ ...content });
    // setTimeout(() => {
    //   (document.querySelector(".entities input") as any)?.focus();
    setItem("item-0", { initialEntered: true });
    // }, 50);
  };

  const addDto = () => {
    setContent((c) => {
      c.dtos = [{ name: "" }, ...(c.dtos || [])];
      return { ...c };
    });

    setTimeout(() => {
      (document.querySelector(".dtos input") as any)?.focus();
    }, 50);
  };

  const addTask = () => {
    setContent((c) => {
      c.tasks = [{ name: "" }, ...(c.tasks || [])];
      return { ...c };
    });

    setTimeout(() => {
      (document.querySelector(".tasks input") as any)?.focus();
    }, 50);
  };

  const addAction = () => {
    setContent((c) => {
      c.actions = [{ name: "" }, ...(c.actions || [])];
      return { ...c };
    });

    setTimeout(() => {
      (document.querySelector(".actions input") as any)?.focus();
    }, 50);
  };

  const setEntityName = (index: number, name: string) => {
    setContent((c) => {
      c.entities = c.entities?.map((entity, index2) => {
        if (index2 === index) {
          return {
            ...entity,
            name,
          };
        }
        return entity;
      });
      return { ...c };
    });
  };

  const setDtoName = (index: number, name: string) => {
    setContent((c) => {
      c.dtos = c.dtos?.map((dto, index2) => {
        if (index2 === index) {
          return {
            ...dto,
            name,
          };
        }
        return dto;
      });
      return { ...c };
    });
  };

  const setTaskName = (index: number, name: string) => {
    setContent((c) => {
      c.tasks = c.tasks?.map((task, index2) => {
        if (index2 === index) {
          return {
            ...task,
            name,
          };
        }
        return task;
      });
      return { ...c };
    });
  };

  const setActionName = (index: number, name: string) => {
    setContent((c) => {
      c.actions = c.actions?.map((action, index2) => {
        if (index2 === index) {
          return {
            ...action,
            name,
          };
        }
        return action;
      });
      return { ...c };
    });
  };

  const deleteEntityAt = (index: number) => {
    content.entities = content.entities?.filter(
      (_, index2) => index !== index2
    );
    setContent({ ...content });
  };

  const deleteDtoAt = (index: number) => {
    setContent((c) => {
      c.dtos = c.dtos?.filter((_, index2) => index !== index2);
      return { ...c };
    });
  };

  const deleteActionAt = (index: number) => {
    setContent((c) => {
      c.actions = c.actions?.filter((_, index2) => index !== index2);
      return { ...c };
    });
  };

  const deleteTaskAt = (index: number) => {
    setContent((c) => {
      c.tasks = c.tasks?.filter((_, index2) => index !== index2);
      return { ...c };
    });
  };

  const updateEntity = (entity: Module2Entity, index: number) => {
    setContent((c) => {
      c.entities = (c.entities || []).map((origin, ind) =>
        ind === index ? entity : origin
      );
      return { ...c };
    });
  };

  const changeModuleName = (name: string) => {
    setContent((c) => {
      c.name = name;
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

  const changeModuleDescription = (description: string) => {
    setContent((c) => {
      c.description = description;
      return { ...c };
    });
  };

  return (
    <div className="interactive container">
      <FormText
        label="Module name"
        value={content.name}
        onChange={(v) => {
          changeModuleName(v);
        }}
      />
      <FormRichText
        label="Module description"
        value={content.description}
        forceBasic
        onChange={(v) => {
          changeModuleDescription(v);
        }}
      />

      <h1 className="section-bar">
        <span>Entities ({content?.entities?.length || 0})</span>{" "}
        <button className="btn btn-sm btn-primary" onClick={addEntity}>
          Add entity
        </button>
      </h1>

      {!content.entities || content.entities?.length === 0 ? (
        <span>No entities in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(content.entities || [])?.map((entity, index) => {
          return (
            <AccordionItem
              key={index}
              itemKey={`item-${index}`}
              itemID={`item-${index}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{entity.name}</div>
                  <div className="description">{entity.cliDescription}</div>
                  <Arrow isUpward={state.isEnter} />
                  <button
                    className="mr-2"
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteEntityAt(index);
                    }}
                  >
                    Del
                  </button>
                </div>
              )}
            >
              <EntityEditor
                entity={entity}
                onChange={(entity) => updateEntity(entity, index)}
              />
            </AccordionItem>
          );
        })}
      </ControlledAccordion>
    </div>
  );
}
