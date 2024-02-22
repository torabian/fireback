import {
  KeyboardShortcutEntity,
  KeyboardShortcutDefaultCombinationEntity,
} from "src/sdk/fireback/modules/keyboardActions";

export interface Shortcut extends KeyboardShortcutDefaultCombinationEntity {}
export interface KeyBinding extends KeyboardShortcutEntity {}
