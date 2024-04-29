import { StyleSheet } from "react-native";

export const theme = {
  activeColor: "#00ff00",
  inactiveColor: "#ff0000",
  idleColor: "#ffff00",
};

export const themeEl = StyleSheet.create({
  textLabel: {
    fontWeight: "bold",
    marginRight: 5,
  },
  keyPairRow: {
    flexDirection: "row",
  },
  h1: {
    fontSize: 26,
    fontWeight: "bold",
    marginLeft: 10,
    marginRight: 10,
    marginTop: 10,
    marginBottom: 16,
  },
});
