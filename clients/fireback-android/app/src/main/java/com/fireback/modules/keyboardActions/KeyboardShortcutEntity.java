/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.keyboardActions;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;
import com.fireback.modules.workspaces.*;

class KeyboardShortcutDefaultCombination extends JsonSerializable {
  public Boolean altKey;
  public String key;
  public Boolean metaKey;
  public Boolean shiftKey;
  public Boolean ctrlKey;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: AltKey altKey
    private MutableLiveData<Boolean> altKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getAltKey() {
      return altKey;
    }

    public void setAltKey(Boolean v) {
      altKey.setValue(v);
    }

    // upper: Key key
    private MutableLiveData<String> key = new MutableLiveData<>();

    public MutableLiveData<String> getKey() {
      return key;
    }

    public void setKey(String v) {
      key.setValue(v);
    }

    // upper: MetaKey metaKey
    private MutableLiveData<Boolean> metaKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getMetaKey() {
      return metaKey;
    }

    public void setMetaKey(Boolean v) {
      metaKey.setValue(v);
    }

    // upper: ShiftKey shiftKey
    private MutableLiveData<Boolean> shiftKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getShiftKey() {
      return shiftKey;
    }

    public void setShiftKey(Boolean v) {
      shiftKey.setValue(v);
    }

    // upper: CtrlKey ctrlKey
    private MutableLiveData<Boolean> ctrlKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getCtrlKey() {
      return ctrlKey;
    }

    public void setCtrlKey(Boolean v) {
      ctrlKey.setValue(v);
    }

    // Handling error message for each field
    // upper: AltKey altKey
    private MutableLiveData<String> altKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getAltKeyMsg() {
      return altKeyMsg;
    }

    public void setAltKeyMsg(String v) {
      altKeyMsg.setValue(v);
    }

    // upper: Key key
    private MutableLiveData<String> keyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getKeyMsg() {
      return keyMsg;
    }

    public void setKeyMsg(String v) {
      keyMsg.setValue(v);
    }

    // upper: MetaKey metaKey
    private MutableLiveData<String> metaKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getMetaKeyMsg() {
      return metaKeyMsg;
    }

    public void setMetaKeyMsg(String v) {
      metaKeyMsg.setValue(v);
    }

    // upper: ShiftKey shiftKey
    private MutableLiveData<String> shiftKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getShiftKeyMsg() {
      return shiftKeyMsg;
    }

    public void setShiftKeyMsg(String v) {
      shiftKeyMsg.setValue(v);
    }

    // upper: CtrlKey ctrlKey
    private MutableLiveData<String> ctrlKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getCtrlKeyMsg() {
      return ctrlKeyMsg;
    }

    public void setCtrlKeyMsg(String v) {
      ctrlKeyMsg.setValue(v);
    }
  }
}

class KeyboardShortcutUserCombination extends JsonSerializable {
  public Boolean altKey;
  public String key;
  public Boolean metaKey;
  public Boolean shiftKey;
  public Boolean ctrlKey;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: AltKey altKey
    private MutableLiveData<Boolean> altKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getAltKey() {
      return altKey;
    }

    public void setAltKey(Boolean v) {
      altKey.setValue(v);
    }

    // upper: Key key
    private MutableLiveData<String> key = new MutableLiveData<>();

    public MutableLiveData<String> getKey() {
      return key;
    }

    public void setKey(String v) {
      key.setValue(v);
    }

    // upper: MetaKey metaKey
    private MutableLiveData<Boolean> metaKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getMetaKey() {
      return metaKey;
    }

    public void setMetaKey(Boolean v) {
      metaKey.setValue(v);
    }

    // upper: ShiftKey shiftKey
    private MutableLiveData<Boolean> shiftKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getShiftKey() {
      return shiftKey;
    }

    public void setShiftKey(Boolean v) {
      shiftKey.setValue(v);
    }

    // upper: CtrlKey ctrlKey
    private MutableLiveData<Boolean> ctrlKey = new MutableLiveData<>();

    public MutableLiveData<Boolean> getCtrlKey() {
      return ctrlKey;
    }

    public void setCtrlKey(Boolean v) {
      ctrlKey.setValue(v);
    }

    // Handling error message for each field
    // upper: AltKey altKey
    private MutableLiveData<String> altKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getAltKeyMsg() {
      return altKeyMsg;
    }

    public void setAltKeyMsg(String v) {
      altKeyMsg.setValue(v);
    }

    // upper: Key key
    private MutableLiveData<String> keyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getKeyMsg() {
      return keyMsg;
    }

