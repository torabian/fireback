package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class HmiComponentTypeEntity extends JsonSerializable {
    public String name;
    public Boolean isDirectInteractable;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: IsDirectInteractable isDirectInteractable
    private MutableLiveData< Boolean > isDirectInteractable = new MutableLiveData<>();
    public MutableLiveData< Boolean > getIsDirectInteractable() {
        return isDirectInteractable;
    }
    public void setIsDirectInteractable( Boolean  v) {
        isDirectInteractable.setValue(v);
    }
    }
}