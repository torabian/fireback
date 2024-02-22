package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class IoWriterDto extends JsonSerializable {
    public String content;
    public String type;
    public String host;
    public String path;
    public static class VM extends ViewModel {
    // upper: Content content
    private MutableLiveData< String > content = new MutableLiveData<>();
    public MutableLiveData< String > getContent() {
        return content;
    }
    public void setContent( String  v) {
        content.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Host host
    private MutableLiveData< String > host = new MutableLiveData<>();
    public MutableLiveData< String > getHost() {
        return host;
    }
    public void setHost( String  v) {
        host.setValue(v);
    }
    // upper: Path path
    private MutableLiveData< String > path = new MutableLiveData<>();
    public MutableLiveData< String > getPath() {
        return path;
    }
    public void setPath( String  v) {
        path.setValue(v);
    }
    }
}