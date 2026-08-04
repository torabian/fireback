import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { Designer } from "./Designer";

export function ModuleScreen() {
  const { push } = useRouter();
  return (
    <>
      {/* <div onClick={() => push("/entity")}>New entity</div> */}
      <Designer />
    </>
  );
}
