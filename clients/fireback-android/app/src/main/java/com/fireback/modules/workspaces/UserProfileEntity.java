package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class UserProfileEntity extends JsonSerializable {
    public String firstName;
    public String lastName;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: FirstName firstName
    private MutableLiveData< String > firstName = new MutableLiveData<>();
    public MutableLiveData< String > getFirstName() {
        return firstName;
    }
    public void setFirstName( String  v) {
        firstName.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData< String > lastName = new MutableLiveData<>();
    public MutableLiveData< String > getLastName() {
        return lastName;
    }
    public void setLastName( String  v) {
        lastName.setValue(v);
    }
    // Handling error message for each field
    // upper: FirstName firstName
    private MutableLiveData<String> firstNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFirstNameMsg() {
        return firstNameMsg;
    }
    public void setFirstNameMsg(String v) {
        firstNameMsg.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData<String> lastNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLastNameMsg() {
        return lastNameMsg;
    }
    public void setLastNameMsg(String v) {
        lastNameMsg.setValue(v);
    }
  }
}