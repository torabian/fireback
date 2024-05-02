package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GeoCountryEntity extends JsonSerializable {
    public String status;
    public String flag;
    public String commonName;
    public String officialName;
    public static class VM extends ViewModel {
    // upper: Status status
    private MutableLiveData< String > status = new MutableLiveData<>();
    public MutableLiveData< String > getStatus() {
        return status;
    }
    public void setStatus( String  v) {
        status.setValue(v);
    }
    // upper: Flag flag
    private MutableLiveData< String > flag = new MutableLiveData<>();
    public MutableLiveData< String > getFlag() {
        return flag;
    }
    public void setFlag( String  v) {
        flag.setValue(v);
    }
    // upper: CommonName commonName
    private MutableLiveData< String > commonName = new MutableLiveData<>();
    public MutableLiveData< String > getCommonName() {
        return commonName;
    }
    public void setCommonName( String  v) {
        commonName.setValue(v);
    }
    // upper: OfficialName officialName
    private MutableLiveData< String > officialName = new MutableLiveData<>();
    public MutableLiveData< String > getOfficialName() {
        return officialName;
    }
    public void setOfficialName( String  v) {
        officialName.setValue(v);
    }
    }
}