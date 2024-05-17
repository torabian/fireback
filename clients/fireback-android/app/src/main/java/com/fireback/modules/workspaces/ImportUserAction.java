package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class ImportUserAction {
    public static class Req extends JsonSerializable {
    public String path;
    // upper: Path path
    private MutableLiveData<String> pathMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPathMsg() {
        return pathMsg;
    }
    public void setPathMsg(String v) {
        pathMsg.setValue(v);
    }
    }
    public static class ReqViewModel extends ViewModel {
    // upper: Path path
    private MutableLiveData< String > path = new MutableLiveData<>();
    public MutableLiveData< String > getPath() {
        return path;
    }
    public void setPath( String  v) {
        path.setValue(v);
    }
    // upper: Path path
    private MutableLiveData<String> pathMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPathMsg() {
        return pathMsg;
    }
    public void setPathMsg(String v) {
        pathMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("path")) {
            this.setPathMsg(item.messageTranslated);
        }
    });
}
    }
}