import { Toast } from "../../hooks/toast";

export const CopyCell = ({ value }: { value?: string }) => {
  const onCopy = (e: any) => {
    e.stopPropagation();
    navigator.clipboard
      .writeText(value)
      .then(() => {
        Toast(`Copied ${value}`, { type: "info", autoClose: 600 });
      })
      .catch((err) => {
        Toast(`Copy failed.`, { type: "error", autoClose: 600 });
      });
  };

  return (
    <div className="table-btn table-copy-btn" onClick={onCopy}>
      <CopyIcon />
    </div>
  );
};

const CopyIcon = ({ size = 16, color = "silver", style = {} }) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    style={style}
  >
    <path
      d="M16 1H6C4.9 1 4 1.9 4 3V17H6V3H16V1ZM18 5H10C8.9 5 8 5.9 8 7V21C8 22.1 8.9 23 10 23H18C19.1 23 20 22.1 20 21V7C20 5.9 19.1 5 18 5ZM18 21H10V7H18V21Z"
      fill={color}
    />
  </svg>
);
