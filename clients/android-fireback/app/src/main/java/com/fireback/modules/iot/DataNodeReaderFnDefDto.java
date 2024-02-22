package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DataNodeReaderFnDefDto extends JsonSerializable {
    public String fn;
    public String description;
    public String reads;
    public String writes;
    public static class VM extends ViewModel {
    // upper: Fn fn
    private MutableLiveData< String > fn = new MutableLiveData<>();
    public MutableLiveData< String > getFn() {
        return fn;
    }
    public void setFn( String  v) {
        fn.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: Reads reads
    private MutableLiveData< String > reads = new MutableLiveData<>();
    public MutableLiveData< String > getReads() {
        return reads;
    }
    public void setReads( String  v) {
        reads.setValue(v);
    }
    // upper: Writes writes
    private MutableLiveData< String > writes = new MutableLiveData<>();
    public MutableLiveData< String > getWrites() {
        return writes;
    }
    public void setWrites( String  v) {
        writes.setValue(v);
    }
    }
}