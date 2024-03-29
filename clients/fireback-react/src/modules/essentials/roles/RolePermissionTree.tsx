import { Checkbox } from "@/components/checkbox/Checkbox";
import { IndeterminateCheck } from "@/definitions/definitions";
import { useGetCapabilitiesTree } from "src/sdk/fireback/modules/workspaces/useGetCapabilitiesTree";
import { useQueryClient } from "react-query";
import { CapabilityChild } from "@/sdk/fireback/core/react-tools";

type NodeChangeFn = (node: string, value: IndeterminateCheck) => void;

export function RolePermissionTree({
  onChange,
  value,
  prefix,
}: {
  value: string[];
  onChange?: (value: string[]) => void;
  prefix?: string;
}) {
  const queryClient = useQueryClient();
  const { query: queryCapabilities } = useGetCapabilitiesTree({
    queryClient,
    query: { uniqueId: "all", itemsPerPage: 999 },
  });
  const items = queryCapabilities.data?.data?.nested || [];

  const onNodeChange: NodeChangeFn = (node, checkValue) => {
    let newValue: string[] = [...(value || [])];
    if (checkValue === "checked") {
      newValue.push(node);
    }
    if (checkValue === "unchecked") {
      newValue = newValue.filter((t) => t !== node);
    }
    onChange && onChange(newValue);
  };

  return (
    <nav className="tree-nav">
      <ul className="list">
        <PermissionTree
          items={items}
          onNodeChange={onNodeChange}
          value={value}
          prefix={prefix}
        />
      </ul>
    </nav>
  );
}

export function PermissionTree({
  items,
  onNodeChange,
  value,
  prefix,
  autoChecked,
}: {
  items: CapabilityChild[];
  value: string[];
  autoChecked?: boolean;
  onNodeChange: NodeChangeFn;
  prefix?: string;
}) {
  const pref = prefix ? prefix + "/" : "";
  return (
    <>
      {items.map((item) => {
        const completeKey = `${pref}${item.uniqueId}${
          item.children?.length ? "/*" : ""
        }`;

        const checkValue: IndeterminateCheck = value.includes(completeKey)
          ? "checked"
          : "unchecked";

        return (
          <li key={item.uniqueId}>
            <span>
              <label className={autoChecked ? "auto-checked" : ""}>
                <Checkbox
                  value={checkValue}
                  onChange={(e) => {
                    onNodeChange(
                      completeKey,
                      checkValue === "checked" ? "unchecked" : "checked"
                    );
                  }}
                />
                {item.uniqueId}
              </label>
            </span>
            {item.children && (
              <ul>
                <PermissionTree
                  autoChecked={autoChecked || checkValue === "checked"}
                  onNodeChange={onNodeChange}
                  value={value}
                  items={item.children}
                  prefix={pref + item.uniqueId}
                />
              </ul>
            )}
          </li>
        );
      })}
    </>
  );
}
