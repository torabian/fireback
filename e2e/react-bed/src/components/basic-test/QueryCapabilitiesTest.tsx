import { useGetCapabilities } from "../../sdk/modules/workspaces/useGetCapabilities";

export function QueryCapabilitiesTest() {
  const { items, query } = useGetCapabilities({});

  return (
    <ul>
      <li>{items.length}</li>
    </ul>
  );
}
