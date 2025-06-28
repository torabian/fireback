// AnimatedRouteWrapper.tsx
import { AnimatePresence, motion } from "framer-motion";
import { useLocation } from "react-router-dom";

const swipeVariants = {
  initial: { x: "100%", opacity: 0 },
  animate: { x: 0, opacity: 1 },
  exit: { x: "100%", opacity: 0 },
};

export function AnimatedRouteWrapper({
  children,
}: {
  children: React.ReactNode;
}) {
  const location = useLocation();
  const animated = location.state?.animated;

  if (!animated) {
    return <>{children}</>; // No animation, just render
  }

  return (
    <AnimatePresence mode="wait">
      <motion.div
        key={location.pathname}
        variants={swipeVariants}
        initial="initial"
        animate="animate"
        exit="exit"
        transition={{ duration: 0.15 }}
        style={{ width: "100%" }}
      >
        {children}
      </motion.div>
    </AnimatePresence>
  );
}
