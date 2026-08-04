export const Arrow = ({ isUpward }: { isUpward: boolean }) => {
  return (
    <div className="arrow-container">
      <div className={`arrow ${isUpward ? "up" : "down"}`} />
    </div>
  );
};
