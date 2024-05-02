package com.fireback.modules.commonprofile;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class CommonProfileEntity extends JsonSerializable {
    public String firstName;
    public String lastName;
    public String phoneNumber;
    public String email;
    public String company;
    public String street;
    public String houseNumber;
    public String zipCode;
    public String city;
    public String gender;
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
    // upper: PhoneNumber phoneNumber
    private MutableLiveData< String > phoneNumber = new MutableLiveData<>();
    public MutableLiveData< String > getPhoneNumber() {
        return phoneNumber;
    }
    public void setPhoneNumber( String  v) {
        phoneNumber.setValue(v);
    }
    // upper: Email email
    private MutableLiveData< String > email = new MutableLiveData<>();
    public MutableLiveData< String > getEmail() {
        return email;
    }
    public void setEmail( String  v) {
        email.setValue(v);
    }
    // upper: Company company
    private MutableLiveData< String > company = new MutableLiveData<>();
    public MutableLiveData< String > getCompany() {
        return company;
    }
    public void setCompany( String  v) {
        company.setValue(v);
    }
    // upper: Street street
    private MutableLiveData< String > street = new MutableLiveData<>();
    public MutableLiveData< String > getStreet() {
        return street;
    }
    public void setStreet( String  v) {
        street.setValue(v);
    }
    // upper: HouseNumber houseNumber
    private MutableLiveData< String > houseNumber = new MutableLiveData<>();
    public MutableLiveData< String > getHouseNumber() {
        return houseNumber;
    }
    public void setHouseNumber( String  v) {
        houseNumber.setValue(v);
    }
    // upper: ZipCode zipCode
    private MutableLiveData< String > zipCode = new MutableLiveData<>();
    public MutableLiveData< String > getZipCode() {
        return zipCode;
    }
    public void setZipCode( String  v) {
        zipCode.setValue(v);
    }
    // upper: City city
    private MutableLiveData< String > city = new MutableLiveData<>();
    public MutableLiveData< String > getCity() {
        return city;
    }
    public void setCity( String  v) {
        city.setValue(v);
    }
    // upper: Gender gender
    private MutableLiveData< String > gender = new MutableLiveData<>();
    public MutableLiveData< String > getGender() {
        return gender;
    }
    public void setGender( String  v) {
        gender.setValue(v);
    }
    }
}