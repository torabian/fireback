package com.fireback.modules.widget;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class WidgetEntity extends JsonSerializable {
    public String name;
    public String family;
    public String providerKey;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Family family
    private MutableLiveData< String > family = new MutableLiveData<>();
    public MutableLiveData< String > getFamily() {
        return family;
    }
    public void setFamily( String  v) {
        family.setValue(v);
    }
    // upper: ProviderKey providerKey
    private MutableLiveData< String > providerKey = new MutableLiveData<>();
    public MutableLiveData< String > getProviderKey() {
        return providerKey;
    }
    public void setProviderKey( String  v) {
        providerKey.setValue(v);
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
    // upper: Family family
    private MutableLiveData<String> familyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFamilyMsg() {
        return familyMsg;
    }
    public void setFamilyMsg(String v) {
        familyMsg.setValue(v);
    }
    // upper: ProviderKey providerKey
    private MutableLiveData<String> providerKeyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProviderKeyMsg() {
        return providerKeyMsg;
    }
    public void setProviderKeyMsg(String v) {
        providerKeyMsg.setValue(v);
    }
  }
}