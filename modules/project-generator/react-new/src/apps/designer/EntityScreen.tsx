import { useRouter } from "@/modules/fireback/hooks/useRouter";

export function EntityScreen() {
  const { push } = useRouter();

  return (
    <>
      <div onClick={() => push("/")}>Modules</div>
      <div>Entity</div>
    </>
  );
}
