package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PersonEntity extends JsonSerializable {
    public String firstName;
    public String lastName;
    public String photo;
    public String gender;
    public String title;
    public java.util.Date birthDate;
    public static class VM extends ViewModel {
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
    }
}