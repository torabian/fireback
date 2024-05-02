package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class PersonEntity extends JsonSerializable {
    public String firstName;
    public String lastName;
    public String photo;
    public String gender;
    public String title;
    public java.util.Date birthDate;
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
    // upper: Photo photo
    private MutableLiveData< String > photo = new MutableLiveData<>();
    public MutableLiveData< String > getPhoto() {
        return photo;
    }
    public void setPhoto( String  v) {
        photo.setValue(v);
    }
    // upper: Gender gender
    private MutableLiveData< String > gender = new MutableLiveData<>();
    public MutableLiveData< String > getGender() {
        return gender;
    }
    public void setGender( String  v) {
        gender.setValue(v);
    }
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: BirthDate birthDate
    private MutableLiveData< java.util.Date > birthDate = new MutableLiveData<>();
    public MutableLiveData< java.util.Date > getBirthDate() {
        return birthDate;
    }
    public void setBirthDate( java.util.Date  v) {
        birthDate.setValue(v);
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
    // upper: Photo photo
    private MutableLiveData<String> photoMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPhotoMsg() {
        return photoMsg;
    }
    public void setPhotoMsg(String v) {
        photoMsg.setValue(v);
    }
    // upper: Gender gender
    private MutableLiveData<String> genderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getGenderMsg() {
        return genderMsg;
    }
    public void setGenderMsg(String v) {
        genderMsg.setValue(v);
    }
    // upper: Title title
    private MutableLiveData<String> titleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTitleMsg() {
        return titleMsg;
    }
    public void setTitleMsg(String v) {
        titleMsg.setValue(v);
    }
    // upper: BirthDate birthDate
    private MutableLiveData<String> birthDateMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBirthDateMsg() {
        return birthDateMsg;
    }
    public void setBirthDateMsg(String v) {
        birthDateMsg.setValue(v);
    }
  }
}