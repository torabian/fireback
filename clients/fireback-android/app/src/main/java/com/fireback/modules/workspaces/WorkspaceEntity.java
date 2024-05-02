package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class WorkspaceEntity extends JsonSerializable {
    public String description;
    public String name;
    public WorkspaceTypeEntity type;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< WorkspaceTypeEntity > type = new MutableLiveData<>();
    public MutableLiveData< WorkspaceTypeEntity > getType() {
        return type;
    }
    public void setType( WorkspaceTypeEntity  v) {
        type.setValue(v);
    }
    // Handling error message for each field
    // upper: Description description
    private MutableLiveData<String> descriptionMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDescriptionMsg() {
        return descriptionMsg;
    }
    public void setDescriptionMsg(String v) {
        descriptionMsg.setValue(v);
    }
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
  }
}