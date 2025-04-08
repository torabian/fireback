import { useGetCapabilities } from "../../sdk/modules/abac/useGetCapabilities";

export function QueryCapabilitiesTest() {
  const { items, query } = useGetCapabilities({});

  return (
    <ul>
      <li>{items.length}</li>
    </ul>
  );
}
