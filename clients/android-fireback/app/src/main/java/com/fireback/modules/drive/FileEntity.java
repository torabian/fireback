package com.fireback.modules.drive;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class FileEntity extends JsonSerializable {
    public String name;
    public String diskPath;
    public int size;
    public String virtualPath;
    public String type;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: DiskPath diskPath
    private MutableLiveData< String > diskPath = new MutableLiveData<>();
    public MutableLiveData< String > getDiskPath() {
        return diskPath;
    }
    public void setDiskPath( String  v) {
        diskPath.setValue(v);
    }
    // upper: Size size
    private MutableLiveData< Integer > size = new MutableLiveData<>();
    public MutableLiveData< Integer > getSize() {
        return size;
    }
    public void setSize( Integer  v) {
        size.setValue(v);
    }
    // upper: VirtualPath virtualPath
    private MutableLiveData< String > virtualPath = new MutableLiveData<>();
    public MutableLiveData< String > getVirtualPath() {
        return virtualPath;
    }
    public void setVirtualPath( String  v) {
        virtualPath.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    }
}