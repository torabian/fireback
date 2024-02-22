package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class SpaceTimeSnapshotEntity extends JsonSerializable {
    public float lat;
    public float lng;
    public float alt;
    public MovableObjectEntity movableObject;
    public static class VM extends ViewModel {
    // upper: Lat lat
    private MutableLiveData< Float > lat = new MutableLiveData<>();
    public MutableLiveData< Float > getLat() {
        return lat;
    }
    public void setLat( Float  v) {
        lat.setValue(v);
    }
    // upper: Lng lng
    private MutableLiveData< Float > lng = new MutableLiveData<>();
    public MutableLiveData< Float > getLng() {
        return lng;
    }
    public void setLng( Float  v) {
        lng.setValue(v);
    }
    // upper: Alt alt
    private MutableLiveData< Float > alt = new MutableLiveData<>();
    public MutableLiveData< Float > getAlt() {
        return alt;
    }
    public void setAlt( Float  v) {
        alt.setValue(v);
    }
    // upper: MovableObject movableObject
    private MutableLiveData< MovableObjectEntity > movableObject = new MutableLiveData<>();
    public MutableLiveData< MovableObjectEntity > getMovableObject() {
        return movableObject;
    }
    public void setMovableObject( MovableObjectEntity  v) {
        movableObject.setValue(v);
    }
    }
}