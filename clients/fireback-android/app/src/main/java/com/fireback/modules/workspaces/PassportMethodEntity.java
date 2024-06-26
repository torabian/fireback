/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class PassportMethodEntity extends JsonSerializable {
  public String name;
  public String type;
  public String region;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData<String> name = new MutableLiveData<>();

    public MutableLiveData<String> getName() {
      return name;
    }

    public void setName(String v) {
      name.setValue(v);
    }

    // upper: Type type
    private MutableLiveData<String> type = new MutableLiveData<>();

    public MutableLiveData<String> getType() {
      return type;
    }

    public void setType(String v) {
      type.setValue(v);
    }

    // upper: Region region
    private MutableLiveData<String> region = new MutableLiveData<>();

    public MutableLiveData<String> getRegion() {
      return region;
    }

    public void setRegion(String v) {
      region.setValue(v);
    }

    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();

    public MutableLiveData<String> getNameMsg() {
      return nameMsg;
    }

    public void setNameMsg(String v) {
      nameMsg.setValue(v);
    }

    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();

    public MutableLiveData<String> getTypeMsg() {
      return typeMsg;
    }

    public void setTypeMsg(String v) {
      typeMsg.setValue(v);
    }

    // upper: Region region
    private MutableLiveData<String> regionMsg = new MutableLiveData<>();

    public MutableLiveData<String> getRegionMsg() {
      return regionMsg;
    }

    public void setRegionMsg(String v) {
      regionMsg.setValue(v);
    }
  }
}
