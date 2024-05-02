package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PaymentStatusEntity extends JsonSerializable {
    public String name;
    public String description;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    }
}