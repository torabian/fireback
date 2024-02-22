package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PassportEntity extends JsonSerializable {
    public String type;
    public UserEntity user;
    public String value;
    public String password;
    public Boolean confirmed;
    public String accessToken;
    public static class VM extends ViewModel {
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    // upper: Confirmed confirmed
    private MutableLiveData< Boolean > confirmed = new MutableLiveData<>();
    public MutableLiveData< Boolean > getConfirmed() {
        return confirmed;
    }
    public void setConfirmed( Boolean  v) {
        confirmed.setValue(v);
    }
    // upper: AccessToken accessToken
    private MutableLiveData< String > accessToken = new MutableLiveData<>();
    public MutableLiveData< String > getAccessToken() {
        return accessToken;
    }
    public void setAccessToken( String  v) {
        accessToken.setValue(v);
    }
    }
}