import { useUiState } from "../../hooks/uiStateContext";

export const OpenInNewRouter = ({ value }: { value?: string }) => {
  const { addRouter } = useUiState();

  const onClick = (e: any) => {
    e.stopPropagation();
    addRouter(value);
  };

  return (
    <div className="table-btn table-open-in-new-router" onClick={onClick}>
      <OpenIcon />
    </div>
  );
};

const OpenIcon = ({ size = 16, color = "silver", style = {} }) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    style={{ cursor: "pointer", ...style }}
  >
    <path
      d="M9 3H5C3.895 3 3 3.895 3 5v14c0 1.105.895 2 2 2h14c1.105 0 2-.895 2-2v-4h-2v4H5V5h4V3ZM21 3h-6v2h3.586l-9.293 9.293 1.414 1.414L20 6.414V10h2V3Z"
      fill={color}
    />
  </svg>
);
