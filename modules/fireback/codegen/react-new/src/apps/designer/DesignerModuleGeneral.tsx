import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module3 } from "./defs";

import { Arrow } from "./Arrow";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

export function DesignerModuleGeneral({
  content,
  setContent,
}: {
  content: Module3;
  setContent: any;
}) {
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

  const changeModuleName = (name: string) => {
    setContent((c) => {
      c.name = name;
      return { ...c };
    });
  };

  return (
    <div className="entities-designer">
      <ControlledAccordion providerValue={providerValue}>
        <AccordionItem
          className={"accordion-item"}
          header={({ state }) => (
            <div className="minified-view">
              <div className="name">{content.name}</div>
              <div className="description">{content.description}</div>
              <Arrow isUpward={state.isEnter} />
            </div>
          )}
        >
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
        </AccordionItem>
      </ControlledAccordion>
    </div>
  );
}
