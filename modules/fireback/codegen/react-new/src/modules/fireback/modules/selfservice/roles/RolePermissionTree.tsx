import { Checkbox } from "@/modules/fireback/components/checkbox/Checkbox";
import { ErrorsView } from "@/modules/fireback/components/error-view/ErrorView";
import { type IndeterminateCheck } from "@/modules/fireback/definitions/definitions";
import { type CapabilityChild } from "@/modules/fireback/sdk/core/react-tools";
import { useCapabilitiesTreeActionQuery } from "@/modules/fireback/sdk/modules/fireback/CapabilitiesTree";

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

  const {data, error} = useCapabilitiesTreeActionQuery({});
  
  const items = data.data?.item?.nested || [];

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
      <ErrorsView error={error} />
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
  const pref = prefix ? prefix + "." : "";
  return (
    <>
      {items.map((item) => {
        const completeKey = `${pref}${item.uniqueId}${
          item.children?.length ? ".*" : ""
        }`;

        const checkValue: IndeterminateCheck = (value || []).includes(
          completeKey
        )
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
