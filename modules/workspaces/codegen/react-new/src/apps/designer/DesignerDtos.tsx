import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module2Dto } from "./defs";
import { DtoEditor } from "./DtoEditor";

import { Arrow } from "./Arrow";

export function DesignerDtos({ content, setContent }: any) {
  const addDto = () => {
    content.dtos = [{ name: "" }, ...(content.dtos || [])];

    setContent({ ...content });
    setItem("item-0", { initialEntered: true });
  };

  const deleteDtoAt = (index: number) => {
    content.dtos = content.dtos?.filter((_, index2) => index !== index2);
    setContent({ ...content });
  };

  const updateDto = (dto: Module2Dto, index: number) => {
    setContent((c) => {
      c.dtos = (c.dtos || []).map((origin, ind) =>
        ind === index ? dto : origin
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
    <div className="dtos-designer">
      <h1 className="section-bar">
        <span>Dtos ({content?.dtos?.length || 0})</span>{" "}
        <button className="btn btn-sm btn-primary" onClick={addDto}>
          Add dto
        </button>
      </h1>

      {!content.dtos || content.dtos?.length === 0 ? (
        <span>No dtos in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(content.dtos || [])?.map((dto, index) => {
          return (
            <AccordionItem
              key={index}
              itemKey={`item-${index}`}
              itemID={`item-${index}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{dto.name}</div>
                  <div className="description">{dto.description}</div>
                  <Arrow isUpward={state.isEnter} />
                  <button
                    className="mr-2"
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteDtoAt(index);
                    }}
                  >
                    Del
                  </button>
                </div>
              )}
            >
              <DtoEditor dto={dto} onChange={(dto) => updateDto(dto, index)} />
            </AccordionItem>
          );
        })}
      </ControlledAccordion>
    </div>
  );
}
