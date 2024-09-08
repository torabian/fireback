import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { DesignerActions } from "./DesignerActions";
import { DesignerDtos } from "./DesignerDtos";
import { DesignerEntities } from "./DesignerEntities";
import { DesignerModuleGeneral } from "./DesignerModuleGeneral";
import { DesignerTasks } from "./DesignerTasks";
import { useDesigner } from "./useDesigner";
import { useState } from "react";
import { Designerconfig } from "./DesignerConfigs";

export function Designer() {
  const { content, setContent } = useDesigner();
  const [search, setSearch] = useState("");

  return (
    <div className="interactive container">
      {/* <FormText
        label="Quick find..."
        onChange={(v) => setSearch(v)}
        value={search}
      ></FormText> */}

      <DesignerModuleGeneral content={content} setContent={setContent} />

      <DesignerEntities
        content={content}
        search={search}
        setContent={setContent}
      ></DesignerEntities>
      <DesignerActions
        search={search}
        content={content}
        setContent={setContent}
      ></DesignerActions>
      <DesignerDtos
        search={search}
        content={content}
        setContent={setContent}
      ></DesignerDtos>
      <DesignerTasks
        search={search}
        content={content}
        setContent={setContent}
      ></DesignerTasks>

      <Designerconfig
        content={content}
        search={search}
        setContent={setContent}
      ></Designerconfig>
    </div>
  );
}
