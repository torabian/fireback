import {
  KeyboardShortcutDefaultCombination,
  KeyboardShortcutEntity,
} from "@/sdk/fireback/modules/keyboardActions/KeyboardShortcutEntity";

export interface Shortcut extends KeyboardShortcutDefaultCombination {}
export interface KeyBinding extends KeyboardShortcutEntity {}
