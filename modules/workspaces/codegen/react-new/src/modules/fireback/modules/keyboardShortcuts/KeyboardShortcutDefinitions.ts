import {
  KeyboardShortcutDefaultCombination,
  KeyboardShortcutEntity,
} from "../../sdk/modules/accessibility/KeyboardShortcutEntity";

export interface Shortcut extends KeyboardShortcutDefaultCombination {}
export interface KeyBinding extends KeyboardShortcutEntity {}
