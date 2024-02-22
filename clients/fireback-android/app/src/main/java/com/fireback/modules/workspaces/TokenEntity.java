package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class TokenEntity extends JsonSerializable {
    public UserEntity user;
    public String validUntil;
    public static class VM extends ViewModel {
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: ValidUntil validUntil
    private MutableLiveData< String > validUntil = new MutableLiveData<>();
    public MutableLiveData< String > getValidUntil() {
        return validUntil;
    }
    public void setValidUntil( String  v) {
        validUntil.setValue(v);
    }
    }
}