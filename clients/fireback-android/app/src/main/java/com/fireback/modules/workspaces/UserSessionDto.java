package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserSessionDto extends JsonSerializable {
    public PassportEntity passport;
    public String token;
    public String exchangeKey;
    public UserWorkspaceEntity[] userWorkspaces;
    public UserEntity user;
    public String userId;
    public static class VM extends ViewModel {
    // upper: Passport passport
    private MutableLiveData< PassportEntity > passport = new MutableLiveData<>();
    public MutableLiveData< PassportEntity > getPassport() {
        return passport;
    }
    public void setPassport( PassportEntity  v) {
        passport.setValue(v);
    }
    // upper: Token token
    private MutableLiveData< String > token = new MutableLiveData<>();
    public MutableLiveData< String > getToken() {
        return token;
    }
    public void setToken( String  v) {
        token.setValue(v);
    }
    // upper: ExchangeKey exchangeKey
    private MutableLiveData< String > exchangeKey = new MutableLiveData<>();
    public MutableLiveData< String > getExchangeKey() {
        return exchangeKey;
    }
    public void setExchangeKey( String  v) {
        exchangeKey.setValue(v);
    }
    // upper: UserWorkspaces userWorkspaces
    private MutableLiveData< UserWorkspaceEntity[] > userWorkspaces = new MutableLiveData<>();
    public MutableLiveData< UserWorkspaceEntity[] > getUserWorkspaces() {
        return userWorkspaces;
    }
    public void setUserWorkspaces( UserWorkspaceEntity[]  v) {
        userWorkspaces.setValue(v);
    }
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: UserId userId
    private MutableLiveData< String > userId = new MutableLiveData<>();
    public MutableLiveData< String > getUserId() {
        return userId;
    }
    public void setUserId( String  v) {
        userId.setValue(v);
    }
    }
}