    public void setKeyMsg(String v) {
      keyMsg.setValue(v);
    }

    // upper: MetaKey metaKey
    private MutableLiveData<String> metaKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getMetaKeyMsg() {
      return metaKeyMsg;
    }

    public void setMetaKeyMsg(String v) {
      metaKeyMsg.setValue(v);
    }

    // upper: ShiftKey shiftKey
    private MutableLiveData<String> shiftKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getShiftKeyMsg() {
      return shiftKeyMsg;
    }

    public void setShiftKeyMsg(String v) {
      shiftKeyMsg.setValue(v);
    }

    // upper: CtrlKey ctrlKey
    private MutableLiveData<String> ctrlKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getCtrlKeyMsg() {
      return ctrlKeyMsg;
    }

    public void setCtrlKeyMsg(String v) {
      ctrlKeyMsg.setValue(v);
    }
  }
}

public class KeyboardShortcutEntity extends JsonSerializable {
  public String os;
  public String host;
  public KeyboardShortcutDefaultCombination defaultCombination;
  public KeyboardShortcutUserCombination userCombination;
  public String action;
  public String actionKey;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Os os
    private MutableLiveData<String> os = new MutableLiveData<>();

    public MutableLiveData<String> getOs() {
      return os;
    }

    public void setOs(String v) {
      os.setValue(v);
    }

    // upper: Host host
    private MutableLiveData<String> host = new MutableLiveData<>();

    public MutableLiveData<String> getHost() {
      return host;
    }

    public void setHost(String v) {
      host.setValue(v);
    }

    // upper: DefaultCombination defaultCombination
    private MutableLiveData<KeyboardShortcutDefaultCombination> defaultCombination =
        new MutableLiveData<>();

    public MutableLiveData<KeyboardShortcutDefaultCombination> getDefaultCombination() {
      return defaultCombination;
    }

    public void setDefaultCombination(KeyboardShortcutDefaultCombination v) {
      defaultCombination.setValue(v);
    }

    // upper: UserCombination userCombination
    private MutableLiveData<KeyboardShortcutUserCombination> userCombination =
        new MutableLiveData<>();

    public MutableLiveData<KeyboardShortcutUserCombination> getUserCombination() {
      return userCombination;
    }

    public void setUserCombination(KeyboardShortcutUserCombination v) {
      userCombination.setValue(v);
    }

    // upper: Action action
    private MutableLiveData<String> action = new MutableLiveData<>();

    public MutableLiveData<String> getAction() {
      return action;
    }

    public void setAction(String v) {
      action.setValue(v);
    }

    // upper: ActionKey actionKey
    private MutableLiveData<String> actionKey = new MutableLiveData<>();

    public MutableLiveData<String> getActionKey() {
      return actionKey;
    }

    public void setActionKey(String v) {
      actionKey.setValue(v);
    }

    // Handling error message for each field
    // upper: Os os
    private MutableLiveData<String> osMsg = new MutableLiveData<>();

    public MutableLiveData<String> getOsMsg() {
      return osMsg;
    }

    public void setOsMsg(String v) {
      osMsg.setValue(v);
    }

    // upper: Host host
    private MutableLiveData<String> hostMsg = new MutableLiveData<>();

    public MutableLiveData<String> getHostMsg() {
      return hostMsg;
    }

    public void setHostMsg(String v) {
      hostMsg.setValue(v);
    }

    // upper: DefaultCombination defaultCombination
    private MutableLiveData<String> defaultCombinationMsg = new MutableLiveData<>();

    public MutableLiveData<String> getDefaultCombinationMsg() {
      return defaultCombinationMsg;
    }

    public void setDefaultCombinationMsg(String v) {
      defaultCombinationMsg.setValue(v);
    }

    // upper: UserCombination userCombination
    private MutableLiveData<String> userCombinationMsg = new MutableLiveData<>();

    public MutableLiveData<String> getUserCombinationMsg() {
      return userCombinationMsg;
    }

    public void setUserCombinationMsg(String v) {
      userCombinationMsg.setValue(v);
    }

    // upper: Action action
    private MutableLiveData<String> actionMsg = new MutableLiveData<>();

    public MutableLiveData<String> getActionMsg() {
      return actionMsg;
    }

    public void setActionMsg(String v) {
      actionMsg.setValue(v);
    }

    // upper: ActionKey actionKey
    private MutableLiveData<String> actionKeyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getActionKeyMsg() {
      return actionKeyMsg;
    }

    public void setActionKeyMsg(String v) {
      actionKeyMsg.setValue(v);
    }
  }
}
