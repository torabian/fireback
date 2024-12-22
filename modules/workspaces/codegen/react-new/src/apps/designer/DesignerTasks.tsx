import {
  AccordionItem,
  ControlledAccordion,
  useAccordionProvider,
} from "@szhsin/react-accordion";
import { Module2Task } from "./defs";
import { TaskEditor } from "./TaskEditor";

import { Arrow } from "./Arrow";

export function DesignerTasks({ content, setContent }: any) {
  const addTask = () => {
    content.tasks = [{ name: "" }, ...(content.tasks || [])];

    setContent({ ...content });
    setItem("item-0", { initialEntered: true });
  };

  const deleteTaskAt = (index: number) => {
    content.tasks = content.tasks?.filter((_, index2) => index !== index2);
    setContent({ ...content });
  };

  const updateTask = (task: Module2Task, index: number) => {
    setContent((c) => {
      c.tasks = (c.tasks || []).map((origin, ind) =>
        ind === index ? task : origin
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
    <div className="tasks-designer">
      <h1 className="section-bar">
        <span>Tasks ({content?.tasks?.length || 0})</span>{" "}
        <button className="btn btn-sm btn-primary" onClick={addTask}>
          Add task
        </button>
      </h1>

      {!content.tasks || content.tasks?.length === 0 ? (
        <span>No tasks in this file</span>
      ) : null}

      <ControlledAccordion providerValue={providerValue}>
        {(content.tasks || [])?.map((task, index) => {
          return (
            <AccordionItem
              key={index}
              itemKey={`item-${index}`}
              itemID={`item-${index}`}
              className={"accordion-item"}
              header={({ state }) => (
                <div className="minified-view">
                  <div className="name">{task.name}</div>
                  <div className="description">{task.description}</div>
                  <Arrow isUpward={state.isEnter} />
                  <button
                    className="mr-2"
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteTaskAt(index);
                    }}
                  >
                    Del
                  </button>
                </div>
              )}
            >
              <TaskEditor
                task={task}
                onChange={(task) => updateTask(task, index)}
              />
            </AccordionItem>
          );
        })}
      </ControlledAccordion>
    </div>
  );
}
