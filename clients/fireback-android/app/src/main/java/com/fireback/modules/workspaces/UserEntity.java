package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class UserEntity extends JsonSerializable {
    public PersonEntity person;
    public String avatar;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Person person
    private MutableLiveData< PersonEntity > person = new MutableLiveData<>();
    public MutableLiveData< PersonEntity > getPerson() {
        return person;
    }
    public void setPerson( PersonEntity  v) {
        person.setValue(v);
    }
    // upper: Avatar avatar
    private MutableLiveData< String > avatar = new MutableLiveData<>();
    public MutableLiveData< String > getAvatar() {
        return avatar;
    }
    public void setAvatar( String  v) {
        avatar.setValue(v);
    }
    // Handling error message for each field
    // upper: Person person
    private MutableLiveData<String> personMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPersonMsg() {
        return personMsg;
    }
    public void setPersonMsg(String v) {
        personMsg.setValue(v);
    }
    // upper: Avatar avatar
    private MutableLiveData<String> avatarMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAvatarMsg() {
        return avatarMsg;
    }
    public void setAvatarMsg(String v) {
        avatarMsg.setValue(v);
    }
  }
}