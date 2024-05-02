package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class TokenEntity extends JsonSerializable {
    public UserEntity user;
    public String validUntil;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
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
    // Handling error message for each field
    // upper: User user
    private MutableLiveData<String> userMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUserMsg() {
        return userMsg;
    }
    public void setUserMsg(String v) {
        userMsg.setValue(v);
    }
    // upper: ValidUntil validUntil
    private MutableLiveData<String> validUntilMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValidUntilMsg() {
        return validUntilMsg;
    }
    public void setValidUntilMsg(String v) {
        validUntilMsg.setValue(v);
    }
  }
}