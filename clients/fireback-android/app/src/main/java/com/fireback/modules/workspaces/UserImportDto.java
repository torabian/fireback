package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class UserImportPassports extends JsonSerializable {
    public String value;
    public String password;
}
class UserImportAddress extends JsonSerializable {
    public String street;
    public String zipCode;
    public String city;
    public String country;
}
public class UserImportDto extends JsonSerializable {
    public String avatar;
    public UserImportPassports[] passports;
    public PersonEntity person;
    public UserImportAddress address;
    public static class VM extends ViewModel {
    // upper: Avatar avatar
    private MutableLiveData< String > avatar = new MutableLiveData<>();
    public MutableLiveData< String > getAvatar() {
        return avatar;
    }
    public void setAvatar( String  v) {
        avatar.setValue(v);
    }
    // upper: Passports passports
    private MutableLiveData< UserImportPassports[] > passports = new MutableLiveData<>();
    public MutableLiveData< UserImportPassports[] > getPassports() {
        return passports;
    }
    public void setPassports( UserImportPassports[]  v) {
        passports.setValue(v);
    }
    // upper: Person person
    private MutableLiveData< PersonEntity > person = new MutableLiveData<>();
    public MutableLiveData< PersonEntity > getPerson() {
        return person;
    }
    public void setPerson( PersonEntity  v) {
        person.setValue(v);
    }
    // upper: Address address
    private MutableLiveData< UserImportAddress > address = new MutableLiveData<>();
    public MutableLiveData< UserImportAddress > getAddress() {
        return address;
    }
    public void setAddress( UserImportAddress  v) {
        address.setValue(v);
    }
    }
}