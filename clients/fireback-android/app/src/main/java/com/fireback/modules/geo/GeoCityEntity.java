package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GeoCityEntity extends JsonSerializable {
    public String name;
    public GeoProvinceEntity province;
    public GeoStateEntity state;
    public GeoCountryEntity country;
    public static class VM extends ViewModel {
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
    }
}