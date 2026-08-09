import { Checkbox } from "../../fireback-ui/components/checkbox/Checkbox";
import { ErrorsView } from "../../fireback-ui/components/error-view/ErrorView";
import { useCapabilitiesTreeActionQuery } from "../../sdk/abac/CapabilitiesTreeAction";
import { MCollection } from "../../sdk/sdk/common/operators";

/**
 * Moved from the old `sdk/core/react-tools.tsx` - this is the only consumer.
 */
interface CapabilityChild {
  uniqueId: string;
  children: CapabilityChild[];
}

type NodeChangeFn = (
  node: string,
  value: "checked" | "unchecked" | "indeterminate",
) => void;

export function RolePermissionTree({
  onChange,
  value,
  prefix,
}: {
  value: string[];
  onChange?: (value: string[]) => void;
  prefix?: string;
}) {
  const { data, error } = useCapabilitiesTreeActionQuery({});

  let items = data?.data?.item?.nested;

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
        {items instanceof MCollection ? (
          <PermissionTree
            items={items}
            onNodeChange={onNodeChange}
            value={value}
            prefix={prefix}
          />
        ) : null}
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
  items: MCollection<CapabilityChild>;
  value: string[];
  autoChecked?: boolean;
  onNodeChange: NodeChangeFn;
  prefix?: string;
}) {
  const pref = prefix ? prefix + "." : "";
  return (
    <>
      {(items.get() || []).map((item) => {
        const completeKey = `${pref}${item.uniqueId}${
          item.children?.length ? ".*" : ""
        }`;

        const checkValue: "checked" | "unchecked" | "indeterminate" = (
          value || []
        ).includes(completeKey)
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
                      checkValue === "checked" ? "unchecked" : "checked",
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
                  items={item.children as any}
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
