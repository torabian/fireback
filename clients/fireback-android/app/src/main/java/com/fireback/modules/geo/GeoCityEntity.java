package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class GeoCityEntity extends JsonSerializable {
    public String name;
    public GeoProvinceEntity province;
    public GeoStateEntity state;
    public GeoCountryEntity country;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Province province
    private MutableLiveData< GeoProvinceEntity > province = new MutableLiveData<>();
    public MutableLiveData< GeoProvinceEntity > getProvince() {
        return province;
    }
    public void setProvince( GeoProvinceEntity  v) {
        province.setValue(v);
    }
    // upper: State state
    private MutableLiveData< GeoStateEntity > state = new MutableLiveData<>();
    public MutableLiveData< GeoStateEntity > getState() {
        return state;
    }
    public void setState( GeoStateEntity  v) {
        state.setValue(v);
    }
    // upper: Country country
    private MutableLiveData< GeoCountryEntity > country = new MutableLiveData<>();
    public MutableLiveData< GeoCountryEntity > getCountry() {
        return country;
    }
    public void setCountry( GeoCountryEntity  v) {
        country.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Province province
    private MutableLiveData<String> provinceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProvinceMsg() {
        return provinceMsg;
    }
    public void setProvinceMsg(String v) {
        provinceMsg.setValue(v);
    }
    // upper: State state
    private MutableLiveData<String> stateMsg = new MutableLiveData<>();
    public MutableLiveData<String> getStateMsg() {
        return stateMsg;
    }
    public void setStateMsg(String v) {
        stateMsg.setValue(v);
    }
    // upper: Country country
    private MutableLiveData<String> countryMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCountryMsg() {
        return countryMsg;
    }
    public void setCountryMsg(String v) {
        countryMsg.setValue(v);
    }
  }
}