import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module2Entity } from "./defs";
import { EntityEditor } from "./EntityEditor";

import { Arrow } from "./Arrow";
import { deburr } from "lodash";

function tokenize(keyword: string): string {
  return deburr(keyword).toLowerCase();
}

function searchInMemory(entities: Module2Entity[], search): Module2Entity[] {
  if (!search) {
    return entities;
  }

  return entities.filter((entity) =>
    [entity?.name, entity.description]
      .map(tokenize)
      .join(" - ")
      .includes(tokenize(search))
  );
}

export function DesignerEntities({ content, setContent, search }: any) {
  const addEntity = () => {
    content.entities = [{ name: "" }, ...(content.entities || [])];

    setContent({ ...content });
    setItem("item-0", { initialEntered: true });
  };

  const deleteEntityAt = (index: number) => {
    content.entities = content.entities?.filter(
      (_, index2) => index !== index2
    );
    setContent({ ...content });
  };

  const updateEntity = (entity: Module2Entity, index: number) => {
    setContent((c) => {
      c.entities = (c.entities || []).map((origin, ind) =>
        ind === index ? entity : origin
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
  const entities = searchInMemory(content.entities, search);

  return (
    <div className="entities-designer">
      <h1 className="section-bar">
        <span>Entities ({content?.entities?.length || 0})</span>{" "}
        <button className="btn btn-sm btn-primary" onClick={addEntity}>
          Add entity
        </button>
      </h1>

      {!entities || entities?.length === 0 ? (
        <span>No entities in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(entities || [])?.map((entity, index) => {
          return (
            <AccordionItem
              key={entity.name + "_" + entity.description}
              itemKey={`item-${entity.name + "_" + entity.description}`}
              itemID={`item-${entity.name + "_" + entity.description}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{entity.name}</div>
                  <div className="description">{entity.description}</div>
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
